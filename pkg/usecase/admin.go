package usecase

import (
	"errors"
	"healy-admin/pkg/config"
	"healy-admin/pkg/domain"
	"healy-admin/pkg/helper"
	interfaces "healy-admin/pkg/repository/interface"
	services "healy-admin/pkg/usecase/interface"
	"healy-admin/pkg/utils/models"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/razorpay/razorpay-go"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repository interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repository,
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
func (ad *adminUseCase) MakePaymentRazorpay(payment models.Paymentreq) (domain.Payment, string, error) {
	cfg, _ := config.LoadConfig()
	paymentdetail, err := ad.adminRepository.AddPaymentDetails(payment)
	if err != nil {
		return domain.Payment{}, "", err
	}
	client := razorpay.NewClient(cfg.KEY_ID_fOR_PAY, cfg.SECRET_KEY_FOR_PAY)
	data := map[string]interface{}{
		"amount":   int(paymentdetail.Fees) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Payment{}, "", err
	}
	RazorpayorderId := body["id"].(string)
	err = ad.adminRepository.AddRazorPayDetails(paymentdetail.PaymentId, RazorpayorderId)
	if err != nil {
		return domain.Payment{}, "", err
	}
	return paymentdetail, RazorpayorderId, nil
}
func (ad *adminUseCase)VerifyPayment(payment_id int)error{
	status,err:=ad.adminRepository.CheckPaymentStatus(payment_id)
	if err!=nil{
		return err
	}
	status = strings.TrimSpace(strings.ToLower(status))
	if status=="pending"{
		if err:= ad.adminRepository.UpdatePaymentStatus(payment_id,"paid");err!=nil{
			return err
		}
		return nil
	}else {
	return errors.New("already paid")
	}
}
