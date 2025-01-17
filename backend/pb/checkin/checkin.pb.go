// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: checkin.proto

package checkin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	identity "hestia/jobfair/api/pb/identity"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckinRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// token from the QR code
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// if event_id is not empty, the check-in will be associated with the event
	EventId       *string `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3,oneof" json:"event_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckinRequest) Reset() {
	*x = CheckinRequest{}
	mi := &file_checkin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinRequest) ProtoMessage() {}

func (x *CheckinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinRequest.ProtoReflect.Descriptor instead.
func (*CheckinRequest) Descriptor() ([]byte, []int) {
	return file_checkin_proto_rawDescGZIP(), []int{0}
}

func (x *CheckinRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CheckinRequest) GetEventId() string {
	if x != nil && x.EventId != nil {
		return *x.EventId
	}
	return ""
}

type CheckinWithCompanyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// token from the QR code
	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId     string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckinWithCompanyRequest) Reset() {
	*x = CheckinWithCompanyRequest{}
	mi := &file_checkin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckinWithCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinWithCompanyRequest) ProtoMessage() {}

func (x *CheckinWithCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinWithCompanyRequest.ProtoReflect.Descriptor instead.
func (*CheckinWithCompanyRequest) Descriptor() ([]byte, []int) {
	return file_checkin_proto_rawDescGZIP(), []int{1}
}

func (x *CheckinWithCompanyRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CheckinWithCompanyRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

var File_checkin_proto protoreflect.FileDescriptor

var file_checkin_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x32, 0x88, 0x02, 0x0a, 0x07, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x12, 0x56, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x12, 0x29, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x6c, 0x0a,
	0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x12, 0x34, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62,
	0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x73, 0x74,
	0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x1f, 0x5a, 0x1d, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2f, 0x6a,
	0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkin_proto_rawDescOnce sync.Once
	file_checkin_proto_rawDescData = file_checkin_proto_rawDesc
)

func file_checkin_proto_rawDescGZIP() []byte {
	file_checkin_proto_rawDescOnce.Do(func() {
		file_checkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkin_proto_rawDescData)
	})
	return file_checkin_proto_rawDescData
}

var file_checkin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_checkin_proto_goTypes = []any{
	(*CheckinRequest)(nil),            // 0: hestia.jobfair.v1.checkin.CheckinRequest
	(*CheckinWithCompanyRequest)(nil), // 1: hestia.jobfair.v1.checkin.CheckinWithCompanyRequest
	(*emptypb.Empty)(nil),             // 2: google.protobuf.Empty
	(*identity.User)(nil),             // 3: hestia.jobfair.v1.identity.User
}
var file_checkin_proto_depIdxs = []int32{
	0, // 0: hestia.jobfair.v1.checkin.Checkin.Checkin:input_type -> hestia.jobfair.v1.checkin.CheckinRequest
	1, // 1: hestia.jobfair.v1.checkin.Checkin.CheckinWithCompany:input_type -> hestia.jobfair.v1.checkin.CheckinWithCompanyRequest
	2, // 2: hestia.jobfair.v1.checkin.Checkin.Empty:input_type -> google.protobuf.Empty
	3, // 3: hestia.jobfair.v1.checkin.Checkin.Checkin:output_type -> hestia.jobfair.v1.identity.User
	3, // 4: hestia.jobfair.v1.checkin.Checkin.CheckinWithCompany:output_type -> hestia.jobfair.v1.identity.User
	2, // 5: hestia.jobfair.v1.checkin.Checkin.Empty:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_checkin_proto_init() }
func file_checkin_proto_init() {
	if File_checkin_proto != nil {
		return
	}
	file_checkin_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_checkin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkin_proto_goTypes,
		DependencyIndexes: file_checkin_proto_depIdxs,
		MessageInfos:      file_checkin_proto_msgTypes,
	}.Build()
	File_checkin_proto = out.File
	file_checkin_proto_rawDesc = nil
	file_checkin_proto_goTypes = nil
	file_checkin_proto_depIdxs = nil
}
