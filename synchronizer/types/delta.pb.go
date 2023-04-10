// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: synchronizer/v1/delta.proto

package types

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

// StorageItem is the item for the storage data
type StorageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skey   string `protobuf:"bytes,1,opt,name=skey,proto3" json:"skey,omitempty"`
	Svalue string `protobuf:"bytes,2,opt,name=svalue,proto3" json:"svalue,omitempty"`
}

func (x *StorageItem) Reset() {
	*x = StorageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronizer_v1_delta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageItem) ProtoMessage() {}

func (x *StorageItem) ProtoReflect() protoreflect.Message {
	mi := &file_synchronizer_v1_delta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageItem.ProtoReflect.Descriptor instead.
func (*StorageItem) Descriptor() ([]byte, []int) {
	return file_synchronizer_v1_delta_proto_rawDescGZIP(), []int{0}
}

func (x *StorageItem) GetSkey() string {
	if x != nil {
		return x.Skey
	}
	return ""
}

func (x *StorageItem) GetSvalue() string {
	if x != nil {
		return x.Svalue
	}
	return ""
}

type StorageItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*StorageItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *StorageItemList) Reset() {
	*x = StorageItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronizer_v1_delta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageItemList) ProtoMessage() {}

func (x *StorageItemList) ProtoReflect() protoreflect.Message {
	mi := &file_synchronizer_v1_delta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageItemList.ProtoReflect.Descriptor instead.
func (*StorageItemList) Descriptor() ([]byte, []int) {
	return file_synchronizer_v1_delta_proto_rawDescGZIP(), []int{1}
}

func (x *StorageItemList) GetItems() []*StorageItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// DeltaItem is the item for the state change
type DeltaItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*DeltaItem_StringValue
	//	*DeltaItem_StorageValue
	Value isDeltaItem_Value `protobuf_oneof:"value"`
}

func (x *DeltaItem) Reset() {
	*x = DeltaItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronizer_v1_delta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaItem) ProtoMessage() {}

func (x *DeltaItem) ProtoReflect() protoreflect.Message {
	mi := &file_synchronizer_v1_delta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaItem.ProtoReflect.Descriptor instead.
func (*DeltaItem) Descriptor() ([]byte, []int) {
	return file_synchronizer_v1_delta_proto_rawDescGZIP(), []int{2}
}

func (x *DeltaItem) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DeltaItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *DeltaItem) GetValue() isDeltaItem_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *DeltaItem) GetStringValue() string {
	if x, ok := x.GetValue().(*DeltaItem_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *DeltaItem) GetStorageValue() *StorageItemList {
	if x, ok := x.GetValue().(*DeltaItem_StorageValue); ok {
		return x.StorageValue
	}
	return nil
}

type isDeltaItem_Value interface {
	isDeltaItem_Value()
}

type DeltaItem_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type DeltaItem_StorageValue struct {
	StorageValue *StorageItemList `protobuf:"bytes,4,opt,name=storage_value,json=storageValue,proto3,oneof"`
}

func (*DeltaItem_StringValue) isDeltaItem_Value() {}

func (*DeltaItem_StorageValue) isDeltaItem_Value() {}

// BlockDelta is the delta data for the given block of the specific chain
type BlockDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64       `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	StateRoot   string       `protobuf:"bytes,2,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Chain       string       `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Delta       []*DeltaItem `protobuf:"bytes,4,rep,name=delta,proto3" json:"delta,omitempty"`
	DeltaHash   string       `protobuf:"bytes,5,opt,name=delta_hash,json=deltaHash,proto3" json:"delta_hash,omitempty"`
}

func (x *BlockDelta) Reset() {
	*x = BlockDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronizer_v1_delta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockDelta) ProtoMessage() {}

func (x *BlockDelta) ProtoReflect() protoreflect.Message {
	mi := &file_synchronizer_v1_delta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockDelta.ProtoReflect.Descriptor instead.
func (*BlockDelta) Descriptor() ([]byte, []int) {
	return file_synchronizer_v1_delta_proto_rawDescGZIP(), []int{3}
}

func (x *BlockDelta) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *BlockDelta) GetStateRoot() string {
	if x != nil {
		return x.StateRoot
	}
	return ""
}

func (x *BlockDelta) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *BlockDelta) GetDelta() []*DeltaItem {
	if x != nil {
		return x.Delta
	}
	return nil
}

func (x *BlockDelta) GetDeltaHash() string {
	if x != nil {
		return x.DeltaHash
	}
	return ""
}

var File_synchronizer_v1_delta_proto protoreflect.FileDescriptor

var file_synchronizer_v1_delta_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73,
	0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x39,
	0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6b, 0x65,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0xae, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x47, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xb5, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x4c, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2d, 0x4e,
	0x6f, 0x64, 0x65, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronizer_v1_delta_proto_rawDescOnce sync.Once
	file_synchronizer_v1_delta_proto_rawDescData = file_synchronizer_v1_delta_proto_rawDesc
)

func file_synchronizer_v1_delta_proto_rawDescGZIP() []byte {
	file_synchronizer_v1_delta_proto_rawDescOnce.Do(func() {
		file_synchronizer_v1_delta_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronizer_v1_delta_proto_rawDescData)
	})
	return file_synchronizer_v1_delta_proto_rawDescData
}

var file_synchronizer_v1_delta_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_synchronizer_v1_delta_proto_goTypes = []interface{}{
	(*StorageItem)(nil),     // 0: synchronizer.v1.StorageItem
	(*StorageItemList)(nil), // 1: synchronizer.v1.StorageItemList
	(*DeltaItem)(nil),       // 2: synchronizer.v1.DeltaItem
	(*BlockDelta)(nil),      // 3: synchronizer.v1.BlockDelta
}
var file_synchronizer_v1_delta_proto_depIdxs = []int32{
	0, // 0: synchronizer.v1.StorageItemList.items:type_name -> synchronizer.v1.StorageItem
	1, // 1: synchronizer.v1.DeltaItem.storage_value:type_name -> synchronizer.v1.StorageItemList
	2, // 2: synchronizer.v1.BlockDelta.delta:type_name -> synchronizer.v1.DeltaItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_synchronizer_v1_delta_proto_init() }
func file_synchronizer_v1_delta_proto_init() {
	if File_synchronizer_v1_delta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synchronizer_v1_delta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageItem); i {
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
		file_synchronizer_v1_delta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageItemList); i {
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
		file_synchronizer_v1_delta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaItem); i {
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
		file_synchronizer_v1_delta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockDelta); i {
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
	file_synchronizer_v1_delta_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DeltaItem_StringValue)(nil),
		(*DeltaItem_StorageValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_synchronizer_v1_delta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronizer_v1_delta_proto_goTypes,
		DependencyIndexes: file_synchronizer_v1_delta_proto_depIdxs,
		MessageInfos:      file_synchronizer_v1_delta_proto_msgTypes,
	}.Build()
	File_synchronizer_v1_delta_proto = out.File
	file_synchronizer_v1_delta_proto_rawDesc = nil
	file_synchronizer_v1_delta_proto_goTypes = nil
	file_synchronizer_v1_delta_proto_depIdxs = nil
}