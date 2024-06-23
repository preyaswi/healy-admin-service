package interfaces

import (
	"healy-admin/pkg/domain"
	"healy-admin/pkg/utils/models"
	"time"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)

	AddToBooking(patientid string,doctordetail models.BookingDoctorDetails)error
	GetBookingByID(bookingid int)(domain.Booking,error)
	RemoveBooking(bookingID int) error

	AddRazorPayDetails(bookingid uint,razorPaypaymentID string) error
	CheckPaymentStatus(bookingid int) (string, error)
	UpdatePaymentStatus(bookingid int,status string)error
	GetPaidBookingsByDoctorID(doctorId int)([]domain.Booking,error)
	CheckPatientPayment(doctorID int, patientID string) (bool, error)
	CreatePrescription(prescription models.PrescriptionRequest) (domain.Prescription, error)

	SetDoctorAvailability(availabiity models.SetAvailability)(string,error)
	GetDoctorAvailability(doctor_id int,date time.Time)([]models.AvailableSlots,error)
}
