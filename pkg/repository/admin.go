package repository

import (
	"errors"
	"healy-admin/pkg/domain"
	interfaces "healy-admin/pkg/repository/interface"
	"healy-admin/pkg/utils/models"

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


func (ad *adminRepository)AddToBooking(patientid int,doctordetail models.BookingDoctorDetails)error  {
	err:=ad.DB.Exec("insert into bookings(patient_id,doctor_id,doctor_name,doctor_email,fees)values(?,?,?,?,?)",patientid,doctordetail.Doctorid,doctordetail.DoctorName,doctordetail.DoctorEmail,doctordetail.Fees).Error
	if err!=nil{
		return err
	}
	return nil
}
func (ad *adminRepository)GetBookingByID(bookingid int)(domain.Booking,error)  {
	var booking domain.Booking
	err:=ad.DB.Where("booking_id=?",bookingid).First(&booking).Error
	if err!=nil{
		return domain.Booking{},err
	}
	return booking,nil
}
func (ad *adminRepository)RemoveBooking(bookingID int) error  {
	err:=ad.DB.Where("booking_id=?",bookingID).Delete(&domain.Booking{}).Error
	if err != nil {
        return err
    }
    return nil
}
func (ad *adminRepository)AddRazorPayDetails(booking_id uint , razorPaypaymentID string) error{
	err:=ad.DB.Exec("insert into razer_pays (booking_id,razor_id) values (?,?)",booking_id,razorPaypaymentID).Error
	if err!=nil{
		return err
	}
	return nil
	
}
func (ad *adminRepository)CheckPaymentStatus(bookingid int)(string,error)  {
	var payment domain.Booking
    if err := ad.DB.First(&payment, bookingid).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return "", errors.New("payment not found")
        }
        return "", err
    }
    return payment.PaymentStatus, nil
}
func (ad *adminRepository)UpdatePaymentStatus(booking_id int,status string)error  {
err:=ad.DB.Model(&domain.Booking{}).Where("booking_id = ?", booking_id).Update("payment_status", status).Error
if err!=nil{
	return err
}
return nil
}