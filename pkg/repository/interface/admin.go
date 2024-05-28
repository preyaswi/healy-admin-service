package interfaces

import (
	"healy-admin/pkg/domain"
	"healy-admin/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)

	AddPaymentDetails(payment models.Paymentreq)(domain.Payment, error)
	AddRazorPayDetails(paymentID uint,razorPaypaymentID string) error
	CheckPaymentStatus(paymentId int)(string,error)
	UpdatePaymentStatus(paymentId int,status string)error
}
