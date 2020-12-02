// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.7.1
// source: protos/hangman_api/v1/hangman_api.proto

package hangman_api_v1

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// TagBodyDetails detail for tag mail body action
type CreateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGameRequest) Reset() {
	*x = CreateGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameRequest) ProtoMessage() {}

func (x *CreateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameRequest.ProtoReflect.Descriptor instead.
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return file_protos_hangman_api_v1_hangman_api_proto_rawDescGZIP(), []int{0}
}

type UpdateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Guess string `protobuf:"bytes,2,opt,name=guess,proto3" json:"guess,omitempty"`
}

func (x *UpdateGameRequest) Reset() {
	*x = UpdateGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGameRequest) ProtoMessage() {}

func (x *UpdateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGameRequest.ProtoReflect.Descriptor instead.
func (*UpdateGameRequest) Descriptor() ([]byte, []int) {
	return file_protos_hangman_api_v1_hangman_api_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateGameRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateGameRequest) GetGuess() string {
	if x != nil {
		return x.Guess
	}
	return ""
}

type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_protos_hangman_api_v1_hangman_api_proto_rawDescGZIP(), []int{2}
}

func (x *GameRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Word         string   `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	GuessesLeft  int32    `protobuf:"varint,3,opt,name=guessesLeft,proto3" json:"guessesLeft,omitempty"`
	WrongGuesses []string `protobuf:"bytes,4,rep,name=wrongGuesses,proto3" json:"wrongGuesses,omitempty"`
	State        string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GameResponse) Reset() {
	*x = GameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameResponse) ProtoMessage() {}

func (x *GameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameResponse.ProtoReflect.Descriptor instead.
func (*GameResponse) Descriptor() ([]byte, []int) {
	return file_protos_hangman_api_v1_hangman_api_proto_rawDescGZIP(), []int{3}
}

func (x *GameResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GameResponse) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *GameResponse) GetGuessesLeft() int32 {
	if x != nil {
		return x.GuessesLeft
	}
	return 0
}

func (x *GameResponse) GetWrongGuesses() []string {
	if x != nil {
		return x.WrongGuesses
	}
	return nil
}

func (x *GameResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type DeleteGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGameRequest) Reset() {
	*x = DeleteGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGameRequest) ProtoMessage() {}

func (x *DeleteGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGameRequest.ProtoReflect.Descriptor instead.
func (*DeleteGameRequest) Descriptor() ([]byte, []int) {
	return file_protos_hangman_api_v1_hangman_api_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGameRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteGameResponse) Reset() {
	*x = DeleteGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGameResponse) ProtoMessage() {}

func (x *DeleteGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hangman_api_v1_hangman_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGameResponse.ProtoReflect.Descriptor instead.
func (*DeleteGameResponse) Descriptor() ([]byte, []int) {
	return file_protos_hangman_api_v1_hangman_api_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGameResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

var File_protos_hangman_api_v1_hangman_api_proto protoreflect.FileDescriptor

var file_protos_hangman_api_v1_hangman_api_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61, 0x6e,
	0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61, 0x6e, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68, 0x61, 0x6e, 0x67, 0x6d,
	0x61, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x67, 0x75, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x05, 0x67, 0x75, 0x65,
	0x73, 0x73, 0x22, 0x26, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x04, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x29, 0x0a, 0x0b, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x4c, 0x65, 0x66, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52,
	0x0b, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x2c, 0x0a, 0x0c,
	0x77, 0x72, 0x6f, 0x6e, 0x67, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x0c, 0x77, 0x72,
	0x6f, 0x6e, 0x67, 0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xc6, 0x03, 0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x67,
	0x6d, 0x61, 0x6e, 0x41, 0x50, 0x49, 0x12, 0x61, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x2e, 0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x67, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e, 0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x61, 0x6e,
	0x67, 0x6d, 0x61, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65,
	0x47, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x61, 0x6e,
	0x67, 0x6d, 0x61, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x32, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x67, 0x75, 0x65, 0x73, 0x73, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x6f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x2e,
	0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x68, 0x61, 0x6e, 0x67, 0x6d, 0x61, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_hangman_api_v1_hangman_api_proto_rawDescOnce sync.Once
	file_protos_hangman_api_v1_hangman_api_proto_rawDescData = file_protos_hangman_api_v1_hangman_api_proto_rawDesc
)

func file_protos_hangman_api_v1_hangman_api_proto_rawDescGZIP() []byte {
	file_protos_hangman_api_v1_hangman_api_proto_rawDescOnce.Do(func() {
		file_protos_hangman_api_v1_hangman_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_hangman_api_v1_hangman_api_proto_rawDescData)
	})
	return file_protos_hangman_api_v1_hangman_api_proto_rawDescData
}

var file_protos_hangman_api_v1_hangman_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_hangman_api_v1_hangman_api_proto_goTypes = []interface{}{
	(*CreateGameRequest)(nil),  // 0: hangman_api.v1.CreateGameRequest
	(*UpdateGameRequest)(nil),  // 1: hangman_api.v1.UpdateGameRequest
	(*GameRequest)(nil),        // 2: hangman_api.v1.GameRequest
	(*GameResponse)(nil),       // 3: hangman_api.v1.GameResponse
	(*DeleteGameRequest)(nil),  // 4: hangman_api.v1.DeleteGameRequest
	(*DeleteGameResponse)(nil), // 5: hangman_api.v1.DeleteGameResponse
}
var file_protos_hangman_api_v1_hangman_api_proto_depIdxs = []int32{
	2, // 0: hangman_api.v1.HangmanAPI.GetGame:input_type -> hangman_api.v1.GameRequest
	0, // 1: hangman_api.v1.HangmanAPI.CreateGame:input_type -> hangman_api.v1.CreateGameRequest
	1, // 2: hangman_api.v1.HangmanAPI.UpdateGameGuesses:input_type -> hangman_api.v1.UpdateGameRequest
	4, // 3: hangman_api.v1.HangmanAPI.DeleteGame:input_type -> hangman_api.v1.DeleteGameRequest
	3, // 4: hangman_api.v1.HangmanAPI.GetGame:output_type -> hangman_api.v1.GameResponse
	3, // 5: hangman_api.v1.HangmanAPI.CreateGame:output_type -> hangman_api.v1.GameResponse
	3, // 6: hangman_api.v1.HangmanAPI.UpdateGameGuesses:output_type -> hangman_api.v1.GameResponse
	5, // 7: hangman_api.v1.HangmanAPI.DeleteGame:output_type -> hangman_api.v1.DeleteGameResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_hangman_api_v1_hangman_api_proto_init() }
func file_protos_hangman_api_v1_hangman_api_proto_init() {
	if File_protos_hangman_api_v1_hangman_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_hangman_api_v1_hangman_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameRequest); i {
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
		file_protos_hangman_api_v1_hangman_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGameRequest); i {
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
		file_protos_hangman_api_v1_hangman_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
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
		file_protos_hangman_api_v1_hangman_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameResponse); i {
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
		file_protos_hangman_api_v1_hangman_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGameRequest); i {
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
		file_protos_hangman_api_v1_hangman_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGameResponse); i {
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
			RawDescriptor: file_protos_hangman_api_v1_hangman_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_hangman_api_v1_hangman_api_proto_goTypes,
		DependencyIndexes: file_protos_hangman_api_v1_hangman_api_proto_depIdxs,
		MessageInfos:      file_protos_hangman_api_v1_hangman_api_proto_msgTypes,
	}.Build()
	File_protos_hangman_api_v1_hangman_api_proto = out.File
	file_protos_hangman_api_v1_hangman_api_proto_rawDesc = nil
	file_protos_hangman_api_v1_hangman_api_proto_goTypes = nil
	file_protos_hangman_api_v1_hangman_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HangmanAPIClient is the client API for HangmanAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HangmanAPIClient interface {
	// GetGame returns the timeline of policy violations
	GetGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameResponse, error)
	// CreateGame creates a new game and returns the id
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*GameResponse, error)
	// UpdateGameGuesses updates the games with guesses
	UpdateGameGuesses(ctx context.Context, in *UpdateGameRequest, opts ...grpc.CallOption) (*GameResponse, error)
	// DeleteGame deletes the game by id
	DeleteGame(ctx context.Context, in *DeleteGameRequest, opts ...grpc.CallOption) (*DeleteGameResponse, error)
}

type hangmanAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewHangmanAPIClient(cc grpc.ClientConnInterface) HangmanAPIClient {
	return &hangmanAPIClient{cc}
}

func (c *hangmanAPIClient) GetGame(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameResponse, error) {
	out := new(GameResponse)
	err := c.cc.Invoke(ctx, "/hangman_api.v1.HangmanAPI/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hangmanAPIClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*GameResponse, error) {
	out := new(GameResponse)
	err := c.cc.Invoke(ctx, "/hangman_api.v1.HangmanAPI/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hangmanAPIClient) UpdateGameGuesses(ctx context.Context, in *UpdateGameRequest, opts ...grpc.CallOption) (*GameResponse, error) {
	out := new(GameResponse)
	err := c.cc.Invoke(ctx, "/hangman_api.v1.HangmanAPI/UpdateGameGuesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hangmanAPIClient) DeleteGame(ctx context.Context, in *DeleteGameRequest, opts ...grpc.CallOption) (*DeleteGameResponse, error) {
	out := new(DeleteGameResponse)
	err := c.cc.Invoke(ctx, "/hangman_api.v1.HangmanAPI/DeleteGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HangmanAPIServer is the server API for HangmanAPI service.
type HangmanAPIServer interface {
	// GetGame returns the timeline of policy violations
	GetGame(context.Context, *GameRequest) (*GameResponse, error)
	// CreateGame creates a new game and returns the id
	CreateGame(context.Context, *CreateGameRequest) (*GameResponse, error)
	// UpdateGameGuesses updates the games with guesses
	UpdateGameGuesses(context.Context, *UpdateGameRequest) (*GameResponse, error)
	// DeleteGame deletes the game by id
	DeleteGame(context.Context, *DeleteGameRequest) (*DeleteGameResponse, error)
}

// UnimplementedHangmanAPIServer can be embedded to have forward compatible implementations.
type UnimplementedHangmanAPIServer struct {
}

func (*UnimplementedHangmanAPIServer) GetGame(context.Context, *GameRequest) (*GameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (*UnimplementedHangmanAPIServer) CreateGame(context.Context, *CreateGameRequest) (*GameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (*UnimplementedHangmanAPIServer) UpdateGameGuesses(context.Context, *UpdateGameRequest) (*GameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGameGuesses not implemented")
}
func (*UnimplementedHangmanAPIServer) DeleteGame(context.Context, *DeleteGameRequest) (*DeleteGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGame not implemented")
}

func RegisterHangmanAPIServer(s *grpc.Server, srv HangmanAPIServer) {
	s.RegisterService(&_HangmanAPI_serviceDesc, srv)
}

func _HangmanAPI_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HangmanAPIServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hangman_api.v1.HangmanAPI/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HangmanAPIServer).GetGame(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HangmanAPI_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HangmanAPIServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hangman_api.v1.HangmanAPI/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HangmanAPIServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HangmanAPI_UpdateGameGuesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HangmanAPIServer).UpdateGameGuesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hangman_api.v1.HangmanAPI/UpdateGameGuesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HangmanAPIServer).UpdateGameGuesses(ctx, req.(*UpdateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HangmanAPI_DeleteGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HangmanAPIServer).DeleteGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hangman_api.v1.HangmanAPI/DeleteGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HangmanAPIServer).DeleteGame(ctx, req.(*DeleteGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HangmanAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hangman_api.v1.HangmanAPI",
	HandlerType: (*HangmanAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGame",
			Handler:    _HangmanAPI_GetGame_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _HangmanAPI_CreateGame_Handler,
		},
		{
			MethodName: "UpdateGameGuesses",
			Handler:    _HangmanAPI_UpdateGameGuesses_Handler,
		},
		{
			MethodName: "DeleteGame",
			Handler:    _HangmanAPI_DeleteGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/hangman_api/v1/hangman_api.proto",
}
