// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: synchronization/configuration.proto

package synchronization

import (
	behavior "github.com/mutagen-io/mutagen/pkg/filesystem/behavior"
	compression "github.com/mutagen-io/mutagen/pkg/synchronization/compression"
	core "github.com/mutagen-io/mutagen/pkg/synchronization/core"
	ignore "github.com/mutagen-io/mutagen/pkg/synchronization/core/ignore"
	hashing "github.com/mutagen-io/mutagen/pkg/synchronization/hashing"
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

// Configuration encodes session configuration parameters. It is used for create
// commands to specify configuration options, for loading global configuration
// options, and for storing a merged configuration inside sessions. It should be
// considered immutable.
type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SynchronizationMode specifies the synchronization mode that should be
	// used in synchronization.
	SynchronizationMode core.SynchronizationMode `protobuf:"varint,11,opt,name=synchronizationMode,proto3,enum=core.SynchronizationMode" json:"synchronizationMode,omitempty"`
	// HashingAlgorithm specifies the content hashing algorithm used to track
	// content and perform differential transfers.
	HashingAlgorithm hashing.Algorithm `protobuf:"varint,17,opt,name=hashingAlgorithm,proto3,enum=hashing.Algorithm" json:"hashingAlgorithm,omitempty"`
	// MaximumEntryCount specifies the maximum number of filesystem entries that
	// endpoints will tolerate managing. A zero value indicates no limit.
	MaximumEntryCount uint64 `protobuf:"varint,12,opt,name=maximumEntryCount,proto3" json:"maximumEntryCount,omitempty"`
	// MaximumStagingFileSize is the maximum (individual) file size that
	// endpoints will stage. A zero value indicates no limit.
	MaximumStagingFileSize uint64 `protobuf:"varint,13,opt,name=maximumStagingFileSize,proto3" json:"maximumStagingFileSize,omitempty"`
	// ProbeMode specifies the filesystem probing mode.
	ProbeMode behavior.ProbeMode `protobuf:"varint,14,opt,name=probeMode,proto3,enum=behavior.ProbeMode" json:"probeMode,omitempty"`
	// ScanMode specifies the synchronization root scanning mode.
	ScanMode ScanMode `protobuf:"varint,15,opt,name=scanMode,proto3,enum=synchronization.ScanMode" json:"scanMode,omitempty"`
	// StageMode specifies the file staging mode.
	StageMode StageMode `protobuf:"varint,16,opt,name=stageMode,proto3,enum=synchronization.StageMode" json:"stageMode,omitempty"`
	// SymbolicLinkMode specifies the symbolic link mode.
	SymbolicLinkMode core.SymbolicLinkMode `protobuf:"varint,1,opt,name=symbolicLinkMode,proto3,enum=core.SymbolicLinkMode" json:"symbolicLinkMode,omitempty"`
	// WatchMode specifies the filesystem watching mode.
	WatchMode WatchMode `protobuf:"varint,21,opt,name=watchMode,proto3,enum=synchronization.WatchMode" json:"watchMode,omitempty"`
	// WatchPollingInterval specifies the interval (in seconds) for poll-based
	// file monitoring. A value of 0 specifies that the default interval should
	// be used.
	WatchPollingInterval uint32 `protobuf:"varint,22,opt,name=watchPollingInterval,proto3" json:"watchPollingInterval,omitempty"`
	// IgnoreSyntax specifies the syntax and semantics to use for ignores.
	// NOTE: This field is out of order due to the historical order in which it
	// was added.
	IgnoreSyntax ignore.Syntax `protobuf:"varint,34,opt,name=ignoreSyntax,proto3,enum=ignore.Syntax" json:"ignoreSyntax,omitempty"`
	// DefaultIgnores specifies the ignore patterns brought in from the global
	// configuration.
	// DEPRECATED: This field is no longer used when loading from global
	// configuration. Instead, ignores provided by global configuration are
	// simply merged into the ignore list of the main configuration. However,
	// older sessions still use this field.
	DefaultIgnores []string `protobuf:"bytes,31,rep,name=defaultIgnores,proto3" json:"defaultIgnores,omitempty"`
	// Ignores specifies the ignore patterns brought in from the create request.
	Ignores []string `protobuf:"bytes,32,rep,name=ignores,proto3" json:"ignores,omitempty"`
	// IgnoreVCSMode specifies the VCS ignore mode that should be used in
	// synchronization.
	IgnoreVCSMode ignore.IgnoreVCSMode `protobuf:"varint,33,opt,name=ignoreVCSMode,proto3,enum=ignore.IgnoreVCSMode" json:"ignoreVCSMode,omitempty"`
	// PermissionsMode species the manner in which permissions should be
	// propagated between endpoints.
	PermissionsMode core.PermissionsMode `protobuf:"varint,61,opt,name=permissionsMode,proto3,enum=core.PermissionsMode" json:"permissionsMode,omitempty"`
	// DefaultFileMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultFileMode uint32 `protobuf:"varint,63,opt,name=defaultFileMode,proto3" json:"defaultFileMode,omitempty"`
	// DefaultDirectoryMode specifies the default permission mode to use for new
	// files in "portable" permission propagation mode.
	DefaultDirectoryMode uint32 `protobuf:"varint,64,opt,name=defaultDirectoryMode,proto3" json:"defaultDirectoryMode,omitempty"`
	// DefaultOwner specifies the default owner identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultOwner string `protobuf:"bytes,65,opt,name=defaultOwner,proto3" json:"defaultOwner,omitempty"`
	// DefaultGroup specifies the default group identifier to use when setting
	// ownership of new files and directories in "portable" permission
	// propagation mode.
	DefaultGroup string `protobuf:"bytes,66,opt,name=defaultGroup,proto3" json:"defaultGroup,omitempty"`
	// CompressionAlgorithm specifies the compression algorithm to use when
	// communicating with the endpoint. This only applies to remote endpoints.
	CompressionAlgorithm compression.Algorithm `protobuf:"varint,81,opt,name=compressionAlgorithm,proto3,enum=compression.Algorithm" json:"compressionAlgorithm,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_synchronization_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetSynchronizationMode() core.SynchronizationMode {
	if x != nil {
		return x.SynchronizationMode
	}
	return core.SynchronizationMode(0)
}

func (x *Configuration) GetHashingAlgorithm() hashing.Algorithm {
	if x != nil {
		return x.HashingAlgorithm
	}
	return hashing.Algorithm(0)
}

func (x *Configuration) GetMaximumEntryCount() uint64 {
	if x != nil {
		return x.MaximumEntryCount
	}
	return 0
}

func (x *Configuration) GetMaximumStagingFileSize() uint64 {
	if x != nil {
		return x.MaximumStagingFileSize
	}
	return 0
}

func (x *Configuration) GetProbeMode() behavior.ProbeMode {
	if x != nil {
		return x.ProbeMode
	}
	return behavior.ProbeMode(0)
}

func (x *Configuration) GetScanMode() ScanMode {
	if x != nil {
		return x.ScanMode
	}
	return ScanMode_ScanModeDefault
}

func (x *Configuration) GetStageMode() StageMode {
	if x != nil {
		return x.StageMode
	}
	return StageMode_StageModeDefault
}

func (x *Configuration) GetSymbolicLinkMode() core.SymbolicLinkMode {
	if x != nil {
		return x.SymbolicLinkMode
	}
	return core.SymbolicLinkMode(0)
}

func (x *Configuration) GetWatchMode() WatchMode {
	if x != nil {
		return x.WatchMode
	}
	return WatchMode_WatchModeDefault
}

func (x *Configuration) GetWatchPollingInterval() uint32 {
	if x != nil {
		return x.WatchPollingInterval
	}
	return 0
}

func (x *Configuration) GetIgnoreSyntax() ignore.Syntax {
	if x != nil {
		return x.IgnoreSyntax
	}
	return ignore.Syntax(0)
}

func (x *Configuration) GetDefaultIgnores() []string {
	if x != nil {
		return x.DefaultIgnores
	}
	return nil
}

func (x *Configuration) GetIgnores() []string {
	if x != nil {
		return x.Ignores
	}
	return nil
}

func (x *Configuration) GetIgnoreVCSMode() ignore.IgnoreVCSMode {
	if x != nil {
		return x.IgnoreVCSMode
	}
	return ignore.IgnoreVCSMode(0)
}

func (x *Configuration) GetPermissionsMode() core.PermissionsMode {
	if x != nil {
		return x.PermissionsMode
	}
	return core.PermissionsMode(0)
}

func (x *Configuration) GetDefaultFileMode() uint32 {
	if x != nil {
		return x.DefaultFileMode
	}
	return 0
}

func (x *Configuration) GetDefaultDirectoryMode() uint32 {
	if x != nil {
		return x.DefaultDirectoryMode
	}
	return 0
}

func (x *Configuration) GetDefaultOwner() string {
	if x != nil {
		return x.DefaultOwner
	}
	return ""
}

func (x *Configuration) GetDefaultGroup() string {
	if x != nil {
		return x.DefaultGroup
	}
	return ""
}

func (x *Configuration) GetCompressionAlgorithm() compression.Algorithm {
	if x != nil {
		return x.CompressionAlgorithm
	}
	return compression.Algorithm(0)
}

var File_synchronization_configuration_proto protoreflect.FileDescriptor

var file_synchronization_configuration_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x24, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73,
	0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x73, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x63, 0x73, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbe, 0x08, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x13, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x13, 0x73, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x3e, 0x0a, 0x10, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x10,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36,
	0x0a, 0x16, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16,
	0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x63, 0x61,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x73, 0x63, 0x61, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x10, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x50, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x77, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x0c,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x74,
	0x61, 0x78, 0x52, 0x0c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78,
	0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x73, 0x18, 0x20, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x56, 0x43, 0x53, 0x4d,
	0x6f, 0x64, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x56, 0x43, 0x53, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0d, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x56, 0x43, 0x53, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x3f, 0x0a, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x18, 0x3f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x18, 0x40, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x41,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x42, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x51,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x14, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_configuration_proto_rawDescOnce sync.Once
	file_synchronization_configuration_proto_rawDescData = file_synchronization_configuration_proto_rawDesc
)

func file_synchronization_configuration_proto_rawDescGZIP() []byte {
	file_synchronization_configuration_proto_rawDescOnce.Do(func() {
		file_synchronization_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_configuration_proto_rawDescData)
	})
	return file_synchronization_configuration_proto_rawDescData
}

var file_synchronization_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_synchronization_configuration_proto_goTypes = []interface{}{
	(*Configuration)(nil),         // 0: synchronization.Configuration
	(core.SynchronizationMode)(0), // 1: core.SynchronizationMode
	(hashing.Algorithm)(0),        // 2: hashing.Algorithm
	(behavior.ProbeMode)(0),       // 3: behavior.ProbeMode
	(ScanMode)(0),                 // 4: synchronization.ScanMode
	(StageMode)(0),                // 5: synchronization.StageMode
	(core.SymbolicLinkMode)(0),    // 6: core.SymbolicLinkMode
	(WatchMode)(0),                // 7: synchronization.WatchMode
	(ignore.Syntax)(0),            // 8: ignore.Syntax
	(ignore.IgnoreVCSMode)(0),     // 9: ignore.IgnoreVCSMode
	(core.PermissionsMode)(0),     // 10: core.PermissionsMode
	(compression.Algorithm)(0),    // 11: compression.Algorithm
}
var file_synchronization_configuration_proto_depIdxs = []int32{
	1,  // 0: synchronization.Configuration.synchronizationMode:type_name -> core.SynchronizationMode
	2,  // 1: synchronization.Configuration.hashingAlgorithm:type_name -> hashing.Algorithm
	3,  // 2: synchronization.Configuration.probeMode:type_name -> behavior.ProbeMode
	4,  // 3: synchronization.Configuration.scanMode:type_name -> synchronization.ScanMode
	5,  // 4: synchronization.Configuration.stageMode:type_name -> synchronization.StageMode
	6,  // 5: synchronization.Configuration.symbolicLinkMode:type_name -> core.SymbolicLinkMode
	7,  // 6: synchronization.Configuration.watchMode:type_name -> synchronization.WatchMode
	8,  // 7: synchronization.Configuration.ignoreSyntax:type_name -> ignore.Syntax
	9,  // 8: synchronization.Configuration.ignoreVCSMode:type_name -> ignore.IgnoreVCSMode
	10, // 9: synchronization.Configuration.permissionsMode:type_name -> core.PermissionsMode
	11, // 10: synchronization.Configuration.compressionAlgorithm:type_name -> compression.Algorithm
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_synchronization_configuration_proto_init() }
func file_synchronization_configuration_proto_init() {
	if File_synchronization_configuration_proto != nil {
		return
	}
	file_synchronization_scan_mode_proto_init()
	file_synchronization_stage_mode_proto_init()
	file_synchronization_watch_mode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_synchronization_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
			RawDescriptor: file_synchronization_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_configuration_proto_goTypes,
		DependencyIndexes: file_synchronization_configuration_proto_depIdxs,
		MessageInfos:      file_synchronization_configuration_proto_msgTypes,
	}.Build()
	File_synchronization_configuration_proto = out.File
	file_synchronization_configuration_proto_rawDesc = nil
	file_synchronization_configuration_proto_goTypes = nil
	file_synchronization_configuration_proto_depIdxs = nil
}
