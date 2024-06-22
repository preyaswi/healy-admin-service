package interfaces

import "healy-admin/pkg/utils/models"
type NewPatientClient interface{
	GetPatientByID(patientid string)(models.Patient,error)
}