//
//Copyright 2021 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: api/grpc/enricher/api.proto

package api_enricher

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

type SyscallsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile string `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *SyscallsRequest) Reset() {
	*x = SyscallsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_enricher_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyscallsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyscallsRequest) ProtoMessage() {}

func (x *SyscallsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_enricher_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyscallsRequest.ProtoReflect.Descriptor instead.
func (*SyscallsRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_enricher_api_proto_rawDescGZIP(), []int{0}
}

func (x *SyscallsRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

type SyscallsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Syscalls []string `protobuf:"bytes,1,rep,name=syscalls,proto3" json:"syscalls,omitempty"`
	GoArch   string   `protobuf:"bytes,2,opt,name=go_arch,json=goArch,proto3" json:"go_arch,omitempty"`
}

func (x *SyscallsResponse) Reset() {
	*x = SyscallsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_enricher_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyscallsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyscallsResponse) ProtoMessage() {}

func (x *SyscallsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_enricher_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyscallsResponse.ProtoReflect.Descriptor instead.
func (*SyscallsResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_enricher_api_proto_rawDescGZIP(), []int{1}
}

func (x *SyscallsResponse) GetSyscalls() []string {
	if x != nil {
		return x.Syscalls
	}
	return nil
}

func (x *SyscallsResponse) GetGoArch() string {
	if x != nil {
		return x.GoArch
	}
	return ""
}

type AvcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile string `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *AvcRequest) Reset() {
	*x = AvcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_enricher_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvcRequest) ProtoMessage() {}

func (x *AvcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_enricher_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvcRequest.ProtoReflect.Descriptor instead.
func (*AvcRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_enricher_api_proto_rawDescGZIP(), []int{2}
}

func (x *AvcRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

type AvcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avc []*AvcResponse_SelinuxAvc `protobuf:"bytes,1,rep,name=avc,proto3" json:"avc,omitempty"`
}

func (x *AvcResponse) Reset() {
	*x = AvcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_enricher_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvcResponse) ProtoMessage() {}

func (x *AvcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_enricher_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvcResponse.ProtoReflect.Descriptor instead.
func (*AvcResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_enricher_api_proto_rawDescGZIP(), []int{3}
}

func (x *AvcResponse) GetAvc() []*AvcResponse_SelinuxAvc {
	if x != nil {
		return x.Avc
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_enricher_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_enricher_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_enricher_api_proto_rawDescGZIP(), []int{4}
}

type AvcResponse_SelinuxAvc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Perm     string `protobuf:"bytes,1,opt,name=perm,proto3" json:"perm,omitempty"`
	Scontext string `protobuf:"bytes,2,opt,name=scontext,proto3" json:"scontext,omitempty"`
	Tcontext string `protobuf:"bytes,3,opt,name=tcontext,proto3" json:"tcontext,omitempty"`
	Tclass   string `protobuf:"bytes,4,opt,name=tclass,proto3" json:"tclass,omitempty"`
}

func (x *AvcResponse_SelinuxAvc) Reset() {
	*x = AvcResponse_SelinuxAvc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_enricher_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvcResponse_SelinuxAvc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvcResponse_SelinuxAvc) ProtoMessage() {}

func (x *AvcResponse_SelinuxAvc) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_enricher_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvcResponse_SelinuxAvc.ProtoReflect.Descriptor instead.
func (*AvcResponse_SelinuxAvc) Descriptor() ([]byte, []int) {
	return file_api_grpc_enricher_api_proto_rawDescGZIP(), []int{3, 0}
}

func (x *AvcResponse_SelinuxAvc) GetPerm() string {
	if x != nil {
		return x.Perm
	}
	return ""
}

func (x *AvcResponse_SelinuxAvc) GetScontext() string {
	if x != nil {
		return x.Scontext
	}
	return ""
}

func (x *AvcResponse_SelinuxAvc) GetTcontext() string {
	if x != nil {
		return x.Tcontext
	}
	return ""
}

func (x *AvcResponse_SelinuxAvc) GetTclass() string {
	if x != nil {
		return x.Tclass
	}
	return ""
}

var File_api_grpc_enricher_api_proto protoreflect.FileDescriptor

var file_api_grpc_enricher_api_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x53,
	0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x47, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x63,
	0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x6f, 0x5f, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x41, 0x72, 0x63,
	0x68, 0x22, 0x26, 0x0a, 0x0a, 0x41, 0x76, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x41, 0x76,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x61, 0x76, 0x63,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x41, 0x76, 0x63, 0x52, 0x03, 0x61, 0x76,
	0x63, 0x1a, 0x70, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x41, 0x76, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xab, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x08, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73,
	0x63, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x63,
	0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x12,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53,
	0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x04, 0x41, 0x76, 0x63, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x41,
	0x76, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x76, 0x63, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_enricher_api_proto_rawDescOnce sync.Once
	file_api_grpc_enricher_api_proto_rawDescData = file_api_grpc_enricher_api_proto_rawDesc
)

func file_api_grpc_enricher_api_proto_rawDescGZIP() []byte {
	file_api_grpc_enricher_api_proto_rawDescOnce.Do(func() {
		file_api_grpc_enricher_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_enricher_api_proto_rawDescData)
	})
	return file_api_grpc_enricher_api_proto_rawDescData
}

var file_api_grpc_enricher_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_grpc_enricher_api_proto_goTypes = []interface{}{
	(*SyscallsRequest)(nil),        // 0: api_enricher.SyscallsRequest
	(*SyscallsResponse)(nil),       // 1: api_enricher.SyscallsResponse
	(*AvcRequest)(nil),             // 2: api_enricher.AvcRequest
	(*AvcResponse)(nil),            // 3: api_enricher.AvcResponse
	(*EmptyResponse)(nil),          // 4: api_enricher.EmptyResponse
	(*AvcResponse_SelinuxAvc)(nil), // 5: api_enricher.AvcResponse.SelinuxAvc
}
var file_api_grpc_enricher_api_proto_depIdxs = []int32{
	5, // 0: api_enricher.AvcResponse.avc:type_name -> api_enricher.AvcResponse.SelinuxAvc
	0, // 1: api_enricher.Enricher.Syscalls:input_type -> api_enricher.SyscallsRequest
	0, // 2: api_enricher.Enricher.ResetSyscalls:input_type -> api_enricher.SyscallsRequest
	2, // 3: api_enricher.Enricher.Avcs:input_type -> api_enricher.AvcRequest
	2, // 4: api_enricher.Enricher.ResetAvcs:input_type -> api_enricher.AvcRequest
	1, // 5: api_enricher.Enricher.Syscalls:output_type -> api_enricher.SyscallsResponse
	4, // 6: api_enricher.Enricher.ResetSyscalls:output_type -> api_enricher.EmptyResponse
	3, // 7: api_enricher.Enricher.Avcs:output_type -> api_enricher.AvcResponse
	4, // 8: api_enricher.Enricher.ResetAvcs:output_type -> api_enricher.EmptyResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpc_enricher_api_proto_init() }
func file_api_grpc_enricher_api_proto_init() {
	if File_api_grpc_enricher_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_enricher_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyscallsRequest); i {
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
		file_api_grpc_enricher_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyscallsResponse); i {
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
		file_api_grpc_enricher_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvcRequest); i {
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
		file_api_grpc_enricher_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvcResponse); i {
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
		file_api_grpc_enricher_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_api_grpc_enricher_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvcResponse_SelinuxAvc); i {
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
			RawDescriptor: file_api_grpc_enricher_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_enricher_api_proto_goTypes,
		DependencyIndexes: file_api_grpc_enricher_api_proto_depIdxs,
		MessageInfos:      file_api_grpc_enricher_api_proto_msgTypes,
	}.Build()
	File_api_grpc_enricher_api_proto = out.File
	file_api_grpc_enricher_api_proto_rawDesc = nil
	file_api_grpc_enricher_api_proto_goTypes = nil
	file_api_grpc_enricher_api_proto_depIdxs = nil
}
