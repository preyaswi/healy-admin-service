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

type Payment struct {
    PaymentId     uint   `json:"payment_id" gorm:"primaryKey;not null"`
    PatientId     uint   `json:"patient_id" gorm:"not null"`
    DoctorId      uint   `json:"doctor_id" gorm:"not null"`
    DoctorName    string `json:"doctor_name" gorm:"not null"`
    Fees          uint64 `json:"fees" gorm:"not null"`
    PaymentStatus string `json:"payment_status" gorm:"not null"`
}

type RazerPay struct {
    ID        uint    `json:"id" gorm:"primaryKey;not null"`
    RazorID   string  `json:"razor_id"`
    PaymentID uint    `json:"payment_id"`
    Payment   Payment `json:"-" gorm:"foreignKey:PaymentID;references:PaymentId"`
}
