// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blob.proto

/*
Package gitaly is a generated protocol buffer package.

It is generated from these files:
	blob.proto
	commit.proto
	deprecated-services.proto
	diff.proto
	notifications.proto
	ref.proto
	repository-service.proto
	shared.proto
	smarthttp.proto
	ssh.proto

It has these top-level messages:
	GetBlobRequest
	GetBlobResponse
	CommitIsAncestorRequest
	CommitIsAncestorResponse
	TreeEntryRequest
	TreeEntryResponse
	CommitsBetweenRequest
	CommitsBetweenResponse
	CountCommitsRequest
	CountCommitsResponse
	CommitDiffRequest
	CommitDiffResponse
	CommitDeltaRequest
	CommitDelta
	CommitDeltaResponse
	PostReceiveRequest
	PostReceiveResponse
	FindDefaultBranchNameRequest
	FindDefaultBranchNameResponse
	FindAllBranchNamesRequest
	FindAllBranchNamesResponse
	FindAllTagNamesRequest
	FindAllTagNamesResponse
	FindRefNameRequest
	FindRefNameResponse
	FindLocalBranchesRequest
	FindLocalBranchesResponse
	FindLocalBranchResponse
	FindLocalBranchCommitAuthor
	RepositoryExistsRequest
	RepositoryExistsResponse
	Repository
	GitCommit
	CommitAuthor
	ExitStatus
	InfoRefsRequest
	InfoRefsResponse
	PostUploadPackRequest
	PostUploadPackResponse
	PostReceivePackRequest
	PostReceivePackResponse
	SSHUploadPackRequest
	SSHUploadPackResponse
	SSHReceivePackRequest
	SSHReceivePackResponse
*/
package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type GetBlobRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Object ID (SHA1) of the blob we want to get
	Oid string `protobuf:"bytes,2,opt,name=oid" json:"oid,omitempty"`
	// Maximum number of bytes we want to receive. Use '-1' to get the full blob no matter how big.
	Limit int64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetBlobRequest) Reset()                    { *m = GetBlobRequest{} }
func (m *GetBlobRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlobRequest) ProtoMessage()               {}
func (*GetBlobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetBlobRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *GetBlobRequest) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func (m *GetBlobRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBlobResponse struct {
	// Blob size; present only in first response message
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Chunk of blob data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Object ID of the actual blob returned. Empty if no blob was found.
	Oid string `protobuf:"bytes,3,opt,name=oid" json:"oid,omitempty"`
}

func (m *GetBlobResponse) Reset()                    { *m = GetBlobResponse{} }
func (m *GetBlobResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlobResponse) ProtoMessage()               {}
func (*GetBlobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetBlobResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetBlobResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlobResponse) GetOid() string {
	if m != nil {
		return m.Oid
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBlobRequest)(nil), "gitaly.GetBlobRequest")
	proto.RegisterType((*GetBlobResponse)(nil), "gitaly.GetBlobResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlobService service

type BlobServiceClient interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error)
}

type blobServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlobServiceClient(cc *grpc.ClientConn) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (BlobService_GetBlobClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlobService_serviceDesc.Streams[0], c.cc, "/gitaly.BlobService/GetBlob", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetBlobClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetBlobClient interface {
	Recv() (*GetBlobResponse, error)
	grpc.ClientStream
}

type blobServiceGetBlobClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetBlobClient) Recv() (*GetBlobResponse, error) {
	m := new(GetBlobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BlobService service

type BlobServiceServer interface {
	// GetBlob returns the contents of a blob object referenced by its object
	// ID. We use a stream to return a chunked arbitrarily large binary
	// response
	GetBlob(*GetBlobRequest, BlobService_GetBlobServer) error
}

func RegisterBlobServiceServer(s *grpc.Server, srv BlobServiceServer) {
	s.RegisterService(&_BlobService_serviceDesc, srv)
}

func _BlobService_GetBlob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlobRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).GetBlob(m, &blobServiceGetBlobServer{stream})
}

type BlobService_GetBlobServer interface {
	Send(*GetBlobResponse) error
	grpc.ServerStream
}

type blobServiceGetBlobServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetBlobServer) Send(m *GetBlobResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlob",
			Handler:       _BlobService_GetBlob_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blob.proto",
}

func init() { proto.RegisterFile("blob.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4b, 0xc7, 0x30,
	0x10, 0xc5, 0x8d, 0xd1, 0xbf, 0x78, 0x2d, 0x2a, 0x87, 0x68, 0xe9, 0x54, 0x3a, 0x75, 0x2a, 0x52,
	0x77, 0x07, 0x17, 0x07, 0x71, 0x89, 0x9f, 0x20, 0xb1, 0x87, 0x06, 0xa2, 0x57, 0x93, 0x28, 0xd4,
	0x4f, 0x2f, 0x4d, 0x6c, 0x51, 0xdc, 0x5e, 0x5e, 0x92, 0xf7, 0x7b, 0x77, 0x00, 0xc6, 0xb1, 0xe9,
	0x27, 0xcf, 0x91, 0x71, 0xf7, 0x6c, 0xa3, 0x76, 0x73, 0x5d, 0x86, 0x17, 0xed, 0x69, 0xcc, 0x6e,
	0xeb, 0xe0, 0xe4, 0x8e, 0xe2, 0xad, 0x63, 0xa3, 0xe8, 0xfd, 0x83, 0x42, 0xc4, 0x01, 0xc0, 0xd3,
	0xc4, 0xc1, 0x46, 0xf6, 0x73, 0x25, 0x1a, 0xd1, 0x15, 0x03, 0xf6, 0xf9, 0x73, 0xaf, 0xb6, 0x1b,
	0xf5, 0xeb, 0x15, 0x9e, 0x81, 0x64, 0x3b, 0x56, 0xfb, 0x8d, 0xe8, 0x8e, 0xd5, 0x22, 0xf1, 0x1c,
	0x0e, 0x9d, 0x7d, 0xb5, 0xb1, 0x92, 0x8d, 0xe8, 0xa4, 0xca, 0x87, 0xf6, 0x1e, 0x4e, 0x37, 0x5a,
	0x98, 0xf8, 0x2d, 0x10, 0x22, 0x1c, 0x04, 0xfb, 0x45, 0x09, 0x24, 0x55, 0xd2, 0x8b, 0x37, 0xea,
	0xa8, 0x53, 0x5e, 0xa9, 0x92, 0x5e, 0x11, 0x72, 0x43, 0x0c, 0x0f, 0x50, 0x2c, 0x49, 0x8f, 0xe4,
	0x3f, 0xed, 0x13, 0xe1, 0x0d, 0x1c, 0xfd, 0x64, 0xe3, 0xc5, 0x5a, 0xf7, 0xef, 0x68, 0xf5, 0xe5,
	0x3f, 0x3f, 0x97, 0x68, 0xf7, 0xae, 0x84, 0xd9, 0xa5, 0x85, 0x5c, 0x7f, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x77, 0x1a, 0x6d, 0x34, 0x01, 0x00, 0x00,
}
