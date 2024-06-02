package domain

import "healy-admin/pkg/utils/models"

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
}

type Booking struct {
	BookingId     uint   `json:"booking_id" gorm:"primaryKey;not null"`
	PatientId     uint   `json:"patient_id" gorm:"not null"`
	DoctorId      uint   `json:"doctor_id" gorm:"not null"`
	DoctorName    string `json:"doctor_name" gorm:"not null"`
	DoctorEmail   string `json:"doctor_email" gorm:"not null"`
	Fees          uint64 `json:"fees" gorm:"not null"`
	PaymentStatus string `json:"payment_status" gorm:"default:'not paid'"`
}


type RazerPay struct {
	ID        uint    `json:"id" gorm:"primaryKey;not null"`
	RazorID   string  `json:"razor_id"`
	BookingID uint    `json:"Booking_id"`
	Booking   Booking `json:"-" gorm:"foreignKey:BookingID;references:BookingId"`
}
