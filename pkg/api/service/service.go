package service

import (
	"context"
	pb "healy-admin/pkg/pb/admin"
	interfaces "healy-admin/pkg/usecase/interface"
	"healy-admin/pkg/utils/models"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {

	return &AdminServer{
		adminUseCase: useCase,
	}

}
func (ad *AdminServer) AdminSignup(ctx context.Context, req *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {
	adminSignup := models.AdminSignUp{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	res, err := ad.adminUseCase.AdminSignUp(adminSignup)
	if err != nil {
		return &pb.AdminSignupResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(res.Admin.ID),
		Firstname: res.Admin.Firstname,
		Lastname:  res.Admin.Lastname,
		Email:     res.Admin.Email,
	}
	return &pb.AdminSignupResponse{
		Status:       201,
		AdminDetails: adminDetails,
		Token:        res.Token,
	}, nil
}

func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := models.AdminLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	admin, err := ad.adminUseCase.LoginHandler(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Lastname:  admin.Admin.Lastname,
		Email:     admin.Admin.Email,
	}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}
func (ad *AdminServer)AddTobookings(ctx context.Context,req *pb.Bookingreq) (*pb.Bookingres, error)  {
	err:=ad.adminUseCase.AddToBooking(int(req.PatientId),int(req.DoctorId))
	if err!=nil{
		return &pb.Bookingres{},err
	}
	return &pb.Bookingres{},nil
}
func (ad *AdminServer)CancelBookings(ctx context.Context,req  *pb.Canbookingreq) (*pb.Bookingres, error)  {
	err:=ad.adminUseCase.CancelBooking(int(req.PatientId),int(req.BookingId))
	if err!=nil{
		return &pb.Bookingres{},err
	}
	return &pb.Bookingres{},nil
}
func (ad *AdminServer)MakePaymentRazorpay(ctx context.Context,req *pb.PaymentReq)(*pb.PaymentRes,error)  {
	bookingid:=models.Paymentreq{
		Bookingid: uint(req.BookingId),
	}
	paymentDetails,razorId,err:=ad.adminUseCase.MakePaymentRazorpay(int(bookingid.Bookingid))
	if err!=nil{
		return &pb.PaymentRes{},err
	}
	paymentDetail:=&pb.PaymentDetails{
		BookingId: uint32(paymentDetails.BookingId),
		PatientId: uint32(paymentDetails.PatientId),
		DoctorId: uint32(paymentDetails.DoctorId),
		DoctorName: paymentDetails.DoctorName,
		DoctorEmail: paymentDetails.DoctorEmail,
		Fees: paymentDetails.Fees,
		PaymentStatus: paymentDetails.PaymentStatus,

	}
	return &pb.PaymentRes{
		PaymentDetails: paymentDetail,
		Razorid: razorId,
	},nil
	
	
}
func (ad *AdminServer)VerifyPayment(ctx context.Context,req *pb.PaymentReq) (*pb.Verifyres, error) {
	err:=ad.adminUseCase.VerifyPayment(int(req.BookingId))
	if err!=nil{
		return &pb.Verifyres{},err
	}
	return &pb.Verifyres{},nil
}
func (ad *AdminServer)GetPaidPatients(ctx context.Context,req *pb.GetPaidPatientsRequest) (*pb.GetPaidPatientsResponse, error) {
	patients,err:=ad.adminUseCase.GetPaidPatients(int(req.DoctorId))
	if err!=nil{
		return &pb.GetPaidPatientsResponse{},err
	}
	patientdetails:=make([]*pb.Patient,len(patients))
	for i,patient:=range patients{
		patientdetails[i]=&pb.Patient{
			Id:uint32(patient.Id) ,
			Fullname: patient.Fullname,
			Email: patient.Email,
			Gender: patient.Gender,
			Contactnumber: patient.Contactnumber,
		}
	}
	return &pb.GetPaidPatientsResponse{
		Patients: patientdetails,
	},nil
}
func (ad *AdminServer)CreatePrescription(ctx context.Context,req *pb.CreatePrescriptionRequest) (*pb.CreatePrescriptionResponse, error)  {
	prescriptionreq:=models.PrescriptionRequest{
		DoctorID: int(req.DoctorId),
		PatientID: int(req.PatientId),
		BookingID: int(req.BookingId),
		Medicine: req.Medicine,
		Dosage: req.Dosage,
		Notes: req.Notes,
	}
	prescription,err:=ad.adminUseCase.CreatePrescription(prescriptionreq)
	if err!=nil{
		return &pb.CreatePrescriptionResponse{},err
	}
	return &pb.CreatePrescriptionResponse{
		Id: uint32(prescription.ID),
		BookingId: uint32(prescription.BookingID),
		DoctorId: uint32(prescription.DoctorID),
		PatientId: uint32(prescription.PatientID),
		Medicine: prescription.Medicine,
		Dosage: prescription.Dosage,
		Notes: prescription.Notes,
	},nil
}