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

package ndc

import (
	"time"

	workflowspb "go.temporal.io/server/api/workflow/v1"
	"go.temporal.io/server/common/namespace"
	"go.temporal.io/server/common/persistence/serialization"
	"go.temporal.io/server/common/util"

	"github.com/pborman/uuid"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	historypb "go.temporal.io/api/history/v1"
	"go.temporal.io/api/serviceerror"

	historyspb "go.temporal.io/server/api/history/v1"
	"go.temporal.io/server/api/historyservice/v1"
	"go.temporal.io/server/common/cluster"
	"go.temporal.io/server/common/log"
	"go.temporal.io/server/common/log/tag"
	"go.temporal.io/server/common/primitives/timestamp"
	"go.temporal.io/server/service/history/consts"
)

type (
	replicationTask interface {
		getNamespaceID() namespace.ID
		getExecution() *commonpb.WorkflowExecution
		getWorkflowID() string
		getRunID() string
		getEventTime() time.Time
		getFirstEvent() *historypb.HistoryEvent
		getLastEvent() *historypb.HistoryEvent
		getVersion() int64
		getSourceCluster() string
		getEvents() []*historypb.HistoryEvent
		getNewEvents() []*historypb.HistoryEvent
		getLogger() log.Logger
		getBaseWorkflowInfo() *workflowspb.BaseExecutionInfo
		getVersionHistory() *historyspb.VersionHistory
		isWorkflowReset() bool

		splitTask(taskStartTime time.Time) (replicationTask, replicationTask, error)
	}

	replicationTaskImpl struct {
		sourceCluster    string
		namespaceID      namespace.ID
		execution        *commonpb.WorkflowExecution
		version          int64
		firstEvent       *historypb.HistoryEvent
		lastEvent        *historypb.HistoryEvent
		eventTime        time.Time
		baseWorkflowInfo *workflowspb.BaseExecutionInfo
		versionHistory   *historyspb.VersionHistory
		events           []*historypb.HistoryEvent
		newEvents        []*historypb.HistoryEvent

		startTime time.Time
		logger    log.Logger
	}
)

var (
	// ErrInvalidNamespaceID is returned if namespace ID is invalid
	ErrInvalidNamespaceID = serviceerror.NewInvalidArgument("invalid namespace ID")
	// ErrInvalidExecution is returned if execution is invalid
	ErrInvalidExecution = serviceerror.NewInvalidArgument("invalid execution")
	// ErrInvalidRunID is returned if run ID is invalid
	ErrInvalidRunID = serviceerror.NewInvalidArgument("invalid run ID")
	// ErrEventIDMismatch is returned if event ID mis-matched
	ErrEventIDMismatch = serviceerror.NewInvalidArgument("event ID mismatch")
	// ErrEventVersionMismatch is returned if event version mis-matched
	ErrEventVersionMismatch = serviceerror.NewInvalidArgument("event version mismatch")
	// ErrNoNewRunHistory is returned if there is no new run history
	ErrNoNewRunHistory = serviceerror.NewInvalidArgument("no new run history events")
	// ErrLastEventIsNotContinueAsNew is returned if the last event is not continue as new
	ErrLastEventIsNotContinueAsNew = serviceerror.NewInvalidArgument("last event is not continue as new")
)

func newReplicationTask(
	clusterMetadata cluster.Metadata,
	historySerializer serialization.Serializer,
	taskStartTime time.Time,
	logger log.Logger,
	request *historyservice.ReplicateEventsV2Request,
) (*replicationTaskImpl, error) {

	events, newEvents, err := validateReplicateEventsRequest(
		historySerializer,
		request,
	)
	if err != nil {
		return nil, err
	}

	namespaceID := namespace.ID(request.GetNamespaceId())
	execution := request.WorkflowExecution
	baseExecutionInfo := request.BaseExecutionInfo
	versionHistory := &historyspb.VersionHistory{
		BranchToken: nil,
		Items:       request.VersionHistoryItems,
	}

	firstEvent := events[0]
	lastEvent := events[len(events)-1]
	version := firstEvent.GetVersion()

	sourceCluster := clusterMetadata.ClusterNameForFailoverVersion(true, version)

	eventTime := time.Time{}
	for _, event := range events {
		eventTime = util.MaxTime(eventTime, timestamp.TimeValue(event.GetEventTime()))
	}
	for _, event := range newEvents {
		eventTime = util.MaxTime(eventTime, timestamp.TimeValue(event.GetEventTime()))
	}

	logger = log.With(
		logger,
		tag.WorkflowID(execution.GetWorkflowId()),
		tag.WorkflowRunID(execution.GetRunId()),
		tag.SourceCluster(sourceCluster),
		tag.IncomingVersion(version),
		tag.WorkflowFirstEventID(firstEvent.GetEventId()),
		tag.WorkflowNextEventID(lastEvent.GetEventId()+1),
	)

	return &replicationTaskImpl{
		sourceCluster:    sourceCluster,
		namespaceID:      namespaceID,
		execution:        execution,
		version:          version,
		firstEvent:       firstEvent,
		lastEvent:        lastEvent,
		eventTime:        eventTime,
		baseWorkflowInfo: baseExecutionInfo,
		versionHistory:   versionHistory,
		events:           events,
		newEvents:        newEvents,

		startTime: taskStartTime,
		logger:    logger,
	}, nil
}

func (t *replicationTaskImpl) getNamespaceID() namespace.ID {
	return t.namespaceID
}

func (t *replicationTaskImpl) getExecution() *commonpb.WorkflowExecution {
	return t.execution
}

func (t *replicationTaskImpl) getWorkflowID() string {
	return t.execution.GetWorkflowId()
}

func (t *replicationTaskImpl) getRunID() string {
	return t.execution.GetRunId()
}

func (t *replicationTaskImpl) getEventTime() time.Time {
	return t.eventTime
}

func (t *replicationTaskImpl) getFirstEvent() *historypb.HistoryEvent {
	return t.firstEvent
}

func (t *replicationTaskImpl) getLastEvent() *historypb.HistoryEvent {
	return t.lastEvent
}

func (t *replicationTaskImpl) getVersion() int64 {
	return t.version
}

func (t *replicationTaskImpl) getSourceCluster() string {
	return t.sourceCluster
}

func (t *replicationTaskImpl) getEvents() []*historypb.HistoryEvent {
	return t.events
}

func (t *replicationTaskImpl) getNewEvents() []*historypb.HistoryEvent {
	return t.newEvents
}

func (t *replicationTaskImpl) getLogger() log.Logger {
	return t.logger
}

func (t *replicationTaskImpl) getBaseWorkflowInfo() *workflowspb.BaseExecutionInfo {
	if t.baseWorkflowInfo != nil {
		return t.baseWorkflowInfo
	}

	// TODO deprecate
	switch t.getFirstEvent().GetEventType() {
	case enumspb.EVENT_TYPE_WORKFLOW_TASK_FAILED:
		workflowTaskFailedEvent := t.getFirstEvent()
		attr := workflowTaskFailedEvent.GetWorkflowTaskFailedEventAttributes()
		baseRunID := attr.GetBaseRunId()
		baseEventID := t.getFirstEvent().EventId - 1
		baseEventVersion := attr.GetForkEventVersion()
		return &workflowspb.BaseExecutionInfo{
			RunId:                            baseRunID,
			LowestCommonAncestorEventId:      baseEventID,
			LowestCommonAncestorEventVersion: baseEventVersion,
		}
	default:
		return nil
	}
}

func (t *replicationTaskImpl) getVersionHistory() *historyspb.VersionHistory {
	return t.versionHistory
}

func (t *replicationTaskImpl) isWorkflowReset() bool {
	if t.baseWorkflowInfo != nil && t.baseWorkflowInfo.LowestCommonAncestorEventId+1 == t.firstEvent.EventId {
		return true
	}

	// TODO deprecate
	switch t.getFirstEvent().GetEventType() {
	case enumspb.EVENT_TYPE_WORKFLOW_TASK_FAILED:
		workflowTaskFailedEvent := t.getFirstEvent()
		attr := workflowTaskFailedEvent.GetWorkflowTaskFailedEventAttributes()
		baseRunID := attr.GetBaseRunId()
		baseEventVersion := attr.GetForkEventVersion()
		newRunID := attr.GetNewRunId()

		return len(baseRunID) > 0 && baseEventVersion != 0 && len(newRunID) > 0

	default:
		return false
	}
}

func (t *replicationTaskImpl) splitTask(
	taskStartTime time.Time,
) (replicationTask, replicationTask, error) {

	if len(t.newEvents) == 0 {
		return nil, nil, ErrNoNewRunHistory
	}
	newHistoryEvents := t.newEvents

	if t.getLastEvent().GetEventType() != enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW ||
		t.getLastEvent().GetWorkflowExecutionContinuedAsNewEventAttributes() == nil {
		return nil, nil, ErrLastEventIsNotContinueAsNew
	}
	newRunID := t.getLastEvent().GetWorkflowExecutionContinuedAsNewEventAttributes().GetNewExecutionRunId()

	newFirstEvent := newHistoryEvents[0]
	newLastEvent := newHistoryEvents[len(newHistoryEvents)-1]

	newEventTime := time.Time{}
	for _, event := range newHistoryEvents {
		newEventTime = util.MaxTime(newEventTime, timestamp.TimeValue(event.GetEventTime()))
	}

	newVersionHistory := &historyspb.VersionHistory{
		BranchToken: nil,
		Items: []*historyspb.VersionHistoryItem{{
			EventId: newLastEvent.GetEventId(),
			Version: newLastEvent.GetVersion(),
		}},
	}

	logger := log.With(
		t.logger,
		tag.WorkflowID(t.getExecution().GetWorkflowId()),
		tag.WorkflowRunID(newRunID),
		tag.SourceCluster(t.sourceCluster),
		tag.IncomingVersion(t.version),
		tag.WorkflowFirstEventID(newFirstEvent.GetEventId()),
		tag.WorkflowNextEventID(newLastEvent.GetEventId()+1),
	)

	newRunTask := &replicationTaskImpl{
		sourceCluster: t.sourceCluster,
		namespaceID:   t.namespaceID,
		execution: &commonpb.WorkflowExecution{
			WorkflowId: t.execution.WorkflowId,
			RunId:      newRunID,
		},
		version:          t.version,
		firstEvent:       newFirstEvent,
		lastEvent:        newLastEvent,
		eventTime:        newEventTime,
		baseWorkflowInfo: nil,
		versionHistory:   newVersionHistory,
		events:           newHistoryEvents,
		newEvents:        []*historypb.HistoryEvent{},

		startTime: taskStartTime,
		logger:    logger,
	}
	t.newEvents = nil

	return t, newRunTask, nil
}

func validateReplicateEventsRequest(
	historySerializer serialization.Serializer,
	request *historyservice.ReplicateEventsV2Request,
) ([]*historypb.HistoryEvent, []*historypb.HistoryEvent, error) {

	// TODO add validation on version history

	if valid := validateUUID(request.GetNamespaceId()); !valid {
		return nil, nil, ErrInvalidNamespaceID
	}
	if request.WorkflowExecution == nil {
		return nil, nil, ErrInvalidExecution
	}
	if valid := validateUUID(request.WorkflowExecution.GetRunId()); !valid {
		return nil, nil, ErrInvalidRunID
	}

	events, err := deserializeBlob(historySerializer, request.Events)
	if err != nil {
		return nil, nil, err
	}
	if len(events) == 0 {
		return nil, nil, consts.ErrEmptyHistoryRawEventBatch
	}

	version, err := validateEvents(events)
	if err != nil {
		return nil, nil, err
	}

	if request.NewRunEvents == nil {
		return events, nil, nil
	}

	newRunEvents, err := deserializeBlob(historySerializer, request.NewRunEvents)
	if err != nil {
		return nil, nil, err
	}

	newRunVersion, err := validateEvents(newRunEvents)
	if err != nil {
		return nil, nil, err
	}
	if version != newRunVersion {
		return nil, nil, ErrEventVersionMismatch
	}
	return events, newRunEvents, nil
}

func validateUUID(input string) bool {
	return uuid.Parse(input) != nil
}

func validateEvents(events []*historypb.HistoryEvent) (int64, error) {

	firstEvent := events[0]
	firstEventID := firstEvent.GetEventId()
	version := firstEvent.GetVersion()

	for index, event := range events {
		if event.GetEventId() != firstEventID+int64(index) {
			return 0, ErrEventIDMismatch
		}
		if event.GetVersion() != version {
			return 0, ErrEventVersionMismatch
		}
	}
	return version, nil
}

func deserializeBlob(
	historySerializer serialization.Serializer,
	blob *commonpb.DataBlob,
) ([]*historypb.HistoryEvent, error) {

	if blob == nil {
		return nil, nil
	}

	events, err := historySerializer.DeserializeEvents(&commonpb.DataBlob{
		EncodingType: enumspb.ENCODING_TYPE_PROTO3,
		Data:         blob.Data,
	})

	return events, err
}
