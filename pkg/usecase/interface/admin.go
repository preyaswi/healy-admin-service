package interfaces

import (
	"healy-admin/pkg/domain"
	"healy-admin/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
	MakePaymentRazorpay(paymentdetail models.Paymentreq)(domain.Payment,string,error)
	VerifyPayment(payment_id int)error

}
