package interfaces

import "healy-admin/pkg/utils/models"
type NewPatientClient interface{
	GetPatientByID(patientid int)(models.Patient,error)
}