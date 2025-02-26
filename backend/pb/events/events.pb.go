// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: hestia/jobfair/events/v1/events.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	common "hestia/jobfair/api/pb/common"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Speaker struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          *string                `protobuf:"bytes,1,opt,name=id"`
	xxx_hidden_Name        *string                `protobuf:"bytes,2,opt,name=name"`
	xxx_hidden_Position    *string                `protobuf:"bytes,3,opt,name=position"`
	xxx_hidden_Company     *string                `protobuf:"bytes,4,opt,name=company"`
	xxx_hidden_Photo       *string                `protobuf:"bytes,5,opt,name=photo"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Speaker) Reset() {
	*x = Speaker{}
	mi := &file_hestia_jobfair_events_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Speaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speaker) ProtoMessage() {}

func (x *Speaker) ProtoReflect() protoreflect.Message {
	mi := &file_hestia_jobfair_events_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Speaker) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *Speaker) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *Speaker) GetPosition() string {
	if x != nil {
		if x.xxx_hidden_Position != nil {
			return *x.xxx_hidden_Position
		}
		return ""
	}
	return ""
}

func (x *Speaker) GetCompany() string {
	if x != nil {
		if x.xxx_hidden_Company != nil {
			return *x.xxx_hidden_Company
		}
		return ""
	}
	return ""
}

func (x *Speaker) GetPhoto() string {
	if x != nil {
		if x.xxx_hidden_Photo != nil {
			return *x.xxx_hidden_Photo
		}
		return ""
	}
	return ""
}

func (x *Speaker) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 5)
}

func (x *Speaker) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 5)
}

func (x *Speaker) SetPosition(v string) {
	x.xxx_hidden_Position = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 5)
}

func (x *Speaker) SetCompany(v string) {
	x.xxx_hidden_Company = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 5)
}

func (x *Speaker) SetPhoto(v string) {
	x.xxx_hidden_Photo = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 5)
}

func (x *Speaker) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Speaker) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *Speaker) HasPosition() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *Speaker) HasCompany() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *Speaker) HasPhoto() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *Speaker) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *Speaker) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Name = nil
}

func (x *Speaker) ClearPosition() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Position = nil
}

func (x *Speaker) ClearCompany() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_Company = nil
}

func (x *Speaker) ClearPhoto() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Photo = nil
}

type Speaker_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id       *string
	Name     *string
	Position *string
	Company  *string
	Photo    *string
}

func (b0 Speaker_builder) Build() *Speaker {
	m0 := &Speaker{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 5)
		x.xxx_hidden_Id = b.Id
	}
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 5)
		x.xxx_hidden_Name = b.Name
	}
	if b.Position != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 5)
		x.xxx_hidden_Position = b.Position
	}
	if b.Company != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 5)
		x.xxx_hidden_Company = b.Company
	}
	if b.Photo != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 5)
		x.xxx_hidden_Photo = b.Photo
	}
	return m0
}

type Event struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          *string                `protobuf:"bytes,1,opt,name=id"`
	xxx_hidden_Name        *string                `protobuf:"bytes,2,opt,name=name"`
	xxx_hidden_Speakers    *[]*Speaker            `protobuf:"bytes,3,rep,name=speakers"`
	xxx_hidden_Description *string                `protobuf:"bytes,4,opt,name=description"`
	xxx_hidden_Location    *string                `protobuf:"bytes,5,opt,name=location"`
	xxx_hidden_StartTime   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime"`
	xxx_hidden_EndTime     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime"`
	xxx_hidden_Seats       int32                  `protobuf:"varint,8,opt,name=seats"`
	xxx_hidden_Info        *string                `protobuf:"bytes,9,opt,name=info"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_hestia_jobfair_events_v1_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_hestia_jobfair_events_v1_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Event) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *Event) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *Event) GetSpeakers() []*Speaker {
	if x != nil {
		if x.xxx_hidden_Speakers != nil {
			return *x.xxx_hidden_Speakers
		}
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		if x.xxx_hidden_Description != nil {
			return *x.xxx_hidden_Description
		}
		return ""
	}
	return ""
}

func (x *Event) GetLocation() string {
	if x != nil {
		if x.xxx_hidden_Location != nil {
			return *x.xxx_hidden_Location
		}
		return ""
	}
	return ""
}

func (x *Event) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_StartTime
	}
	return nil
}

func (x *Event) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_EndTime
	}
	return nil
}

func (x *Event) GetSeats() int32 {
	if x != nil {
		return x.xxx_hidden_Seats
	}
	return 0
}

func (x *Event) GetInfo() string {
	if x != nil {
		if x.xxx_hidden_Info != nil {
			return *x.xxx_hidden_Info
		}
		return ""
	}
	return ""
}

func (x *Event) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 9)
}

func (x *Event) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 9)
}

func (x *Event) SetSpeakers(v []*Speaker) {
	x.xxx_hidden_Speakers = &v
}

func (x *Event) SetDescription(v string) {
	x.xxx_hidden_Description = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 9)
}

func (x *Event) SetLocation(v string) {
	x.xxx_hidden_Location = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 9)
}

func (x *Event) SetStartTime(v *timestamppb.Timestamp) {
	x.xxx_hidden_StartTime = v
}

func (x *Event) SetEndTime(v *timestamppb.Timestamp) {
	x.xxx_hidden_EndTime = v
}

func (x *Event) SetSeats(v int32) {
	x.xxx_hidden_Seats = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 7, 9)
}

func (x *Event) SetInfo(v string) {
	x.xxx_hidden_Info = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 8, 9)
}

func (x *Event) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Event) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *Event) HasDescription() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *Event) HasLocation() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *Event) HasStartTime() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_StartTime != nil
}

func (x *Event) HasEndTime() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_EndTime != nil
}

func (x *Event) HasSeats() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 7)
}

func (x *Event) HasInfo() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 8)
}

func (x *Event) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *Event) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Name = nil
}

func (x *Event) ClearDescription() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_Description = nil
}

func (x *Event) ClearLocation() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Location = nil
}

func (x *Event) ClearStartTime() {
	x.xxx_hidden_StartTime = nil
}

func (x *Event) ClearEndTime() {
	x.xxx_hidden_EndTime = nil
}

func (x *Event) ClearSeats() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 7)
	x.xxx_hidden_Seats = 0
}

func (x *Event) ClearInfo() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 8)
	x.xxx_hidden_Info = nil
}

type Event_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id          *string
	Name        *string
	Speakers    []*Speaker
	Description *string
	Location    *string
	StartTime   *timestamppb.Timestamp
	EndTime     *timestamppb.Timestamp
	Seats       *int32
	Info        *string
}

func (b0 Event_builder) Build() *Event {
	m0 := &Event{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 9)
		x.xxx_hidden_Id = b.Id
	}
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 9)
		x.xxx_hidden_Name = b.Name
	}
	x.xxx_hidden_Speakers = &b.Speakers
	if b.Description != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 9)
		x.xxx_hidden_Description = b.Description
	}
	if b.Location != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 9)
		x.xxx_hidden_Location = b.Location
	}
	x.xxx_hidden_StartTime = b.StartTime
	x.xxx_hidden_EndTime = b.EndTime
	if b.Seats != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 7, 9)
		x.xxx_hidden_Seats = *b.Seats
	}
	if b.Info != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 8, 9)
		x.xxx_hidden_Info = b.Info
	}
	return m0
}

type EventList struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Events *[]*Event              `protobuf:"bytes,1,rep,name=events"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *EventList) Reset() {
	*x = EventList{}
	mi := &file_hestia_jobfair_events_v1_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventList) ProtoMessage() {}

func (x *EventList) ProtoReflect() protoreflect.Message {
	mi := &file_hestia_jobfair_events_v1_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *EventList) GetEvents() []*Event {
	if x != nil {
		if x.xxx_hidden_Events != nil {
			return *x.xxx_hidden_Events
		}
	}
	return nil
}

func (x *EventList) SetEvents(v []*Event) {
	x.xxx_hidden_Events = &v
}

type EventList_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Events []*Event
}

func (b0 EventList_builder) Build() *EventList {
	m0 := &EventList{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Events = &b.Events
	return m0
}

var File_hestia_jobfair_events_v1_events_proto protoreflect.FileDescriptor

var file_hestia_jobfair_events_v1_events_proto_rawDesc = string([]byte{
	0x0a, 0x25, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2f, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e,
	0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2f, 0x6a, 0x6f, 0x62, 0x66, 0x61,
	0x69, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x53, 0x70,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x70, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x73,
	0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x08, 0x73,
	0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x44, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x95, 0x04, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69,
	0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x73, 0x74,
	0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x68,
	0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x49, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e,
	0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x1a, 0x1f, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62,
	0x66, 0x61, 0x69, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x10, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e,
	0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x73, 0x74,
	0x69, 0x61, 0x2e, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x26, 0x5a, 0x1c, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2f, 0x6a, 0x6f, 0x62, 0x66, 0x61, 0x69,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x92,
	0x03, 0x05, 0xd2, 0x3e, 0x02, 0x10, 0x03, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x70, 0xe8, 0x07,
})

var file_hestia_jobfair_events_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hestia_jobfair_events_v1_events_proto_goTypes = []any{
	(*Speaker)(nil),               // 0: hestia.jobfair.events.v1.Speaker
	(*Event)(nil),                 // 1: hestia.jobfair.events.v1.Event
	(*EventList)(nil),             // 2: hestia.jobfair.events.v1.EventList
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
	(*common.Id)(nil),             // 5: hestia.jobfair.common.v1.Id
}
var file_hestia_jobfair_events_v1_events_proto_depIdxs = []int32{
	0,  // 0: hestia.jobfair.events.v1.Event.speakers:type_name -> hestia.jobfair.events.v1.Speaker
	3,  // 1: hestia.jobfair.events.v1.Event.start_time:type_name -> google.protobuf.Timestamp
	3,  // 2: hestia.jobfair.events.v1.Event.end_time:type_name -> google.protobuf.Timestamp
	1,  // 3: hestia.jobfair.events.v1.EventList.events:type_name -> hestia.jobfair.events.v1.Event
	1,  // 4: hestia.jobfair.events.v1.EventsService.CreateEvent:input_type -> hestia.jobfair.events.v1.Event
	4,  // 5: hestia.jobfair.events.v1.EventsService.GetEvents:input_type -> google.protobuf.Empty
	5,  // 6: hestia.jobfair.events.v1.EventsService.GetEvent:input_type -> hestia.jobfair.common.v1.Id
	1,  // 7: hestia.jobfair.events.v1.EventsService.UpdateEvent:input_type -> hestia.jobfair.events.v1.Event
	5,  // 8: hestia.jobfair.events.v1.EventsService.DeleteEvent:input_type -> hestia.jobfair.common.v1.Id
	5,  // 9: hestia.jobfair.events.v1.EventsService.SubscribeToEvent:input_type -> hestia.jobfair.common.v1.Id
	5,  // 10: hestia.jobfair.events.v1.EventsService.UnsubscribeToEvent:input_type -> hestia.jobfair.common.v1.Id
	5,  // 11: hestia.jobfair.events.v1.EventsService.CreateEvent:output_type -> hestia.jobfair.common.v1.Id
	2,  // 12: hestia.jobfair.events.v1.EventsService.GetEvents:output_type -> hestia.jobfair.events.v1.EventList
	1,  // 13: hestia.jobfair.events.v1.EventsService.GetEvent:output_type -> hestia.jobfair.events.v1.Event
	4,  // 14: hestia.jobfair.events.v1.EventsService.UpdateEvent:output_type -> google.protobuf.Empty
	4,  // 15: hestia.jobfair.events.v1.EventsService.DeleteEvent:output_type -> google.protobuf.Empty
	4,  // 16: hestia.jobfair.events.v1.EventsService.SubscribeToEvent:output_type -> google.protobuf.Empty
	4,  // 17: hestia.jobfair.events.v1.EventsService.UnsubscribeToEvent:output_type -> google.protobuf.Empty
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_hestia_jobfair_events_v1_events_proto_init() }
func file_hestia_jobfair_events_v1_events_proto_init() {
	if File_hestia_jobfair_events_v1_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_hestia_jobfair_events_v1_events_proto_rawDesc), len(file_hestia_jobfair_events_v1_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hestia_jobfair_events_v1_events_proto_goTypes,
		DependencyIndexes: file_hestia_jobfair_events_v1_events_proto_depIdxs,
		MessageInfos:      file_hestia_jobfair_events_v1_events_proto_msgTypes,
	}.Build()
	File_hestia_jobfair_events_v1_events_proto = out.File
	file_hestia_jobfair_events_v1_events_proto_goTypes = nil
	file_hestia_jobfair_events_v1_events_proto_depIdxs = nil
}
