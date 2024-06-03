package interfaces

import (
	"healy-admin/pkg/domain"
	"healy-admin/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
	AddToBooking(patientid,doctorid int)error
	CancelBooking(patientid,bookingid int)error
	MakePaymentRazorpay(bookingid int )(domain.Booking,string,error)
	VerifyPayment(booking_id int)error
	GetPaidPatients(doctor_id int)([]models.Patient,error)
}
