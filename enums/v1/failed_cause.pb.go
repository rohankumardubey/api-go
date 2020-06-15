// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/enums/v1/failed_cause.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DecisionTaskFailedCause int32

const (
	DECISION_TASK_FAILED_CAUSE_UNSPECIFIED                                               DecisionTaskFailedCause = 0
	DECISION_TASK_FAILED_CAUSE_UNHANDLED_DECISION                                        DecisionTaskFailedCause = 1
	DECISION_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES                          DecisionTaskFailedCause = 2
	DECISION_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES                    DecisionTaskFailedCause = 3
	DECISION_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES                                DecisionTaskFailedCause = 4
	DECISION_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES                               DecisionTaskFailedCause = 5
	DECISION_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES                              DecisionTaskFailedCause = 6
	DECISION_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES                DecisionTaskFailedCause = 7
	DECISION_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES                    DecisionTaskFailedCause = 8
	DECISION_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES                  DecisionTaskFailedCause = 9
	DECISION_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES DecisionTaskFailedCause = 10
	DECISION_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES                            DecisionTaskFailedCause = 11
	DECISION_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID                                  DecisionTaskFailedCause = 12
	DECISION_TASK_FAILED_CAUSE_RESET_STICKY_TASKLIST                                     DecisionTaskFailedCause = 13
	DECISION_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE                         DecisionTaskFailedCause = 14
	DECISION_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES                  DecisionTaskFailedCause = 15
	DECISION_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES                      DecisionTaskFailedCause = 16
	DECISION_TASK_FAILED_CAUSE_FORCE_CLOSE_DECISION                                      DecisionTaskFailedCause = 17
	DECISION_TASK_FAILED_CAUSE_FAILOVER_CLOSE_DECISION                                   DecisionTaskFailedCause = 18
	DECISION_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE                                     DecisionTaskFailedCause = 19
	DECISION_TASK_FAILED_CAUSE_RESET_WORKFLOW                                            DecisionTaskFailedCause = 20
	DECISION_TASK_FAILED_CAUSE_BAD_BINARY                                                DecisionTaskFailedCause = 21
	DECISION_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID                            DecisionTaskFailedCause = 22
	DECISION_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES                                     DecisionTaskFailedCause = 23
)

var DecisionTaskFailedCause_name = map[int32]string{
	0:  "Unspecified",
	1:  "UnhandledDecision",
	2:  "BadScheduleActivityAttributes",
	3:  "BadRequestCancelActivityAttributes",
	4:  "BadStartTimerAttributes",
	5:  "BadCancelTimerAttributes",
	6:  "BadRecordMarkerAttributes",
	7:  "BadCompleteWorkflowExecutionAttributes",
	8:  "BadFailWorkflowExecutionAttributes",
	9:  "BadCancelWorkflowExecutionAttributes",
	10: "BadRequestCancelExternalWorkflowExecutionAttributes",
	11: "BadContinueAsNewAttributes",
	12: "StartTimerDuplicateId",
	13: "ResetStickyTasklist",
	14: "WorkflowWorkerUnhandledFailure",
	15: "BadSignalWorkflowExecutionAttributes",
	16: "BadStartChildExecutionAttributes",
	17: "ForceCloseDecision",
	18: "FailoverCloseDecision",
	19: "BadSignalInputSize",
	20: "ResetWorkflow",
	21: "BadBinary",
	22: "ScheduleActivityDuplicateId",
	23: "BadSearchAttributes",
}

var DecisionTaskFailedCause_value = map[string]int32{
	"Unspecified":                                         0,
	"UnhandledDecision":                                   1,
	"BadScheduleActivityAttributes":                       2,
	"BadRequestCancelActivityAttributes":                  3,
	"BadStartTimerAttributes":                             4,
	"BadCancelTimerAttributes":                            5,
	"BadRecordMarkerAttributes":                           6,
	"BadCompleteWorkflowExecutionAttributes":              7,
	"BadFailWorkflowExecutionAttributes":                  8,
	"BadCancelWorkflowExecutionAttributes":                9,
	"BadRequestCancelExternalWorkflowExecutionAttributes": 10,
	"BadContinueAsNewAttributes":                          11,
	"StartTimerDuplicateId":                               12,
	"ResetStickyTasklist":                                 13,
	"WorkflowWorkerUnhandledFailure":                      14,
	"BadSignalWorkflowExecutionAttributes":                15,
	"BadStartChildExecutionAttributes":                    16,
	"ForceCloseDecision":                                  17,
	"FailoverCloseDecision":                               18,
	"BadSignalInputSize":                                  19,
	"ResetWorkflow":                                       20,
	"BadBinary":                                           21,
	"ScheduleActivityDuplicateId":                         22,
	"BadSearchAttributes":                                 23,
}

func (DecisionTaskFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6aa03c6c7587cda2, []int{0}
}

type StartChildWorkflowExecutionFailedCause int32

const (
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED             StartChildWorkflowExecutionFailedCause = 0
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS StartChildWorkflowExecutionFailedCause = 1
)

var StartChildWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "WorkflowAlreadyExists",
}

var StartChildWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":           0,
	"WorkflowAlreadyExists": 1,
}

func (StartChildWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6aa03c6c7587cda2, []int{1}
}

type CancelExternalWorkflowExecutionFailedCause int32

const (
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           CancelExternalWorkflowExecutionFailedCause = 0
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND CancelExternalWorkflowExecutionFailedCause = 1
)

var CancelExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
}

var CancelExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
}

func (CancelExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6aa03c6c7587cda2, []int{2}
}

type SignalExternalWorkflowExecutionFailedCause int32

const (
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           SignalExternalWorkflowExecutionFailedCause = 0
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND SignalExternalWorkflowExecutionFailedCause = 1
)

var SignalExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
}

var SignalExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
}

func (SignalExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6aa03c6c7587cda2, []int{3}
}

func init() {
	proto.RegisterEnum("temporal.enums.v1.DecisionTaskFailedCause", DecisionTaskFailedCause_name, DecisionTaskFailedCause_value)
	proto.RegisterEnum("temporal.enums.v1.StartChildWorkflowExecutionFailedCause", StartChildWorkflowExecutionFailedCause_name, StartChildWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.enums.v1.CancelExternalWorkflowExecutionFailedCause", CancelExternalWorkflowExecutionFailedCause_name, CancelExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.enums.v1.SignalExternalWorkflowExecutionFailedCause", SignalExternalWorkflowExecutionFailedCause_name, SignalExternalWorkflowExecutionFailedCause_value)
}

func init() {
	proto.RegisterFile("temporal/enums/v1/failed_cause.proto", fileDescriptor_6aa03c6c7587cda2)
}

var fileDescriptor_6aa03c6c7587cda2 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xcf, 0x52, 0xdb, 0x3a,
	0x14, 0xc6, 0xe3, 0xfb, 0x87, 0x7b, 0xaf, 0xb8, 0xf7, 0xd6, 0xa8, 0xb4, 0xec, 0xbc, 0x6a, 0x99,
	0x21, 0x53, 0x92, 0x52, 0x5a, 0x18, 0x9a, 0x76, 0xa8, 0x22, 0x9f, 0x10, 0x4d, 0x8c, 0x9d, 0x4a,
	0x32, 0x49, 0xd8, 0x68, 0xdc, 0x60, 0xa8, 0x87, 0x10, 0x33, 0x49, 0xa0, 0x2c, 0xfb, 0x00, 0x5d,
	0xf4, 0x31, 0xba, 0xea, 0x73, 0xb0, 0x64, 0x49, 0x77, 0x25, 0x6c, 0xba, 0xe4, 0x11, 0x3a, 0xc9,
	0x10, 0x6a, 0x0a, 0x63, 0x3b, 0xdd, 0xc9, 0xf6, 0xf9, 0x7d, 0xd2, 0x77, 0x74, 0x7c, 0xe6, 0xa0,
	0x07, 0x3d, 0x7f, 0x6f, 0x3f, 0xec, 0x78, 0xad, 0xbc, 0xdf, 0x3e, 0xd8, 0xeb, 0xe6, 0x0f, 0x17,
	0xf2, 0xdb, 0x5e, 0xd0, 0xf2, 0xb7, 0x54, 0xd3, 0x3b, 0xe8, 0xfa, 0xb9, 0xfd, 0x4e, 0xd8, 0x0b,
	0xf1, 0xd4, 0x28, 0x2a, 0x37, 0x8c, 0xca, 0x1d, 0x2e, 0x64, 0xbf, 0x4c, 0xa2, 0x19, 0xd3, 0x6f,
	0x06, 0xdd, 0x20, 0x6c, 0x4b, 0xaf, 0xbb, 0x5b, 0x1a, 0x52, 0x74, 0x00, 0xe1, 0x2c, 0x9a, 0x35,
	0x81, 0x32, 0xc1, 0x1c, 0x5b, 0x49, 0x22, 0x2a, 0xaa, 0x44, 0x98, 0x05, 0xa6, 0xa2, 0xc4, 0x15,
	0xa0, 0x5c, 0x5b, 0x54, 0x81, 0xb2, 0x12, 0x03, 0x53, 0xcf, 0xe0, 0x05, 0x34, 0x1f, 0x1b, 0x5b,
	0x26, 0xb6, 0x39, 0x78, 0x1e, 0x05, 0xe9, 0x1a, 0x5e, 0x45, 0x85, 0x18, 0xa4, 0x48, 0x4c, 0x25,
	0x68, 0x19, 0x4c, 0xd7, 0x02, 0x45, 0xa8, 0x64, 0x1b, 0x4c, 0x36, 0x14, 0x91, 0x92, 0xb3, 0xa2,
	0x2b, 0x41, 0xe8, 0xbf, 0x61, 0x40, 0x24, 0x41, 0x80, 0xc3, 0x6b, 0x17, 0x84, 0x54, 0x94, 0xd8,
	0x14, 0xac, 0x5b, 0x65, 0x7e, 0xc7, 0x2b, 0xe8, 0x59, 0xd2, 0x39, 0x24, 0xe1, 0x52, 0x49, 0xb6,
	0x0e, 0x3c, 0x8a, 0xfe, 0x81, 0x9f, 0xa3, 0xa5, 0x04, 0xf4, 0x72, 0xe7, 0x1b, 0xec, 0x9f, 0xb8,
	0x80, 0x96, 0x13, 0x4f, 0x4f, 0x1d, 0x6e, 0xaa, 0x75, 0xc2, 0x2b, 0xd7, 0xe1, 0x09, 0xcc, 0x10,
	0x24, 0x6d, 0xec, 0xac, 0x57, 0x2d, 0x90, 0xa0, 0x6a, 0x0e, 0xaf, 0x94, 0x2c, 0xa7, 0xa6, 0xa0,
	0x0e, 0xd4, 0x95, 0x03, 0x22, 0x22, 0xf5, 0x57, 0x8a, 0x2c, 0x0e, 0x5e, 0x24, 0xc8, 0xfc, 0x8d,
	0xd7, 0x10, 0x4d, 0x97, 0x8a, 0x78, 0xa1, 0x7f, 0x70, 0x1d, 0xc9, 0xf1, 0x6e, 0x15, 0xea, 0x12,
	0xb8, 0x4d, 0x92, 0x94, 0x11, 0x7e, 0x89, 0x56, 0x12, 0x93, 0x66, 0x4b, 0x66, 0xbb, 0xa0, 0x88,
	0x50, 0x36, 0xd4, 0xa2, 0xf8, 0x24, 0x5e, 0x46, 0x8b, 0x31, 0x78, 0xb4, 0x46, 0x4c, 0xb7, 0x6a,
	0x31, 0x4a, 0x24, 0x28, 0x66, 0xea, 0xff, 0xe2, 0xa7, 0xe8, 0x71, 0x0c, 0xc8, 0x41, 0x80, 0x54,
	0x42, 0x32, 0x5a, 0x69, 0x0c, 0x3f, 0x5b, 0x4c, 0x48, 0xfd, 0x3f, 0xfc, 0x0a, 0xbd, 0x88, 0xa1,
	0xae, 0xbc, 0x0e, 0x16, 0xc0, 0x23, 0x7f, 0xd8, 0x20, 0xcc, 0xe5, 0xa0, 0xff, 0x9f, 0xe2, 0x4a,
	0x04, 0x5b, 0x4b, 0x4e, 0xdc, 0x1d, 0x4c, 0xd1, 0x6a, 0xaa, 0x3f, 0x84, 0x96, 0x99, 0x65, 0xde,
	0x2e, 0xa2, 0xe3, 0x45, 0x94, 0x8f, 0x11, 0x29, 0x39, 0x9c, 0x82, 0xa2, 0x96, 0x23, 0xe0, 0x47,
	0x8f, 0x98, 0xc2, 0x4b, 0xe8, 0x49, 0x1c, 0x44, 0x98, 0xe5, 0x6c, 0x00, 0xff, 0x99, 0xc3, 0x09,
	0x29, 0x8f, 0x58, 0x67, 0x76, 0xd5, 0x95, 0x4a, 0xb0, 0x4d, 0xd0, 0xef, 0xe2, 0x79, 0x34, 0x97,
	0x78, 0x51, 0xa3, 0x5c, 0xe9, 0xd3, 0x78, 0x0e, 0x3d, 0x4c, 0xd8, 0xa4, 0xc8, 0x6c, 0xc2, 0x1b,
	0xfa, 0xbd, 0x84, 0xd2, 0xbb, 0xd9, 0xe7, 0xae, 0x55, 0xd0, 0xfd, 0x34, 0x76, 0x80, 0x70, 0x5a,
	0x8e, 0x66, 0x7c, 0x26, 0xfb, 0x59, 0x43, 0xb3, 0xa2, 0xe7, 0x75, 0x7a, 0xf4, 0x6d, 0xd0, 0xda,
	0xaa, 0x85, 0x9d, 0xdd, 0xed, 0x56, 0xf8, 0x0e, 0x8e, 0xfc, 0xe6, 0x41, 0x2f, 0x08, 0xdb, 0xd1,
	0x56, 0x5f, 0x40, 0xcb, 0xd1, 0x2b, 0xbc, 0xa5, 0x20, 0x62, 0x7a, 0xff, 0x1a, 0xa2, 0xe3, 0xc0,
	0x57, 0xdf, 0x89, 0xc5, 0x81, 0x98, 0x0d, 0x05, 0x75, 0x26, 0xa4, 0xd0, 0xb5, 0xec, 0xb1, 0x86,
	0xb2, 0xd4, 0x6b, 0x37, 0xfd, 0x16, 0x1c, 0xf5, 0xfc, 0x4e, 0xdb, 0x6b, 0xc5, 0x1e, 0x7a, 0x15,
	0x15, 0x52, 0xb4, 0x80, 0x98, 0x83, 0x37, 0x90, 0x3b, 0xae, 0x40, 0x5c, 0xa0, 0xed, 0x48, 0x55,
	0x72, 0x5c, 0xdb, 0xbc, 0xb4, 0x22, 0x82, 0x9d, 0xb6, 0x97, 0xda, 0xca, 0x65, 0x41, 0xfe, 0xba,
	0x95, 0x71, 0x05, 0x52, 0x5a, 0x29, 0x7e, 0xd0, 0x4e, 0xce, 0x8c, 0xcc, 0xe9, 0x99, 0x91, 0xb9,
	0x38, 0x33, 0xb4, 0xf7, 0x7d, 0x43, 0xfb, 0xd4, 0x37, 0xb4, 0xe3, 0xbe, 0xa1, 0x9d, 0xf4, 0x0d,
	0xed, 0x6b, 0xdf, 0xd0, 0xbe, 0xf5, 0x8d, 0xcc, 0x45, 0xdf, 0xd0, 0x3e, 0x9e, 0x1b, 0x99, 0x93,
	0x73, 0x23, 0x73, 0x7a, 0x6e, 0x64, 0xd0, 0x74, 0x10, 0xe6, 0x6e, 0xcc, 0x1b, 0x45, 0x3d, 0xe2,
	0xba, 0x3a, 0x18, 0x4a, 0xaa, 0xda, 0xe6, 0xa3, 0x9d, 0x48, 0x64, 0x10, 0xe6, 0x47, 0xeb, 0xf9,
	0xe1, 0xd4, 0x72, 0x35, 0xd1, 0x14, 0x86, 0x8b, 0x37, 0x13, 0xc3, 0xb7, 0x8b, 0xdf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x48, 0x9d, 0xfe, 0xf3, 0x08, 0x00, 0x00,
}

func (x DecisionTaskFailedCause) String() string {
	s, ok := DecisionTaskFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x StartChildWorkflowExecutionFailedCause) String() string {
	s, ok := StartChildWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CancelExternalWorkflowExecutionFailedCause) String() string {
	s, ok := CancelExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x SignalExternalWorkflowExecutionFailedCause) String() string {
	s, ok := SignalExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}