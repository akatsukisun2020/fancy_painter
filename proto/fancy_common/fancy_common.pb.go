// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: fancy_common.proto

package fancy_common

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

// 智能画图任务 zset结构： key: SmartDrawImgTask score: 毫秒时间戳
type SmartDrawImgTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId      string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`                // 任务id， ${phone_number}_${毫秒时间戳}
	TextDesc    string `protobuf:"bytes,2,opt,name=text_desc,json=textDesc,proto3" json:"text_desc,omitempty"`          // 描述
	Artist      string `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`                              // 艺术家
	TargetStyle string `protobuf:"bytes,4,opt,name=target_style,json=targetStyle,proto3" json:"target_style,omitempty"` // 目标分割
	ImgWidth    int64  `protobuf:"varint,5,opt,name=img_width,json=imgWidth,proto3" json:"img_width,omitempty"`         // 宽
	ImgHeight   int64  `protobuf:"varint,7,opt,name=img_height,json=imgHeight,proto3" json:"img_height,omitempty"`      // 高
}

func (x *SmartDrawImgTask) Reset() {
	*x = SmartDrawImgTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fancy_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartDrawImgTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartDrawImgTask) ProtoMessage() {}

func (x *SmartDrawImgTask) ProtoReflect() protoreflect.Message {
	mi := &file_fancy_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartDrawImgTask.ProtoReflect.Descriptor instead.
func (*SmartDrawImgTask) Descriptor() ([]byte, []int) {
	return file_fancy_common_proto_rawDescGZIP(), []int{0}
}

func (x *SmartDrawImgTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *SmartDrawImgTask) GetTextDesc() string {
	if x != nil {
		return x.TextDesc
	}
	return ""
}

func (x *SmartDrawImgTask) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *SmartDrawImgTask) GetTargetStyle() string {
	if x != nil {
		return x.TargetStyle
	}
	return ""
}

func (x *SmartDrawImgTask) GetImgWidth() int64 {
	if x != nil {
		return x.ImgWidth
	}
	return 0
}

func (x *SmartDrawImgTask) GetImgHeight() int64 {
	if x != nil {
		return x.ImgHeight
	}
	return 0
}

// 用户生成的图片记录
type UserImgRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImgTask *SmartDrawImgTask `protobuf:"bytes,1,opt,name=img_task,json=imgTask,proto3" json:"img_task,omitempty"` // 任务信息
	ImgUrl  string            `protobuf:"bytes,2,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`    // 生成的图片信息
	Time    int64             `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`                     // 时间
}

func (x *UserImgRecord) Reset() {
	*x = UserImgRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fancy_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserImgRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserImgRecord) ProtoMessage() {}

func (x *UserImgRecord) ProtoReflect() protoreflect.Message {
	mi := &file_fancy_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserImgRecord.ProtoReflect.Descriptor instead.
func (*UserImgRecord) Descriptor() ([]byte, []int) {
	return file_fancy_common_proto_rawDescGZIP(), []int{1}
}

func (x *UserImgRecord) GetImgTask() *SmartDrawImgTask {
	if x != nil {
		return x.ImgTask
	}
	return nil
}

func (x *UserImgRecord) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *UserImgRecord) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// 用户图片资产 kv结构: key: UserAssets_${手机号}
type UserAssets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string           `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"` // 手机号
	Records     []*UserImgRecord `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`                            // 任务记录列表，按照时间戳排序
}

func (x *UserAssets) Reset() {
	*x = UserAssets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fancy_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAssets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAssets) ProtoMessage() {}

func (x *UserAssets) ProtoReflect() protoreflect.Message {
	mi := &file_fancy_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAssets.ProtoReflect.Descriptor instead.
func (*UserAssets) Descriptor() ([]byte, []int) {
	return file_fancy_common_proto_rawDescGZIP(), []int{2}
}

func (x *UserAssets) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserAssets) GetRecords() []*UserImgRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_fancy_common_proto protoreflect.FileDescriptor

var file_fancy_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x22, 0xbf, 0x01, 0x0a, 0x10, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x44, 0x72, 0x61, 0x77, 0x49,
	0x6d, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73,
	0x74, 0x79, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x67, 0x5f, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6d, 0x67, 0x57,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x67, 0x5f, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6d, 0x67, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x76, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x69, 0x6d, 0x67, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x44, 0x72, 0x61, 0x77, 0x49, 0x6d,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x07, 0x69, 0x6d, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x66, 0x61, 0x6e, 0x63, 0x79, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6d, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fancy_common_proto_rawDescOnce sync.Once
	file_fancy_common_proto_rawDescData = file_fancy_common_proto_rawDesc
)

func file_fancy_common_proto_rawDescGZIP() []byte {
	file_fancy_common_proto_rawDescOnce.Do(func() {
		file_fancy_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_fancy_common_proto_rawDescData)
	})
	return file_fancy_common_proto_rawDescData
}

var file_fancy_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fancy_common_proto_goTypes = []interface{}{
	(*SmartDrawImgTask)(nil), // 0: fancycommon.SmartDrawImgTask
	(*UserImgRecord)(nil),    // 1: fancycommon.UserImgRecord
	(*UserAssets)(nil),       // 2: fancycommon.UserAssets
}
var file_fancy_common_proto_depIdxs = []int32{
	0, // 0: fancycommon.UserImgRecord.img_task:type_name -> fancycommon.SmartDrawImgTask
	1, // 1: fancycommon.UserAssets.records:type_name -> fancycommon.UserImgRecord
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fancy_common_proto_init() }
func file_fancy_common_proto_init() {
	if File_fancy_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fancy_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartDrawImgTask); i {
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
		file_fancy_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserImgRecord); i {
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
		file_fancy_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAssets); i {
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
			RawDescriptor: file_fancy_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fancy_common_proto_goTypes,
		DependencyIndexes: file_fancy_common_proto_depIdxs,
		MessageInfos:      file_fancy_common_proto_msgTypes,
	}.Build()
	File_fancy_common_proto = out.File
	file_fancy_common_proto_rawDesc = nil
	file_fancy_common_proto_goTypes = nil
	file_fancy_common_proto_depIdxs = nil
}
