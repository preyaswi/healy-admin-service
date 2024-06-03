package usecase

import (
	"errors"
	"fmt"
	clientinterface "healy-admin/pkg/client/interface"
	"healy-admin/pkg/config"
	"healy-admin/pkg/domain"
	"healy-admin/pkg/helper"
	interfaces "healy-admin/pkg/repository/interface"
	services "healy-admin/pkg/usecase/interface"
	"healy-admin/pkg/utils/models"
	"strings"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/razorpay/razorpay-go"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository   interfaces.AdminRepository
	doctorRepository  clientinterface.NewDoctorClient
	patientRepository clientinterface.NewPatientClient
}

func NewAdminUseCase(repository interfaces.AdminRepository, doctorRepo clientinterface.NewDoctorClient, patientrepo clientinterface.NewPatientClient) services.AdminUseCase {
	return &adminUseCase{
		adminRepository:   repository,
		doctorRepository:  doctorRepo,
		patientRepository: patientrepo,
	}
}
func (ad *adminUseCase) AdminSignUp(admin models.AdminSignUp) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistsByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenAdmin{}, errors.New("admin with this email is already exists")
	}
	hashPassword, err := helper.PasswordHash(admin.Password)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error in hashing password")
	}
	admin.Password = hashPassword
	admindata, err := ad.adminRepository.AdminSignUp(admin)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("could not add the user")
	}
	tokenString, err := helper.GenerateTokenAdmin(admindata)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: admindata,
		Token: tokenString,
	}, nil
}

func (ad *adminUseCase) LoginHandler(admin models.AdminLogin) (*domain.TokenAdmin, error) {
	email, err := ad.adminRepository.CheckAdminExistsByEmail(admin.Email)
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenAdmin{}, errors.New("email doesn't exist")
	}
	admindeatils, err := ad.adminRepository.FindAdminByEmail(admin)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}
	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(admindeatils.Password), []byte(admin.Password))
	if err != nil {
		return &domain.TokenAdmin{}, errors.New("password not matching")
	}
	var adminDetailsResponse models.AdminDetailsResponse
	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &admindeatils)
	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return &domain.TokenAdmin{}, err
	}

	return &domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil
}
func (ad *adminUseCase) AddToBooking(patientid, doctorid int) error {
	ok, err := ad.doctorRepository.CheckDoctor(doctorid)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("doctor doesn't exist")
	}
	doctordetail, err := ad.doctorRepository.DoctorDetailforBooking(doctorid)
	if err != nil {
		return err
	}
	err = ad.adminRepository.AddToBooking(patientid, doctordetail)
	if err != nil {
		return err
	}
	return nil

}
func (ad *adminUseCase) CancelBooking(patientid, bookingid int) error {
	booking, err := ad.adminRepository.GetBookingByID(bookingid)
	if err != nil {
		return err
	}
	if booking.PatientId != uint(patientid) {
		return errors.New("unauthorized: patient ID does not match booking")
	}
	return ad.adminRepository.RemoveBooking(bookingid)

}
func (ad *adminUseCase) MakePaymentRazorpay(bookingid int) (domain.Booking, string, error) {
	cfg, _ := config.LoadConfig()
	paymentdetail, err := ad.adminRepository.GetBookingByID(bookingid)
	if err != nil {
		return domain.Booking{}, "", err
	}
	client := razorpay.NewClient(cfg.KEY_ID_fOR_PAY, cfg.SECRET_KEY_FOR_PAY)
	data := map[string]interface{}{
		"amount":   int(paymentdetail.Fees) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Booking{}, "", err
	}
	RazorpayorderId := body["id"].(string)
	err = ad.adminRepository.AddRazorPayDetails(paymentdetail.BookingId, RazorpayorderId)
	if err != nil {
		return domain.Booking{}, "", err
	}
	return paymentdetail, RazorpayorderId, nil
}
func (ad *adminUseCase) VerifyPayment(booking_id int) error {
	status, err := ad.adminRepository.CheckPaymentStatus(booking_id)
	if err != nil {
		return err
	}
	status = strings.TrimSpace(strings.ToLower(status))
	if status == "not paid" {
		if err := ad.adminRepository.UpdatePaymentStatus(booking_id, "paid"); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("already paid")
	}
}

func (ad *adminUseCase) GetPaidPatients(doctor_id int) ([]models.Patient, error) {
	bookings, err := ad.adminRepository.GetPaidBookingsByDoctorID(doctor_id)
	if err != nil {
		return nil, err
	}
	patientChan := make(chan models.Patient)
	errorChan := make(chan error)
	var wg sync.WaitGroup

	for _, booking := range bookings {
		wg.Add(1)
		go func(patientid int) {
			defer wg.Done()
			patient, err := ad.patientRepository.GetPatientByID(patientid)

			if err != nil {
				errorChan <- err
				return
			}
			patientChan <- patient

		}(int(booking.PatientId))
	}
	go func() {
		wg.Wait()
		close(patientChan)
		close(errorChan)
	}()
	var patients []models.Patient
	for {
		select {
		case patient, ok := <-patientChan:
			if ok {
				patients = append(patients, patient)
			}
		case err, ok := <-errorChan:
			if ok {
				return nil, err
			}
		}
		if len(patientChan) == 0 && len(errorChan) == 0 {
			break
		}
	}

	return patients, nil
}
func (ad *adminUseCase) CreatePrescription(prescription models.PrescriptionRequest) (domain.Prescription, error) {
	paid, err := ad.adminRepository.CheckPatientPayment(prescription.DoctorID, prescription.PatientID)
	if err != nil {
		return domain.Prescription{}, fmt.Errorf("error checking payment status")
	}

	if !paid {
		return domain.Prescription{}, fmt.Errorf("patient has not paid the booking fee")
	}

	createdPrescription, err := ad.adminRepository.CreatePrescription(prescription)
    if err != nil {
        return domain.Prescription{}, fmt.Errorf("error creating prescription")
    }

    return createdPrescription, nil
}
