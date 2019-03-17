// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stackfinder.proto

package stackfinder_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{0}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type Category struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{1}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HealthResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	ChromeProcesses      int32    `protobuf:"varint,3,opt,name=ChromeProcesses,proto3" json:"ChromeProcesses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{2}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *HealthResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *HealthResponse) GetChromeProcesses() int32 {
	if m != nil {
		return m.ChromeProcesses
	}
	return 0
}

type StackFinderPageRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	IncludeUnknown       bool     `protobuf:"varint,2,opt,name=IncludeUnknown,proto3" json:"IncludeUnknown,omitempty"`
	ScreenShot           bool     `protobuf:"varint,3,opt,name=ScreenShot,proto3" json:"ScreenShot,omitempty"`
	HeaderBidInfo        bool     `protobuf:"varint,4,opt,name=HeaderBidInfo,proto3" json:"HeaderBidInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StackFinderPageRequest) Reset()         { *m = StackFinderPageRequest{} }
func (m *StackFinderPageRequest) String() string { return proto.CompactTextString(m) }
func (*StackFinderPageRequest) ProtoMessage()    {}
func (*StackFinderPageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{3}
}

func (m *StackFinderPageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackFinderPageRequest.Unmarshal(m, b)
}
func (m *StackFinderPageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackFinderPageRequest.Marshal(b, m, deterministic)
}
func (m *StackFinderPageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackFinderPageRequest.Merge(m, src)
}
func (m *StackFinderPageRequest) XXX_Size() int {
	return xxx_messageInfo_StackFinderPageRequest.Size(m)
}
func (m *StackFinderPageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StackFinderPageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StackFinderPageRequest proto.InternalMessageInfo

func (m *StackFinderPageRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *StackFinderPageRequest) GetIncludeUnknown() bool {
	if m != nil {
		return m.IncludeUnknown
	}
	return false
}

func (m *StackFinderPageRequest) GetScreenShot() bool {
	if m != nil {
		return m.ScreenShot
	}
	return false
}

func (m *StackFinderPageRequest) GetHeaderBidInfo() bool {
	if m != nil {
		return m.HeaderBidInfo
	}
	return false
}

type StackFinderHTMLRequest struct {
	HTML                 string   `protobuf:"bytes,1,opt,name=HTML,proto3" json:"HTML,omitempty"`
	Width                int64    `protobuf:"varint,2,opt,name=Width,proto3" json:"Width,omitempty"`
	Height               int64    `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`
	IncludeUnknown       bool     `protobuf:"varint,4,opt,name=IncludeUnknown,proto3" json:"IncludeUnknown,omitempty"`
	ScreenShot           bool     `protobuf:"varint,5,opt,name=ScreenShot,proto3" json:"ScreenShot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StackFinderHTMLRequest) Reset()         { *m = StackFinderHTMLRequest{} }
func (m *StackFinderHTMLRequest) String() string { return proto.CompactTextString(m) }
func (*StackFinderHTMLRequest) ProtoMessage()    {}
func (*StackFinderHTMLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{4}
}

func (m *StackFinderHTMLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackFinderHTMLRequest.Unmarshal(m, b)
}
func (m *StackFinderHTMLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackFinderHTMLRequest.Marshal(b, m, deterministic)
}
func (m *StackFinderHTMLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackFinderHTMLRequest.Merge(m, src)
}
func (m *StackFinderHTMLRequest) XXX_Size() int {
	return xxx_messageInfo_StackFinderHTMLRequest.Size(m)
}
func (m *StackFinderHTMLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StackFinderHTMLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StackFinderHTMLRequest proto.InternalMessageInfo

func (m *StackFinderHTMLRequest) GetHTML() string {
	if m != nil {
		return m.HTML
	}
	return ""
}

func (m *StackFinderHTMLRequest) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *StackFinderHTMLRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StackFinderHTMLRequest) GetIncludeUnknown() bool {
	if m != nil {
		return m.IncludeUnknown
	}
	return false
}

func (m *StackFinderHTMLRequest) GetScreenShot() bool {
	if m != nil {
		return m.ScreenShot
	}
	return false
}

type StackFinderResponse struct {
	Count                int32                       `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Items                []*StackFinderResponse_Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	BidResponses         []byte                      `protobuf:"bytes,3,opt,name=BidResponses,proto3" json:"BidResponses,omitempty"`
	WinningBids          []byte                      `protobuf:"bytes,4,opt,name=WinningBids,proto3" json:"WinningBids,omitempty"`
	Screenshot           []byte                      `protobuf:"bytes,5,opt,name=Screenshot,proto3" json:"Screenshot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StackFinderResponse) Reset()         { *m = StackFinderResponse{} }
func (m *StackFinderResponse) String() string { return proto.CompactTextString(m) }
func (*StackFinderResponse) ProtoMessage()    {}
func (*StackFinderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{5}
}

func (m *StackFinderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackFinderResponse.Unmarshal(m, b)
}
func (m *StackFinderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackFinderResponse.Marshal(b, m, deterministic)
}
func (m *StackFinderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackFinderResponse.Merge(m, src)
}
func (m *StackFinderResponse) XXX_Size() int {
	return xxx_messageInfo_StackFinderResponse.Size(m)
}
func (m *StackFinderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StackFinderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StackFinderResponse proto.InternalMessageInfo

func (m *StackFinderResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *StackFinderResponse) GetItems() []*StackFinderResponse_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *StackFinderResponse) GetBidResponses() []byte {
	if m != nil {
		return m.BidResponses
	}
	return nil
}

func (m *StackFinderResponse) GetWinningBids() []byte {
	if m != nil {
		return m.WinningBids
	}
	return nil
}

func (m *StackFinderResponse) GetScreenshot() []byte {
	if m != nil {
		return m.Screenshot
	}
	return nil
}

type StackFinderResponse_Item struct {
	ItemId               int32       `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	ItemName             string      `protobuf:"bytes,2,opt,name=ItemName,proto3" json:"ItemName,omitempty"`
	Url                  string      `protobuf:"bytes,3,opt,name=Url,proto3" json:"Url,omitempty"`
	Type                 string      `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	ItemCategory         []*Category `protobuf:"bytes,5,rep,name=ItemCategory,proto3" json:"ItemCategory,omitempty"`
	InitiatorItemId      int32       `protobuf:"varint,6,opt,name=InitiatorItemId,proto3" json:"InitiatorItemId,omitempty"`
	InitiatorItemName    string      `protobuf:"bytes,7,opt,name=InitiatorItemName,proto3" json:"InitiatorItemName,omitempty"`
	InitiatorCategory    []*Category `protobuf:"bytes,8,rep,name=InitiatorCategory,proto3" json:"InitiatorCategory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StackFinderResponse_Item) Reset()         { *m = StackFinderResponse_Item{} }
func (m *StackFinderResponse_Item) String() string { return proto.CompactTextString(m) }
func (*StackFinderResponse_Item) ProtoMessage()    {}
func (*StackFinderResponse_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{5, 0}
}

func (m *StackFinderResponse_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackFinderResponse_Item.Unmarshal(m, b)
}
func (m *StackFinderResponse_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackFinderResponse_Item.Marshal(b, m, deterministic)
}
func (m *StackFinderResponse_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackFinderResponse_Item.Merge(m, src)
}
func (m *StackFinderResponse_Item) XXX_Size() int {
	return xxx_messageInfo_StackFinderResponse_Item.Size(m)
}
func (m *StackFinderResponse_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_StackFinderResponse_Item.DiscardUnknown(m)
}

var xxx_messageInfo_StackFinderResponse_Item proto.InternalMessageInfo

func (m *StackFinderResponse_Item) GetItemId() int32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *StackFinderResponse_Item) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *StackFinderResponse_Item) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *StackFinderResponse_Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StackFinderResponse_Item) GetItemCategory() []*Category {
	if m != nil {
		return m.ItemCategory
	}
	return nil
}

func (m *StackFinderResponse_Item) GetInitiatorItemId() int32 {
	if m != nil {
		return m.InitiatorItemId
	}
	return 0
}

func (m *StackFinderResponse_Item) GetInitiatorItemName() string {
	if m != nil {
		return m.InitiatorItemName
	}
	return ""
}

func (m *StackFinderResponse_Item) GetInitiatorCategory() []*Category {
	if m != nil {
		return m.InitiatorCategory
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthRequest)(nil), "stackfinder_service.HealthRequest")
	proto.RegisterType((*Category)(nil), "stackfinder_service.Category")
	proto.RegisterType((*HealthResponse)(nil), "stackfinder_service.HealthResponse")
	proto.RegisterType((*StackFinderPageRequest)(nil), "stackfinder_service.StackFinderPageRequest")
	proto.RegisterType((*StackFinderHTMLRequest)(nil), "stackfinder_service.StackFinderHTMLRequest")
	proto.RegisterType((*StackFinderResponse)(nil), "stackfinder_service.StackFinderResponse")
	proto.RegisterType((*StackFinderResponse_Item)(nil), "stackfinder_service.StackFinderResponse.Item")
}

func init() { proto.RegisterFile("stackfinder.proto", fileDescriptor_3b833ff3c53c1a5e) }

var fileDescriptor_3b833ff3c53c1a5e = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0xfc, 0xd3, 0x26, 0xd3, 0x90, 0xd0, 0x69, 0x55, 0x59, 0x91, 0x40, 0x91, 0x41, 0x28,
	0x12, 0x90, 0x43, 0x79, 0x02, 0x12, 0xa9, 0x8a, 0xc5, 0x8f, 0xca, 0xa6, 0xa5, 0x47, 0x64, 0xe2,
	0x6d, 0xbc, 0x6a, 0xb2, 0x1b, 0x76, 0x37, 0xa0, 0xbe, 0x09, 0x57, 0xb8, 0xf2, 0x22, 0x3c, 0x16,
	0xda, 0xb5, 0x9d, 0xda, 0x6e, 0x44, 0x73, 0xf2, 0xcc, 0xe7, 0x9d, 0x99, 0x6f, 0x66, 0xbf, 0x59,
	0x38, 0x54, 0x3a, 0x9e, 0xdd, 0x5c, 0x33, 0x9e, 0x50, 0x39, 0x5c, 0x49, 0xa1, 0x05, 0x1e, 0x95,
	0xa0, 0x2f, 0x8a, 0xca, 0xef, 0x6c, 0x46, 0xc3, 0x2e, 0x3c, 0x9a, 0xd0, 0x78, 0xa1, 0x53, 0x42,
	0xbf, 0xad, 0xa9, 0xd2, 0xe1, 0x10, 0x9a, 0xe3, 0x58, 0xd3, 0xb9, 0x90, 0xb7, 0xd8, 0x01, 0x27,
	0x4a, 0x82, 0x46, 0xbf, 0x31, 0xf0, 0x89, 0x13, 0x25, 0x88, 0xe0, 0x7d, 0x8c, 0x97, 0x34, 0x70,
	0xfa, 0x8d, 0x41, 0x8b, 0x58, 0x3b, 0x5c, 0x40, 0xa7, 0x48, 0xa0, 0x56, 0x82, 0x2b, 0x8a, 0x27,
	0xb0, 0x37, 0xd5, 0xb1, 0x5e, 0x2b, 0x1b, 0xd9, 0x22, 0xb9, 0x87, 0x01, 0xec, 0x7f, 0xa6, 0x52,
	0x31, 0xc1, 0xf3, 0x04, 0x85, 0x8b, 0x03, 0xe8, 0x8e, 0x53, 0x29, 0x96, 0xf4, 0x5c, 0x8a, 0x19,
	0x55, 0x8a, 0xaa, 0xc0, 0xb5, 0x45, 0xeb, 0x70, 0xf8, 0xb3, 0x01, 0x27, 0x53, 0xd3, 0xc6, 0x99,
	0x6d, 0xe3, 0x3c, 0x9e, 0xd3, 0x9c, 0x38, 0x3e, 0x06, 0xf7, 0x52, 0x2e, 0xf2, 0x9a, 0xc6, 0xc4,
	0x17, 0xd0, 0x89, 0xf8, 0x6c, 0xb1, 0x4e, 0xe8, 0x25, 0xbf, 0xe1, 0xe2, 0x47, 0x56, 0xb7, 0x49,
	0x6a, 0x28, 0x3e, 0x05, 0x98, 0xce, 0x24, 0xa5, 0x7c, 0x9a, 0x0a, 0x6d, 0x2b, 0x37, 0x49, 0x09,
	0xc1, 0xe7, 0x76, 0x46, 0x09, 0x95, 0x23, 0x96, 0x44, 0xfc, 0x5a, 0x04, 0x9e, 0x3d, 0x52, 0x05,
	0xc3, 0xdf, 0x55, 0x6a, 0x93, 0x8b, 0x0f, 0xef, 0x0b, 0x6a, 0x08, 0x9e, 0x71, 0x73, 0x6e, 0xd6,
	0xc6, 0x63, 0xf0, 0xaf, 0x58, 0xa2, 0x53, 0xcb, 0xc9, 0x25, 0x99, 0x63, 0x66, 0x37, 0xa1, 0x6c,
	0x9e, 0x66, 0x34, 0x5c, 0x92, 0x7b, 0x5b, 0x5a, 0xf1, 0x76, 0x68, 0xc5, 0xaf, 0xb7, 0x12, 0xfe,
	0xf2, 0xe0, 0xa8, 0x44, 0x72, 0x73, 0x67, 0xc7, 0xe0, 0x8f, 0xc5, 0x9a, 0xeb, 0xfc, 0xb2, 0x33,
	0x07, 0xc7, 0xe0, 0x33, 0x4d, 0x97, 0x2a, 0x70, 0xfa, 0xee, 0xe0, 0xe0, 0xf4, 0xf5, 0x70, 0x8b,
	0x82, 0x86, 0x5b, 0xd2, 0x0d, 0x23, 0x4d, 0x97, 0x24, 0x8b, 0xc5, 0x10, 0xda, 0x23, 0x96, 0x14,
	0xbf, 0xb2, 0x9b, 0x6d, 0x93, 0x0a, 0x86, 0x7d, 0x38, 0xb8, 0x62, 0x9c, 0x33, 0x3e, 0x1f, 0xb1,
	0x44, 0xd9, 0xde, 0xda, 0xa4, 0x0c, 0xdd, 0x35, 0xa6, 0x8a, 0xc6, 0xda, 0xa4, 0x84, 0xf4, 0xfe,
	0x3a, 0xe0, 0x99, 0xaa, 0x66, 0x82, 0xe6, 0xbb, 0xd1, 0x6d, 0xee, 0x61, 0x0f, 0x9a, 0xc6, 0x2a,
	0xe9, 0x77, 0xe3, 0x17, 0xd2, 0x71, 0xef, 0xa4, 0x83, 0xe0, 0x5d, 0xdc, 0xae, 0xa8, 0x65, 0xd2,
	0x22, 0xd6, 0xc6, 0xb7, 0xd0, 0x36, 0x11, 0xc5, 0x76, 0x04, 0xbe, 0x1d, 0xca, 0x93, 0xad, 0x43,
	0x29, 0x0e, 0x91, 0x4a, 0x88, 0x11, 0x7a, 0xc4, 0x99, 0x66, 0xb1, 0x16, 0x32, 0x67, 0xb9, 0x97,
	0x09, 0xbd, 0x06, 0xe3, 0x2b, 0x38, 0xac, 0x40, 0x96, 0xf7, 0xbe, 0x65, 0x73, 0xff, 0x07, 0xbe,
	0x2b, 0x9d, 0xde, 0xf0, 0x6b, 0xee, 0xc2, 0xef, 0x7e, 0xdc, 0xe9, 0x1f, 0x07, 0x9c, 0xe9, 0x19,
	0x7e, 0x32, 0x52, 0x34, 0x8b, 0x8d, 0xe1, 0xd6, 0x14, 0x95, 0x67, 0xa3, 0xf7, 0xec, 0xbf, 0x67,
	0x72, 0x95, 0xa5, 0xd0, 0xad, 0x2d, 0x2f, 0xbe, 0x7c, 0x48, 0x53, 0xa5, 0x15, 0xef, 0x0d, 0x76,
	0x15, 0x60, 0xad, 0x92, 0x5d, 0xb8, 0x07, 0x2b, 0x95, 0x36, 0x76, 0xf7, 0x4a, 0x5f, 0xf7, 0xec,
	0xe3, 0xfa, 0xe6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x7f, 0xae, 0xe0, 0x71, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SFClient is the client API for SF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SFClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	StackFinderPage(ctx context.Context, in *StackFinderPageRequest, opts ...grpc.CallOption) (*StackFinderResponse, error)
	StackFinderHTML(ctx context.Context, in *StackFinderHTMLRequest, opts ...grpc.CallOption) (*StackFinderResponse, error)
}

type sFClient struct {
	cc *grpc.ClientConn
}

func NewSFClient(cc *grpc.ClientConn) SFClient {
	return &sFClient{cc}
}

func (c *sFClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/stackfinder_service.SF/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sFClient) StackFinderPage(ctx context.Context, in *StackFinderPageRequest, opts ...grpc.CallOption) (*StackFinderResponse, error) {
	out := new(StackFinderResponse)
	err := c.cc.Invoke(ctx, "/stackfinder_service.SF/StackFinderPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sFClient) StackFinderHTML(ctx context.Context, in *StackFinderHTMLRequest, opts ...grpc.CallOption) (*StackFinderResponse, error) {
	out := new(StackFinderResponse)
	err := c.cc.Invoke(ctx, "/stackfinder_service.SF/StackFinderHTML", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SFServer is the server API for SF service.
type SFServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	StackFinderPage(context.Context, *StackFinderPageRequest) (*StackFinderResponse, error)
	StackFinderHTML(context.Context, *StackFinderHTMLRequest) (*StackFinderResponse, error)
}

func RegisterSFServer(s *grpc.Server, srv SFServer) {
	s.RegisterService(&_SF_serviceDesc, srv)
}

func _SF_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SFServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stackfinder_service.SF/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SFServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SF_StackFinderPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackFinderPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SFServer).StackFinderPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stackfinder_service.SF/StackFinderPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SFServer).StackFinderPage(ctx, req.(*StackFinderPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SF_StackFinderHTML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackFinderHTMLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SFServer).StackFinderHTML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stackfinder_service.SF/StackFinderHTML",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SFServer).StackFinderHTML(ctx, req.(*StackFinderHTMLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SF_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stackfinder_service.SF",
	HandlerType: (*SFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _SF_Health_Handler,
		},
		{
			MethodName: "StackFinderPage",
			Handler:    _SF_StackFinderPage_Handler,
		},
		{
			MethodName: "StackFinderHTML",
			Handler:    _SF_StackFinderHTML_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stackfinder.proto",
}
