// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jobs/service.proto

package jobspb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type JobStatusRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobStatusRequest) Reset()         { *m = JobStatusRequest{} }
func (m *JobStatusRequest) String() string { return proto.CompactTextString(m) }
func (*JobStatusRequest) ProtoMessage()    {}
func (*JobStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_886578d0f940e42c, []int{0}
}

func (m *JobStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobStatusRequest.Unmarshal(m, b)
}
func (m *JobStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobStatusRequest.Marshal(b, m, deterministic)
}
func (m *JobStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobStatusRequest.Merge(m, src)
}
func (m *JobStatusRequest) XXX_Size() int {
	return xxx_messageInfo_JobStatusRequest.Size(m)
}
func (m *JobStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobStatusRequest proto.InternalMessageInfo

func (m *JobStatusRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type JobStatusResponse struct {
	JobId                string               `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status               string               `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JobStatusResponse) Reset()         { *m = JobStatusResponse{} }
func (m *JobStatusResponse) String() string { return proto.CompactTextString(m) }
func (*JobStatusResponse) ProtoMessage()    {}
func (*JobStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_886578d0f940e42c, []int{1}
}

func (m *JobStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobStatusResponse.Unmarshal(m, b)
}
func (m *JobStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobStatusResponse.Marshal(b, m, deterministic)
}
func (m *JobStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobStatusResponse.Merge(m, src)
}
func (m *JobStatusResponse) XXX_Size() int {
	return xxx_messageInfo_JobStatusResponse.Size(m)
}
func (m *JobStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobStatusResponse proto.InternalMessageInfo

func (m *JobStatusResponse) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *JobStatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *JobStatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *JobStatusResponse) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func init() {
	proto.RegisterType((*JobStatusRequest)(nil), "jobs.JobStatusRequest")
	proto.RegisterType((*JobStatusResponse)(nil), "jobs.JobStatusResponse")
}

func init() { proto.RegisterFile("jobs/service.proto", fileDescriptor_886578d0f940e42c) }

var fileDescriptor_886578d0f940e42c = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x49, 0x5f, 0x5f, 0xdf, 0x7b, 0xdb, 0x52, 0xfa, 0x16, 0xac, 0x21, 0x08, 0x86, 0x9c,
	0x5a, 0xb0, 0x09, 0xd4, 0xb3, 0x07, 0x7b, 0x29, 0xf6, 0x54, 0xaa, 0x5e, 0xbc, 0x94, 0xdd, 0x66,
	0x0c, 0x09, 0x4d, 0x66, 0xed, 0x6c, 0x54, 0x10, 0x2f, 0x7e, 0x04, 0x3d, 0xfa, 0xb1, 0xfc, 0x0a,
	0x7e, 0x10, 0xc9, 0x6e, 0x2a, 0x62, 0xf1, 0xb4, 0xcc, 0xcc, 0x8f, 0xe1, 0xff, 0x9b, 0x65, 0x3c,
	0x43, 0x49, 0x11, 0xc1, 0xe6, 0x36, 0x5d, 0x41, 0xa8, 0x36, 0xa8, 0x91, 0x37, 0xab, 0x9e, 0x77,
	0x90, 0x20, 0x26, 0x6b, 0x88, 0x84, 0x4a, 0x23, 0x51, 0x14, 0xa8, 0x85, 0x4e, 0xb1, 0x20, 0xcb,
	0x78, 0x87, 0xf5, 0xd4, 0x54, 0xb2, 0xbc, 0x8e, 0x74, 0x9a, 0x03, 0x69, 0x91, 0xab, 0x1a, 0x38,
	0x32, 0xcf, 0x6a, 0x94, 0x40, 0x31, 0xa2, 0x3b, 0x91, 0x24, 0xb0, 0x89, 0x50, 0x99, 0x15, 0xbb,
	0xeb, 0x82, 0x21, 0xeb, 0xcd, 0x50, 0x9e, 0x6b, 0xa1, 0x4b, 0x5a, 0xc0, 0x4d, 0x09, 0xa4, 0xf9,
	0x1e, 0x6b, 0x65, 0x28, 0x97, 0x69, 0xec, 0x3a, 0xbe, 0x33, 0xf8, 0xb7, 0xf8, 0x9d, 0xa1, 0x3c,
	0x8b, 0x83, 0x57, 0x87, 0xfd, 0xff, 0xc2, 0x92, 0xc2, 0x82, 0xe0, 0x07, 0x98, 0xf7, 0x59, 0x8b,
	0x0c, 0xe8, 0x36, 0x4c, 0xbb, 0xae, 0xb8, 0xcb, 0xfe, 0xe4, 0x40, 0x24, 0x12, 0x70, 0x7f, 0x99,
	0xc1, 0xb6, 0xe4, 0x27, 0xac, 0xb3, 0x16, 0xa4, 0x97, 0xa5, 0x8a, 0x85, 0x86, 0xd8, 0x6d, 0xfa,
	0xce, 0xa0, 0x3d, 0xf6, 0x42, 0xeb, 0x1b, 0x6e, 0x7d, 0xc3, 0x8b, 0xad, 0xef, 0xa2, 0x5d, 0xf1,
	0x97, 0x16, 0x1f, 0xdf, 0x33, 0x56, 0x85, 0xb3, 0xf7, 0xe4, 0x19, 0xeb, 0x4c, 0x41, 0x7f, 0xa6,
	0xe5, 0xfd, 0xb0, 0x3a, 0x6d, 0xf8, 0x5d, 0xd5, 0xdb, 0xdf, 0xe9, 0x5b, 0xad, 0x60, 0xf8, 0x7c,
	0xda, 0xf3, 0xba, 0x53, 0xd0, 0xfe, 0x0c, 0xa5, 0x6f, 0x87, 0x4f, 0x6f, 0xef, 0x2f, 0x8d, 0x1e,
	0xef, 0x46, 0xe6, 0xeb, 0x1e, 0xac, 0xf6, 0xe3, 0x24, 0x60, 0x7f, 0x57, 0x98, 0x9b, 0x45, 0x93,
	0x4e, 0x1d, 0x60, 0x5e, 0xa5, 0x9d, 0x3b, 0x57, 0xd5, 0x6d, 0x48, 0x49, 0xd9, 0x32, 0xf1, 0x8f,
	0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x5d, 0xec, 0x2c, 0xf6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	GetJobStatus(ctx context.Context, in *JobStatusRequest, opts ...grpc.CallOption) (*JobStatusResponse, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) GetJobStatus(ctx context.Context, in *JobStatusRequest, opts ...grpc.CallOption) (*JobStatusResponse, error) {
	out := new(JobStatusResponse)
	err := c.cc.Invoke(ctx, "/jobs.JobService/GetJobStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	GetJobStatus(context.Context, *JobStatusRequest) (*JobStatusResponse, error)
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_GetJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.JobService/GetJobStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobStatus(ctx, req.(*JobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jobs.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobStatus",
			Handler:    _JobService_GetJobStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jobs/service.proto",
}
