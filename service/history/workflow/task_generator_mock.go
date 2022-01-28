// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: task_generator.go

// Package workflow is a generated GoMock package.
package workflow

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	history "go.temporal.io/api/history/v1"
	tasks "go.temporal.io/server/service/history/tasks"
)

// MockTaskGenerator is a mock of TaskGenerator interface.
type MockTaskGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockTaskGeneratorMockRecorder
}

// MockTaskGeneratorMockRecorder is the mock recorder for MockTaskGenerator.
type MockTaskGeneratorMockRecorder struct {
	mock *MockTaskGenerator
}

// NewMockTaskGenerator creates a new mock instance.
func NewMockTaskGenerator(ctrl *gomock.Controller) *MockTaskGenerator {
	mock := &MockTaskGenerator{ctrl: ctrl}
	mock.recorder = &MockTaskGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskGenerator) EXPECT() *MockTaskGeneratorMockRecorder {
	return m.recorder
}

// GenerateActivityRetryTasks mocks base method.
func (m *MockTaskGenerator) GenerateActivityRetryTasks(activityScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateActivityRetryTasks", activityScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateActivityRetryTasks indicates an expected call of GenerateActivityRetryTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateActivityRetryTasks(activityScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateActivityRetryTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateActivityRetryTasks), activityScheduleID)
}

// GenerateActivityTimerTasks mocks base method.
func (m *MockTaskGenerator) GenerateActivityTimerTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateActivityTimerTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateActivityTimerTasks indicates an expected call of GenerateActivityTimerTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateActivityTimerTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateActivityTimerTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateActivityTimerTasks), now)
}

// GenerateActivityTransferTasks mocks base method.
func (m *MockTaskGenerator) GenerateActivityTransferTasks(now time.Time, event *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateActivityTransferTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateActivityTransferTasks indicates an expected call of GenerateActivityTransferTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateActivityTransferTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateActivityTransferTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateActivityTransferTasks), now, event)
}

// GenerateChildWorkflowTasks mocks base method.
func (m *MockTaskGenerator) GenerateChildWorkflowTasks(now time.Time, event *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateChildWorkflowTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateChildWorkflowTasks indicates an expected call of GenerateChildWorkflowTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateChildWorkflowTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateChildWorkflowTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateChildWorkflowTasks), now, event)
}

// GenerateDelayedWorkflowTasks mocks base method.
func (m *MockTaskGenerator) GenerateDelayedWorkflowTasks(now time.Time, startEvent *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDelayedWorkflowTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateDelayedWorkflowTasks indicates an expected call of GenerateDelayedWorkflowTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateDelayedWorkflowTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDelayedWorkflowTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateDelayedWorkflowTasks), now, startEvent)
}

// GenerateDeleteExecutionTask mocks base method.
func (m *MockTaskGenerator) GenerateDeleteExecutionTask(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDeleteExecutionTask", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateDeleteExecutionTask indicates an expected call of GenerateDeleteExecutionTask.
func (mr *MockTaskGeneratorMockRecorder) GenerateDeleteExecutionTask(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDeleteExecutionTask", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateDeleteExecutionTask), now)
}

// GenerateHistoryReplicationTasks mocks base method.
func (m *MockTaskGenerator) GenerateHistoryReplicationTasks(now time.Time, branchToken []byte, events []*history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateHistoryReplicationTasks", now, branchToken, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateHistoryReplicationTasks indicates an expected call of GenerateHistoryReplicationTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateHistoryReplicationTasks(now, branchToken, events interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateHistoryReplicationTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateHistoryReplicationTasks), now, branchToken, events)
}

// GenerateLastHistoryReplicationTasks mocks base method.
func (m *MockTaskGenerator) GenerateLastHistoryReplicationTasks(now time.Time) (*tasks.HistoryReplicationTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateLastHistoryReplicationTasks", now)
	ret0, _ := ret[0].(*tasks.HistoryReplicationTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateLastHistoryReplicationTasks indicates an expected call of GenerateLastHistoryReplicationTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateLastHistoryReplicationTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateLastHistoryReplicationTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateLastHistoryReplicationTasks), now)
}

// GenerateRecordWorkflowStartedTasks mocks base method.
func (m *MockTaskGenerator) GenerateRecordWorkflowStartedTasks(now time.Time, startEvent *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRecordWorkflowStartedTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateRecordWorkflowStartedTasks indicates an expected call of GenerateRecordWorkflowStartedTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateRecordWorkflowStartedTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRecordWorkflowStartedTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateRecordWorkflowStartedTasks), now, startEvent)
}

// GenerateRequestCancelExternalTasks mocks base method.
func (m *MockTaskGenerator) GenerateRequestCancelExternalTasks(now time.Time, event *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRequestCancelExternalTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateRequestCancelExternalTasks indicates an expected call of GenerateRequestCancelExternalTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateRequestCancelExternalTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRequestCancelExternalTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateRequestCancelExternalTasks), now, event)
}

// GenerateScheduleWorkflowTaskTasks mocks base method.
func (m *MockTaskGenerator) GenerateScheduleWorkflowTaskTasks(now time.Time, workflowTaskScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateScheduleWorkflowTaskTasks", now, workflowTaskScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateScheduleWorkflowTaskTasks indicates an expected call of GenerateScheduleWorkflowTaskTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateScheduleWorkflowTaskTasks(now, workflowTaskScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateScheduleWorkflowTaskTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateScheduleWorkflowTaskTasks), now, workflowTaskScheduleID)
}

// GenerateSignalExternalTasks mocks base method.
func (m *MockTaskGenerator) GenerateSignalExternalTasks(now time.Time, event *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSignalExternalTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateSignalExternalTasks indicates an expected call of GenerateSignalExternalTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateSignalExternalTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSignalExternalTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateSignalExternalTasks), now, event)
}

// GenerateStartWorkflowTaskTasks mocks base method.
func (m *MockTaskGenerator) GenerateStartWorkflowTaskTasks(now time.Time, workflowTaskScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateStartWorkflowTaskTasks", now, workflowTaskScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateStartWorkflowTaskTasks indicates an expected call of GenerateStartWorkflowTaskTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateStartWorkflowTaskTasks(now, workflowTaskScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateStartWorkflowTaskTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateStartWorkflowTaskTasks), now, workflowTaskScheduleID)
}

// GenerateUserTimerTasks mocks base method.
func (m *MockTaskGenerator) GenerateUserTimerTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUserTimerTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateUserTimerTasks indicates an expected call of GenerateUserTimerTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateUserTimerTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUserTimerTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateUserTimerTasks), now)
}

// GenerateWorkflowCloseTasks mocks base method.
func (m *MockTaskGenerator) GenerateWorkflowCloseTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowCloseTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowCloseTasks indicates an expected call of GenerateWorkflowCloseTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateWorkflowCloseTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowCloseTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateWorkflowCloseTasks), now)
}

// GenerateWorkflowResetTasks mocks base method.
func (m *MockTaskGenerator) GenerateWorkflowResetTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowResetTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowResetTasks indicates an expected call of GenerateWorkflowResetTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateWorkflowResetTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowResetTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateWorkflowResetTasks), now)
}

// GenerateWorkflowSearchAttrTasks mocks base method.
func (m *MockTaskGenerator) GenerateWorkflowSearchAttrTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowSearchAttrTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowSearchAttrTasks indicates an expected call of GenerateWorkflowSearchAttrTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateWorkflowSearchAttrTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowSearchAttrTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateWorkflowSearchAttrTasks), now)
}

// GenerateWorkflowStartTasks mocks base method.
func (m *MockTaskGenerator) GenerateWorkflowStartTasks(now time.Time, startEvent *history.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowStartTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowStartTasks indicates an expected call of GenerateWorkflowStartTasks.
func (mr *MockTaskGeneratorMockRecorder) GenerateWorkflowStartTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowStartTasks", reflect.TypeOf((*MockTaskGenerator)(nil).GenerateWorkflowStartTasks), now, startEvent)
}
