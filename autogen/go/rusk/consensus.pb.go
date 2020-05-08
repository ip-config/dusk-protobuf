// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consensus.proto

package rusk

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Call during block generation
type DistributeTransaction struct {
	TotalReward           uint64       `protobuf:"fixed64,1,opt,name=total_reward,json=totalReward,proto3" json:"total_reward,omitempty"`
	ProvisionersAddresses []byte       `protobuf:"bytes,2,opt,name=provisioners_addresses,json=provisionersAddresses,proto3" json:"provisioners_addresses,omitempty"`
	BgPk                  *PublicKey   `protobuf:"bytes,3,opt,name=bg_pk,json=bgPk,proto3" json:"bg_pk,omitempty"`
	Tx                    *Transaction `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}     `json:"-"`
	XXX_unrecognized      []byte       `json:"-"`
	XXX_sizecache         int32        `json:"-"`
}

func (m *DistributeTransaction) Reset()         { *m = DistributeTransaction{} }
func (m *DistributeTransaction) String() string { return proto.CompactTextString(m) }
func (*DistributeTransaction) ProtoMessage()    {}
func (*DistributeTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{0}
}
func (m *DistributeTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeTransaction.Unmarshal(m, b)
}
func (m *DistributeTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeTransaction.Marshal(b, m, deterministic)
}
func (m *DistributeTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeTransaction.Merge(m, src)
}
func (m *DistributeTransaction) XXX_Size() int {
	return xxx_messageInfo_DistributeTransaction.Size(m)
}
func (m *DistributeTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeTransaction proto.InternalMessageInfo

func (m *DistributeTransaction) GetTotalReward() uint64 {
	if m != nil {
		return m.TotalReward
	}
	return 0
}

func (m *DistributeTransaction) GetProvisionersAddresses() []byte {
	if m != nil {
		return m.ProvisionersAddresses
	}
	return nil
}

func (m *DistributeTransaction) GetBgPk() *PublicKey {
	if m != nil {
		return m.BgPk
	}
	return nil
}

func (m *DistributeTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// WithdrawFeesTransaction is the transaction for a Provisioner to withdraw the fees using the smart contract
type WithdrawFeesTransaction struct {
	BlsKey               []byte       `protobuf:"bytes,1,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	Sig                  []byte       `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	Msg                  []byte       `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WithdrawFeesTransaction) Reset()         { *m = WithdrawFeesTransaction{} }
func (m *WithdrawFeesTransaction) String() string { return proto.CompactTextString(m) }
func (*WithdrawFeesTransaction) ProtoMessage()    {}
func (*WithdrawFeesTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{1}
}
func (m *WithdrawFeesTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawFeesTransaction.Unmarshal(m, b)
}
func (m *WithdrawFeesTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawFeesTransaction.Marshal(b, m, deterministic)
}
func (m *WithdrawFeesTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawFeesTransaction.Merge(m, src)
}
func (m *WithdrawFeesTransaction) XXX_Size() int {
	return xxx_messageInfo_WithdrawFeesTransaction.Size(m)
}
func (m *WithdrawFeesTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawFeesTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawFeesTransaction proto.InternalMessageInfo

func (m *WithdrawFeesTransaction) GetBlsKey() []byte {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *WithdrawFeesTransaction) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *WithdrawFeesTransaction) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *WithdrawFeesTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// Call from consensus
type SlashTransaction struct {
	BlsKey               []byte       `protobuf:"bytes,1,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	Step                 uint32       `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	Round                uint64       `protobuf:"fixed64,3,opt,name=round,proto3" json:"round,omitempty"`
	FirstMsg             []byte       `protobuf:"bytes,4,opt,name=first_msg,json=firstMsg,proto3" json:"first_msg,omitempty"`
	FirstSig             []byte       `protobuf:"bytes,5,opt,name=first_sig,json=firstSig,proto3" json:"first_sig,omitempty"`
	SecondMsg            []byte       `protobuf:"bytes,6,opt,name=second_msg,json=secondMsg,proto3" json:"second_msg,omitempty"`
	SecondSig            []byte       `protobuf:"bytes,7,opt,name=second_sig,json=secondSig,proto3" json:"second_sig,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,8,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SlashTransaction) Reset()         { *m = SlashTransaction{} }
func (m *SlashTransaction) String() string { return proto.CompactTextString(m) }
func (*SlashTransaction) ProtoMessage()    {}
func (*SlashTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{2}
}
func (m *SlashTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlashTransaction.Unmarshal(m, b)
}
func (m *SlashTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlashTransaction.Marshal(b, m, deterministic)
}
func (m *SlashTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlashTransaction.Merge(m, src)
}
func (m *SlashTransaction) XXX_Size() int {
	return xxx_messageInfo_SlashTransaction.Size(m)
}
func (m *SlashTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SlashTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SlashTransaction proto.InternalMessageInfo

func (m *SlashTransaction) GetBlsKey() []byte {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *SlashTransaction) GetStep() uint32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *SlashTransaction) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *SlashTransaction) GetFirstMsg() []byte {
	if m != nil {
		return m.FirstMsg
	}
	return nil
}

func (m *SlashTransaction) GetFirstSig() []byte {
	if m != nil {
		return m.FirstSig
	}
	return nil
}

func (m *SlashTransaction) GetSecondMsg() []byte {
	if m != nil {
		return m.SecondMsg
	}
	return nil
}

func (m *SlashTransaction) GetSecondSig() []byte {
	if m != nil {
		return m.SecondSig
	}
	return nil
}

func (m *SlashTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// StakeTransaction is the transaction for the Stake used by the Provisioners to engage in committees
type StakeTransaction struct {
	BlsKey               []byte       `protobuf:"bytes,1,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	Value                uint64       `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	ExpirationHeight     uint64       `protobuf:"fixed64,3,opt,name=expiration_height,json=expirationHeight,proto3" json:"expiration_height,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StakeTransaction) Reset()         { *m = StakeTransaction{} }
func (m *StakeTransaction) String() string { return proto.CompactTextString(m) }
func (*StakeTransaction) ProtoMessage()    {}
func (*StakeTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{3}
}
func (m *StakeTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeTransaction.Unmarshal(m, b)
}
func (m *StakeTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeTransaction.Marshal(b, m, deterministic)
}
func (m *StakeTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeTransaction.Merge(m, src)
}
func (m *StakeTransaction) XXX_Size() int {
	return xxx_messageInfo_StakeTransaction.Size(m)
}
func (m *StakeTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_StakeTransaction proto.InternalMessageInfo

func (m *StakeTransaction) GetBlsKey() []byte {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *StakeTransaction) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *StakeTransaction) GetExpirationHeight() uint64 {
	if m != nil {
		return m.ExpirationHeight
	}
	return 0
}

func (m *StakeTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// BidTransaction is the transaction for the Bid used by the Block Generator to create a Score
type BidTransaction struct {
	M                    []byte       `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
	Commitment           []byte       `protobuf:"bytes,2,opt,name=commitment,proto3" json:"commitment,omitempty"`
	EncryptedValue       []byte       `protobuf:"bytes,3,opt,name=encrypted_value,json=encryptedValue,proto3" json:"encrypted_value,omitempty"`
	EncryptedBlinder     []byte       `protobuf:"bytes,4,opt,name=encrypted_blinder,json=encryptedBlinder,proto3" json:"encrypted_blinder,omitempty"`
	ExpirationHeight     uint64       `protobuf:"fixed64,5,opt,name=expiration_height,json=expirationHeight,proto3" json:"expiration_height,omitempty"`
	Pk                   []byte       `protobuf:"bytes,6,opt,name=pk,proto3" json:"pk,omitempty"`
	R                    []byte       `protobuf:"bytes,7,opt,name=r,proto3" json:"r,omitempty"`
	Seed                 []byte       `protobuf:"bytes,8,opt,name=seed,proto3" json:"seed,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,9,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BidTransaction) Reset()         { *m = BidTransaction{} }
func (m *BidTransaction) String() string { return proto.CompactTextString(m) }
func (*BidTransaction) ProtoMessage()    {}
func (*BidTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{4}
}
func (m *BidTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidTransaction.Unmarshal(m, b)
}
func (m *BidTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidTransaction.Marshal(b, m, deterministic)
}
func (m *BidTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidTransaction.Merge(m, src)
}
func (m *BidTransaction) XXX_Size() int {
	return xxx_messageInfo_BidTransaction.Size(m)
}
func (m *BidTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_BidTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_BidTransaction proto.InternalMessageInfo

func (m *BidTransaction) GetM() []byte {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *BidTransaction) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *BidTransaction) GetEncryptedValue() []byte {
	if m != nil {
		return m.EncryptedValue
	}
	return nil
}

func (m *BidTransaction) GetEncryptedBlinder() []byte {
	if m != nil {
		return m.EncryptedBlinder
	}
	return nil
}

func (m *BidTransaction) GetExpirationHeight() uint64 {
	if m != nil {
		return m.ExpirationHeight
	}
	return 0
}

func (m *BidTransaction) GetPk() []byte {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (m *BidTransaction) GetR() []byte {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *BidTransaction) GetSeed() []byte {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *BidTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// WithdrawBidTransaction is the transaction to withdraw a bid
type WithdrawBidTransaction struct {
	Commitment           []byte       `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	EncryptedValue       []byte       `protobuf:"bytes,2,opt,name=encrypted_value,json=encryptedValue,proto3" json:"encrypted_value,omitempty"`
	EncryptedBlinder     []byte       `protobuf:"bytes,3,opt,name=encrypted_blinder,json=encryptedBlinder,proto3" json:"encrypted_blinder,omitempty"`
	Bid                  []byte       `protobuf:"bytes,4,opt,name=bid,proto3" json:"bid,omitempty"`
	Sig                  []byte       `protobuf:"bytes,5,opt,name=sig,proto3" json:"sig,omitempty"`
	EdPk                 []byte       `protobuf:"bytes,6,opt,name=ed_pk,json=edPk,proto3" json:"ed_pk,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,7,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WithdrawBidTransaction) Reset()         { *m = WithdrawBidTransaction{} }
func (m *WithdrawBidTransaction) String() string { return proto.CompactTextString(m) }
func (*WithdrawBidTransaction) ProtoMessage()    {}
func (*WithdrawBidTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{5}
}
func (m *WithdrawBidTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawBidTransaction.Unmarshal(m, b)
}
func (m *WithdrawBidTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawBidTransaction.Marshal(b, m, deterministic)
}
func (m *WithdrawBidTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawBidTransaction.Merge(m, src)
}
func (m *WithdrawBidTransaction) XXX_Size() int {
	return xxx_messageInfo_WithdrawBidTransaction.Size(m)
}
func (m *WithdrawBidTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawBidTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawBidTransaction proto.InternalMessageInfo

func (m *WithdrawBidTransaction) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *WithdrawBidTransaction) GetEncryptedValue() []byte {
	if m != nil {
		return m.EncryptedValue
	}
	return nil
}

func (m *WithdrawBidTransaction) GetEncryptedBlinder() []byte {
	if m != nil {
		return m.EncryptedBlinder
	}
	return nil
}

func (m *WithdrawBidTransaction) GetBid() []byte {
	if m != nil {
		return m.Bid
	}
	return nil
}

func (m *WithdrawBidTransaction) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *WithdrawBidTransaction) GetEdPk() []byte {
	if m != nil {
		return m.EdPk
	}
	return nil
}

func (m *WithdrawBidTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// WithdrawStakeTransaction is the transaction to withdraw a stake
type WithdrawStakeTransaction struct {
	BlsKey               []byte       `protobuf:"bytes,1,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	Sig                  []byte       `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WithdrawStakeTransaction) Reset()         { *m = WithdrawStakeTransaction{} }
func (m *WithdrawStakeTransaction) String() string { return proto.CompactTextString(m) }
func (*WithdrawStakeTransaction) ProtoMessage()    {}
func (*WithdrawStakeTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{6}
}
func (m *WithdrawStakeTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawStakeTransaction.Unmarshal(m, b)
}
func (m *WithdrawStakeTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawStakeTransaction.Marshal(b, m, deterministic)
}
func (m *WithdrawStakeTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawStakeTransaction.Merge(m, src)
}
func (m *WithdrawStakeTransaction) XXX_Size() int {
	return xxx_messageInfo_WithdrawStakeTransaction.Size(m)
}
func (m *WithdrawStakeTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawStakeTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawStakeTransaction proto.InternalMessageInfo

func (m *WithdrawStakeTransaction) GetBlsKey() []byte {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *WithdrawStakeTransaction) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *WithdrawStakeTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterType((*DistributeTransaction)(nil), "rusk.DistributeTransaction")
	proto.RegisterType((*WithdrawFeesTransaction)(nil), "rusk.WithdrawFeesTransaction")
	proto.RegisterType((*SlashTransaction)(nil), "rusk.SlashTransaction")
	proto.RegisterType((*StakeTransaction)(nil), "rusk.StakeTransaction")
	proto.RegisterType((*BidTransaction)(nil), "rusk.BidTransaction")
	proto.RegisterType((*WithdrawBidTransaction)(nil), "rusk.WithdrawBidTransaction")
	proto.RegisterType((*WithdrawStakeTransaction)(nil), "rusk.WithdrawStakeTransaction")
}

func init() { proto.RegisterFile("consensus.proto", fileDescriptor_56f0f2c53b3de771) }

var fileDescriptor_56f0f2c53b3de771 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x26, 0xd9, 0xcd, 0xb6, 0x3d, 0x5d, 0xdb, 0xed, 0xd8, 0x9f, 0xa5, 0xa2, 0xb4, 0x45, 0x70,
	0x51, 0xdc, 0x05, 0xc5, 0x07, 0xb0, 0x88, 0x08, 0x45, 0x28, 0xa9, 0x28, 0x78, 0x13, 0x92, 0xcc,
	0x69, 0x32, 0xe4, 0x67, 0xc2, 0xcc, 0xa4, 0x6d, 0x1e, 0xc3, 0x1b, 0x1f, 0xc6, 0x47, 0xf2, 0xca,
	0x47, 0x90, 0x99, 0x64, 0xb3, 0xf1, 0xa7, 0x4b, 0xee, 0x4e, 0xce, 0x77, 0x4e, 0xce, 0xf7, 0x7d,
	0x67, 0x66, 0x60, 0x37, 0xe4, 0xb9, 0xc4, 0x5c, 0x96, 0x72, 0x5e, 0x08, 0xae, 0x38, 0x19, 0x8a,
	0x52, 0x26, 0xc7, 0x7b, 0x4a, 0xf8, 0xb9, 0xf4, 0x43, 0xc5, 0x78, 0x5e, 0x03, 0xc7, 0x90, 0x60,
	0xd5, 0x14, 0x9d, 0xfd, 0xb0, 0xe0, 0xe0, 0x1d, 0x93, 0x4a, 0xb0, 0xa0, 0x54, 0xf8, 0x69, 0x55,
	0x4b, 0x4e, 0x61, 0xac, 0xb8, 0xf2, 0x53, 0x4f, 0xe0, 0xad, 0x2f, 0xe8, 0xd4, 0x3a, 0xb1, 0x66,
	0x23, 0x77, 0xdb, 0xe4, 0x5c, 0x93, 0x22, 0x6f, 0xe0, 0xb0, 0x10, 0xfc, 0x86, 0x49, 0xc6, 0x73,
	0x14, 0xd2, 0xf3, 0x29, 0x15, 0x28, 0x25, 0xca, 0xa9, 0x7d, 0x62, 0xcd, 0xc6, 0xee, 0x41, 0x17,
	0x7d, 0xbb, 0x04, 0xc9, 0x53, 0x70, 0x82, 0xc8, 0x2b, 0x92, 0xe9, 0xe0, 0xc4, 0x9a, 0x6d, 0xbf,
	0xda, 0x9d, 0x6b, 0xa2, 0xf3, 0xcb, 0x32, 0x48, 0x59, 0x78, 0x81, 0x95, 0x3b, 0x0c, 0xa2, 0xcb,
	0x84, 0x9c, 0x82, 0xad, 0xee, 0xa6, 0x43, 0x53, 0xb2, 0x57, 0x97, 0x74, 0xe8, 0xb9, 0xb6, 0xba,
	0x3b, 0xab, 0xe0, 0xe8, 0x0b, 0x53, 0x31, 0x15, 0xfe, 0xed, 0x7b, 0x44, 0xd9, 0x65, 0x7f, 0x04,
	0x1b, 0x41, 0x2a, 0xbd, 0x04, 0x2b, 0x43, 0x7c, 0xec, 0x8e, 0x82, 0x54, 0x5e, 0x60, 0x45, 0x26,
	0x30, 0x90, 0x2c, 0x6a, 0x08, 0xea, 0x50, 0x67, 0x32, 0x19, 0x19, 0x32, 0x63, 0x57, 0x87, 0x7d,
	0x46, 0xff, 0xb2, 0x60, 0x72, 0x95, 0xfa, 0x32, 0xee, 0x35, 0x94, 0xc0, 0x50, 0x2a, 0x2c, 0xcc,
	0xd4, 0x07, 0xae, 0x89, 0xc9, 0x3e, 0x38, 0x82, 0x97, 0x39, 0x35, 0x83, 0x47, 0x6e, 0xfd, 0x41,
	0x1e, 0xc1, 0xd6, 0x35, 0x13, 0x52, 0x79, 0x9a, 0xd2, 0xd0, 0xfc, 0x64, 0xd3, 0x24, 0x3e, 0xca,
	0x68, 0x05, 0x6a, 0x05, 0x4e, 0x07, 0xbc, 0x62, 0x11, 0x79, 0x0c, 0x20, 0x31, 0xe4, 0x39, 0x35,
	0xad, 0x23, 0x83, 0x6e, 0xd5, 0x19, 0xdd, 0xbb, 0x82, 0x75, 0xf3, 0x46, 0x17, 0xd6, 0xdd, 0xb5,
	0xe4, 0xcd, 0x75, 0x92, 0xbf, 0x69, 0xc9, 0xca, 0x4f, 0xb0, 0x97, 0xe4, 0x7d, 0x70, 0x6e, 0xfc,
	0xb4, 0x44, 0xa3, 0x79, 0xe4, 0xd6, 0x1f, 0xe4, 0x05, 0xec, 0xe1, 0x5d, 0xc1, 0x84, 0xaf, 0x9b,
	0xbd, 0x18, 0x59, 0x14, 0xab, 0xc6, 0x80, 0xc9, 0x0a, 0xf8, 0x60, 0xf2, 0x7d, 0xd6, 0xf0, 0xdd,
	0x86, 0x9d, 0x73, 0x46, 0xbb, 0x8c, 0xc6, 0x60, 0x65, 0x0d, 0x17, 0x2b, 0x23, 0x4f, 0x00, 0x42,
	0x9e, 0x65, 0x4c, 0x65, 0x98, 0xab, 0x66, 0xeb, 0x9d, 0x0c, 0x79, 0x06, 0xbb, 0x98, 0x87, 0xa2,
	0x2a, 0x14, 0x52, 0xaf, 0x26, 0x5c, 0x1f, 0x84, 0x9d, 0x36, 0xfd, 0xb9, 0x65, 0xde, 0x16, 0x06,
	0x29, 0xcb, 0x29, 0x8a, 0x66, 0x41, 0x93, 0x16, 0x38, 0xaf, 0xf3, 0xff, 0x97, 0xe9, 0xdc, 0x23,
	0x73, 0x07, 0xec, 0x22, 0x69, 0x16, 0x66, 0x17, 0x89, 0x16, 0x20, 0x9a, 0x05, 0x59, 0xc2, 0x1c,
	0x1d, 0x44, 0x6a, 0x56, 0x33, 0x76, 0x4d, 0xdc, 0x18, 0xb3, 0xb5, 0xce, 0x98, 0x9f, 0x16, 0x1c,
	0x2e, 0xef, 0xc6, 0x5f, 0x06, 0xfd, 0x69, 0x89, 0xd5, 0xc7, 0x12, 0xbb, 0xbf, 0x25, 0x83, 0x7b,
	0x2c, 0x99, 0xc0, 0x20, 0x60, 0xb4, 0x71, 0x4c, 0x87, 0xcb, 0x9b, 0xe8, 0xac, 0x6e, 0xe2, 0x43,
	0x70, 0x90, 0x7a, 0xad, 0x19, 0x43, 0xa4, 0xed, 0x3b, 0xb0, 0xb1, 0x4e, 0x6c, 0x0c, 0xd3, 0xa5,
	0xd6, 0xfe, 0x07, 0xf4, 0xdf, 0x87, 0xa0, 0x9e, 0x34, 0x58, 0x33, 0xe9, 0xfc, 0xf9, 0xd7, 0x59,
	0xc4, 0x54, 0x5c, 0x06, 0xf3, 0x90, 0x67, 0x0b, 0x5a, 0xca, 0xe4, 0xa5, 0x79, 0x47, 0x83, 0xf2,
	0x7a, 0xe1, 0x97, 0x8a, 0x47, 0x98, 0x2f, 0x22, 0xbe, 0xd0, 0xbd, 0xc1, 0xc8, 0x20, 0xaf, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xdb, 0x80, 0x7e, 0x99, 0x05, 0x00, 0x00,
}
