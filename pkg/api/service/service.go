package service

import (
	"context"
	interfaces "healy-admin/pkg/usecase/interface"
	"healy-admin/pkg/utils/models"
	pb "healy-admin/pkg/pb/admin"
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
func (ad *AdminServer)MakePaymentRazorpay(ctx context.Context,req *pb.PaymentReq)(*pb.PaymentRes,error)  {
	paymentdetail:=models.Paymentreq{
		PatientId: uint(req.PatientId),
		DoctorId: uint(req.DoctorId),
		DoctorName: req.DoctorName,
		Fees: req.Fees,
	}
	paymentDetails,razorId,err:=ad.adminUseCase.MakePaymentRazorpay(paymentdetail)
	if err!=nil{
		return &pb.PaymentRes{},err
	}
	paymentDetail:=&pb.PaymentDetails{
		PaymentId: uint32(paymentDetails.PaymentId),
		PatientId: uint32(paymentDetails.PatientId),
		DoctorId: uint32(paymentDetails.DoctorId),
		DoctorName: paymentDetails.DoctorName,
		Fees: paymentDetails.Fees,
		PaymentStatus: paymentDetails.PaymentStatus,

	}
	return &pb.PaymentRes{
		PaymentDetails: paymentDetail,
		Razorid: razorId,
	},nil
	
	
}
func (ad *AdminServer)VerifyPayment(ctx context.Context,req *pb.Verifyreq) (*pb.Verifyres, error) {
	err:=ad.adminUseCase.VerifyPayment(int(req.PaymentId))
	if err!=nil{
		return &pb.Verifyres{},err
	}
	return &pb.Verifyres{},nil
}