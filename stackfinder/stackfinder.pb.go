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

type StackFinderSiteRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	IncludeUnknown       bool     `protobuf:"varint,2,opt,name=IncludeUnknown,proto3" json:"IncludeUnknown,omitempty"`
	ScreenShot           bool     `protobuf:"varint,3,opt,name=ScreenShot,proto3" json:"ScreenShot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StackFinderSiteRequest) Reset()         { *m = StackFinderSiteRequest{} }
func (m *StackFinderSiteRequest) String() string { return proto.CompactTextString(m) }
func (*StackFinderSiteRequest) ProtoMessage()    {}
func (*StackFinderSiteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{3}
}

func (m *StackFinderSiteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackFinderSiteRequest.Unmarshal(m, b)
}
func (m *StackFinderSiteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackFinderSiteRequest.Marshal(b, m, deterministic)
}
func (m *StackFinderSiteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackFinderSiteRequest.Merge(m, src)
}
func (m *StackFinderSiteRequest) XXX_Size() int {
	return xxx_messageInfo_StackFinderSiteRequest.Size(m)
}
func (m *StackFinderSiteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StackFinderSiteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StackFinderSiteRequest proto.InternalMessageInfo

func (m *StackFinderSiteRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *StackFinderSiteRequest) GetIncludeUnknown() bool {
	if m != nil {
		return m.IncludeUnknown
	}
	return false
}

func (m *StackFinderSiteRequest) GetScreenShot() bool {
	if m != nil {
		return m.ScreenShot
	}
	return false
}

type StackFinderHTMLRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
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

func (m *StackFinderHTMLRequest) GetCode() string {
	if m != nil {
		return m.Code
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
	Screenshot           []byte                      `protobuf:"bytes,3,opt,name=Screenshot,proto3" json:"Screenshot,omitempty"`
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
	ItemCategory         []*Category `protobuf:"bytes,4,rep,name=ItemCategory,proto3" json:"ItemCategory,omitempty"`
	InitiatorItemId      int32       `protobuf:"varint,5,opt,name=InitiatorItemId,proto3" json:"InitiatorItemId,omitempty"`
	InitiatorItemName    string      `protobuf:"bytes,6,opt,name=InitiatorItemName,proto3" json:"InitiatorItemName,omitempty"`
	InitiatorCategory    []*Category `protobuf:"bytes,7,rep,name=InitiatorCategory,proto3" json:"InitiatorCategory,omitempty"`
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
	proto.RegisterType((*StackFinderSiteRequest)(nil), "stackfinder_service.StackFinderSiteRequest")
	proto.RegisterType((*StackFinderHTMLRequest)(nil), "stackfinder_service.StackFinderHTMLRequest")
	proto.RegisterType((*StackFinderResponse)(nil), "stackfinder_service.StackFinderResponse")
	proto.RegisterType((*StackFinderResponse_Item)(nil), "stackfinder_service.StackFinderResponse.Item")
}

func init() { proto.RegisterFile("stackfinder.proto", fileDescriptor_3b833ff3c53c1a5e) }

var fileDescriptor_3b833ff3c53c1a5e = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0xec, 0x38, 0x4d, 0x87, 0x92, 0xd0, 0x69, 0x55, 0x59, 0x91, 0x40, 0x91, 0x91, 0x50,
	0x24, 0x20, 0x87, 0xf2, 0x04, 0xc8, 0x52, 0x15, 0x8b, 0x1f, 0xc1, 0x9a, 0xc2, 0x11, 0x19, 0x7b,
	0xa8, 0xad, 0x26, 0xde, 0xb2, 0xbb, 0x01, 0xf1, 0x3a, 0x5c, 0xb9, 0xf2, 0x0e, 0xbc, 0x16, 0xda,
	0xf5, 0x3a, 0x75, 0x1c, 0x8b, 0xe6, 0xe4, 0x99, 0xd9, 0x99, 0xf9, 0xbe, 0xdd, 0xf9, 0xc6, 0x70,
	0x2c, 0x55, 0x92, 0x5e, 0x7f, 0x2d, 0xca, 0x8c, 0xc4, 0xfc, 0x46, 0x70, 0xc5, 0xf1, 0xa4, 0x11,
	0xfa, 0x2c, 0x49, 0x7c, 0x2f, 0x52, 0x0a, 0xc6, 0x70, 0x7f, 0x41, 0xc9, 0x52, 0xe5, 0x8c, 0xbe,
	0xad, 0x49, 0xaa, 0x60, 0x0e, 0xc3, 0x30, 0x51, 0x74, 0xc5, 0xc5, 0x4f, 0x1c, 0x81, 0x13, 0x65,
	0x7e, 0x6f, 0xda, 0x9b, 0x79, 0xcc, 0x89, 0x32, 0x44, 0xe8, 0xbf, 0x4d, 0x56, 0xe4, 0x3b, 0xd3,
	0xde, 0xec, 0x90, 0x19, 0x3b, 0x58, 0xc2, 0xa8, 0x6e, 0x20, 0x6f, 0x78, 0x29, 0x09, 0xcf, 0x60,
	0x10, 0xab, 0x44, 0xad, 0xa5, 0xa9, 0x3c, 0x64, 0xd6, 0x43, 0x1f, 0x0e, 0x3e, 0x92, 0x90, 0x05,
	0x2f, 0x6d, 0x83, 0xda, 0xc5, 0x19, 0x8c, 0xc3, 0x5c, 0xf0, 0x15, 0xbd, 0x13, 0x3c, 0x25, 0x29,
	0x49, 0xfa, 0xae, 0x01, 0x6d, 0x87, 0x03, 0x01, 0x67, 0xb1, 0xbe, 0xc5, 0x85, 0xb9, 0x45, 0x5c,
	0x28, 0xb2, 0xbc, 0xf1, 0x01, 0xb8, 0x97, 0x62, 0x69, 0x21, 0xb5, 0x89, 0x4f, 0x60, 0x14, 0x95,
	0xe9, 0x72, 0x9d, 0xd1, 0x65, 0x79, 0x5d, 0xf2, 0x1f, 0x15, 0xec, 0x90, 0xb5, 0xa2, 0xf8, 0x08,
	0x20, 0x4e, 0x05, 0x51, 0x19, 0xe7, 0x5c, 0x19, 0xe0, 0x21, 0x6b, 0x44, 0x82, 0x5f, 0xbd, 0x2d,
	0xd0, 0xc5, 0x87, 0x37, 0xaf, 0x6b, 0x50, 0x84, 0x7e, 0xc8, 0x33, 0xb2, 0xa8, 0xc6, 0xc6, 0x53,
	0xf0, 0x3e, 0x15, 0x99, 0xca, 0x0d, 0x9a, 0xcb, 0x2a, 0x47, 0x3f, 0xca, 0x82, 0x8a, 0xab, 0xbc,
	0x02, 0x70, 0x99, 0xf5, 0x3a, 0x48, 0xf6, 0xf7, 0x20, 0xe9, 0xed, 0x90, 0xfc, 0xeb, 0xc2, 0x49,
	0x83, 0xe4, 0x66, 0x18, 0xa7, 0xe0, 0x85, 0x7c, 0x5d, 0x2a, 0x3b, 0xc5, 0xca, 0xc1, 0x10, 0xbc,
	0x42, 0xd1, 0x4a, 0xfa, 0xce, 0xd4, 0x9d, 0xdd, 0x3b, 0x7f, 0x3e, 0xef, 0x90, 0xc6, 0xbc, 0xa3,
	0xdd, 0x3c, 0x52, 0xb4, 0x62, 0x55, 0xed, 0x2d, 0x25, 0x59, 0xbf, 0xdb, 0x11, 0x6b, 0x44, 0x26,
	0x7f, 0x1c, 0xe8, 0xeb, 0x7c, 0x7d, 0x77, 0xfd, 0xdd, 0x48, 0xc9, 0x7a, 0x38, 0x81, 0xa1, 0xb6,
	0x1a, 0x92, 0xda, 0xf8, 0xf5, 0x38, 0xdd, 0xdb, 0x71, 0xbe, 0x84, 0x23, 0x7d, 0x5a, 0x8b, 0xd3,
	0xef, 0x1b, 0xea, 0x0f, 0x3b, 0xa9, 0xd7, 0x49, 0x6c, 0xab, 0x44, 0xeb, 0x2c, 0x2a, 0x0b, 0x55,
	0x24, 0x8a, 0x0b, 0xcb, 0xc8, 0xab, 0x74, 0xd6, 0x0a, 0xe3, 0x33, 0x38, 0xde, 0x0a, 0x19, 0x8e,
	0x03, 0x43, 0x66, 0xf7, 0x00, 0x5f, 0x35, 0xb2, 0x37, 0xfc, 0x0e, 0xf6, 0xe1, 0xb7, 0x5b, 0x77,
	0xfe, 0xdb, 0x01, 0x27, 0xbe, 0xc0, 0xf7, 0x5a, 0x30, 0x7a, 0xaf, 0x30, 0xe8, 0x6c, 0xb1, 0xb5,
	0xb5, 0x93, 0xc7, 0xff, 0xcd, 0xb1, 0x5a, 0xc8, 0x61, 0xdc, 0x5a, 0x1e, 0x7c, 0x7a, 0xd7, 0xe4,
	0x1b, 0x2b, 0x36, 0x99, 0xed, 0x2b, 0x93, 0x16, 0x92, 0xde, 0x98, 0xbb, 0x91, 0x1a, 0x7b, 0xb5,
	0x3f, 0xd2, 0x97, 0x81, 0xf9, 0xb7, 0xbd, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xac, 0x6c, 0x47,
	0xa1, 0xf0, 0x04, 0x00, 0x00,
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
	StackFinderSite(ctx context.Context, in *StackFinderSiteRequest, opts ...grpc.CallOption) (*StackFinderResponse, error)
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

func (c *sFClient) StackFinderSite(ctx context.Context, in *StackFinderSiteRequest, opts ...grpc.CallOption) (*StackFinderResponse, error) {
	out := new(StackFinderResponse)
	err := c.cc.Invoke(ctx, "/stackfinder_service.SF/StackFinderSite", in, out, opts...)
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
	StackFinderSite(context.Context, *StackFinderSiteRequest) (*StackFinderResponse, error)
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

func _SF_StackFinderSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackFinderSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SFServer).StackFinderSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stackfinder_service.SF/StackFinderSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SFServer).StackFinderSite(ctx, req.(*StackFinderSiteRequest))
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
			MethodName: "StackFinderSite",
			Handler:    _SF_StackFinderSite_Handler,
		},
		{
			MethodName: "StackFinderHTML",
			Handler:    _SF_StackFinderHTML_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stackfinder.proto",
}
