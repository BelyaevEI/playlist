// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: playlist.proto

package playlist_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongName string `protobuf:"bytes,1,opt,name=song_name,json=songName,proto3" json:"song_name,omitempty"`
	Artist   string `protobuf:"bytes,2,opt,name=artist,proto3" json:"artist,omitempty"`
	Duration int32  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *AddSongRequest) Reset() {
	*x = AddSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongRequest) ProtoMessage() {}

func (x *AddSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongRequest.ProtoReflect.Descriptor instead.
func (*AddSongRequest) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{0}
}

func (x *AddSongRequest) GetSongName() string {
	if x != nil {
		return x.SongName
	}
	return ""
}

func (x *AddSongRequest) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *AddSongRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_playlist_proto protoreflect.FileDescriptor

var file_playlist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x6e, 0x67,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6e,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xe8, 0x02, 0x0a, 0x0a, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x31, 0x12, 0x3a, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x53,
	0x6f, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x04,
	0x50, 0x72, 0x65, 0x76, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x42, 0x65, 0x6c, 0x79, 0x61, 0x65, 0x76, 0x45, 0x49, 0x2f, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x76, 0x31, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playlist_proto_rawDescOnce sync.Once
	file_playlist_proto_rawDescData = file_playlist_proto_rawDesc
)

func file_playlist_proto_rawDescGZIP() []byte {
	file_playlist_proto_rawDescOnce.Do(func() {
		file_playlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_playlist_proto_rawDescData)
	})
	return file_playlist_proto_rawDescData
}

var file_playlist_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_playlist_proto_goTypes = []any{
	(*AddSongRequest)(nil), // 0: auth_v1.AddSongRequest
	(*emptypb.Empty)(nil),  // 1: google.protobuf.Empty
}
var file_playlist_proto_depIdxs = []int32{
	0, // 0: auth_v1.PlaylistV1.AddSong:input_type -> auth_v1.AddSongRequest
	0, // 1: auth_v1.PlaylistV1.DeleteSong:input_type -> auth_v1.AddSongRequest
	1, // 2: auth_v1.PlaylistV1.Play:input_type -> google.protobuf.Empty
	1, // 3: auth_v1.PlaylistV1.Pause:input_type -> google.protobuf.Empty
	1, // 4: auth_v1.PlaylistV1.Next:input_type -> google.protobuf.Empty
	1, // 5: auth_v1.PlaylistV1.Prev:input_type -> google.protobuf.Empty
	1, // 6: auth_v1.PlaylistV1.AddSong:output_type -> google.protobuf.Empty
	1, // 7: auth_v1.PlaylistV1.DeleteSong:output_type -> google.protobuf.Empty
	1, // 8: auth_v1.PlaylistV1.Play:output_type -> google.protobuf.Empty
	1, // 9: auth_v1.PlaylistV1.Pause:output_type -> google.protobuf.Empty
	1, // 10: auth_v1.PlaylistV1.Next:output_type -> google.protobuf.Empty
	1, // 11: auth_v1.PlaylistV1.Prev:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_playlist_proto_init() }
func file_playlist_proto_init() {
	if File_playlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_playlist_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddSongRequest); i {
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
			RawDescriptor: file_playlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playlist_proto_goTypes,
		DependencyIndexes: file_playlist_proto_depIdxs,
		MessageInfos:      file_playlist_proto_msgTypes,
	}.Build()
	File_playlist_proto = out.File
	file_playlist_proto_rawDesc = nil
	file_playlist_proto_goTypes = nil
	file_playlist_proto_depIdxs = nil
}
