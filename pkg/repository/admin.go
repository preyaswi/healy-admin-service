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
func  (ad *adminRepository)AddPaymentDetails(payment models.Paymentreq)(domain.Payment, error) {
	newPayment := domain.Payment{
        PatientId:  payment.PatientId,
        DoctorId:   payment.DoctorId,
        DoctorName: payment.DoctorName,
        Fees:       payment.Fees,
        PaymentStatus: "Pending", // Assuming initial status is pending
    }

    // Insert the payment record
    if err := ad.DB.Create(&newPayment).Error; err != nil {
        return domain.Payment{}, err
    }

    // Return the inserted payment record
    return newPayment, nil
}
func (ad *adminRepository)AddRazorPayDetails(paymentID uint , razorPaypaymentID string) error{
	err:=ad.DB.Exec("insert into razer_pays (payment_id,razor_id) values (?,?)",paymentID,razorPaypaymentID).Error
	if err!=nil{
		return err
	}
	return nil
	
}
func (ad *adminRepository)CheckPaymentStatus(paymentId int)(string,error)  {
	var payment domain.Payment
    if err := ad.DB.First(&payment, paymentId).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return "", errors.New("payment not found")
        }
        return "", err
    }
    return payment.PaymentStatus, nil
}
func (ad *adminRepository)UpdatePaymentStatus(paymentId int,status string)error  {
err:=ad.DB.Model(&domain.Payment{}).Where("payment_id = ?", paymentId).Update("payment_status", status).Error
if err!=nil{
	return err
}
return nil
}