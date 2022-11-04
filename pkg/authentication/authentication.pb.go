// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: authentication.proto

package authentication

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

// Represents the possible authentication outcomes
type AuthenticationResponse_Outcome int32

const (
	// The authentication token was empty and the user should be redirected to an external identity provider to obtain a token
	AuthenticationResponse_REDIRECT AuthenticationResponse_Outcome = 0
	// The supplied authentication token was invalid
	AuthenticationResponse_INVALID_TOKEN AuthenticationResponse_Outcome = 1
	// The user was successfully authenticated
	AuthenticationResponse_AUTHENTICATED AuthenticationResponse_Outcome = 2
	// An error occurred when attempting to perform authentication
	AuthenticationResponse_ERROR AuthenticationResponse_Outcome = 3
)

// Enum value maps for AuthenticationResponse_Outcome.
var (
	AuthenticationResponse_Outcome_name = map[int32]string{
		0: "REDIRECT",
		1: "INVALID_TOKEN",
		2: "AUTHENTICATED",
		3: "ERROR",
	}
	AuthenticationResponse_Outcome_value = map[string]int32{
		"REDIRECT":      0,
		"INVALID_TOKEN": 1,
		"AUTHENTICATED": 2,
		"ERROR":         3,
	}
)

func (x AuthenticationResponse_Outcome) Enum() *AuthenticationResponse_Outcome {
	p := new(AuthenticationResponse_Outcome)
	*p = x
	return p
}

func (x AuthenticationResponse_Outcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthenticationResponse_Outcome) Descriptor() protoreflect.EnumDescriptor {
	return file_authentication_proto_enumTypes[0].Descriptor()
}

func (AuthenticationResponse_Outcome) Type() protoreflect.EnumType {
	return &file_authentication_proto_enumTypes[0]
}

func (x AuthenticationResponse_Outcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthenticationResponse_Outcome.Descriptor instead.
func (AuthenticationResponse_Outcome) EnumDescriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{1, 0}
}

type AuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The authentication token, which we expect to have been issued by an external identity provider.
	// This will be empty in the initial (pre-redirect) request for workflows where users are redirected to a URL to obtain a token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Used to indicate to the Authentication Plugin which identity provider should be used when the plugin implementation supports interacting with multiple
	// identity providers. This field is ignored when the plugin implementation supports only a single identity provider and should be left as an empty string.
	// For workflows where users are redirected to a URL to obtain a token, this value must be the same in both the pre-redirect and post-redirect requests.
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthenticationRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type AuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The outcome of the authentication request
	Outcome AuthenticationResponse_Outcome `protobuf:"varint,1,opt,name=outcome,proto3,enum=proto.AuthenticationResponse_Outcome" json:"outcome,omitempty"`
	// At most one of these fields will ever be set
	//
	// Types that are assignable to Payload:
	//	*AuthenticationResponse_Url
	//	*AuthenticationResponse_Id
	//	*AuthenticationResponse_Error
	Payload isAuthenticationResponse_Payload `protobuf_oneof:"payload"`
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationResponse) GetOutcome() AuthenticationResponse_Outcome {
	if x != nil {
		return x.Outcome
	}
	return AuthenticationResponse_REDIRECT
}

func (m *AuthenticationResponse) GetPayload() isAuthenticationResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *AuthenticationResponse) GetUrl() string {
	if x, ok := x.GetPayload().(*AuthenticationResponse_Url); ok {
		return x.Url
	}
	return ""
}

func (x *AuthenticationResponse) GetId() string {
	if x, ok := x.GetPayload().(*AuthenticationResponse_Id); ok {
		return x.Id
	}
	return ""
}

func (x *AuthenticationResponse) GetError() string {
	if x, ok := x.GetPayload().(*AuthenticationResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isAuthenticationResponse_Payload interface {
	isAuthenticationResponse_Payload()
}

type AuthenticationResponse_Url struct {
	// The URL of an external identity provider to which the user should be redirected to obtain an authentication token
	// (Set when the outcome is REDIRECT)
	Url string `protobuf:"bytes,2,opt,name=url,proto3,oneof"`
}

type AuthenticationResponse_Id struct {
	// The unique ID representing the authenticated user
	// (Set when the outcome is AUTHENTICATED)
	Id string `protobuf:"bytes,3,opt,name=id,proto3,oneof"`
}

type AuthenticationResponse_Error struct {
	// The message providing details of the error that occurred
	// (Set when the outcome is ERROR)
	Error string `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*AuthenticationResponse_Url) isAuthenticationResponse_Payload() {}

func (*AuthenticationResponse_Id) isAuthenticationResponse_Payload() {}

func (*AuthenticationResponse_Error) isAuthenticationResponse_Payload() {}

var File_authentication_proto protoreflect.FileDescriptor

var file_authentication_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xec, 0x01, 0x0a, 0x16, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x07, 0x6f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x48, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x09, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x65, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x4d, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_authentication_proto_rawDescOnce sync.Once
	file_authentication_proto_rawDescData = file_authentication_proto_rawDesc
)

func file_authentication_proto_rawDescGZIP() []byte {
	file_authentication_proto_rawDescOnce.Do(func() {
		file_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_authentication_proto_rawDescData)
	})
	return file_authentication_proto_rawDescData
}

var file_authentication_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_authentication_proto_goTypes = []interface{}{
	(AuthenticationResponse_Outcome)(0), // 0: proto.AuthenticationResponse.Outcome
	(*AuthenticationRequest)(nil),       // 1: proto.AuthenticationRequest
	(*AuthenticationResponse)(nil),      // 2: proto.AuthenticationResponse
}
var file_authentication_proto_depIdxs = []int32{
	0, // 0: proto.AuthenticationResponse.outcome:type_name -> proto.AuthenticationResponse.Outcome
	1, // 1: proto.AuthenticationPlugin.Authenticate:input_type -> proto.AuthenticationRequest
	2, // 2: proto.AuthenticationPlugin.Authenticate:output_type -> proto.AuthenticationResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_authentication_proto_init() }
func file_authentication_proto_init() {
	if File_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationRequest); i {
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
		file_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationResponse); i {
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
	file_authentication_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AuthenticationResponse_Url)(nil),
		(*AuthenticationResponse_Id)(nil),
		(*AuthenticationResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authentication_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authentication_proto_goTypes,
		DependencyIndexes: file_authentication_proto_depIdxs,
		EnumInfos:         file_authentication_proto_enumTypes,
		MessageInfos:      file_authentication_proto_msgTypes,
	}.Build()
	File_authentication_proto = out.File
	file_authentication_proto_rawDesc = nil
	file_authentication_proto_goTypes = nil
	file_authentication_proto_depIdxs = nil
}
