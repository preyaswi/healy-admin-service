// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pkg/pb/patient/patient.proto

package patient

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Idreq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Idreq) Reset() {
	*x = Idreq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Idreq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idreq) ProtoMessage() {}

func (x *Idreq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idreq.ProtoReflect.Descriptor instead.
func (*Idreq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{0}
}

func (x *Idreq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type PatientDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Fullname      string `protobuf:"bytes,2,opt,name=Fullname,proto3" json:"Fullname,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Gender        string `protobuf:"bytes,4,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Contactnumber string `protobuf:"bytes,5,opt,name=Contactnumber,proto3" json:"Contactnumber,omitempty"`
}

func (x *PatientDetails) Reset() {
	*x = PatientDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientDetails) ProtoMessage() {}

func (x *PatientDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientDetails.ProtoReflect.Descriptor instead.
func (*PatientDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{1}
}

func (x *PatientDetails) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatientDetails) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *PatientDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PatientDetails) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PatientDetails) GetContactnumber() string {
	if x != nil {
		return x.Contactnumber
	}
	return ""
}

type GooglePatientDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Googleid     string `protobuf:"bytes,1,opt,name=googleid,proto3" json:"googleid,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Accesstoken  string `protobuf:"bytes,3,opt,name=accesstoken,proto3" json:"accesstoken,omitempty"`
	Refreshtoken string `protobuf:"bytes,4,opt,name=refreshtoken,proto3" json:"refreshtoken,omitempty"`
	Tokenexpiry  string `protobuf:"bytes,5,opt,name=tokenexpiry,proto3" json:"tokenexpiry,omitempty"`
}

func (x *GooglePatientDetails) Reset() {
	*x = GooglePatientDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GooglePatientDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GooglePatientDetails) ProtoMessage() {}

func (x *GooglePatientDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GooglePatientDetails.ProtoReflect.Descriptor instead.
func (*GooglePatientDetails) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{2}
}

func (x *GooglePatientDetails) GetGoogleid() string {
	if x != nil {
		return x.Googleid
	}
	return ""
}

func (x *GooglePatientDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GooglePatientDetails) GetAccesstoken() string {
	if x != nil {
		return x.Accesstoken
	}
	return ""
}

func (x *GooglePatientDetails) GetRefreshtoken() string {
	if x != nil {
		return x.Refreshtoken
	}
	return ""
}

func (x *GooglePatientDetails) GetTokenexpiry() string {
	if x != nil {
		return x.Tokenexpiry
	}
	return ""
}

type UpdateGoogleTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoogleID     string `protobuf:"bytes,1,opt,name=googleID,proto3" json:"googleID,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	TokenExpiry  string `protobuf:"bytes,4,opt,name=tokenExpiry,proto3" json:"tokenExpiry,omitempty"`
}

func (x *UpdateGoogleTokenReq) Reset() {
	*x = UpdateGoogleTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoogleTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoogleTokenReq) ProtoMessage() {}

func (x *UpdateGoogleTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoogleTokenReq.ProtoReflect.Descriptor instead.
func (*UpdateGoogleTokenReq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGoogleTokenReq) GetGoogleID() string {
	if x != nil {
		return x.GoogleID
	}
	return ""
}

func (x *UpdateGoogleTokenReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UpdateGoogleTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *UpdateGoogleTokenReq) GetTokenExpiry() string {
	if x != nil {
		return x.TokenExpiry
	}
	return ""
}

type UpdateGoogleTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGoogleTokenRes) Reset() {
	*x = UpdateGoogleTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_patient_patient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoogleTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoogleTokenRes) ProtoMessage() {}

func (x *UpdateGoogleTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_patient_patient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoogleTokenRes.ProtoReflect.Descriptor instead.
func (*UpdateGoogleTokenRes) Descriptor() ([]byte, []int) {
	return file_pkg_pb_patient_patient_proto_rawDescGZIP(), []int{4}
}

var File_pkg_pb_patient_patient_proto protoreflect.FileDescriptor

var file_pkg_pb_patient_patient_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x0a, 0x05, 0x69, 0x64, 0x72, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xb0, 0x01, 0x0a,
	0x14, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22,
	0x9a, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0x16, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x32, 0xf5, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x3e, 0x0a, 0x11, 0x49, 0x6e, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x69, 0x64, 0x72, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x0e, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x64, 0x72, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00,
	0x12, 0x5a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_patient_patient_proto_rawDescOnce sync.Once
	file_pkg_pb_patient_patient_proto_rawDescData = file_pkg_pb_patient_patient_proto_rawDesc
)

func file_pkg_pb_patient_patient_proto_rawDescGZIP() []byte {
	file_pkg_pb_patient_patient_proto_rawDescOnce.Do(func() {
		file_pkg_pb_patient_patient_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_patient_patient_proto_rawDescData)
	})
	return file_pkg_pb_patient_patient_proto_rawDescData
}

var file_pkg_pb_patient_patient_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_pb_patient_patient_proto_goTypes = []interface{}{
	(*Idreq)(nil),                // 0: patient.idreq
	(*PatientDetails)(nil),       // 1: patient.PatientDetails
	(*GooglePatientDetails)(nil), // 2: patient.GooglePatientDetails
	(*UpdateGoogleTokenReq)(nil), // 3: patient.UpdateGoogleTokenReq
	(*UpdateGoogleTokenRes)(nil), // 4: patient.UpdateGoogleTokenRes
}
var file_pkg_pb_patient_patient_proto_depIdxs = []int32{
	0, // 0: patient.Patient.IndPatientDetails:input_type -> patient.idreq
	0, // 1: patient.Patient.GetPatientGoogleDetailsByID:input_type -> patient.idreq
	3, // 2: patient.Patient.UpdatePatientGoogleToken:input_type -> patient.UpdateGoogleTokenReq
	1, // 3: patient.Patient.IndPatientDetails:output_type -> patient.PatientDetails
	2, // 4: patient.Patient.GetPatientGoogleDetailsByID:output_type -> patient.GooglePatientDetails
	4, // 5: patient.Patient.UpdatePatientGoogleToken:output_type -> patient.UpdateGoogleTokenRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_patient_patient_proto_init() }
func file_pkg_pb_patient_patient_proto_init() {
	if File_pkg_pb_patient_patient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_patient_patient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Idreq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_patient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_patient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GooglePatientDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_patient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoogleTokenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_pb_patient_patient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoogleTokenRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_patient_patient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_patient_patient_proto_goTypes,
		DependencyIndexes: file_pkg_pb_patient_patient_proto_depIdxs,
		MessageInfos:      file_pkg_pb_patient_patient_proto_msgTypes,
	}.Build()
	File_pkg_pb_patient_patient_proto = out.File
	file_pkg_pb_patient_patient_proto_rawDesc = nil
	file_pkg_pb_patient_patient_proto_goTypes = nil
	file_pkg_pb_patient_patient_proto_depIdxs = nil
}
