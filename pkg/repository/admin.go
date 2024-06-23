package repository

import (
	"errors"
	"fmt"
	"healy-admin/pkg/domain"
	interfaces "healy-admin/pkg/repository/interface"
	"healy-admin/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var model models.AdminDetailsResponse
	if err := ad.DB.Raw("INSERT INTO admins (firstname,lastname,email,password) VALUES (?, ?, ? ,?) RETURNING id,firstname,lastname,email", adminDetails.Firstname, adminDetails.Lastname, adminDetails.Email, adminDetails.Password).Scan(&model).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}
	return model, nil
}
func (ad *adminRepository) CheckAdminExistsByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	res := ad.DB.Where(&domain.Admin{Email: email}).First(&admin)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Admin{}, res.Error
	}
	return &admin, nil
}

func (ad *adminRepository) FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error) {
	var user models.AdminSignUp
	err := ad.DB.Raw("SELECT * FROM admins WHERE email=? ", admin.Email).Scan(&user).Error
	if err != nil {
		return models.AdminSignUp{}, errors.New("error checking user details")
	}
	return user, nil
}

func (ad *adminRepository) AddToBooking(patientid string, doctordetail models.BookingDoctorDetails) error {
	err := ad.DB.Exec("insert into bookings(patient_id,doctor_id,doctor_name,doctor_email,fees)values(?,?,?,?,?)", patientid, doctordetail.Doctorid, doctordetail.DoctorName, doctordetail.DoctorEmail, doctordetail.Fees).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminRepository) GetBookingByID(bookingid int) (domain.Booking, error) {
	var booking domain.Booking
	err := ad.DB.Where("booking_id=?", bookingid).First(&booking).Error
	if err != nil {
		return domain.Booking{}, err
	}
	return booking, nil
}
func (ad *adminRepository) RemoveBooking(bookingID int) error {
	err := ad.DB.Where("booking_id=?", bookingID).Delete(&domain.Booking{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminRepository) AddRazorPayDetails(booking_id uint, razorPaypaymentID string) error {
	err := ad.DB.Exec("insert into razer_pays (booking_id,razor_id) values (?,?)", booking_id, razorPaypaymentID).Error
	if err != nil {
		return err
	}
	return nil

}
func (ad *adminRepository) CheckPaymentStatus(bookingid int) (string, error) {
	var payment domain.Booking
	if err := ad.DB.First(&payment, bookingid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("payment not found")
		}
		return "", err
	}
	return payment.PaymentStatus, nil
}
func (ad *adminRepository) UpdatePaymentStatus(booking_id int, status string) error {
	err := ad.DB.Model(&domain.Booking{}).Where("booking_id = ?", booking_id).Update("payment_status", status).Error
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminRepository) GetPaidBookingsByDoctorID(doctorId int) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := ad.DB.Where("doctor_id = ? AND payment_status = ?", doctorId, "paid").Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
func (ad *adminRepository) CheckPatientPayment(doctorID int, patientID string) (bool, error) {
	var booking domain.Booking
	err := ad.DB.Where("doctor_id = ? AND patient_id = ? AND payment_status = ?", doctorID, patientID, "paid").First(&booking).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("error checking payment status")
	}

	return true, nil
}

func (ad *adminRepository) CreatePrescription(prescription models.PrescriptionRequest) (domain.Prescription, error) {
	prescriptionModel := domain.Prescription{
		BookingID: uint(prescription.BookingID),
		DoctorID:  uint(prescription.DoctorID),
		PatientID: prescription.PatientID,
		Medicine:  prescription.Medicine,
		Dosage:    prescription.Dosage,
		Notes:     prescription.Notes,
	}

	if err := ad.DB.Create(&prescriptionModel).Error; err != nil {
		return domain.Prescription{}, fmt.Errorf("error saving prescription")
	}

	// Retrieve the created prescription with all fields populated
	var createdPrescription domain.Prescription
	if err := ad.DB.First(&createdPrescription, prescriptionModel.ID).Error; err != nil {
		return domain.Prescription{}, fmt.Errorf("error retrieving created prescription")
	}

	return createdPrescription, nil
}
func (ad *adminRepository) SetDoctorAvailability(availability models.SetAvailability) (string, error) {
	var slots []domain.Availability
	currentTime := availability.StartTime
	for currentTime.Before(availability.EndTime) {
		// Define the end time for the current slot
		slotEndTime := currentTime.Add(30 * time.Minute)

		// Check for overlapping slots
		var existingSlots []domain.Availability
		err := ad.DB.Where("doctor_id = ? AND date = ? AND ((start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?) OR (start_time >= ? AND end_time <= ?))",
			availability.DoctorId, availability.Date, slotEndTime, currentTime, currentTime, slotEndTime, currentTime, slotEndTime).Find(&existingSlots).Error

		if err != nil {
			return "", err
		}

		// If no overlapping slots found, add the new slot to the list
		if len(existingSlots) == 0 {
			slots = append(slots, domain.Availability{
				DoctorID:  uint(availability.DoctorId),
				Date:      availability.Date,
				StartTime: currentTime,
				EndTime:   slotEndTime,
				IsBooked:  false,
			})
		}

		// Move to the next 30-minute slot
		currentTime = slotEndTime
	}

	// Check if there are slots to be inserted
	if len(slots) == 0 {
		return "", errors.New("no available slots to be added")
	}
	
	// Save the non-overlapping slots to the database
	if err := ad.DB.Create(&slots).Error; err != nil {
		return "", err
	}
	return "success", nil

}
func (ad *adminRepository)GetDoctorAvailability(doctor_id int,date time.Time)([]models.AvailableSlots,error)  {
	var slots []domain.Availability
	if err:=ad.DB.Where("doctor_id = ? AND date = ?", doctor_id, date).Find(&slots).Error;err!=nil{
		return []models.AvailableSlots{},err
	}
	var newslots []models.AvailableSlots
	for _,slot:=range slots{
		newslots = append(newslots, models.AvailableSlots{
			Slot_id: int(slot.ID),
			Time: fmt.Sprintf("%s-%s",slot.StartTime.Format("15:04"),slot.EndTime.Format("15:04")),
			IsBooked: slot.IsBooked,
		})
	}
	return newslots,nil
}