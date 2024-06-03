package interfaces

import (
	"healy-admin/pkg/domain"
	"healy-admin/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)

	AddToBooking(patientid int,doctordetail models.BookingDoctorDetails)error
	GetBookingByID(bookingid int)(domain.Booking,error)
	RemoveBooking(bookingID int) error

	AddRazorPayDetails(bookingid uint,razorPaypaymentID string) error
	CheckPaymentStatus(bookingid int) (string, error)
	UpdatePaymentStatus(bookingid int,status string)error
	GetPaidBookingsByDoctorID(doctorId int)([]domain.Booking,error)
}
