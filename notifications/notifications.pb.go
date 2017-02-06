// Code generated by protoc-gen-go.
// source: notifications.proto
// DO NOT EDIT!

/*
Package notifications is a generated protocol buffer package.

It is generated from these files:
	notifications.proto

It has these top-level messages:
	NotificationRequest
	NotificationDataRequest
	NotificationResponse
	NotificationUpdateResponse
*/
package notifications

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NotificationRequest struct {
	NotificationUid string `protobuf:"bytes,1,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
}

func (m *NotificationRequest) Reset()                    { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()               {}
func (*NotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NotificationRequest) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

type NotificationDataRequest struct {
	NotificationId  string            `protobuf:"bytes,1,opt,name=notification_id,json=notificationId" json:"notification_id,omitempty"`
	NotificationUid string            `protobuf:"bytes,2,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
	ProjectId       string            `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	TargetId        []string          `protobuf:"bytes,4,rep,name=target_id,json=targetId" json:"target_id,omitempty"`
	Percentage      int32             `protobuf:"varint,5,opt,name=percentage" json:"percentage,omitempty"`
	Data            map[string]string `protobuf:"bytes,6,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NotificationDataRequest) Reset()                    { *m = NotificationDataRequest{} }
func (m *NotificationDataRequest) String() string            { return proto.CompactTextString(m) }
func (*NotificationDataRequest) ProtoMessage()               {}
func (*NotificationDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NotificationDataRequest) GetNotificationId() string {
	if m != nil {
		return m.NotificationId
	}
	return ""
}

func (m *NotificationDataRequest) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *NotificationDataRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *NotificationDataRequest) GetTargetId() []string {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *NotificationDataRequest) GetPercentage() int32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *NotificationDataRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type NotificationResponse struct {
	NotificationUid string                     `protobuf:"bytes,1,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
	NotificationId  string                     `protobuf:"bytes,2,opt,name=notification_id,json=notificationId" json:"notification_id,omitempty"`
	ProjectId       string                     `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Percentage      int32                      `protobuf:"varint,4,opt,name=percentage" json:"percentage,omitempty"`
	IsComplete      bool                       `protobuf:"varint,5,opt,name=is_complete,json=isComplete" json:"is_complete,omitempty"`
	Started         *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=started" json:"started,omitempty"`
	LastUpdate      *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=last_update,json=lastUpdate" json:"last_update,omitempty"`
}

func (m *NotificationResponse) Reset()                    { *m = NotificationResponse{} }
func (m *NotificationResponse) String() string            { return proto.CompactTextString(m) }
func (*NotificationResponse) ProtoMessage()               {}
func (*NotificationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NotificationResponse) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *NotificationResponse) GetNotificationId() string {
	if m != nil {
		return m.NotificationId
	}
	return ""
}

func (m *NotificationResponse) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *NotificationResponse) GetPercentage() int32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func (m *NotificationResponse) GetIsComplete() bool {
	if m != nil {
		return m.IsComplete
	}
	return false
}

func (m *NotificationResponse) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *NotificationResponse) GetLastUpdate() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

type NotificationUpdateResponse struct {
	NotificationUid string `protobuf:"bytes,1,opt,name=notification_uid,json=notificationUid" json:"notification_uid,omitempty"`
	Message         string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Success         bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
}

func (m *NotificationUpdateResponse) Reset()                    { *m = NotificationUpdateResponse{} }
func (m *NotificationUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*NotificationUpdateResponse) ProtoMessage()               {}
func (*NotificationUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NotificationUpdateResponse) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *NotificationUpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NotificationUpdateResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*NotificationRequest)(nil), "notifications.NotificationRequest")
	proto.RegisterType((*NotificationDataRequest)(nil), "notifications.NotificationDataRequest")
	proto.RegisterType((*NotificationResponse)(nil), "notifications.NotificationResponse")
	proto.RegisterType((*NotificationUpdateResponse)(nil), "notifications.NotificationUpdateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Notifications service

type NotificationsClient interface {
	Send(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error)
	Update(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error)
	Retrieve(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error)
}

type notificationsClient struct {
	cc *grpc.ClientConn
}

func NewNotificationsClient(cc *grpc.ClientConn) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) Send(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error) {
	out := new(NotificationUpdateResponse)
	err := grpc.Invoke(ctx, "/notifications.Notifications/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) Update(ctx context.Context, in *NotificationDataRequest, opts ...grpc.CallOption) (*NotificationUpdateResponse, error) {
	out := new(NotificationUpdateResponse)
	err := grpc.Invoke(ctx, "/notifications.Notifications/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) Retrieve(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationResponse, error) {
	out := new(NotificationResponse)
	err := grpc.Invoke(ctx, "/notifications.Notifications/Retrieve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notifications service

type NotificationsServer interface {
	Send(context.Context, *NotificationDataRequest) (*NotificationUpdateResponse, error)
	Update(context.Context, *NotificationDataRequest) (*NotificationUpdateResponse, error)
	Retrieve(context.Context, *NotificationRequest) (*NotificationResponse, error)
}

func RegisterNotificationsServer(s *grpc.Server, srv NotificationsServer) {
	s.RegisterService(&_Notifications_serviceDesc, srv)
}

func _Notifications_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Send(ctx, req.(*NotificationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Update(ctx, req.(*NotificationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Retrieve(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Notifications_Send_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Notifications_Update_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _Notifications_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications.proto",
}

func init() { proto.RegisterFile("notifications.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x27, 0x4d, 0xe2, 0x89, 0x0a, 0xd5, 0xb6, 0x12, 0x96, 0x11, 0xd4, 0x32, 0x12,
	0x75, 0x0e, 0x38, 0x28, 0x20, 0x81, 0xe0, 0x82, 0xa0, 0x1c, 0x72, 0xe1, 0xe0, 0xd2, 0x03, 0x70,
	0x88, 0x36, 0xf6, 0xd4, 0x2c, 0xc4, 0x5e, 0xe3, 0x1d, 0x57, 0x54, 0xbc, 0x0a, 0xaf, 0xc2, 0x9d,
	0xc7, 0x42, 0xf6, 0xc6, 0x95, 0x0d, 0x21, 0x05, 0x24, 0x6e, 0xde, 0x7f, 0xff, 0xf9, 0x47, 0xf3,
	0xed, 0xc8, 0xb0, 0x9f, 0x49, 0x12, 0x67, 0x22, 0xe2, 0x24, 0x64, 0xa6, 0x82, 0xbc, 0x90, 0x24,
	0xd9, 0x6e, 0x47, 0x74, 0x0e, 0x13, 0x29, 0x93, 0x15, 0x4e, 0xeb, 0xcb, 0x65, 0x79, 0x36, 0x25,
	0x91, 0xa2, 0x22, 0x9e, 0xe6, 0xda, 0xef, 0x3d, 0x83, 0xfd, 0x57, 0xad, 0x8a, 0x10, 0x3f, 0x95,
	0xa8, 0x88, 0x4d, 0x60, 0xaf, 0x1d, 0xb4, 0x28, 0x45, 0x6c, 0x1b, 0xae, 0xe1, 0x5b, 0xe1, 0xf5,
	0xb6, 0x7e, 0x2a, 0x62, 0xef, 0xbb, 0x09, 0x37, 0xda, 0x11, 0xc7, 0x9c, 0x78, 0x13, 0x73, 0x04,
	0x1d, 0xfb, 0xe2, 0x32, 0xe5, 0x5a, 0x5b, 0x9e, 0xc7, 0x1b, 0xfb, 0x99, 0x1b, 0xfb, 0xb1, 0x5b,
	0x00, 0x79, 0x21, 0x3f, 0x60, 0x44, 0x55, 0x5c, 0xaf, 0x36, 0x59, 0x6b, 0x65, 0x1e, 0xb3, 0x9b,
	0x60, 0x11, 0x2f, 0x12, 0xac, 0x6f, 0xfb, 0x6e, 0xcf, 0xb7, 0xc2, 0x91, 0x16, 0xe6, 0x31, 0xbb,
	0x0d, 0x90, 0x63, 0x11, 0x61, 0x46, 0x3c, 0x41, 0x7b, 0xc7, 0x35, 0xfc, 0x9d, 0xb0, 0xa5, 0xb0,
	0x63, 0xe8, 0xc7, 0x9c, 0xb8, 0x3d, 0x70, 0x7b, 0xfe, 0x78, 0x76, 0x3f, 0xe8, 0x12, 0xfe, 0xcd,
	0x94, 0x41, 0xf5, 0xfd, 0x32, 0xa3, 0xe2, 0x22, 0xac, 0xab, 0x9d, 0x47, 0x60, 0x5d, 0x4a, 0x6c,
	0x0f, 0x7a, 0x1f, 0xf1, 0x62, 0x3d, 0x76, 0xf5, 0xc9, 0x0e, 0x60, 0xe7, 0x9c, 0xaf, 0x4a, 0x5c,
	0x0f, 0xa8, 0x0f, 0x4f, 0xcc, 0xc7, 0x86, 0xf7, 0xcd, 0x84, 0x83, 0xee, 0x6b, 0xa8, 0x5c, 0x66,
	0x0a, 0xff, 0xe2, 0x39, 0x36, 0x21, 0x37, 0x37, 0x22, 0xbf, 0x82, 0x63, 0x17, 0x55, 0xff, 0x17,
	0x54, 0x87, 0x30, 0x16, 0x6a, 0x11, 0xc9, 0x34, 0x5f, 0x21, 0x69, 0x96, 0xa3, 0x10, 0x84, 0x7a,
	0xb1, 0x56, 0xd8, 0x43, 0x18, 0x2a, 0xe2, 0x05, 0x61, 0x6c, 0x0f, 0x5c, 0xc3, 0x1f, 0xcf, 0x9c,
	0x40, 0x2f, 0x63, 0xd0, 0x2c, 0x63, 0xf0, 0xba, 0x59, 0xc6, 0xb0, 0xb1, 0xb2, 0xa7, 0x30, 0x5e,
	0x71, 0x45, 0x8b, 0x32, 0x8f, 0x39, 0xa1, 0x3d, 0xbc, 0xb2, 0x12, 0x2a, 0xfb, 0x69, 0xed, 0xf6,
	0xbe, 0x80, 0xd3, 0xc6, 0xa7, 0xd5, 0x7f, 0x81, 0x68, 0xc3, 0x30, 0x45, 0xa5, 0xaa, 0xc9, 0x35,
	0xbc, 0xe6, 0x58, 0xdd, 0xa8, 0x32, 0x8a, 0x50, 0xa9, 0x1a, 0xd9, 0x28, 0x6c, 0x8e, 0xb3, 0xaf,
	0x26, 0xec, 0xb6, 0xbb, 0x2b, 0xf6, 0x06, 0xfa, 0x27, 0x98, 0xc5, 0xec, 0xee, 0x9f, 0xed, 0x91,
	0x33, 0xd9, 0xe2, 0xfb, 0x69, 0x96, 0x77, 0x30, 0xd0, 0xca, 0xff, 0x08, 0x3f, 0x81, 0x51, 0x88,
	0x54, 0x08, 0x3c, 0x47, 0xe6, 0x6d, 0x29, 0x6b, 0xa2, 0xef, 0x6c, 0xf5, 0xe8, 0xd0, 0xe7, 0x93,
	0xb7, 0x47, 0x89, 0xa0, 0xf7, 0xe5, 0x32, 0x88, 0x64, 0x3a, 0x8d, 0xca, 0x25, 0x7e, 0xd6, 0x7f,
	0xa5, 0x7b, 0x89, 0x9c, 0x76, 0xea, 0x97, 0x83, 0x5a, 0x7f, 0xf0, 0x23, 0x00, 0x00, 0xff, 0xff,
	0x45, 0x79, 0x03, 0x8b, 0xe1, 0x04, 0x00, 0x00,
}
