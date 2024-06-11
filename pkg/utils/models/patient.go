package models

type Patient struct {
	Id            uint
	Fullname      string
	Email         string
	Gender        string
	Contactnumber string
}
type BookedPatient struct{
	BookingId int
	PaymentStatus string
	Patientdetail Patient
}

