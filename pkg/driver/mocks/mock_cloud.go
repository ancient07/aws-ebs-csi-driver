// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ancient07/aws-ebs-csi-driver/pkg/cloud (interfaces: Cloud)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	cloud "github.com/ancient07/aws-ebs-csi-driver/pkg/cloud"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCloud is a mock of Cloud interface.
type MockCloud struct {
	ctrl     *gomock.Controller
	recorder *MockCloudMockRecorder
}

// MockCloudMockRecorder is the mock recorder for MockCloud.
type MockCloudMockRecorder struct {
	mock *MockCloud
}

// NewMockCloud creates a new mock instance.
func NewMockCloud(ctrl *gomock.Controller) *MockCloud {
	mock := &MockCloud{ctrl: ctrl}
	mock.recorder = &MockCloudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloud) EXPECT() *MockCloudMockRecorder {
	return m.recorder
}

// AttachDisk mocks base method.
func (m *MockCloud) AttachDisk(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachDisk", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachDisk indicates an expected call of AttachDisk.
func (mr *MockCloudMockRecorder) AttachDisk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachDisk", reflect.TypeOf((*MockCloud)(nil).AttachDisk), arg0, arg1, arg2)
}

// CreateDisk mocks base method.
func (m *MockCloud) CreateDisk(arg0 context.Context, arg1 string, arg2 *cloud.DiskOptions) (*cloud.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDisk", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cloud.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDisk indicates an expected call of CreateDisk.
func (mr *MockCloudMockRecorder) CreateDisk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDisk", reflect.TypeOf((*MockCloud)(nil).CreateDisk), arg0, arg1, arg2)
}

// CreateSnapshot mocks base method.
func (m *MockCloud) CreateSnapshot(arg0 context.Context, arg1 string, arg2 *cloud.SnapshotOptions) (*cloud.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cloud.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockCloudMockRecorder) CreateSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockCloud)(nil).CreateSnapshot), arg0, arg1, arg2)
}

// DeleteDisk mocks base method.
func (m *MockCloud) DeleteDisk(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDisk", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDisk indicates an expected call of DeleteDisk.
func (mr *MockCloudMockRecorder) DeleteDisk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDisk", reflect.TypeOf((*MockCloud)(nil).DeleteDisk), arg0, arg1)
}

// DeleteSnapshot mocks base method.
func (m *MockCloud) DeleteSnapshot(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockCloudMockRecorder) DeleteSnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockCloud)(nil).DeleteSnapshot), arg0, arg1)
}

// DetachDisk mocks base method.
func (m *MockCloud) DetachDisk(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachDisk", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachDisk indicates an expected call of DetachDisk.
func (mr *MockCloudMockRecorder) DetachDisk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachDisk", reflect.TypeOf((*MockCloud)(nil).DetachDisk), arg0, arg1, arg2)
}

// GetDiskByID mocks base method.
func (m *MockCloud) GetDiskByID(arg0 context.Context, arg1 string) (*cloud.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskByID", arg0, arg1)
	ret0, _ := ret[0].(*cloud.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskByID indicates an expected call of GetDiskByID.
func (mr *MockCloudMockRecorder) GetDiskByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskByID", reflect.TypeOf((*MockCloud)(nil).GetDiskByID), arg0, arg1)
}

// GetDiskByName mocks base method.
func (m *MockCloud) GetDiskByName(arg0 context.Context, arg1 string, arg2 int64) (*cloud.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskByName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cloud.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskByName indicates an expected call of GetDiskByName.
func (mr *MockCloudMockRecorder) GetDiskByName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskByName", reflect.TypeOf((*MockCloud)(nil).GetDiskByName), arg0, arg1, arg2)
}

// GetSnapshotByID mocks base method.
func (m *MockCloud) GetSnapshotByID(arg0 context.Context, arg1 string) (*cloud.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotByID", arg0, arg1)
	ret0, _ := ret[0].(*cloud.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshotByID indicates an expected call of GetSnapshotByID.
func (mr *MockCloudMockRecorder) GetSnapshotByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotByID", reflect.TypeOf((*MockCloud)(nil).GetSnapshotByID), arg0, arg1)
}

// GetSnapshotByName mocks base method.
func (m *MockCloud) GetSnapshotByName(arg0 context.Context, arg1 string) (*cloud.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotByName", arg0, arg1)
	ret0, _ := ret[0].(*cloud.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshotByName indicates an expected call of GetSnapshotByName.
func (mr *MockCloudMockRecorder) GetSnapshotByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotByName", reflect.TypeOf((*MockCloud)(nil).GetSnapshotByName), arg0, arg1)
}

// IsExistInstance mocks base method.
func (m *MockCloud) IsExistInstance(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistInstance", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExistInstance indicates an expected call of IsExistInstance.
func (mr *MockCloudMockRecorder) IsExistInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistInstance", reflect.TypeOf((*MockCloud)(nil).IsExistInstance), arg0, arg1)
}

// ListSnapshots mocks base method.
func (m *MockCloud) ListSnapshots(arg0 context.Context, arg1 string, arg2 int64, arg3 string) (*cloud.ListSnapshotsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSnapshots", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*cloud.ListSnapshotsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSnapshots indicates an expected call of ListSnapshots.
func (mr *MockCloudMockRecorder) ListSnapshots(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnapshots", reflect.TypeOf((*MockCloud)(nil).ListSnapshots), arg0, arg1, arg2, arg3)
}

// ResizeDisk mocks base method.
func (m *MockCloud) ResizeDisk(arg0 context.Context, arg1 string, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeDisk", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResizeDisk indicates an expected call of ResizeDisk.
func (mr *MockCloudMockRecorder) ResizeDisk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeDisk", reflect.TypeOf((*MockCloud)(nil).ResizeDisk), arg0, arg1, arg2)
}

// WaitForAttachmentState mocks base method.
func (m *MockCloud) WaitForAttachmentState(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 bool) (*ec2.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForAttachmentState", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*ec2.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForAttachmentState indicates an expected call of WaitForAttachmentState.
func (mr *MockCloudMockRecorder) WaitForAttachmentState(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForAttachmentState", reflect.TypeOf((*MockCloud)(nil).WaitForAttachmentState), arg0, arg1, arg2, arg3, arg4, arg5)
}
