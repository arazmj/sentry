// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/proto/sentry.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartJobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Command       string                 `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	CommandArgs   []string               `protobuf:"bytes,2,rep,name=commandArgs,proto3" json:"commandArgs,omitempty"`
	MemoryLimit   string                 `protobuf:"bytes,3,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	CpuLimit      string                 `protobuf:"bytes,4,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	Mount         string                 `protobuf:"bytes,5,opt,name=mount,proto3" json:"mount,omitempty"`
	WriteBps      string                 `protobuf:"bytes,6,opt,name=write_bps,json=writeBps,proto3" json:"write_bps,omitempty"`
	ReadBps       string                 `protobuf:"bytes,7,opt,name=read_bps,json=readBps,proto3" json:"read_bps,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartJobRequest) Reset() {
	*x = StartJobRequest{}
	mi := &file_api_proto_sentry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobRequest) ProtoMessage() {}

func (x *StartJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobRequest.ProtoReflect.Descriptor instead.
func (*StartJobRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{0}
}

func (x *StartJobRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *StartJobRequest) GetCommandArgs() []string {
	if x != nil {
		return x.CommandArgs
	}
	return nil
}

func (x *StartJobRequest) GetMemoryLimit() string {
	if x != nil {
		return x.MemoryLimit
	}
	return ""
}

func (x *StartJobRequest) GetCpuLimit() string {
	if x != nil {
		return x.CpuLimit
	}
	return ""
}

func (x *StartJobRequest) GetMount() string {
	if x != nil {
		return x.Mount
	}
	return ""
}

func (x *StartJobRequest) GetWriteBps() string {
	if x != nil {
		return x.WriteBps
	}
	return ""
}

func (x *StartJobRequest) GetReadBps() string {
	if x != nil {
		return x.ReadBps
	}
	return ""
}

type JobOutput struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	IsStderr      bool                   `protobuf:"varint,3,opt,name=is_stderr,json=isStderr,proto3" json:"is_stderr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobOutput) Reset() {
	*x = JobOutput{}
	mi := &file_api_proto_sentry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobOutput) ProtoMessage() {}

func (x *JobOutput) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobOutput.ProtoReflect.Descriptor instead.
func (*JobOutput) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{1}
}

func (x *JobOutput) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobOutput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *JobOutput) GetIsStderr() bool {
	if x != nil {
		return x.IsStderr
	}
	return false
}

type StartJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartJobResponse) Reset() {
	*x = StartJobResponse{}
	mi := &file_api_proto_sentry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobResponse) ProtoMessage() {}

func (x *StartJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobResponse.ProtoReflect.Descriptor instead.
func (*StartJobResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{2}
}

func (x *StartJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type StopJobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopJobRequest) Reset() {
	*x = StopJobRequest{}
	mi := &file_api_proto_sentry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobRequest) ProtoMessage() {}

func (x *StopJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobRequest.ProtoReflect.Descriptor instead.
func (*StopJobRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{3}
}

func (x *StopJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type StopJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopJobResponse) Reset() {
	*x = StopJobResponse{}
	mi := &file_api_proto_sentry_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobResponse) ProtoMessage() {}

func (x *StopJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobResponse.ProtoReflect.Descriptor instead.
func (*StopJobResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{4}
}

func (x *StopJobResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *StopJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JobStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobStatusRequest) Reset() {
	*x = JobStatusRequest{}
	mi := &file_api_proto_sentry_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusRequest) ProtoMessage() {}

func (x *JobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusRequest.ProtoReflect.Descriptor instead.
func (*JobStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{5}
}

func (x *JobStatusRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type JobStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsRunning     bool                   `protobuf:"varint,1,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobStatusResponse) Reset() {
	*x = JobStatusResponse{}
	mi := &file_api_proto_sentry_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusResponse) ProtoMessage() {}

func (x *JobStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusResponse.ProtoReflect.Descriptor instead.
func (*JobStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{6}
}

func (x *JobStatusResponse) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

type JobLogsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobLogsRequest) Reset() {
	*x = JobLogsRequest{}
	mi := &file_api_proto_sentry_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogsRequest) ProtoMessage() {}

func (x *JobLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogsRequest.ProtoReflect.Descriptor instead.
func (*JobLogsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{7}
}

func (x *JobLogsRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type JobLogsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Logs          []byte                 `protobuf:"bytes,1,opt,name=logs,proto3" json:"logs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobLogsResponse) Reset() {
	*x = JobLogsResponse{}
	mi := &file_api_proto_sentry_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogsResponse) ProtoMessage() {}

func (x *JobLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogsResponse.ProtoReflect.Descriptor instead.
func (*JobLogsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{8}
}

func (x *JobLogsResponse) GetLogs() []byte {
	if x != nil {
		return x.Logs
	}
	return nil
}

type ListJobsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	mi := &file_api_proto_sentry_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{9}
}

type JobInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Command       string                 `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	IsRunning     bool                   `protobuf:"varint,3,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
	MemoryLimit   string                 `protobuf:"bytes,4,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	CpuLimit      string                 `protobuf:"bytes,5,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	Mount         string                 `protobuf:"bytes,6,opt,name=mount,proto3" json:"mount,omitempty"`
	WriteBps      string                 `protobuf:"bytes,7,opt,name=write_bps,json=writeBps,proto3" json:"write_bps,omitempty"`
	ReadBps       string                 `protobuf:"bytes,8,opt,name=read_bps,json=readBps,proto3" json:"read_bps,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobInfo) Reset() {
	*x = JobInfo{}
	mi := &file_api_proto_sentry_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobInfo) ProtoMessage() {}

func (x *JobInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobInfo.ProtoReflect.Descriptor instead.
func (*JobInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{10}
}

func (x *JobInfo) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobInfo) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *JobInfo) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

func (x *JobInfo) GetMemoryLimit() string {
	if x != nil {
		return x.MemoryLimit
	}
	return ""
}

func (x *JobInfo) GetCpuLimit() string {
	if x != nil {
		return x.CpuLimit
	}
	return ""
}

func (x *JobInfo) GetMount() string {
	if x != nil {
		return x.Mount
	}
	return ""
}

func (x *JobInfo) GetWriteBps() string {
	if x != nil {
		return x.WriteBps
	}
	return ""
}

func (x *JobInfo) GetReadBps() string {
	if x != nil {
		return x.ReadBps
	}
	return ""
}

type ListJobsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Jobs          []*JobInfo             `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	mi := &file_api_proto_sentry_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{11}
}

func (x *ListJobsResponse) GetJobs() []*JobInfo {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type KillJobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KillJobRequest) Reset() {
	*x = KillJobRequest{}
	mi := &file_api_proto_sentry_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KillJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillJobRequest) ProtoMessage() {}

func (x *KillJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillJobRequest.ProtoReflect.Descriptor instead.
func (*KillJobRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{12}
}

func (x *KillJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type KillJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KillJobResponse) Reset() {
	*x = KillJobResponse{}
	mi := &file_api_proto_sentry_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KillJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillJobResponse) ProtoMessage() {}

func (x *KillJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_sentry_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillJobResponse.ProtoReflect.Descriptor instead.
func (*KillJobResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_sentry_proto_rawDescGZIP(), []int{13}
}

func (x *KillJobResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *KillJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_proto_sentry_proto protoreflect.FileDescriptor

var file_api_proto_sentry_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x22, 0xdb, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x62, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x42, 0x70, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x70, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x42, 0x70, 0x73, 0x22, 0x53,
	0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x22, 0x29, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x27,
	0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29,
	0x0a, 0x10, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x4a, 0x6f, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x27, 0x0a,
	0x0e, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x11, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xe7, 0x01, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x70, 0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x70, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x70, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x70, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x42, 0x70, 0x73, 0x22, 0x37, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6a,
	0x6f, 0x62, 0x73, 0x22, 0x27, 0x0a, 0x0e, 0x4b, 0x69, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0f,
	0x4b, 0x69, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xd6, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f,
	0x62, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x4b, 0x69, 0x6c, 0x6c, 0x4a, 0x6f,
	0x62, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x73,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4a, 0x6f,
	0x62, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f,
	0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_proto_sentry_proto_rawDescOnce sync.Once
	file_api_proto_sentry_proto_rawDescData []byte
)

func file_api_proto_sentry_proto_rawDescGZIP() []byte {
	file_api_proto_sentry_proto_rawDescOnce.Do(func() {
		file_api_proto_sentry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_sentry_proto_rawDesc), len(file_api_proto_sentry_proto_rawDesc)))
	})
	return file_api_proto_sentry_proto_rawDescData
}

var file_api_proto_sentry_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_proto_sentry_proto_goTypes = []any{
	(*StartJobRequest)(nil),   // 0: sentry.StartJobRequest
	(*JobOutput)(nil),         // 1: sentry.JobOutput
	(*StartJobResponse)(nil),  // 2: sentry.StartJobResponse
	(*StopJobRequest)(nil),    // 3: sentry.StopJobRequest
	(*StopJobResponse)(nil),   // 4: sentry.StopJobResponse
	(*JobStatusRequest)(nil),  // 5: sentry.JobStatusRequest
	(*JobStatusResponse)(nil), // 6: sentry.JobStatusResponse
	(*JobLogsRequest)(nil),    // 7: sentry.JobLogsRequest
	(*JobLogsResponse)(nil),   // 8: sentry.JobLogsResponse
	(*ListJobsRequest)(nil),   // 9: sentry.ListJobsRequest
	(*JobInfo)(nil),           // 10: sentry.JobInfo
	(*ListJobsResponse)(nil),  // 11: sentry.ListJobsResponse
	(*KillJobRequest)(nil),    // 12: sentry.KillJobRequest
	(*KillJobResponse)(nil),   // 13: sentry.KillJobResponse
}
var file_api_proto_sentry_proto_depIdxs = []int32{
	10, // 0: sentry.ListJobsResponse.jobs:type_name -> sentry.JobInfo
	0,  // 1: sentry.SentryService.StartJob:input_type -> sentry.StartJobRequest
	12, // 2: sentry.SentryService.KillJob:input_type -> sentry.KillJobRequest
	5,  // 3: sentry.SentryService.GetJobStatus:input_type -> sentry.JobStatusRequest
	7,  // 4: sentry.SentryService.StreamJobLogs:input_type -> sentry.JobLogsRequest
	9,  // 5: sentry.SentryService.ListJobs:input_type -> sentry.ListJobsRequest
	2,  // 6: sentry.SentryService.StartJob:output_type -> sentry.StartJobResponse
	13, // 7: sentry.SentryService.KillJob:output_type -> sentry.KillJobResponse
	6,  // 8: sentry.SentryService.GetJobStatus:output_type -> sentry.JobStatusResponse
	1,  // 9: sentry.SentryService.StreamJobLogs:output_type -> sentry.JobOutput
	11, // 10: sentry.SentryService.ListJobs:output_type -> sentry.ListJobsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_sentry_proto_init() }
func file_api_proto_sentry_proto_init() {
	if File_api_proto_sentry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_sentry_proto_rawDesc), len(file_api_proto_sentry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_sentry_proto_goTypes,
		DependencyIndexes: file_api_proto_sentry_proto_depIdxs,
		MessageInfos:      file_api_proto_sentry_proto_msgTypes,
	}.Build()
	File_api_proto_sentry_proto = out.File
	file_api_proto_sentry_proto_goTypes = nil
	file_api_proto_sentry_proto_depIdxs = nil
}
