// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: club_membership_service.proto

package membership

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	v1 "github.com/superp00t/gophercraft/bnet/bgs/protocol/account/v1"
	v11 "github.com/superp00t/gophercraft/bnet/bgs/protocol/club/v1"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SubscribeRequest struct {
	AgentId              *v1.AccountId                       `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Options              *v11.ClubMembershipSubscribeOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{0}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubscribeRequest) GetOptions() *v11.ClubMembershipSubscribeOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type SubscribeResponse struct {
	State                *v11.ClubMembershipState `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{1}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetState() *v11.ClubMembershipState {
	if m != nil {
		return m.State
	}
	return nil
}

type UnsubscribeRequest struct {
	AgentId              *v1.AccountId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UnsubscribeRequest) Reset()         { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()    {}
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{2}
}

func (m *UnsubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsubscribeRequest.Unmarshal(m, b)
}
func (m *UnsubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsubscribeRequest.Marshal(b, m, deterministic)
}
func (m *UnsubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsubscribeRequest.Merge(m, src)
}
func (m *UnsubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_UnsubscribeRequest.Size(m)
}
func (m *UnsubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnsubscribeRequest proto.InternalMessageInfo

func (m *UnsubscribeRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

type GetStateRequest struct {
	AgentId              *v1.AccountId                      `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Options              *v11.ClubMembershipGetStateOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GetStateRequest) Reset()         { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()    {}
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{3}
}

func (m *GetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateRequest.Unmarshal(m, b)
}
func (m *GetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateRequest.Marshal(b, m, deterministic)
}
func (m *GetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateRequest.Merge(m, src)
}
func (m *GetStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetStateRequest.Size(m)
}
func (m *GetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateRequest proto.InternalMessageInfo

func (m *GetStateRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *GetStateRequest) GetOptions() *v11.ClubMembershipGetStateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetStateResponse struct {
	State                *v11.ClubMembershipState `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetStateResponse) Reset()         { *m = GetStateResponse{} }
func (m *GetStateResponse) String() string { return proto.CompactTextString(m) }
func (*GetStateResponse) ProtoMessage()    {}
func (*GetStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{4}
}

func (m *GetStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateResponse.Unmarshal(m, b)
}
func (m *GetStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateResponse.Marshal(b, m, deterministic)
}
func (m *GetStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateResponse.Merge(m, src)
}
func (m *GetStateResponse) XXX_Size() int {
	return xxx_messageInfo_GetStateResponse.Size(m)
}
func (m *GetStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateResponse proto.InternalMessageInfo

func (m *GetStateResponse) GetState() *v11.ClubMembershipState {
	if m != nil {
		return m.State
	}
	return nil
}

type UpdateClubSharedSettingsRequest struct {
	AgentId              *v1.AccountId                  `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Options              *v11.ClubSharedSettingsOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UpdateClubSharedSettingsRequest) Reset()         { *m = UpdateClubSharedSettingsRequest{} }
func (m *UpdateClubSharedSettingsRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateClubSharedSettingsRequest) ProtoMessage()    {}
func (*UpdateClubSharedSettingsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{5}
}

func (m *UpdateClubSharedSettingsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateClubSharedSettingsRequest.Unmarshal(m, b)
}
func (m *UpdateClubSharedSettingsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateClubSharedSettingsRequest.Marshal(b, m, deterministic)
}
func (m *UpdateClubSharedSettingsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateClubSharedSettingsRequest.Merge(m, src)
}
func (m *UpdateClubSharedSettingsRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateClubSharedSettingsRequest.Size(m)
}
func (m *UpdateClubSharedSettingsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateClubSharedSettingsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateClubSharedSettingsRequest proto.InternalMessageInfo

func (m *UpdateClubSharedSettingsRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateClubSharedSettingsRequest) GetOptions() *v11.ClubSharedSettingsOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetStreamMentionsRequest struct {
	AgentId              *v1.AccountId             `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Options              *protocol.GetEventOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	FetchMessages        *bool                     `protobuf:"varint,3,opt,name=fetch_messages,json=fetchMessages" json:"fetch_messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetStreamMentionsRequest) Reset()         { *m = GetStreamMentionsRequest{} }
func (m *GetStreamMentionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStreamMentionsRequest) ProtoMessage()    {}
func (*GetStreamMentionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{6}
}

func (m *GetStreamMentionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStreamMentionsRequest.Unmarshal(m, b)
}
func (m *GetStreamMentionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStreamMentionsRequest.Marshal(b, m, deterministic)
}
func (m *GetStreamMentionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStreamMentionsRequest.Merge(m, src)
}
func (m *GetStreamMentionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStreamMentionsRequest.Size(m)
}
func (m *GetStreamMentionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStreamMentionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStreamMentionsRequest proto.InternalMessageInfo

func (m *GetStreamMentionsRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *GetStreamMentionsRequest) GetOptions() *protocol.GetEventOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *GetStreamMentionsRequest) GetFetchMessages() bool {
	if m != nil && m.FetchMessages != nil {
		return *m.FetchMessages
	}
	return false
}

type GetStreamMentionsResponse struct {
	Mention              []*v11.StreamMention `protobuf:"bytes,1,rep,name=mention" json:"mention,omitempty"`
	Continuation         *uint64              `protobuf:"varint,2,opt,name=continuation" json:"continuation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetStreamMentionsResponse) Reset()         { *m = GetStreamMentionsResponse{} }
func (m *GetStreamMentionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStreamMentionsResponse) ProtoMessage()    {}
func (*GetStreamMentionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{7}
}

func (m *GetStreamMentionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStreamMentionsResponse.Unmarshal(m, b)
}
func (m *GetStreamMentionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStreamMentionsResponse.Marshal(b, m, deterministic)
}
func (m *GetStreamMentionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStreamMentionsResponse.Merge(m, src)
}
func (m *GetStreamMentionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStreamMentionsResponse.Size(m)
}
func (m *GetStreamMentionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStreamMentionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStreamMentionsResponse proto.InternalMessageInfo

func (m *GetStreamMentionsResponse) GetMention() []*v11.StreamMention {
	if m != nil {
		return m.Mention
	}
	return nil
}

func (m *GetStreamMentionsResponse) GetContinuation() uint64 {
	if m != nil && m.Continuation != nil {
		return *m.Continuation
	}
	return 0
}

type RemoveStreamMentionsRequest struct {
	AgentId              *v1.AccountId            `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MentionId            []*protocol.TimeSeriesId `protobuf:"bytes,2,rep,name=mention_id,json=mentionId" json:"mention_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RemoveStreamMentionsRequest) Reset()         { *m = RemoveStreamMentionsRequest{} }
func (m *RemoveStreamMentionsRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveStreamMentionsRequest) ProtoMessage()    {}
func (*RemoveStreamMentionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{8}
}

func (m *RemoveStreamMentionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveStreamMentionsRequest.Unmarshal(m, b)
}
func (m *RemoveStreamMentionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveStreamMentionsRequest.Marshal(b, m, deterministic)
}
func (m *RemoveStreamMentionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveStreamMentionsRequest.Merge(m, src)
}
func (m *RemoveStreamMentionsRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveStreamMentionsRequest.Size(m)
}
func (m *RemoveStreamMentionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveStreamMentionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveStreamMentionsRequest proto.InternalMessageInfo

func (m *RemoveStreamMentionsRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *RemoveStreamMentionsRequest) GetMentionId() []*protocol.TimeSeriesId {
	if m != nil {
		return m.MentionId
	}
	return nil
}

type AdvanceStreamMentionViewTimeRequest struct {
	AgentId              *v1.AccountId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AdvanceStreamMentionViewTimeRequest) Reset()         { *m = AdvanceStreamMentionViewTimeRequest{} }
func (m *AdvanceStreamMentionViewTimeRequest) String() string { return proto.CompactTextString(m) }
func (*AdvanceStreamMentionViewTimeRequest) ProtoMessage()    {}
func (*AdvanceStreamMentionViewTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70572c63cf771381, []int{9}
}

func (m *AdvanceStreamMentionViewTimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvanceStreamMentionViewTimeRequest.Unmarshal(m, b)
}
func (m *AdvanceStreamMentionViewTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvanceStreamMentionViewTimeRequest.Marshal(b, m, deterministic)
}
func (m *AdvanceStreamMentionViewTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvanceStreamMentionViewTimeRequest.Merge(m, src)
}
func (m *AdvanceStreamMentionViewTimeRequest) XXX_Size() int {
	return xxx_messageInfo_AdvanceStreamMentionViewTimeRequest.Size(m)
}
func (m *AdvanceStreamMentionViewTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvanceStreamMentionViewTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdvanceStreamMentionViewTimeRequest proto.InternalMessageInfo

func (m *AdvanceStreamMentionViewTimeRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "bgs.protocol.club.v1.membership.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "bgs.protocol.club.v1.membership.SubscribeResponse")
	proto.RegisterType((*UnsubscribeRequest)(nil), "bgs.protocol.club.v1.membership.UnsubscribeRequest")
	proto.RegisterType((*GetStateRequest)(nil), "bgs.protocol.club.v1.membership.GetStateRequest")
	proto.RegisterType((*GetStateResponse)(nil), "bgs.protocol.club.v1.membership.GetStateResponse")
	proto.RegisterType((*UpdateClubSharedSettingsRequest)(nil), "bgs.protocol.club.v1.membership.UpdateClubSharedSettingsRequest")
	proto.RegisterType((*GetStreamMentionsRequest)(nil), "bgs.protocol.club.v1.membership.GetStreamMentionsRequest")
	proto.RegisterType((*GetStreamMentionsResponse)(nil), "bgs.protocol.club.v1.membership.GetStreamMentionsResponse")
	proto.RegisterType((*RemoveStreamMentionsRequest)(nil), "bgs.protocol.club.v1.membership.RemoveStreamMentionsRequest")
	proto.RegisterType((*AdvanceStreamMentionViewTimeRequest)(nil), "bgs.protocol.club.v1.membership.AdvanceStreamMentionViewTimeRequest")
}

func init() { proto.RegisterFile("club_membership_service.proto", fileDescriptor_70572c63cf771381) }

var fileDescriptor_70572c63cf771381 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4d, 0x6b, 0x13, 0x6d,
	0x14, 0x7d, 0x9f, 0xf4, 0x23, 0xe9, 0xed, 0xfb, 0xbe, 0x4d, 0x1e, 0x2a, 0xc4, 0xd1, 0xd2, 0x32,
	0x45, 0xa8, 0x06, 0x67, 0x9a, 0x54, 0x90, 0x8a, 0x45, 0xab, 0x2d, 0x25, 0x42, 0x2a, 0x4c, 0x5a,
	0x17, 0x6e, 0xe2, 0x7c, 0xdc, 0x26, 0x03, 0x9d, 0x0f, 0xe7, 0x79, 0x26, 0xc5, 0x2e, 0x44, 0xb2,
	0x74, 0xe9, 0xc6, 0xa5, 0xb8, 0xf6, 0x17, 0xf8, 0x1f, 0xfc, 0x43, 0xd9, 0xc9, 0x7c, 0xa4, 0xed,
	0x4c, 0xd2, 0x4c, 0x2a, 0x71, 0x13, 0x86, 0x9b, 0x7b, 0xce, 0x3d, 0xf7, 0x0c, 0xf7, 0x24, 0xb0,
	0xa2, 0x9f, 0xfa, 0x5a, 0xcb, 0x42, 0x4b, 0x43, 0x8f, 0x75, 0x4c, 0xb7, 0xc5, 0xd0, 0xeb, 0x9a,
	0x3a, 0x4a, 0xae, 0xe7, 0x70, 0x87, 0xae, 0x6a, 0x6d, 0x16, 0x3d, 0xea, 0xce, 0xa9, 0x14, 0xf4,
	0x4a, 0xdd, 0xaa, 0x74, 0xd9, 0x2e, 0x14, 0x43, 0x3c, 0xff, 0xe0, 0x62, 0xdc, 0x27, 0x94, 0xc2,
	0x0a, 0xe3, 0x1e, 0xaa, 0x56, 0x54, 0x12, 0xbf, 0x13, 0x28, 0x36, 0x7d, 0x8d, 0xe9, 0x9e, 0xa9,
	0xa1, 0x82, 0xef, 0x7d, 0x64, 0x9c, 0xee, 0x40, 0x41, 0x6d, 0xa3, 0xcd, 0x5b, 0xa6, 0x51, 0x26,
	0x6b, 0x64, 0x63, 0xb1, 0x26, 0x4a, 0x89, 0x69, 0xaa, 0xae, 0x3b, 0xbe, 0xcd, 0x83, 0x81, 0xbb,
	0xd1, 0x63, 0xdd, 0x50, 0xf2, 0x21, 0xa6, 0x6e, 0xd0, 0x43, 0xc8, 0x3b, 0x2e, 0x37, 0x1d, 0x9b,
	0x95, 0x73, 0x21, 0xfa, 0x91, 0x34, 0x52, 0xeb, 0xcb, 0x53, 0x5f, 0x6b, 0x5c, 0xe8, 0xbd, 0x50,
	0xf1, 0x3a, 0xc2, 0x2a, 0x03, 0x12, 0xf1, 0x08, 0x4a, 0x57, 0x24, 0x32, 0xd7, 0xb1, 0x19, 0xd2,
	0x67, 0x30, 0xc7, 0xb8, 0xca, 0x31, 0x16, 0x78, 0x7f, 0xa2, 0x11, 0x01, 0x40, 0x89, 0x70, 0x62,
	0x13, 0xe8, 0xb1, 0xcd, 0xa6, 0xbb, 0xba, 0xf8, 0x8d, 0xc0, 0xd2, 0x01, 0xf2, 0x68, 0xd0, 0x74,
	0xdc, 0x6c, 0xa4, 0xdd, 0xdc, 0x9a, 0x64, 0xd5, 0x81, 0x88, 0x21, 0x33, 0x9b, 0x50, 0xbc, 0x14,
	0x38, 0x2d, 0x2f, 0x7f, 0x10, 0x58, 0x3d, 0x76, 0x0d, 0x95, 0x63, 0xd0, 0xd4, 0xec, 0xa8, 0x1e,
	0x1a, 0x4d, 0xe4, 0xdc, 0xb4, 0xdb, 0x6c, 0x4a, 0x36, 0xd4, 0xd3, 0x36, 0xc8, 0xd7, 0xab, 0x4c,
	0x0a, 0x18, 0xb2, 0xe0, 0x27, 0x81, 0x72, 0xe8, 0x41, 0x70, 0x07, 0x0d, 0xb4, 0xa3, 0xaf, 0xa7,
	0x23, 0xf3, 0x71, 0x5a, 0xe6, 0x4a, 0x12, 0x7d, 0x80, 0x7c, 0xbf, 0x8b, 0x36, 0x4f, 0x8b, 0xa2,
	0xf7, 0xe0, 0xff, 0x13, 0xe4, 0x7a, 0xa7, 0x65, 0x21, 0x63, 0x6a, 0x1b, 0x59, 0x79, 0x66, 0x8d,
	0x6c, 0x14, 0x94, 0xff, 0xc2, 0x6a, 0x23, 0x2e, 0x8a, 0x1f, 0xe1, 0xf6, 0x08, 0xe9, 0xf1, 0x7b,
	0xdc, 0x81, 0xbc, 0x15, 0xd5, 0xca, 0x64, 0x6d, 0x66, 0x63, 0xb1, 0xb6, 0x3e, 0xda, 0xa3, 0x04,
	0x5c, 0x19, 0x60, 0xa8, 0x08, 0xff, 0xea, 0x8e, 0xcd, 0x4d, 0xdb, 0x57, 0x43, 0x8e, 0x60, 0x81,
	0x59, 0x25, 0x51, 0x13, 0xbf, 0x12, 0xb8, 0xa3, 0xa0, 0xe5, 0x74, 0xf1, 0xaf, 0xd8, 0xb7, 0x0d,
	0x10, 0xab, 0x09, 0x08, 0x72, 0xe1, 0x12, 0x42, 0x92, 0xe0, 0xc8, 0xb4, 0xb0, 0x89, 0x9e, 0x89,
	0xac, 0x6e, 0x28, 0x0b, 0x71, 0x77, 0xdd, 0x10, 0x0d, 0x58, 0xdf, 0x35, 0xba, 0xaa, 0xad, 0x27,
	0x95, 0xbd, 0x31, 0xf1, 0x2c, 0x80, 0x4c, 0x47, 0x60, 0xed, 0x57, 0x1e, 0x6e, 0xa5, 0x0e, 0x21,
	0x4a, 0x65, 0x7a, 0x0e, 0x0b, 0x17, 0x29, 0x45, 0xab, 0x52, 0x46, 0x3a, 0x4b, 0xe9, 0xd0, 0x15,
	0x6a, 0x37, 0x81, 0x44, 0x2f, 0x5c, 0x9c, 0xef, 0xf5, 0x2b, 0xb9, 0x02, 0xa1, 0xef, 0x60, 0xf1,
	0x4a, 0x96, 0xd1, 0xad, 0x4c, 0xaa, 0xe1, 0xe4, 0x13, 0x96, 0x93, 0xa0, 0x43, 0x67, 0x4f, 0xe5,
	0x6a, 0x3c, 0x21, 0x47, 0xcf, 0xa0, 0x30, 0x88, 0x0d, 0xba, 0x99, 0x49, 0x9f, 0x8a, 0x40, 0xa1,
	0x7a, 0x03, 0x44, 0x62, 0xb5, 0x19, 0x7a, 0x0e, 0xe5, 0xeb, 0x92, 0x85, 0x3e, 0xcf, 0xde, 0x73,
	0x7c, 0x28, 0x8d, 0x5d, 0x7a, 0x96, 0x7e, 0x21, 0x50, 0x1a, 0xba, 0x36, 0xba, 0x3d, 0xd9, 0x32,
	0x23, 0xae, 0x43, 0x78, 0xf2, 0x27, 0xd0, 0x84, 0x21, 0x73, 0xd4, 0x83, 0xe5, 0x51, 0x07, 0x48,
	0x9f, 0x66, 0x72, 0x8f, 0xb9, 0xdb, 0xb1, 0x46, 0xcc, 0xd3, 0x1e, 0x81, 0xbb, 0xe3, 0x8e, 0x8b,
	0xee, 0x65, 0x0e, 0x9f, 0xe0, 0x36, 0xc7, 0x8a, 0xc8, 0x0b, 0x4a, 0xaf, 0x5f, 0x29, 0x3d, 0x58,
	0x4a, 0xfd, 0x2b, 0xea, 0xf5, 0x2b, 0x0f, 0xa1, 0xa2, 0xd9, 0xc8, 0x33, 0x7f, 0xb1, 0xa2, 0x43,
	0xfd, 0xdc, 0xaf, 0xe4, 0xca, 0x24, 0xf8, 0x2c, 0x92, 0x17, 0xaf, 0xde, 0xee, 0xb7, 0x4d, 0xde,
	0xf1, 0x35, 0x49, 0x77, 0x2c, 0x99, 0xf9, 0x2e, 0x7a, 0xee, 0xe6, 0x26, 0x97, 0xdb, 0x8e, 0xdb,
	0x41, 0x4f, 0xf7, 0xd4, 0x13, 0x2e, 0x07, 0xdc, 0xb2, 0xd6, 0x66, 0xf2, 0x80, 0x5f, 0x0e, 0xf8,
	0xe5, 0x6e, 0x55, 0xbe, 0x54, 0xf1, 0x89, 0xfc, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x63,
	0x4f, 0x6a, 0xb0, 0x09, 0x00, 0x00,
}
