//*
// Scheduler
//
// Create and manage scheduling jobs.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.4
// source: ct/scheduler/scheduler.proto

package ct

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobStatus int32

const (
	JobStatus_JOB_STATUS_NONE JobStatus = 0
	// A job has been successfully created.
	JobStatus_JOB_STATUS_CREATED JobStatus = 1
	// A job has a next run datetime and in standby mode.
	JobStatus_JOB_STATUS_READY JobStatus = 2
	// When a job is flagged with inflight, the job has been dispatched and in process.
	JobStatus_JOB_STATUS_INFLIGHT JobStatus = 3
	// A job has been completed and waiting to be assigned the next run datetime.
	JobStatus_JOB_STATUS_SUCCEEDED JobStatus = 4
	// A job has been failed. The history logging is available for the details.
	JobStatus_JOB_STATUS_FAILED JobStatus = 5
	// A job has been paused until resumed datetime or when resumed.
	JobStatus_JOB_STATUS_PAUSED JobStatus = 6
	// A job has been expired and ready to be deleted. Expired jobs will be deleted.
	JobStatus_JOB_STATUS_EXPIRED JobStatus = 7
	// A job has been deleted and the SchedulingJob record does not exist.
	JobStatus_JOB_STATUS_DELETED JobStatus = 8
)

// Enum value maps for JobStatus.
var (
	JobStatus_name = map[int32]string{
		0: "JOB_STATUS_NONE",
		1: "JOB_STATUS_CREATED",
		2: "JOB_STATUS_READY",
		3: "JOB_STATUS_INFLIGHT",
		4: "JOB_STATUS_SUCCEEDED",
		5: "JOB_STATUS_FAILED",
		6: "JOB_STATUS_PAUSED",
		7: "JOB_STATUS_EXPIRED",
		8: "JOB_STATUS_DELETED",
	}
	JobStatus_value = map[string]int32{
		"JOB_STATUS_NONE":      0,
		"JOB_STATUS_CREATED":   1,
		"JOB_STATUS_READY":     2,
		"JOB_STATUS_INFLIGHT":  3,
		"JOB_STATUS_SUCCEEDED": 4,
		"JOB_STATUS_FAILED":    5,
		"JOB_STATUS_PAUSED":    6,
		"JOB_STATUS_EXPIRED":   7,
		"JOB_STATUS_DELETED":   8,
	}
)

func (x JobStatus) Enum() *JobStatus {
	p := new(JobStatus)
	*p = x
	return p
}

func (x JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ct_scheduler_scheduler_proto_enumTypes[0].Descriptor()
}

func (JobStatus) Type() protoreflect.EnumType {
	return &file_ct_scheduler_scheduler_proto_enumTypes[0]
}

func (x JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatus.Descriptor instead.
func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_ct_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

type SchedulingJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auto generated 22 char identifier. Read only.
	// @tag: validateCreate:"-" validateUpdate:"required" validateGeneric:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"-" validateUpdate:"required" validateGeneric:"required"`
	// HttpMethod and Endpoint with comma separation.
	// Eg. POST,https://api.pub1.passkit.io/members/count/05QFES9HNRMZb1gwcsepww
	// @tag: validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"
	JobFunction string `protobuf:"bytes,2,opt,name=jobFunction,proto3" json:"jobFunction,omitempty" validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"`
	// Json string of the job function payload.
	// @tag: validateCreate:"omitempty" validateUpdate:"omitempty" validateGeneric:"omitempty"
	JobPayload string `protobuf:"bytes,3,opt,name=jobPayload,proto3" json:"jobPayload,omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateGeneric:"omitempty"`
	// Job name.
	// @tag: validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"`
	// Job description.
	// @tag: validateCreate:"omitempty" validateUpdate:"omitempty" validateGeneric:"omitempty"
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateGeneric:"omitempty"`
	// Job schedule. You can set recursive or one-off job.
	// @tag: validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"
	Schedule *Schedule `protobuf:"bytes,6,opt,name=schedule,proto3" json:"schedule,omitempty" validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"`
	// Job status. Writable values are JOB_STATUS_READY and JOB_STATUS_PAUSED only. Default is JOB_STATUS_READY.
	// Setting status=JOB_STATUS_PAUSED sets nextRunDatetime=null only when Schedule.StartsAt is not set.
	// Status will be always JOB_STATUS_READY, if you set Schedule.StartsAt.
	// @tag: validateCreate:"omitempty" validateUpdate:"omitempty" validateGeneric:"omitempty"
	Status JobStatus `protobuf:"varint,7,opt,name=status,proto3,enum=ct.JobStatus" json:"status,omitempty" validateCreate:"omitempty" validateUpdate:"omitempty" validateGeneric:"omitempty"`
	// Next scheduled run datetime. Read only.
	// @tag: validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"
	NextRunDatetime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=nextRunDatetime,proto3" json:"nextRunDatetime,omitempty" validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"`
	// The datetime when the job was executed last time. Read only.
	// @tag: validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"
	LastRunDatetime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=lastRunDatetime,proto3" json:"lastRunDatetime,omitempty" validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"`
	// 5 latest logs of executed jobs. Read only.
	// @tag: validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"
	Logs []*JobHistory `protobuf:"bytes,10,rep,name=logs,proto3" json:"logs,omitempty" validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"`
	// Datetime the job was created. Read only.
	// @tag: validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"
	Created *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created,proto3" json:"created,omitempty" validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"`
	// Datetime the job was last updated. Read only.
	// @tag: validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"
	Updated *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated,proto3" json:"updated,omitempty" validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"`
	// Datetime when the job will for the last time and expires. Read only. Expiry setting is available in the Schedule field.
	// @tag: validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"
	Expiry *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=expiry,proto3" json:"expiry,omitempty" validateCreate:"isdefault" validateUpdate:"isdefault" validateGeneric:"isdefault"`
}

func (x *SchedulingJob) Reset() {
	*x = SchedulingJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ct_scheduler_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulingJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulingJob) ProtoMessage() {}

func (x *SchedulingJob) ProtoReflect() protoreflect.Message {
	mi := &file_ct_scheduler_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulingJob.ProtoReflect.Descriptor instead.
func (*SchedulingJob) Descriptor() ([]byte, []int) {
	return file_ct_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *SchedulingJob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SchedulingJob) GetJobFunction() string {
	if x != nil {
		return x.JobFunction
	}
	return ""
}

func (x *SchedulingJob) GetJobPayload() string {
	if x != nil {
		return x.JobPayload
	}
	return ""
}

func (x *SchedulingJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SchedulingJob) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SchedulingJob) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *SchedulingJob) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_JOB_STATUS_NONE
}

func (x *SchedulingJob) GetNextRunDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.NextRunDatetime
	}
	return nil
}

func (x *SchedulingJob) GetLastRunDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRunDatetime
	}
	return nil
}

func (x *SchedulingJob) GetLogs() []*JobHistory {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *SchedulingJob) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *SchedulingJob) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *SchedulingJob) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Crontab schedule expression. e.g. '0 * * * *', '@hourly'
	// For month end job, '@monthend' is available.
	// To run a one-off job now, '@now' is available.
	// Crontab generator tool https://crontab.guru
	// Reference https://en.wikipedia.org/wiki/Cron
	// @tag: validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"
	Schedule string `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty" validateCreate:"required" validateUpdate:"required" validateGeneric:"omitempty"`
	// Iana timezone. Default is UTC.
	Timezone string `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// Datetime to start the first job or resume the paused job.
	StartsAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=startsAt,proto3" json:"startsAt,omitempty"`
	// Types that are assignable to Expiry:
	//
	//	*Schedule_RepeatCount
	//	*Schedule_ExpiryDate
	Expiry isSchedule_Expiry `protobuf_oneof:"expiry"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ct_scheduler_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_ct_scheduler_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_ct_scheduler_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *Schedule) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Schedule) GetStartsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartsAt
	}
	return nil
}

func (m *Schedule) GetExpiry() isSchedule_Expiry {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func (x *Schedule) GetRepeatCount() int32 {
	if x, ok := x.GetExpiry().(*Schedule_RepeatCount); ok {
		return x.RepeatCount
	}
	return 0
}

func (x *Schedule) GetExpiryDate() *timestamppb.Timestamp {
	if x, ok := x.GetExpiry().(*Schedule_ExpiryDate); ok {
		return x.ExpiryDate
	}
	return nil
}

type isSchedule_Expiry interface {
	isSchedule_Expiry()
}

type Schedule_RepeatCount struct {
	// Set repeat count to repeat the job for certain times (eg. 3 will repeat the job for 3 times).
	// If job is one off, set -1.
	RepeatCount int32 `protobuf:"zigzag32,15,opt,name=repeatCount,proto3,oneof"`
}

type Schedule_ExpiryDate struct {
	// The date which job will be expired.
	ExpiryDate *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=expiryDate,proto3,oneof"`
}

func (*Schedule_RepeatCount) isSchedule_Expiry() {}

func (*Schedule_ExpiryDate) isSchedule_Expiry() {}

type SchedulingJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auto generated 22 char identifier. Read only.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Job status. Writable values are JOB_STATUS_READY and JOB_STATUS_PAUSED only.
	Status JobStatus `protobuf:"varint,2,opt,name=status,proto3,enum=ct.JobStatus" json:"status,omitempty"`
	// Next scheduled run datetime. Read only.
	NextRunDatetime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=nextRunDatetime,proto3" json:"nextRunDatetime,omitempty"`
	// Datetime when the job will for the last time and expires. Read only. Expiry setting is available in the Schedule field.
	Expiry *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *SchedulingJobResponse) Reset() {
	*x = SchedulingJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ct_scheduler_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulingJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulingJobResponse) ProtoMessage() {}

func (x *SchedulingJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ct_scheduler_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulingJobResponse.ProtoReflect.Descriptor instead.
func (*SchedulingJobResponse) Descriptor() ([]byte, []int) {
	return file_ct_scheduler_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *SchedulingJobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SchedulingJobResponse) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_JOB_STATUS_NONE
}

func (x *SchedulingJobResponse) GetNextRunDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.NextRunDatetime
	}
	return nil
}

func (x *SchedulingJobResponse) GetExpiry() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

type JobHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Job history id.
	// @tag: validateCreate:"-"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validateCreate:"-"`
	// Job ID.
	// @tag: validateCreate:"required"
	JobId string `protobuf:"bytes,2,opt,name=jobId,proto3" json:"jobId,omitempty" validateCreate:"required"`
	// True if the job has succeeded. False if the job has failed.
	// @tag: validateCreate:"omitempty"
	Success bool `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty" validateCreate:"omitempty"`
	// The job output log.
	// @tag: validateCreate:"required"
	Log string `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" validateCreate:"required"`
	// The status code.
	// @tag: validateCreate:"omitempty"
	StatusCode uint32 `protobuf:"varint,5,opt,name=statusCode,proto3" json:"statusCode,omitempty" validateCreate:"omitempty"`
	// Scheduled run datetime for the job.
	// @tag: validateCreate:"omitempty"
	ScheduledRunDateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=scheduledRunDateTime,proto3" json:"scheduledRunDateTime,omitempty" validateCreate:"omitempty"`
	// Job completion datetime.
	// @tag: validateCreate:"required"
	CompletedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=completedAt,proto3" json:"completedAt,omitempty" validateCreate:"required"`
}

func (x *JobHistory) Reset() {
	*x = JobHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ct_scheduler_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobHistory) ProtoMessage() {}

func (x *JobHistory) ProtoReflect() protoreflect.Message {
	mi := &file_ct_scheduler_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobHistory.ProtoReflect.Descriptor instead.
func (*JobHistory) Descriptor() ([]byte, []int) {
	return file_ct_scheduler_scheduler_proto_rawDescGZIP(), []int{3}
}

func (x *JobHistory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobHistory) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobHistory) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *JobHistory) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *JobHistory) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *JobHistory) GetScheduledRunDateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledRunDateTime
	}
	return nil
}

func (x *JobHistory) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

var File_ct_scheduler_scheduler_proto protoreflect.FileDescriptor

var file_ct_scheduler_scheduler_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x74, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x63, 0x74, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x05, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69,
	0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x62,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x74, 0x2e, 0x4a, 0x6f, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x74, 0x2e, 0x4a, 0x6f,
	0x62, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x34,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x3a, 0x50,
	0x92, 0x41, 0x4d, 0x0a, 0x4b, 0x2a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e,
	0x67, 0x20, 0x4a, 0x6f, 0x62, 0x32, 0x19, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65,
	0x20, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x65, 0x2d, 0x6f, 0x66, 0x66, 0x20, 0x6a, 0x6f, 0x62, 0x2e,
	0xd2, 0x01, 0x0b, 0x6a, 0x6f, 0x62, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xd2, 0x01,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x22, 0xc8, 0x02, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x12, 0x22, 0x0a,
	0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x11, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x3a,
	0x60, 0x92, 0x41, 0x5d, 0x0a, 0x5b, 0x2a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x32, 0x44, 0x53, 0x65, 0x74, 0x20, 0x61, 0x20, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6a, 0x6f, 0x62, 0x2e, 0x20, 0x53, 0x65,
	0x74, 0x20, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x74, 0x6f,
	0x20, 0x2d, 0x31, 0x20, 0x69, 0x66, 0x20, 0x6a, 0x6f, 0x62, 0x20, 0x69, 0x73, 0x20, 0x6f, 0x6e,
	0x65, 0x2d, 0x6f, 0x66, 0x66, 0x2e, 0xd2, 0x01, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0xa1, 0x02, 0x0a, 0x15,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x3a, 0x57, 0x92, 0x41, 0x54, 0x0a, 0x52, 0x2a, 0x17, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x4a, 0x6f, 0x62, 0x20, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x37, 0x41, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x6a, 0x6f, 0x62, 0x2e, 0x22,
	0xd7, 0x02, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x4e, 0x0a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x52, 0x75, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x49,
	0x92, 0x41, 0x46, 0x0a, 0x44, 0x2a, 0x0b, 0x4a, 0x6f, 0x62, 0x20, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x32, 0x19, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x20, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0xd2, 0x01, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0xd2, 0x01, 0x03, 0x6c, 0x6f, 0x67, 0xd2, 0x01, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0xdf, 0x01, 0x0a, 0x09, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x4a, 0x4f, 0x42, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f,
	0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x46, 0x4c, 0x49, 0x47, 0x48,
	0x54, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a,
	0x11, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x4a,
	0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x08, 0x42, 0x4d, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x74, 0x5a, 0x24, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x74, 0xaa, 0x02, 0x0f, 0x50, 0x61, 0x73, 0x73, 0x4b,
	0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ct_scheduler_scheduler_proto_rawDescOnce sync.Once
	file_ct_scheduler_scheduler_proto_rawDescData = file_ct_scheduler_scheduler_proto_rawDesc
)

func file_ct_scheduler_scheduler_proto_rawDescGZIP() []byte {
	file_ct_scheduler_scheduler_proto_rawDescOnce.Do(func() {
		file_ct_scheduler_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_ct_scheduler_scheduler_proto_rawDescData)
	})
	return file_ct_scheduler_scheduler_proto_rawDescData
}

var file_ct_scheduler_scheduler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ct_scheduler_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ct_scheduler_scheduler_proto_goTypes = []interface{}{
	(JobStatus)(0),                // 0: ct.JobStatus
	(*SchedulingJob)(nil),         // 1: ct.SchedulingJob
	(*Schedule)(nil),              // 2: ct.Schedule
	(*SchedulingJobResponse)(nil), // 3: ct.SchedulingJobResponse
	(*JobHistory)(nil),            // 4: ct.JobHistory
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_ct_scheduler_scheduler_proto_depIdxs = []int32{
	2,  // 0: ct.SchedulingJob.schedule:type_name -> ct.Schedule
	0,  // 1: ct.SchedulingJob.status:type_name -> ct.JobStatus
	5,  // 2: ct.SchedulingJob.nextRunDatetime:type_name -> google.protobuf.Timestamp
	5,  // 3: ct.SchedulingJob.lastRunDatetime:type_name -> google.protobuf.Timestamp
	4,  // 4: ct.SchedulingJob.logs:type_name -> ct.JobHistory
	5,  // 5: ct.SchedulingJob.created:type_name -> google.protobuf.Timestamp
	5,  // 6: ct.SchedulingJob.updated:type_name -> google.protobuf.Timestamp
	5,  // 7: ct.SchedulingJob.expiry:type_name -> google.protobuf.Timestamp
	5,  // 8: ct.Schedule.startsAt:type_name -> google.protobuf.Timestamp
	5,  // 9: ct.Schedule.expiryDate:type_name -> google.protobuf.Timestamp
	0,  // 10: ct.SchedulingJobResponse.status:type_name -> ct.JobStatus
	5,  // 11: ct.SchedulingJobResponse.nextRunDatetime:type_name -> google.protobuf.Timestamp
	5,  // 12: ct.SchedulingJobResponse.expiry:type_name -> google.protobuf.Timestamp
	5,  // 13: ct.JobHistory.scheduledRunDateTime:type_name -> google.protobuf.Timestamp
	5,  // 14: ct.JobHistory.completedAt:type_name -> google.protobuf.Timestamp
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_ct_scheduler_scheduler_proto_init() }
func file_ct_scheduler_scheduler_proto_init() {
	if File_ct_scheduler_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ct_scheduler_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulingJob); i {
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
		file_ct_scheduler_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_ct_scheduler_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulingJobResponse); i {
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
		file_ct_scheduler_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobHistory); i {
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
	file_ct_scheduler_scheduler_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Schedule_RepeatCount)(nil),
		(*Schedule_ExpiryDate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ct_scheduler_scheduler_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ct_scheduler_scheduler_proto_goTypes,
		DependencyIndexes: file_ct_scheduler_scheduler_proto_depIdxs,
		EnumInfos:         file_ct_scheduler_scheduler_proto_enumTypes,
		MessageInfos:      file_ct_scheduler_scheduler_proto_msgTypes,
	}.Build()
	File_ct_scheduler_scheduler_proto = out.File
	file_ct_scheduler_scheduler_proto_rawDesc = nil
	file_ct_scheduler_scheduler_proto_goTypes = nil
	file_ct_scheduler_scheduler_proto_depIdxs = nil
}
