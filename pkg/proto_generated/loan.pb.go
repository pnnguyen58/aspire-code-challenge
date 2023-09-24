// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/proto/loan.proto

package protogen

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Loan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    uint64                 `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	RepaymentType string                 `protobuf:"bytes,3,opt,name=repayment_type,json=repaymentType,proto3" json:"repayment_type,omitempty"`
	Amount        float32                `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Term          int32                  `protobuf:"varint,5,opt,name=term,proto3" json:"term,omitempty"`
	State         string                 `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Repayments    []*Repayment           `protobuf:"bytes,9,rep,name=repayments,proto3" json:"repayments,omitempty"`
}

func (x *Loan) Reset() {
	*x = Loan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_loan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loan) ProtoMessage() {}

func (x *Loan) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_loan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loan.ProtoReflect.Descriptor instead.
func (*Loan) Descriptor() ([]byte, []int) {
	return file_pkg_proto_loan_proto_rawDescGZIP(), []int{0}
}

func (x *Loan) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Loan) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Loan) GetRepaymentType() string {
	if x != nil {
		return x.RepaymentType
	}
	return ""
}

func (x *Loan) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Loan) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Loan) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Loan) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Loan) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Loan) GetRepayments() []*Repayment {
	if x != nil {
		return x.Repayments
	}
	return nil
}

type LoanCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount     float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Term       int32   `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *LoanCreateRequest) Reset() {
	*x = LoanCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_loan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanCreateRequest) ProtoMessage() {}

func (x *LoanCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_loan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanCreateRequest.ProtoReflect.Descriptor instead.
func (*LoanCreateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_loan_proto_rawDescGZIP(), []int{1}
}

func (x *LoanCreateRequest) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *LoanCreateRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *LoanCreateRequest) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

type LoanCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *Loan  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Code  int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LoanCreateResponse) Reset() {
	*x = LoanCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_loan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanCreateResponse) ProtoMessage() {}

func (x *LoanCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_loan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanCreateResponse.ProtoReflect.Descriptor instead.
func (*LoanCreateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_loan_proto_rawDescGZIP(), []int{2}
}

func (x *LoanCreateResponse) GetData() *Loan {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoanCreateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LoanCreateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type LoanApproveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LoanApproveRequest) Reset() {
	*x = LoanApproveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_loan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanApproveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanApproveRequest) ProtoMessage() {}

func (x *LoanApproveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_loan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanApproveRequest.ProtoReflect.Descriptor instead.
func (*LoanApproveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_loan_proto_rawDescGZIP(), []int{3}
}

func (x *LoanApproveRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LoanApproveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  uint64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Code  int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LoanApproveResponse) Reset() {
	*x = LoanApproveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_loan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanApproveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanApproveResponse) ProtoMessage() {}

func (x *LoanApproveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_loan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanApproveResponse.ProtoReflect.Descriptor instead.
func (*LoanApproveResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_loan_proto_rawDescGZIP(), []int{4}
}

func (x *LoanApproveResponse) GetData() uint64 {
	if x != nil {
		return x.Data
	}
	return 0
}

func (x *LoanApproveResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LoanApproveResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type LoanGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LoanGetRequest) Reset() {
	*x = LoanGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_loan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanGetRequest) ProtoMessage() {}

func (x *LoanGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_loan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanGetRequest.ProtoReflect.Descriptor instead.
func (*LoanGetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_loan_proto_rawDescGZIP(), []int{5}
}

func (x *LoanGetRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LoanGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *Loan  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Code  int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LoanGetResponse) Reset() {
	*x = LoanGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_loan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanGetResponse) ProtoMessage() {}

func (x *LoanGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_loan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanGetResponse.ProtoReflect.Descriptor instead.
func (*LoanGetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_loan_proto_rawDescGZIP(), []int{6}
}

func (x *LoanGetResponse) GetData() *Loan {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *LoanGetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *LoanGetResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_pkg_proto_loan_proto protoreflect.FileDescriptor

var file_pkg_proto_loan_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x61, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x61, 0x6e, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x02, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x60, 0x0a, 0x11, 0x4c, 0x6f, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x22, 0x5e, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x4c, 0x6f,
	0x61, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x13, 0x4c, 0x6f, 0x61, 0x6e,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x20, 0x0a,
	0x0e, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5b, 0x0a, 0x0f, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x87, 0x03, 0x0a,
	0x0b, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x17, 0x2e, 0x6c, 0x6f, 0x61,
	0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x73,
	0x12, 0x62, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x12,
	0x18, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x6f, 0x61, 0x6e,
	0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x1a,
	0x13, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x12,
	0x14, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61,
	0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x73, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x6c, 0x6f, 0x61, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6e, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x35, 0x38, 0x2f,
	0x61, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_loan_proto_rawDescOnce sync.Once
	file_pkg_proto_loan_proto_rawDescData = file_pkg_proto_loan_proto_rawDesc
)

func file_pkg_proto_loan_proto_rawDescGZIP() []byte {
	file_pkg_proto_loan_proto_rawDescOnce.Do(func() {
		file_pkg_proto_loan_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_loan_proto_rawDescData)
	})
	return file_pkg_proto_loan_proto_rawDescData
}

var file_pkg_proto_loan_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_loan_proto_goTypes = []interface{}{
	(*Loan)(nil),                    // 0: loan.Loan
	(*LoanCreateRequest)(nil),       // 1: loan.LoanCreateRequest
	(*LoanCreateResponse)(nil),      // 2: loan.LoanCreateResponse
	(*LoanApproveRequest)(nil),      // 3: loan.LoanApproveRequest
	(*LoanApproveResponse)(nil),     // 4: loan.LoanApproveResponse
	(*LoanGetRequest)(nil),          // 5: loan.LoanGetRequest
	(*LoanGetResponse)(nil),         // 6: loan.LoanGetResponse
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
	(*Repayment)(nil),               // 8: loan.Repayment
	(*RepaymentCreateRequest)(nil),  // 9: loan.RepaymentCreateRequest
	(*RepaymentCreateResponse)(nil), // 10: loan.RepaymentCreateResponse
}
var file_pkg_proto_loan_proto_depIdxs = []int32{
	7,  // 0: loan.Loan.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: loan.Loan.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 2: loan.Loan.repayments:type_name -> loan.Repayment
	0,  // 3: loan.LoanCreateResponse.data:type_name -> loan.Loan
	0,  // 4: loan.LoanGetResponse.data:type_name -> loan.Loan
	1,  // 5: loan.LoanService.CreateLoan:input_type -> loan.LoanCreateRequest
	3,  // 6: loan.LoanService.ApproveLoan:input_type -> loan.LoanApproveRequest
	5,  // 7: loan.LoanService.GetLoan:input_type -> loan.LoanGetRequest
	9,  // 8: loan.LoanService.AddRepayment:input_type -> loan.RepaymentCreateRequest
	2,  // 9: loan.LoanService.CreateLoan:output_type -> loan.LoanCreateResponse
	4,  // 10: loan.LoanService.ApproveLoan:output_type -> loan.LoanApproveResponse
	6,  // 11: loan.LoanService.GetLoan:output_type -> loan.LoanGetResponse
	10, // 12: loan.LoanService.AddRepayment:output_type -> loan.RepaymentCreateResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_proto_loan_proto_init() }
func file_pkg_proto_loan_proto_init() {
	if File_pkg_proto_loan_proto != nil {
		return
	}
	file_pkg_proto_repayment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_loan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loan); i {
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
		file_pkg_proto_loan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanCreateRequest); i {
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
		file_pkg_proto_loan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanCreateResponse); i {
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
		file_pkg_proto_loan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanApproveRequest); i {
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
		file_pkg_proto_loan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanApproveResponse); i {
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
		file_pkg_proto_loan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanGetRequest); i {
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
		file_pkg_proto_loan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanGetResponse); i {
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
			RawDescriptor: file_pkg_proto_loan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_loan_proto_goTypes,
		DependencyIndexes: file_pkg_proto_loan_proto_depIdxs,
		MessageInfos:      file_pkg_proto_loan_proto_msgTypes,
	}.Build()
	File_pkg_proto_loan_proto = out.File
	file_pkg_proto_loan_proto_rawDesc = nil
	file_pkg_proto_loan_proto_goTypes = nil
	file_pkg_proto_loan_proto_depIdxs = nil
}
