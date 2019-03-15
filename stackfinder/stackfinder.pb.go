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

type StackFinderCodeRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Width                int64    `protobuf:"varint,2,opt,name=Width,proto3" json:"Width,omitempty"`
	Height               int64    `protobuf:"varint,3,opt,name=Height,proto3" json:"Height,omitempty"`
	IncludeUnknown       bool     `protobuf:"varint,4,opt,name=IncludeUnknown,proto3" json:"IncludeUnknown,omitempty"`
	ScreenShot           bool     `protobuf:"varint,5,opt,name=ScreenShot,proto3" json:"ScreenShot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StackFinderCodeRequest) Reset()         { *m = StackFinderCodeRequest{} }
func (m *StackFinderCodeRequest) String() string { return proto.CompactTextString(m) }
func (*StackFinderCodeRequest) ProtoMessage()    {}
func (*StackFinderCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b833ff3c53c1a5e, []int{4}
}

func (m *StackFinderCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StackFinderCodeRequest.Unmarshal(m, b)
}
func (m *StackFinderCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StackFinderCodeRequest.Marshal(b, m, deterministic)
}
func (m *StackFinderCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StackFinderCodeRequest.Merge(m, src)
}
func (m *StackFinderCodeRequest) XXX_Size() int {
	return xxx_messageInfo_StackFinderCodeRequest.Size(m)
}
func (m *StackFinderCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StackFinderCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StackFinderCodeRequest proto.InternalMessageInfo

func (m *StackFinderCodeRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *StackFinderCodeRequest) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *StackFinderCodeRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StackFinderCodeRequest) GetIncludeUnknown() bool {
	if m != nil {
		return m.IncludeUnknown
	}
	return false
}

func (m *StackFinderCodeRequest) GetScreenShot() bool {
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
	proto.RegisterType((*StackFinderCodeRequest)(nil), "stackfinder_service.StackFinderCodeRequest")
	proto.RegisterType((*StackFinderResponse)(nil), "stackfinder_service.StackFinderResponse")
	proto.RegisterType((*StackFinderResponse_Item)(nil), "stackfinder_service.StackFinderResponse.Item")
}

func init() { proto.RegisterFile("stackfinder.proto", fileDescriptor_3b833ff3c53c1a5e) }

var fileDescriptor_3b833ff3c53c1a5e = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0xec, 0x38, 0x4d, 0xa7, 0x25, 0xa1, 0xd3, 0xaa, 0xb2, 0x22, 0x81, 0x22, 0x23, 0xa1,
	0x48, 0x40, 0x3e, 0xca, 0x09, 0x90, 0x25, 0x54, 0x0b, 0x09, 0xc1, 0x5a, 0x85, 0x4f, 0x64, 0xec,
	0xa1, 0xb6, 0x9a, 0x78, 0xcb, 0xee, 0x06, 0xc4, 0x75, 0x38, 0x43, 0xef, 0xc0, 0xb5, 0xd0, 0xae,
	0xd7, 0xae, 0xe3, 0x58, 0x6d, 0xbe, 0xb2, 0x6f, 0x3c, 0x33, 0xef, 0xed, 0xce, 0x9b, 0xc0, 0x89,
	0x54, 0x49, 0x7a, 0xf3, 0xa3, 0x28, 0x33, 0x12, 0xcb, 0x5b, 0xc1, 0x15, 0xc7, 0xd3, 0x56, 0xe8,
	0x9b, 0x24, 0xf1, 0xab, 0x48, 0x29, 0x98, 0xc2, 0x93, 0x4b, 0x4a, 0x56, 0x2a, 0x67, 0xf4, 0x73,
	0x43, 0x52, 0x05, 0x4b, 0x18, 0x87, 0x89, 0xa2, 0x6b, 0x2e, 0xfe, 0xe0, 0x04, 0x9c, 0x28, 0xf3,
	0x07, 0xf3, 0xc1, 0xc2, 0x63, 0x4e, 0x94, 0x21, 0xc2, 0xf0, 0x63, 0xb2, 0x26, 0xdf, 0x99, 0x0f,
	0x16, 0x87, 0xcc, 0x9c, 0x83, 0x15, 0x4c, 0xea, 0x06, 0xf2, 0x96, 0x97, 0x92, 0xf0, 0x1c, 0x46,
	0xb1, 0x4a, 0xd4, 0x46, 0x9a, 0xca, 0x43, 0x66, 0x11, 0xfa, 0x70, 0xf0, 0x85, 0x84, 0x2c, 0x78,
	0x69, 0x1b, 0xd4, 0x10, 0x17, 0x30, 0x0d, 0x73, 0xc1, 0xd7, 0xf4, 0x49, 0xf0, 0x94, 0xa4, 0x24,
	0xe9, 0xbb, 0x86, 0xb4, 0x1b, 0x0e, 0x04, 0x9c, 0xc7, 0xfa, 0x16, 0xef, 0xcd, 0x2d, 0xe2, 0x42,
	0x91, 0xd5, 0x8d, 0x4f, 0xc1, 0xbd, 0x12, 0x2b, 0x4b, 0xa9, 0x8f, 0xf8, 0x12, 0x26, 0x51, 0x99,
	0xae, 0x36, 0x19, 0x5d, 0x95, 0x37, 0x25, 0xff, 0x5d, 0xd1, 0x8e, 0x59, 0x27, 0x8a, 0xcf, 0x01,
	0xe2, 0x54, 0x10, 0x95, 0x71, 0xce, 0x95, 0x21, 0x1e, 0xb3, 0x56, 0x24, 0xf8, 0x3b, 0xd8, 0x22,
	0x0d, 0x79, 0xd6, 0x90, 0x22, 0x0c, 0x35, 0xb4, 0xac, 0xe6, 0x8c, 0x67, 0xe0, 0x7d, 0x2d, 0x32,
	0x95, 0x1b, 0x36, 0x97, 0x55, 0x40, 0x3f, 0xca, 0x25, 0x15, 0xd7, 0x79, 0x45, 0xe0, 0x32, 0x8b,
	0x7a, 0x44, 0x0e, 0xf7, 0x10, 0xe9, 0xed, 0x88, 0xfc, 0xe7, 0xc2, 0x69, 0x4b, 0x64, 0x33, 0x8c,
	0x33, 0xf0, 0x42, 0xbe, 0x29, 0x95, 0x9d, 0x62, 0x05, 0x30, 0x04, 0xaf, 0x50, 0xb4, 0x96, 0xbe,
	0x33, 0x77, 0x17, 0x47, 0x17, 0x6f, 0x96, 0x3d, 0xd6, 0x58, 0xf6, 0xb4, 0x5b, 0x46, 0x8a, 0xd6,
	0xac, 0xaa, 0xbd, 0x97, 0x24, 0xeb, 0x77, 0x3b, 0x66, 0xad, 0xc8, 0xec, 0xce, 0x81, 0xa1, 0xce,
	0xd7, 0x77, 0xd7, 0xbf, 0x8d, 0x95, 0x2c, 0xc2, 0x19, 0x8c, 0xf5, 0xa9, 0x65, 0xa9, 0x06, 0xd7,
	0xe3, 0x74, 0xef, 0xc7, 0xf9, 0x0e, 0x8e, 0xf5, 0xd7, 0xda, 0x9c, 0xfe, 0xd0, 0x48, 0x7f, 0xd6,
	0x2b, 0xbd, 0x4e, 0x62, 0x5b, 0x25, 0xda, 0x67, 0x51, 0x59, 0xa8, 0x22, 0x51, 0x5c, 0x58, 0x45,
	0x5e, 0xe5, 0xb3, 0x4e, 0x18, 0x5f, 0xc3, 0xc9, 0x56, 0xc8, 0x68, 0x1c, 0x19, 0x31, 0xbb, 0x1f,
	0xf0, 0x43, 0x2b, 0xbb, 0xd1, 0x77, 0xb0, 0x8f, 0xbe, 0xdd, 0xba, 0x8b, 0x3b, 0x07, 0x8e, 0x5a,
	0x4f, 0x8f, 0x9f, 0xb5, 0x73, 0xf4, 0x82, 0x61, 0xd0, 0xdb, 0x6b, 0x6b, 0x7d, 0x67, 0x2f, 0x1e,
	0xcc, 0xb1, 0xa6, 0xc8, 0x61, 0xda, 0xd9, 0x22, 0x7c, 0xf5, 0x98, 0x05, 0x5a, 0xbb, 0x36, 0x5b,
	0xec, 0xeb, 0x97, 0x0e, 0x93, 0xd9, 0x8f, 0x47, 0x99, 0x5a, 0x0b, 0xb6, 0x3f, 0xd3, 0xf7, 0x91,
	0xf9, 0x93, 0x7b, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xf3, 0x8c, 0x0d, 0xf9, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StackFinderClient is the client API for StackFinder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StackFinderClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	StackFinderSite(ctx context.Context, in *StackFinderSiteRequest, opts ...grpc.CallOption) (*StackFinderResponse, error)
	StackFinderCode(ctx context.Context, in *StackFinderCodeRequest, opts ...grpc.CallOption) (*StackFinderResponse, error)
}

type stackFinderClient struct {
	cc *grpc.ClientConn
}

func NewStackFinderClient(cc *grpc.ClientConn) StackFinderClient {
	return &stackFinderClient{cc}
}

func (c *stackFinderClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/stackfinder_service.StackFinder/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackFinderClient) StackFinderSite(ctx context.Context, in *StackFinderSiteRequest, opts ...grpc.CallOption) (*StackFinderResponse, error) {
	out := new(StackFinderResponse)
	err := c.cc.Invoke(ctx, "/stackfinder_service.StackFinder/StackFinderSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stackFinderClient) StackFinderCode(ctx context.Context, in *StackFinderCodeRequest, opts ...grpc.CallOption) (*StackFinderResponse, error) {
	out := new(StackFinderResponse)
	err := c.cc.Invoke(ctx, "/stackfinder_service.StackFinder/StackFinderCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StackFinderServer is the server API for StackFinder service.
type StackFinderServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	StackFinderSite(context.Context, *StackFinderSiteRequest) (*StackFinderResponse, error)
	StackFinderCode(context.Context, *StackFinderCodeRequest) (*StackFinderResponse, error)
}

func RegisterStackFinderServer(s *grpc.Server, srv StackFinderServer) {
	s.RegisterService(&_StackFinder_serviceDesc, srv)
}

func _StackFinder_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackFinderServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stackfinder_service.StackFinder/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackFinderServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackFinder_StackFinderSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackFinderSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackFinderServer).StackFinderSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stackfinder_service.StackFinder/StackFinderSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackFinderServer).StackFinderSite(ctx, req.(*StackFinderSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StackFinder_StackFinderCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackFinderCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StackFinderServer).StackFinderCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stackfinder_service.StackFinder/StackFinderCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StackFinderServer).StackFinderCode(ctx, req.(*StackFinderCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StackFinder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stackfinder_service.StackFinder",
	HandlerType: (*StackFinderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _StackFinder_Health_Handler,
		},
		{
			MethodName: "StackFinderSite",
			Handler:    _StackFinder_StackFinderSite_Handler,
		},
		{
			MethodName: "StackFinderCode",
			Handler:    _StackFinder_StackFinderCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stackfinder.proto",
}
