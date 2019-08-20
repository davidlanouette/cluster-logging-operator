// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/hotel_performance_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [HotelPerformanceViewService.GetHotelPerformanceView][google.ads.googleads.v2.services.HotelPerformanceViewService.GetHotelPerformanceView].
type GetHotelPerformanceViewRequest struct {
	// Resource name of the Hotel Performance View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotelPerformanceViewRequest) Reset()         { *m = GetHotelPerformanceViewRequest{} }
func (m *GetHotelPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotelPerformanceViewRequest) ProtoMessage()    {}
func (*GetHotelPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13bf764e5e04adb1, []int{0}
}

func (m *GetHotelPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetHotelPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetHotelPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotelPerformanceViewRequest.Merge(m, src)
}
func (m *GetHotelPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Size(m)
}
func (m *GetHotelPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotelPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotelPerformanceViewRequest proto.InternalMessageInfo

func (m *GetHotelPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHotelPerformanceViewRequest)(nil), "google.ads.googleads.v2.services.GetHotelPerformanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/hotel_performance_view_service.proto", fileDescriptor_13bf764e5e04adb1)
}

var fileDescriptor_13bf764e5e04adb1 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x25, 0x11, 0x04, 0x83, 0x6e, 0xb2, 0x69, 0x49, 0x45, 0x42, 0xed, 0x42, 0x5c, 0xcc, 0x40,
	0x14, 0x85, 0xf1, 0x07, 0x53, 0x28, 0xed, 0x4a, 0x4a, 0x85, 0x2c, 0x24, 0x10, 0xc6, 0xe4, 0x1a,
	0x03, 0xc9, 0x4c, 0x9c, 0x99, 0xa6, 0x0b, 0x71, 0xe3, 0xc6, 0xa5, 0x0b, 0xdf, 0xc0, 0xa5, 0x8f,
	0xd2, 0xa5, 0xbe, 0x82, 0x2b, 0x9f, 0x42, 0xd2, 0xc9, 0xa4, 0x7e, 0x1f, 0x4d, 0xbb, 0x3b, 0xcc,
	0x3d, 0xf7, 0x9c, 0x7b, 0xcf, 0x1d, 0x67, 0x91, 0x73, 0x9e, 0x97, 0x80, 0x69, 0x26, 0xb1, 0x86,
	0x2d, 0x6a, 0x02, 0x2c, 0x41, 0x34, 0x45, 0x0a, 0x12, 0x7f, 0xe0, 0x0a, 0xca, 0xa4, 0x06, 0xf1,
	0x9e, 0x8b, 0x8a, 0xb2, 0x14, 0x92, 0xa6, 0x80, 0x5d, 0xd2, 0xd5, 0x51, 0x2d, 0xb8, 0xe2, 0xae,
	0xaf, 0x7b, 0x11, 0xcd, 0x24, 0xea, 0x65, 0x50, 0x13, 0x20, 0x23, 0xe3, 0xbd, 0x1c, 0x32, 0x12,
	0x20, 0xf9, 0x56, 0x0c, 0x3b, 0x69, 0x07, 0xef, 0xae, 0xe9, 0xaf, 0x0b, 0x4c, 0x19, 0xe3, 0x8a,
	0xaa, 0x82, 0x33, 0xd9, 0x55, 0x47, 0xff, 0x55, 0xd3, 0xb2, 0x00, 0xa6, 0x74, 0x61, 0xba, 0x70,
	0xee, 0x2d, 0x41, 0xad, 0x5a, 0xe5, 0xf5, 0x51, 0x38, 0x2a, 0x60, 0xb7, 0x81, 0x8f, 0x5b, 0x90,
	0xca, 0xbd, 0xef, 0xdc, 0x31, 0x23, 0x24, 0x8c, 0x56, 0x30, 0xb6, 0x7c, 0xeb, 0xc1, 0xad, 0xcd,
	0x6d, 0xf3, 0xf8, 0x9a, 0x56, 0x10, 0x7c, 0xb3, 0x9d, 0xc9, 0x29, 0x91, 0x37, 0x7a, 0x3d, 0xf7,
	0x97, 0xe5, 0x8c, 0x06, 0x7c, 0xdc, 0x57, 0xe8, 0x52, 0x38, 0xe8, 0xfc, 0x88, 0xde, 0xd3, 0x41,
	0x85, 0x3e, 0x3c, 0x74, 0xaa, 0x7f, 0xfa, 0xfc, 0xcb, 0xef, 0x3f, 0xdf, 0xed, 0x27, 0xee, 0xe3,
	0x36, 0xe8, 0x4f, 0x57, 0xd6, 0x7c, 0x91, 0x6e, 0xa5, 0xe2, 0x15, 0x08, 0x89, 0x1f, 0xea, 0xe4,
	0xaf, 0x35, 0x7f, 0xf6, 0x26, 0xfb, 0x70, 0x7c, 0x74, 0xeb, 0x50, 0x5d, 0x48, 0x94, 0xf2, 0x6a,
	0xfe, 0xd5, 0x76, 0x66, 0x29, 0xaf, 0x2e, 0xee, 0x36, 0xf7, 0xcf, 0xe4, 0xb6, 0x6e, 0x6f, 0xb4,
	0xb6, 0xde, 0xae, 0x3a, 0x95, 0x9c, 0x97, 0x94, 0xe5, 0x88, 0x8b, 0x1c, 0xe7, 0xc0, 0x0e, 0x17,
	0xc4, 0x47, 0xdf, 0xe1, 0x4f, 0xfa, 0xcc, 0x80, 0x1f, 0xf6, 0x8d, 0x65, 0x18, 0xfe, 0xb4, 0xfd,
	0xa5, 0x16, 0x0c, 0x33, 0x89, 0x34, 0x6c, 0x51, 0x14, 0xa0, 0xce, 0x58, 0xee, 0x0d, 0x25, 0x0e,
	0x33, 0x19, 0xf7, 0x94, 0x38, 0x0a, 0x62, 0x43, 0xf9, 0x6b, 0xcf, 0xf4, 0x3b, 0x21, 0x61, 0x26,
	0x09, 0xe9, 0x49, 0x84, 0x44, 0x01, 0x21, 0x86, 0xf6, 0xee, 0xe6, 0x61, 0xce, 0x47, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x54, 0x77, 0xab, 0xaa, 0x4b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotelPerformanceViewServiceClient is the client API for HotelPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotelPerformanceViewServiceClient interface {
	// Returns the requested Hotel Performance View in full detail.
	GetHotelPerformanceView(ctx context.Context, in *GetHotelPerformanceViewRequest, opts ...grpc.CallOption) (*resources.HotelPerformanceView, error)
}

type hotelPerformanceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewHotelPerformanceViewServiceClient(cc *grpc.ClientConn) HotelPerformanceViewServiceClient {
	return &hotelPerformanceViewServiceClient{cc}
}

func (c *hotelPerformanceViewServiceClient) GetHotelPerformanceView(ctx context.Context, in *GetHotelPerformanceViewRequest, opts ...grpc.CallOption) (*resources.HotelPerformanceView, error) {
	out := new(resources.HotelPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.HotelPerformanceViewService/GetHotelPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelPerformanceViewServiceServer is the server API for HotelPerformanceViewService service.
type HotelPerformanceViewServiceServer interface {
	// Returns the requested Hotel Performance View in full detail.
	GetHotelPerformanceView(context.Context, *GetHotelPerformanceViewRequest) (*resources.HotelPerformanceView, error)
}

func RegisterHotelPerformanceViewServiceServer(s *grpc.Server, srv HotelPerformanceViewServiceServer) {
	s.RegisterService(&_HotelPerformanceViewService_serviceDesc, srv)
}

func _HotelPerformanceViewService_GetHotelPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelPerformanceViewServiceServer).GetHotelPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.HotelPerformanceViewService/GetHotelPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelPerformanceViewServiceServer).GetHotelPerformanceView(ctx, req.(*GetHotelPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.HotelPerformanceViewService",
	HandlerType: (*HotelPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelPerformanceView",
			Handler:    _HotelPerformanceViewService_GetHotelPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/hotel_performance_view_service.proto",
}