// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: user.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define the message
type UpdatePrimaryMemberDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title                   string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Gender                  string                 `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	FirstName               string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName                string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PictureUrl              string                 `protobuf:"bytes,5,opt,name=picture_url,json=pictureUrl,proto3" json:"picture_url,omitempty"`
	Email                   string                 `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	SecondaryEmail          string                 `protobuf:"bytes,7,opt,name=secondary_email,json=secondaryEmail,proto3" json:"secondary_email,omitempty"`
	LandlinePhoneNumber     string                 `protobuf:"bytes,8,opt,name=landline_phone_number,json=landlinePhoneNumber,proto3" json:"landline_phone_number,omitempty"`
	LandlinePhoneCountry    string                 `protobuf:"bytes,9,opt,name=landline_phone_country,json=landlinePhoneCountry,proto3" json:"landline_phone_country,omitempty"`
	LandlinePhoneNumberAlt  string                 `protobuf:"bytes,10,opt,name=landline_phone_number_alt,json=landlinePhoneNumberAlt,proto3" json:"landline_phone_number_alt,omitempty"`
	LandlinePhoneCountryAlt string                 `protobuf:"bytes,11,opt,name=landline_phone_country_alt,json=landlinePhoneCountryAlt,proto3" json:"landline_phone_country_alt,omitempty"`
	MobilePhoneNumber       string                 `protobuf:"bytes,12,opt,name=mobile_phone_number,json=mobilePhoneNumber,proto3" json:"mobile_phone_number,omitempty"`
	MobilePhoneCountry      string                 `protobuf:"bytes,13,opt,name=mobile_phone_country,json=mobilePhoneCountry,proto3" json:"mobile_phone_country,omitempty"`
	SmsCheckboxYn           string                 `protobuf:"bytes,14,opt,name=sms_checkbox_yn,json=smsCheckboxYn,proto3" json:"sms_checkbox_yn,omitempty"`
	ShirtSize               string                 `protobuf:"bytes,15,opt,name=shirt_size,json=shirtSize,proto3" json:"shirt_size,omitempty"`
	BirthDate               *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	WeddingDate             *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=wedding_date,json=weddingDate,proto3" json:"wedding_date,omitempty"`
	ResetWeddingDate        bool                   `protobuf:"varint,18,opt,name=reset_wedding_date,json=resetWeddingDate,proto3" json:"reset_wedding_date,omitempty"`
}

func (x *UpdatePrimaryMemberDTO) Reset() {
	*x = UpdatePrimaryMemberDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrimaryMemberDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrimaryMemberDTO) ProtoMessage() {}

func (x *UpdatePrimaryMemberDTO) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrimaryMemberDTO.ProtoReflect.Descriptor instead.
func (*UpdatePrimaryMemberDTO) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePrimaryMemberDTO) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetSecondaryEmail() string {
	if x != nil {
		return x.SecondaryEmail
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetLandlinePhoneNumber() string {
	if x != nil {
		return x.LandlinePhoneNumber
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetLandlinePhoneCountry() string {
	if x != nil {
		return x.LandlinePhoneCountry
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetLandlinePhoneNumberAlt() string {
	if x != nil {
		return x.LandlinePhoneNumberAlt
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetLandlinePhoneCountryAlt() string {
	if x != nil {
		return x.LandlinePhoneCountryAlt
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetMobilePhoneNumber() string {
	if x != nil {
		return x.MobilePhoneNumber
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetMobilePhoneCountry() string {
	if x != nil {
		return x.MobilePhoneCountry
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetSmsCheckboxYn() string {
	if x != nil {
		return x.SmsCheckboxYn
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetShirtSize() string {
	if x != nil {
		return x.ShirtSize
	}
	return ""
}

func (x *UpdatePrimaryMemberDTO) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *UpdatePrimaryMemberDTO) GetWeddingDate() *timestamppb.Timestamp {
	if x != nil {
		return x.WeddingDate
	}
	return nil
}

func (x *UpdatePrimaryMemberDTO) GetResetWeddingDate() bool {
	if x != nil {
		return x.ResetWeddingDate
	}
	return false
}

// Define the request message
type UpdatePrimaryMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member *UpdatePrimaryMemberDTO `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *UpdatePrimaryMemberRequest) Reset() {
	*x = UpdatePrimaryMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrimaryMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrimaryMemberRequest) ProtoMessage() {}

func (x *UpdatePrimaryMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrimaryMemberRequest.ProtoReflect.Descriptor instead.
func (*UpdatePrimaryMemberRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePrimaryMemberRequest) GetMember() *UpdatePrimaryMemberDTO {
	if x != nil {
		return x.Member
	}
	return nil
}

// Define the response message
type UpdatePrimaryMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdatePrimaryMemberResponse) Reset() {
	*x = UpdatePrimaryMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePrimaryMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePrimaryMemberResponse) ProtoMessage() {}

func (x *UpdatePrimaryMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePrimaryMemberResponse.ProtoReflect.Descriptor instead.
func (*UpdatePrimaryMemberResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePrimaryMemberResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdatePrimaryMemberResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x95, 0x06, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x61, 0x6e, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16,
	0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6c, 0x61,
	0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x39, 0x0a, 0x19, 0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x6c, 0x74, 0x12, 0x3b, 0x0a,
	0x1a, 0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x17, 0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0f,
	0x73, 0x6d, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x5f, 0x79, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6d, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x62,
	0x6f, 0x78, 0x59, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x72, 0x74, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x72, 0x74, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d,
	0x0a, 0x0c, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x57, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x65, 0x22, 0x52, 0x0a, 0x1a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x51, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x69, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x65, 0x6e,
	0x69, 0x65, 0x72, 0x4f, 0x73, 0x6f, 0x72, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_proto_goTypes = []interface{}{
	(*UpdatePrimaryMemberDTO)(nil),      // 0: grpc.UpdatePrimaryMemberDTO
	(*UpdatePrimaryMemberRequest)(nil),  // 1: grpc.UpdatePrimaryMemberRequest
	(*UpdatePrimaryMemberResponse)(nil), // 2: grpc.UpdatePrimaryMemberResponse
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	3, // 0: grpc.UpdatePrimaryMemberDTO.birth_date:type_name -> google.protobuf.Timestamp
	3, // 1: grpc.UpdatePrimaryMemberDTO.wedding_date:type_name -> google.protobuf.Timestamp
	0, // 2: grpc.UpdatePrimaryMemberRequest.member:type_name -> grpc.UpdatePrimaryMemberDTO
	1, // 3: grpc.UserService.UpdatePrimaryMember:input_type -> grpc.UpdatePrimaryMemberRequest
	2, // 4: grpc.UserService.UpdatePrimaryMember:output_type -> grpc.UpdatePrimaryMemberResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePrimaryMemberDTO); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePrimaryMemberRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePrimaryMemberResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
