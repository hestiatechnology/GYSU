// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.2
// source: events.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	common "hestia/jobfair/api/pb/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Speaker struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Photo         string                 `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Speaker) Reset() {
	*x = Speaker{}
	mi := &file_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Speaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speaker) ProtoMessage() {}

func (x *Speaker) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Speaker.ProtoReflect.Descriptor instead.
func (*Speaker) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *Speaker) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Speaker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Speaker) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Speaker) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Speakers      []*Speaker             `protobuf:"bytes,3,rep,name=speakers,proto3" json:"speakers,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Location      string                 `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Seats         int32                  `protobuf:"varint,8,opt,name=seats,proto3" json:"seats,omitempty"`
	Info          string                 `protobuf:"bytes,9,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetSpeakers() []*Speaker {
	if x != nil {
		return x.Speakers
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Event) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Event) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Event) GetSeats() int32 {
	if x != nil {
		return x.Seats
	}
	return 0
}

func (x *Event) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type EventList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Events        []*Event               `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventList) Reset() {
	*x = EventList{}
	mi := &file_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventList) ProtoMessage() {}

func (x *EventList) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventList.ProtoReflect.Descriptor instead.
func (*EventList) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *EventList) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x07, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0xd0, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68,
	0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52,
	0x08, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61,
	0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61,
	0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x32, 0x8e, 0x04, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x73,
	0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23,
	0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x1f, 0x2e,
	0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x46,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e,
	0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a,
	0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x1e, 0x5a, 0x1c, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2f, 0x6a, 0x6f, 0x62, 0x66,
	0x61, 0x69, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_events_proto_goTypes = []any{
	(*Speaker)(nil),               // 0: hestia.jobfair.v1.events.Speaker
	(*Event)(nil),                 // 1: hestia.jobfair.v1.events.Event
	(*EventList)(nil),             // 2: hestia.jobfair.v1.events.EventList
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
	(*common.Id)(nil),             // 5: hestia.jobfair.v1.common.Id
}
var file_events_proto_depIdxs = []int32{
	0,  // 0: hestia.jobfair.v1.events.Event.speakers:type_name -> hestia.jobfair.v1.events.Speaker
	3,  // 1: hestia.jobfair.v1.events.Event.start_time:type_name -> google.protobuf.Timestamp
	3,  // 2: hestia.jobfair.v1.events.Event.end_time:type_name -> google.protobuf.Timestamp
	1,  // 3: hestia.jobfair.v1.events.EventList.events:type_name -> hestia.jobfair.v1.events.Event
	1,  // 4: hestia.jobfair.v1.events.Events.CreateEvent:input_type -> hestia.jobfair.v1.events.Event
	4,  // 5: hestia.jobfair.v1.events.Events.GetEvents:input_type -> google.protobuf.Empty
	5,  // 6: hestia.jobfair.v1.events.Events.GetEvent:input_type -> hestia.jobfair.v1.common.Id
	1,  // 7: hestia.jobfair.v1.events.Events.UpdateEvent:input_type -> hestia.jobfair.v1.events.Event
	5,  // 8: hestia.jobfair.v1.events.Events.DeleteEvent:input_type -> hestia.jobfair.v1.common.Id
	5,  // 9: hestia.jobfair.v1.events.Events.SubscribeToEvent:input_type -> hestia.jobfair.v1.common.Id
	5,  // 10: hestia.jobfair.v1.events.Events.UnsubscribeToEvent:input_type -> hestia.jobfair.v1.common.Id
	5,  // 11: hestia.jobfair.v1.events.Events.CreateEvent:output_type -> hestia.jobfair.v1.common.Id
	2,  // 12: hestia.jobfair.v1.events.Events.GetEvents:output_type -> hestia.jobfair.v1.events.EventList
	1,  // 13: hestia.jobfair.v1.events.Events.GetEvent:output_type -> hestia.jobfair.v1.events.Event
	4,  // 14: hestia.jobfair.v1.events.Events.UpdateEvent:output_type -> google.protobuf.Empty
	4,  // 15: hestia.jobfair.v1.events.Events.DeleteEvent:output_type -> google.protobuf.Empty
	4,  // 16: hestia.jobfair.v1.events.Events.SubscribeToEvent:output_type -> google.protobuf.Empty
	4,  // 17: hestia.jobfair.v1.events.Events.UnsubscribeToEvent:output_type -> google.protobuf.Empty
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	file_events_proto_msgTypes[0].OneofWrappers = []any{}
	file_events_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}