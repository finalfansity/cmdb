// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pkg/resource/pb/resource.proto

package resource

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

type Vendor int32

const (
	Vendor_NULL    Vendor = 0
	Vendor_ALIYUN  Vendor = 1
	Vendor_TENCENT Vendor = 2
	Vendor_HUAWEI  Vendor = 3
	Vendor_VSPHERE Vendor = 4
)

// Enum value maps for Vendor.
var (
	Vendor_name = map[int32]string{
		0: "NULL",
		1: "ALIYUN",
		2: "TENCENT",
		3: "HUAWEI",
		4: "VSPHERE",
	}
	Vendor_value = map[string]int32{
		"NULL":    0,
		"ALIYUN":  1,
		"TENCENT": 2,
		"HUAWEI":  3,
		"VSPHERE": 4,
	}
)

func (x Vendor) Enum() *Vendor {
	p := new(Vendor)
	*p = x
	return p
}

func (x Vendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vendor) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_resource_pb_resource_proto_enumTypes[0].Descriptor()
}

func (Vendor) Type() protoreflect.EnumType {
	return &file_pkg_resource_pb_resource_proto_enumTypes[0]
}

func (x Vendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vendor.Descriptor instead.
func (Vendor) EnumDescriptor() ([]byte, []int) {
	return file_pkg_resource_pb_resource_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_UNSUPORT Type = 0
	Type_HOST     Type = 1
	Type_RDS      Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "UNSUPORT",
		1: "HOST",
		2: "RDS",
	}
	Type_value = map[string]int32{
		"UNSUPORT": 0,
		"HOST":     1,
		"RDS":      2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_resource_pb_resource_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_pkg_resource_pb_resource_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_pkg_resource_pb_resource_proto_rawDescGZIP(), []int{1}
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源元数据信息
	// @gotags: json:"base"
	Base *Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base"`
	// 资源信息
	// @gotags: json:"information"
	Information *Information `protobuf:"bytes,2,opt,name=information,proto3" json:"information"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_resource_pb_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_resource_pb_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_pkg_resource_pb_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Resource) GetInformation() *Information {
	if x != nil {
		return x.Information
	}
	return nil
}

type ResourceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"items"
	Items []*Resource `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *ResourceSet) Reset() {
	*x = ResourceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_resource_pb_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSet) ProtoMessage() {}

func (x *ResourceSet) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_resource_pb_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSet.ProtoReflect.Descriptor instead.
func (*ResourceSet) Descriptor() ([]byte, []int) {
	return file_pkg_resource_pb_resource_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceSet) GetItems() []*Resource {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ResourceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 全局唯一Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 同步时间
	// @gotags: json:"sync_at"
	SyncAt int64 `protobuf:"varint,2,opt,name=sync_at,json=syncAt,proto3" json:"sync_at"`
	// 用于同步的凭证ID
	// @gotags: json:"secret_id"
	SecretId string `protobuf:"bytes,3,opt,name=secret_id,json=secretId,proto3" json:"secret_id"`
	// 厂商
	// @gotags: json:"vendor"
	Vendor Vendor `protobuf:"varint,4,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.Vendor" json:"vendor"`
	// 资源类型
	// @gotags: json:"resource_type"
	ResourceType Type `protobuf:"varint,5,opt,name=resource_type,json=resourceType,proto3,enum=infraboard.cmdb.resource.Type" json:"resource_type"`
	// 地域
	// @gotags: json:"region"
	Region string `protobuf:"bytes,6,opt,name=region,proto3" json:"region"`
	// 区域
	// @gotags: json:"zone"
	Zone string `protobuf:"bytes,7,opt,name=zone,proto3" json:"zone"`
	// 创建时间
	// @gotags: json:"create_at"
	CreateAt int64 `protobuf:"varint,8,opt,name=create_at,json=createAt,proto3" json:"create_at"`
	// 实例ID
	// @gotags: json:"instance_id"
	InstanceId string `protobuf:"bytes,9,opt,name=instance_id,json=instanceId,proto3" json:"instance_id"`
	// 基础数据Hash
	// @gotags: json:"resource_hash"
	ResourceHash string `protobuf:"bytes,10,opt,name=resource_hash,json=resourceHash,proto3" json:"resource_hash"`
	// 描述数据Hash
	// @gotags: json:"describe_hash"
	DescribeHash string `protobuf:"bytes,11,opt,name=describe_hash,json=describeHash,proto3" json:"describe_hash"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_resource_pb_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_resource_pb_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_pkg_resource_pb_resource_proto_rawDescGZIP(), []int{2}
}

func (x *Base) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Base) GetSyncAt() int64 {
	if x != nil {
		return x.SyncAt
	}
	return 0
}

func (x *Base) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *Base) GetVendor() Vendor {
	if x != nil {
		return x.Vendor
	}
	return Vendor_NULL
}

func (x *Base) GetResourceType() Type {
	if x != nil {
		return x.ResourceType
	}
	return Type_UNSUPORT
}

func (x *Base) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Base) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Base) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Base) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *Base) GetResourceHash() string {
	if x != nil {
		return x.ResourceHash
	}
	return ""
}

func (x *Base) GetDescribeHash() string {
	if x != nil {
		return x.DescribeHash
	}
	return ""
}

type Information struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 过期时间
	// @gotags: json:"expire_at"
	ExpireAt int64 `protobuf:"varint,1,opt,name=expire_at,json=expireAt,proto3" json:"expire_at"`
	// 种类
	// @gotags: json:"category"
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category"`
	// 规格
	// @gotags: json:"type"
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	// 名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// 描述
	// @gotags: json:"description"
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description"`
	// 服务商中的状态
	// @gotags: json:"status"
	Status string `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	// 标签
	// @gotags: json:"tags"
	Tags map[string]string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 更新时间
	// @gotags: json:"update_at"
	UpdateAt int64 `protobuf:"varint,8,opt,name=update_at,json=updateAt,proto3" json:"update_at"`
	// 同步的账号
	// @gotags: json:"sync_account"
	SyncAccount string `protobuf:"bytes,9,opt,name=sync_account,json=syncAccount,proto3" json:"sync_account"`
	// 公网IP
	// @gotags: json:"public_ip"
	PublicIp []string `protobuf:"bytes,10,rep,name=public_ip,json=publicIp,proto3" json:"public_ip"`
	// 内网IP
	// @gotags: json:"private_ip"
	PrivateIp []string `protobuf:"bytes,11,rep,name=private_ip,json=privateIp,proto3" json:"private_ip"`
	// 实例付费方式
	// @gotags: json:"pay_type"
	PayType string `protobuf:"bytes,12,opt,name=pay_type,json=payType,proto3" json:"pay_type"`
}

func (x *Information) Reset() {
	*x = Information{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_resource_pb_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Information) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Information) ProtoMessage() {}

func (x *Information) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_resource_pb_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Information.ProtoReflect.Descriptor instead.
func (*Information) Descriptor() ([]byte, []int) {
	return file_pkg_resource_pb_resource_proto_rawDescGZIP(), []int{3}
}

func (x *Information) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *Information) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Information) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Information) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Information) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Information) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Information) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Information) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Information) GetSyncAccount() string {
	if x != nil {
		return x.SyncAccount
	}
	return ""
}

func (x *Information) GetPublicIp() []string {
	if x != nil {
		return x.PublicIp
	}
	return nil
}

func (x *Information) GetPrivateIp() []string {
	if x != nil {
		return x.PrivateIp
	}
	return nil
}

func (x *Information) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   uint64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber uint64 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Vendor     Vendor `protobuf:"varint,3,opt,name=vendor,proto3,enum=infraboard.cmdb.resource.Vendor" json:"vendor,omitempty"`
	Type       Type   `protobuf:"varint,4,opt,name=type,proto3,enum=infraboard.cmdb.resource.Type" json:"type,omitempty"`
	Keywords   string `protobuf:"bytes,5,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_resource_pb_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_resource_pb_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_pkg_resource_pb_resource_proto_rawDescGZIP(), []int{4}
}

func (x *SearchRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchRequest) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchRequest) GetVendor() Vendor {
	if x != nil {
		return x.Vendor
	}
	return Vendor_NULL
}

func (x *SearchRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_UNSUPORT
}

func (x *SearchRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

var File_pkg_resource_pb_resource_proto protoreflect.FileDescriptor

var file_pkg_resource_pb_resource_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d,
	0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0xff, 0x02, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x79, 0x6e, 0x63, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0xbd, 0x03, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x43, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x70,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x49,
	0x70, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x37, 0x0a, 0x09,
	0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd7, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12,
	0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2a,
	0x44, 0x0a, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c,
	0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c, 0x49, 0x59, 0x55, 0x4e, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x45, 0x4e, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x48, 0x55, 0x41, 0x57, 0x45, 0x49, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x53, 0x50, 0x48,
	0x45, 0x52, 0x45, 0x10, 0x04, 0x2a, 0x27, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x44, 0x53, 0x10, 0x02, 0x32, 0x63,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6d, 0x64,
	0x62, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_resource_pb_resource_proto_rawDescOnce sync.Once
	file_pkg_resource_pb_resource_proto_rawDescData = file_pkg_resource_pb_resource_proto_rawDesc
)

func file_pkg_resource_pb_resource_proto_rawDescGZIP() []byte {
	file_pkg_resource_pb_resource_proto_rawDescOnce.Do(func() {
		file_pkg_resource_pb_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_resource_pb_resource_proto_rawDescData)
	})
	return file_pkg_resource_pb_resource_proto_rawDescData
}

var file_pkg_resource_pb_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_resource_pb_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_resource_pb_resource_proto_goTypes = []interface{}{
	(Vendor)(0),           // 0: infraboard.cmdb.resource.Vendor
	(Type)(0),             // 1: infraboard.cmdb.resource.Type
	(*Resource)(nil),      // 2: infraboard.cmdb.resource.Resource
	(*ResourceSet)(nil),   // 3: infraboard.cmdb.resource.ResourceSet
	(*Base)(nil),          // 4: infraboard.cmdb.resource.Base
	(*Information)(nil),   // 5: infraboard.cmdb.resource.Information
	(*SearchRequest)(nil), // 6: infraboard.cmdb.resource.SearchRequest
	nil,                   // 7: infraboard.cmdb.resource.Information.TagsEntry
}
var file_pkg_resource_pb_resource_proto_depIdxs = []int32{
	4, // 0: infraboard.cmdb.resource.Resource.base:type_name -> infraboard.cmdb.resource.Base
	5, // 1: infraboard.cmdb.resource.Resource.information:type_name -> infraboard.cmdb.resource.Information
	2, // 2: infraboard.cmdb.resource.ResourceSet.items:type_name -> infraboard.cmdb.resource.Resource
	0, // 3: infraboard.cmdb.resource.Base.vendor:type_name -> infraboard.cmdb.resource.Vendor
	1, // 4: infraboard.cmdb.resource.Base.resource_type:type_name -> infraboard.cmdb.resource.Type
	7, // 5: infraboard.cmdb.resource.Information.tags:type_name -> infraboard.cmdb.resource.Information.TagsEntry
	0, // 6: infraboard.cmdb.resource.SearchRequest.vendor:type_name -> infraboard.cmdb.resource.Vendor
	1, // 7: infraboard.cmdb.resource.SearchRequest.type:type_name -> infraboard.cmdb.resource.Type
	6, // 8: infraboard.cmdb.resource.Service.Search:input_type -> infraboard.cmdb.resource.SearchRequest
	3, // 9: infraboard.cmdb.resource.Service.Search:output_type -> infraboard.cmdb.resource.ResourceSet
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_resource_pb_resource_proto_init() }
func file_pkg_resource_pb_resource_proto_init() {
	if File_pkg_resource_pb_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_resource_pb_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_pkg_resource_pb_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceSet); i {
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
		file_pkg_resource_pb_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_pkg_resource_pb_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Information); i {
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
		file_pkg_resource_pb_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
			RawDescriptor: file_pkg_resource_pb_resource_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_resource_pb_resource_proto_goTypes,
		DependencyIndexes: file_pkg_resource_pb_resource_proto_depIdxs,
		EnumInfos:         file_pkg_resource_pb_resource_proto_enumTypes,
		MessageInfos:      file_pkg_resource_pb_resource_proto_msgTypes,
	}.Build()
	File_pkg_resource_pb_resource_proto = out.File
	file_pkg_resource_pb_resource_proto_rawDesc = nil
	file_pkg_resource_pb_resource_proto_goTypes = nil
	file_pkg_resource_pb_resource_proto_depIdxs = nil
}
