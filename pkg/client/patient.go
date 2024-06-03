package client

import (
	"context"
	"fmt"
	"healy-admin/pkg/config"
	pbpa "healy-admin/pkg/pb/patient"
	"healy-admin/pkg/utils/models"

	"google.golang.org/grpc"
)

type patientClient struct {
	Client pbpa.PatientClient
}

func NewPatientClient(cfg *config.Config) *patientClient {
	cc, err := grpc.Dial(cfg.PATIENT_SVC, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	pbclient := pbpa.NewPatientClient(cc)
	return &patientClient{
		Client: pbclient,
	}

}

func (c *patientClient) GetPatientByID(patientid int) (models.Patient, error) {
	res, err := c.Client.IndPatientDetails(context.Background(),&pbpa.Idreq{UserId: int32(patientid)})
	if err != nil {
		return models.Patient{}, err
	}
	return models.Patient{
		Id:            uint(res.Id),
		Fullname:      res.Fullname,
		Email:         res.Email,
		Gender:        res.Gender,
		Contactnumber: res.Contactnumber,
	}, nil
}
