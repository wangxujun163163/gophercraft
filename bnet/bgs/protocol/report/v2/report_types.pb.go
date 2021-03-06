// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: api/client/v2/report_types.proto

package v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	v1 "github.com/superp00t/gophercraft/bnet/bgs/protocol/account/v1"
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

type IssueType int32

const (
	IssueType_ISSUE_TYPE_SPAM              IssueType = 0
	IssueType_ISSUE_TYPE_HARASSMENT        IssueType = 1
	IssueType_ISSUE_TYPE_OFFENSIVE_CONTENT IssueType = 3
	IssueType_ISSUE_TYPE_HACKING           IssueType = 4
	IssueType_ISSUE_TYPE_BOTTING           IssueType = 5
)

var IssueType_name = map[int32]string{
	0: "ISSUE_TYPE_SPAM",
	1: "ISSUE_TYPE_HARASSMENT",
	3: "ISSUE_TYPE_OFFENSIVE_CONTENT",
	4: "ISSUE_TYPE_HACKING",
	5: "ISSUE_TYPE_BOTTING",
}

var IssueType_value = map[string]int32{
	"ISSUE_TYPE_SPAM":              0,
	"ISSUE_TYPE_HARASSMENT":        1,
	"ISSUE_TYPE_OFFENSIVE_CONTENT": 3,
	"ISSUE_TYPE_HACKING":           4,
	"ISSUE_TYPE_BOTTING":           5,
}

func (x IssueType) Enum() *IssueType {
	p := new(IssueType)
	*p = x
	return p
}

func (x IssueType) String() string {
	return proto.EnumName(IssueType_name, int32(x))
}

func (x *IssueType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IssueType_value, data, "IssueType")
	if err != nil {
		return err
	}
	*x = IssueType(value)
	return nil
}

func (IssueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_585b08f3e02b5193, []int{0}
}

type UserSource int32

const (
	UserSource_USER_SOURCE_OTHER             UserSource = 0
	UserSource_USER_SOURCE_WHISPER           UserSource = 1
	UserSource_USER_SOURCE_PROFILE           UserSource = 2
	UserSource_USER_SOURCE_BATTLE_TAG        UserSource = 3
	UserSource_USER_SOURCE_CHAT              UserSource = 4
	UserSource_USER_SOURCE_FRIEND_INVITATION UserSource = 5
	UserSource_USER_SOURCE_VOICE             UserSource = 6
)

var UserSource_name = map[int32]string{
	0: "USER_SOURCE_OTHER",
	1: "USER_SOURCE_WHISPER",
	2: "USER_SOURCE_PROFILE",
	3: "USER_SOURCE_BATTLE_TAG",
	4: "USER_SOURCE_CHAT",
	5: "USER_SOURCE_FRIEND_INVITATION",
	6: "USER_SOURCE_VOICE",
}

var UserSource_value = map[string]int32{
	"USER_SOURCE_OTHER":             0,
	"USER_SOURCE_WHISPER":           1,
	"USER_SOURCE_PROFILE":           2,
	"USER_SOURCE_BATTLE_TAG":        3,
	"USER_SOURCE_CHAT":              4,
	"USER_SOURCE_FRIEND_INVITATION": 5,
	"USER_SOURCE_VOICE":             6,
}

func (x UserSource) Enum() *UserSource {
	p := new(UserSource)
	*p = x
	return p
}

func (x UserSource) String() string {
	return proto.EnumName(UserSource_name, int32(x))
}

func (x *UserSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserSource_value, data, "UserSource")
	if err != nil {
		return err
	}
	*x = UserSource(value)
	return nil
}

func (UserSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_585b08f3e02b5193, []int{1}
}

type ClubSource int32

const (
	ClubSource_CLUB_SOURCE_OTHER       ClubSource = 0
	ClubSource_CLUB_SOURCE_MESSAGE     ClubSource = 1
	ClubSource_CLUB_SOURCE_CLUB_NAME   ClubSource = 2
	ClubSource_CLUB_SOURCE_STREAM_NAME ClubSource = 3
)

var ClubSource_name = map[int32]string{
	0: "CLUB_SOURCE_OTHER",
	1: "CLUB_SOURCE_MESSAGE",
	2: "CLUB_SOURCE_CLUB_NAME",
	3: "CLUB_SOURCE_STREAM_NAME",
}

var ClubSource_value = map[string]int32{
	"CLUB_SOURCE_OTHER":       0,
	"CLUB_SOURCE_MESSAGE":     1,
	"CLUB_SOURCE_CLUB_NAME":   2,
	"CLUB_SOURCE_STREAM_NAME": 3,
}

func (x ClubSource) Enum() *ClubSource {
	p := new(ClubSource)
	*p = x
	return p
}

func (x ClubSource) String() string {
	return proto.EnumName(ClubSource_name, int32(x))
}

func (x *ClubSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClubSource_value, data, "ClubSource")
	if err != nil {
		return err
	}
	*x = ClubSource(value)
	return nil
}

func (ClubSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_585b08f3e02b5193, []int{2}
}

type ReportItem struct {
	// Types that are valid to be assigned to Type:
	//	*ReportItem_MessageId
	Type                 isReportItem_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReportItem) Reset()         { *m = ReportItem{} }
func (m *ReportItem) String() string { return proto.CompactTextString(m) }
func (*ReportItem) ProtoMessage()    {}
func (*ReportItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_585b08f3e02b5193, []int{0}
}

func (m *ReportItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportItem.Unmarshal(m, b)
}
func (m *ReportItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportItem.Marshal(b, m, deterministic)
}
func (m *ReportItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportItem.Merge(m, src)
}
func (m *ReportItem) XXX_Size() int {
	return xxx_messageInfo_ReportItem.Size(m)
}
func (m *ReportItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportItem.DiscardUnknown(m)
}

var xxx_messageInfo_ReportItem proto.InternalMessageInfo

type isReportItem_Type interface {
	isReportItem_Type()
}

type ReportItem_MessageId struct {
	MessageId *protocol.MessageId `protobuf:"bytes,1,opt,name=message_id,json=messageId,oneof"`
}

func (*ReportItem_MessageId) isReportItem_Type() {}

func (m *ReportItem) GetType() isReportItem_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ReportItem) GetMessageId() *protocol.MessageId {
	if x, ok := m.GetType().(*ReportItem_MessageId); ok {
		return x.MessageId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReportItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReportItem_MessageId)(nil),
	}
}

type UserOptions struct {
	TargetId             *v1.AccountId `protobuf:"bytes,1,opt,name=target_id,json=targetId" json:"target_id,omitempty"`
	Type                 *IssueType    `protobuf:"varint,2,opt,name=type,enum=bgs.protocol.report.v2.IssueType" json:"type,omitempty"`
	Source               *UserSource   `protobuf:"varint,3,opt,name=source,enum=bgs.protocol.report.v2.UserSource" json:"source,omitempty"`
	Item                 *ReportItem   `protobuf:"bytes,4,opt,name=item" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserOptions) Reset()         { *m = UserOptions{} }
func (m *UserOptions) String() string { return proto.CompactTextString(m) }
func (*UserOptions) ProtoMessage()    {}
func (*UserOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_585b08f3e02b5193, []int{1}
}

func (m *UserOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOptions.Unmarshal(m, b)
}
func (m *UserOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOptions.Marshal(b, m, deterministic)
}
func (m *UserOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOptions.Merge(m, src)
}
func (m *UserOptions) XXX_Size() int {
	return xxx_messageInfo_UserOptions.Size(m)
}
func (m *UserOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOptions.DiscardUnknown(m)
}

var xxx_messageInfo_UserOptions proto.InternalMessageInfo

func (m *UserOptions) GetTargetId() *v1.AccountId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *UserOptions) GetType() IssueType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return IssueType_ISSUE_TYPE_SPAM
}

func (m *UserOptions) GetSource() UserSource {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return UserSource_USER_SOURCE_OTHER
}

func (m *UserOptions) GetItem() *ReportItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type ClubOptions struct {
	ClubId               *uint64     `protobuf:"varint,1,opt,name=club_id,json=clubId" json:"club_id,omitempty"`
	StreamId             *uint64     `protobuf:"varint,2,opt,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	Type                 *IssueType  `protobuf:"varint,3,opt,name=type,enum=bgs.protocol.report.v2.IssueType" json:"type,omitempty"`
	Source               *ClubSource `protobuf:"varint,4,opt,name=source,enum=bgs.protocol.report.v2.ClubSource" json:"source,omitempty"`
	Item                 *ReportItem `protobuf:"bytes,5,opt,name=item" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClubOptions) Reset()         { *m = ClubOptions{} }
func (m *ClubOptions) String() string { return proto.CompactTextString(m) }
func (*ClubOptions) ProtoMessage()    {}
func (*ClubOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_585b08f3e02b5193, []int{2}
}

func (m *ClubOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClubOptions.Unmarshal(m, b)
}
func (m *ClubOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClubOptions.Marshal(b, m, deterministic)
}
func (m *ClubOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClubOptions.Merge(m, src)
}
func (m *ClubOptions) XXX_Size() int {
	return xxx_messageInfo_ClubOptions.Size(m)
}
func (m *ClubOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ClubOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ClubOptions proto.InternalMessageInfo

func (m *ClubOptions) GetClubId() uint64 {
	if m != nil && m.ClubId != nil {
		return *m.ClubId
	}
	return 0
}

func (m *ClubOptions) GetStreamId() uint64 {
	if m != nil && m.StreamId != nil {
		return *m.StreamId
	}
	return 0
}

func (m *ClubOptions) GetType() IssueType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return IssueType_ISSUE_TYPE_SPAM
}

func (m *ClubOptions) GetSource() ClubSource {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ClubSource_CLUB_SOURCE_OTHER
}

func (m *ClubOptions) GetItem() *ReportItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterEnum("bgs.protocol.report.v2.IssueType", IssueType_name, IssueType_value)
	proto.RegisterEnum("bgs.protocol.report.v2.UserSource", UserSource_name, UserSource_value)
	proto.RegisterEnum("bgs.protocol.report.v2.ClubSource", ClubSource_name, ClubSource_value)
	proto.RegisterType((*ReportItem)(nil), "bgs.protocol.report.v2.ReportItem")
	proto.RegisterType((*UserOptions)(nil), "bgs.protocol.report.v2.UserOptions")
	proto.RegisterType((*ClubOptions)(nil), "bgs.protocol.report.v2.ClubOptions")
}

func init() { proto.RegisterFile("api/client/v2/report_types.proto", fileDescriptor_585b08f3e02b5193) }

var fileDescriptor_585b08f3e02b5193 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x6f, 0xd2, 0x40,
	0x1c, 0xa7, 0xa3, 0xc3, 0xf1, 0x5d, 0xe2, 0xea, 0xe1, 0x06, 0x6e, 0x9a, 0x30, 0x9e, 0x96, 0x3d,
	0xb4, 0x93, 0x44, 0x63, 0x8c, 0x89, 0x29, 0xf5, 0x18, 0x17, 0x47, 0x4b, 0xee, 0x8e, 0x19, 0x7d,
	0x69, 0xa0, 0x9c, 0x8c, 0x84, 0xd2, 0xa6, 0xbd, 0x92, 0xec, 0xcd, 0x7f, 0xc0, 0x3f, 0xca, 0x7f,
	0xc8, 0x77, 0xdf, 0x4c, 0x5b, 0x90, 0xb2, 0x6c, 0x26, 0xfa, 0x76, 0xf7, 0xf9, 0xf1, 0xe5, 0xf3,
	0xf9, 0x1e, 0x29, 0x34, 0x47, 0xe1, 0xcc, 0xf0, 0xe6, 0x33, 0xb1, 0x90, 0xc6, 0xb2, 0x6d, 0x44,
	0x22, 0x0c, 0x22, 0xe9, 0xca, 0xdb, 0x50, 0xc4, 0x7a, 0x18, 0x05, 0x32, 0x40, 0x47, 0xe3, 0xe9,
	0xea, 0xe8, 0x05, 0x73, 0x3d, 0x17, 0xe8, 0xcb, 0xf6, 0x71, 0x6d, 0xe4, 0x79, 0x41, 0xb2, 0xd8,
	0x12, 0x1f, 0x1f, 0x44, 0xa1, 0xb7, 0x05, 0xd4, 0x7c, 0x11, 0xc7, 0xa3, 0xa9, 0x28, 0x82, 0x2d,
	0x1b, 0x80, 0x66, 0x73, 0x88, 0x14, 0x3e, 0x7a, 0x03, 0xb0, 0x16, 0xcd, 0x26, 0x0d, 0xa5, 0xa9,
	0x9c, 0xed, 0xb7, 0xeb, 0xfa, 0xd6, 0xaf, 0xf6, 0x73, 0x9e, 0x4c, 0x7a, 0x25, 0x5a, 0xf5, 0xd7,
	0x97, 0x4e, 0x05, 0xd4, 0x74, 0x6c, 0xeb, 0x97, 0x02, 0xfb, 0xc3, 0x58, 0x44, 0x4e, 0x28, 0x67,
	0xc1, 0x22, 0x46, 0xef, 0xa1, 0x2a, 0x47, 0xd1, 0x54, 0xc8, 0xcd, 0xc0, 0xd6, 0xf6, 0xc0, 0x55,
	0x76, 0x7d, 0xf9, 0x52, 0x37, 0xf3, 0x23, 0x99, 0xd0, 0xbd, 0xdc, 0x44, 0x26, 0xe8, 0x55, 0x3e,
	0xb8, 0xb1, 0xd3, 0x54, 0xce, 0x1e, 0xb7, 0x4f, 0xf5, 0xfb, 0x57, 0xa0, 0x93, 0x38, 0x4e, 0x04,
	0xbf, 0x0d, 0x05, 0xcd, 0xe4, 0xe8, 0x2d, 0x54, 0xe2, 0x20, 0x89, 0x3c, 0xd1, 0x28, 0x67, 0xc6,
	0xd6, 0x43, 0xc6, 0x34, 0x2c, 0xcb, 0x94, 0x74, 0xe5, 0x40, 0xaf, 0x41, 0x9d, 0x49, 0xe1, 0x37,
	0xd4, 0xfb, 0xe2, 0x6e, 0x9c, 0x9b, 0xbd, 0xd1, 0x4c, 0xdf, 0xfa, 0xa9, 0xc0, 0xbe, 0x35, 0x4f,
	0xc6, 0xeb, 0xee, 0x75, 0x78, 0xe4, 0xcd, 0x93, 0xf1, 0xba, 0xb9, 0x4a, 0x2b, 0xe9, 0x95, 0x4c,
	0xd0, 0x09, 0x54, 0x63, 0x19, 0x89, 0x91, 0x9f, 0x52, 0x3b, 0x19, 0xb5, 0x97, 0x03, 0x85, 0xc2,
	0xe5, 0xff, 0x2d, 0xac, 0xfe, 0xbd, 0x70, 0x9a, 0xf0, 0x81, 0xc2, 0xbb, 0xff, 0x56, 0xf8, 0xfc,
	0xbb, 0x02, 0xd5, 0x3f, 0x39, 0x50, 0x0d, 0x0e, 0x08, 0x63, 0x43, 0xec, 0xf2, 0xcf, 0x03, 0xec,
	0xb2, 0x81, 0xd9, 0xd7, 0x4a, 0xe8, 0x19, 0x1c, 0x16, 0xc0, 0x9e, 0x49, 0x4d, 0xc6, 0xfa, 0xd8,
	0xe6, 0x9a, 0x82, 0x9a, 0xf0, 0xbc, 0x40, 0x39, 0xdd, 0x2e, 0xb6, 0x19, 0xb9, 0xc6, 0xae, 0xe5,
	0xd8, 0x3c, 0x55, 0x94, 0xd1, 0x11, 0xa0, 0x2d, 0xb3, 0xf5, 0x91, 0xd8, 0x97, 0x9a, 0x7a, 0x07,
	0xef, 0x38, 0x9c, 0xa7, 0xf8, 0xee, 0xf9, 0x0f, 0x05, 0x60, 0xf3, 0x9e, 0xe8, 0x10, 0x9e, 0x0c,
	0x19, 0xa6, 0x2e, 0x73, 0x86, 0xd4, 0xc2, 0xae, 0xc3, 0x7b, 0x98, 0x6a, 0x25, 0x54, 0x87, 0x5a,
	0x11, 0xfe, 0xd4, 0x23, 0x6c, 0x80, 0xa9, 0xa6, 0xdc, 0x25, 0x06, 0xd4, 0xe9, 0x92, 0x2b, 0xac,
	0xed, 0xa0, 0x63, 0x38, 0x2a, 0x12, 0x1d, 0x93, 0xf3, 0x2b, 0xec, 0x72, 0xf3, 0x52, 0x2b, 0xa3,
	0xa7, 0xa0, 0x15, 0x39, 0xab, 0x67, 0x72, 0x4d, 0x45, 0xa7, 0xf0, 0xa2, 0x88, 0x76, 0x29, 0xc1,
	0xf6, 0x07, 0x97, 0xd8, 0xd7, 0x84, 0x9b, 0x9c, 0x38, 0xb6, 0xb6, 0x7b, 0x37, 0xdd, 0xb5, 0x43,
	0x2c, 0xac, 0x55, 0xce, 0x25, 0xc0, 0xe6, 0x85, 0x52, 0x91, 0x75, 0x35, 0xec, 0xdc, 0x53, 0xa1,
	0x08, 0xf7, 0x31, 0x63, 0xe6, 0x25, 0xd6, 0x94, 0x74, 0xdd, 0x45, 0x22, 0x3b, 0xdb, 0x66, 0x3f,
	0x2d, 0x71, 0x02, 0xf5, 0x22, 0xc5, 0x38, 0xc5, 0x66, 0x3f, 0x27, 0xcb, 0x1d, 0xf3, 0xcb, 0xbb,
	0xe9, 0x4c, 0xde, 0x24, 0x63, 0xdd, 0x0b, 0x7c, 0x23, 0x4e, 0x42, 0x11, 0x85, 0x17, 0x17, 0xd2,
	0x98, 0x06, 0xe1, 0x8d, 0x88, 0xbc, 0x68, 0xf4, 0x55, 0x1a, 0xe3, 0x85, 0x90, 0xc6, 0x78, 0x1a,
	0x1b, 0xeb, 0xbf, 0xc6, 0xea, 0x13, 0x65, 0x2c, 0xdb, 0xdf, 0x94, 0xd2, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x01, 0x5a, 0x7e, 0xaf, 0xbf, 0x04, 0x00, 0x00,
}
