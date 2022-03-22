// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gitpod-io/gitpod/ws-daemon/api (interfaces: WorkspaceContentServiceClient,WorkspaceContentServiceServer,InWorkspaceServiceClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	api "github.com/gitpod-io/gitpod/ws-daemon/api"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockWorkspaceContentServiceClient is a mock of WorkspaceContentServiceClient interface.
type MockWorkspaceContentServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceContentServiceClientMockRecorder
}

// MockWorkspaceContentServiceClientMockRecorder is the mock recorder for MockWorkspaceContentServiceClient.
type MockWorkspaceContentServiceClientMockRecorder struct {
	mock *MockWorkspaceContentServiceClient
}

// NewMockWorkspaceContentServiceClient creates a new mock instance.
func NewMockWorkspaceContentServiceClient(ctrl *gomock.Controller) *MockWorkspaceContentServiceClient {
	mock := &MockWorkspaceContentServiceClient{ctrl: ctrl}
	mock.recorder = &MockWorkspaceContentServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceContentServiceClient) EXPECT() *MockWorkspaceContentServiceClientMockRecorder {
	return m.recorder
}

// BackupWorkspace mocks base method.
func (m *MockWorkspaceContentServiceClient) BackupWorkspace(arg0 context.Context, arg1 *api.BackupWorkspaceRequest, arg2 ...grpc.CallOption) (*api.BackupWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BackupWorkspace", varargs...)
	ret0, _ := ret[0].(*api.BackupWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackupWorkspace indicates an expected call of BackupWorkspace.
func (mr *MockWorkspaceContentServiceClientMockRecorder) BackupWorkspace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).BackupWorkspace), varargs...)
}

// DisposeWorkspace mocks base method.
func (m *MockWorkspaceContentServiceClient) DisposeWorkspace(arg0 context.Context, arg1 *api.DisposeWorkspaceRequest, arg2 ...grpc.CallOption) (*api.DisposeWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisposeWorkspace", varargs...)
	ret0, _ := ret[0].(*api.DisposeWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisposeWorkspace indicates an expected call of DisposeWorkspace.
func (mr *MockWorkspaceContentServiceClientMockRecorder) DisposeWorkspace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisposeWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).DisposeWorkspace), varargs...)
}

// InitWorkspace mocks base method.
func (m *MockWorkspaceContentServiceClient) InitWorkspace(arg0 context.Context, arg1 *api.InitWorkspaceRequest, arg2 ...grpc.CallOption) (*api.InitWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InitWorkspace", varargs...)
	ret0, _ := ret[0].(*api.InitWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitWorkspace indicates an expected call of InitWorkspace.
func (mr *MockWorkspaceContentServiceClientMockRecorder) InitWorkspace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).InitWorkspace), varargs...)
}

// TakeSnapshot mocks base method.
func (m *MockWorkspaceContentServiceClient) TakeSnapshot(arg0 context.Context, arg1 *api.TakeSnapshotRequest, arg2 ...grpc.CallOption) (*api.TakeSnapshotResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TakeSnapshot", varargs...)
	ret0, _ := ret[0].(*api.TakeSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeSnapshot indicates an expected call of TakeSnapshot.
func (mr *MockWorkspaceContentServiceClientMockRecorder) TakeSnapshot(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeSnapshot", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).TakeSnapshot), varargs...)
}

// WaitForInit mocks base method.
func (m *MockWorkspaceContentServiceClient) WaitForInit(arg0 context.Context, arg1 *api.WaitForInitRequest, arg2 ...grpc.CallOption) (*api.WaitForInitResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitForInit", varargs...)
	ret0, _ := ret[0].(*api.WaitForInitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForInit indicates an expected call of WaitForInit.
func (mr *MockWorkspaceContentServiceClientMockRecorder) WaitForInit(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForInit", reflect.TypeOf((*MockWorkspaceContentServiceClient)(nil).WaitForInit), varargs...)
}

// MockWorkspaceContentServiceServer is a mock of WorkspaceContentServiceServer interface.
type MockWorkspaceContentServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceContentServiceServerMockRecorder
	api.UnimplementedWorkspaceContentServiceServer
}

// MockWorkspaceContentServiceServerMockRecorder is the mock recorder for MockWorkspaceContentServiceServer.
type MockWorkspaceContentServiceServerMockRecorder struct {
	mock *MockWorkspaceContentServiceServer
}

// NewMockWorkspaceContentServiceServer creates a new mock instance.
func NewMockWorkspaceContentServiceServer(ctrl *gomock.Controller) *MockWorkspaceContentServiceServer {
	mock := &MockWorkspaceContentServiceServer{ctrl: ctrl}
	mock.recorder = &MockWorkspaceContentServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceContentServiceServer) EXPECT() *MockWorkspaceContentServiceServerMockRecorder {
	return m.recorder
}

// BackupWorkspace mocks base method.
func (m *MockWorkspaceContentServiceServer) BackupWorkspace(arg0 context.Context, arg1 *api.BackupWorkspaceRequest) (*api.BackupWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupWorkspace", arg0, arg1)
	ret0, _ := ret[0].(*api.BackupWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackupWorkspace indicates an expected call of BackupWorkspace.
func (mr *MockWorkspaceContentServiceServerMockRecorder) BackupWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).BackupWorkspace), arg0, arg1)
}

// DisposeWorkspace mocks base method.
func (m *MockWorkspaceContentServiceServer) DisposeWorkspace(arg0 context.Context, arg1 *api.DisposeWorkspaceRequest) (*api.DisposeWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisposeWorkspace", arg0, arg1)
	ret0, _ := ret[0].(*api.DisposeWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisposeWorkspace indicates an expected call of DisposeWorkspace.
func (mr *MockWorkspaceContentServiceServerMockRecorder) DisposeWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisposeWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).DisposeWorkspace), arg0, arg1)
}

// InitWorkspace mocks base method.
func (m *MockWorkspaceContentServiceServer) InitWorkspace(arg0 context.Context, arg1 *api.InitWorkspaceRequest) (*api.InitWorkspaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitWorkspace", arg0, arg1)
	ret0, _ := ret[0].(*api.InitWorkspaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitWorkspace indicates an expected call of InitWorkspace.
func (mr *MockWorkspaceContentServiceServerMockRecorder) InitWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitWorkspace", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).InitWorkspace), arg0, arg1)
}

// TakeSnapshot mocks base method.
func (m *MockWorkspaceContentServiceServer) TakeSnapshot(arg0 context.Context, arg1 *api.TakeSnapshotRequest) (*api.TakeSnapshotResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeSnapshot", arg0, arg1)
	ret0, _ := ret[0].(*api.TakeSnapshotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeSnapshot indicates an expected call of TakeSnapshot.
func (mr *MockWorkspaceContentServiceServerMockRecorder) TakeSnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeSnapshot", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).TakeSnapshot), arg0, arg1)
}

// WaitForInit mocks base method.
func (m *MockWorkspaceContentServiceServer) WaitForInit(arg0 context.Context, arg1 *api.WaitForInitRequest) (*api.WaitForInitResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForInit", arg0, arg1)
	ret0, _ := ret[0].(*api.WaitForInitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForInit indicates an expected call of WaitForInit.
func (mr *MockWorkspaceContentServiceServerMockRecorder) WaitForInit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForInit", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).WaitForInit), arg0, arg1)
}

// mustEmbedUnimplementedWorkspaceContentServiceServer mocks base method.
func (m *MockWorkspaceContentServiceServer) mustEmbedUnimplementedWorkspaceContentServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedWorkspaceContentServiceServer")
}

// mustEmbedUnimplementedWorkspaceContentServiceServer indicates an expected call of mustEmbedUnimplementedWorkspaceContentServiceServer.
func (mr *MockWorkspaceContentServiceServerMockRecorder) mustEmbedUnimplementedWorkspaceContentServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedWorkspaceContentServiceServer", reflect.TypeOf((*MockWorkspaceContentServiceServer)(nil).mustEmbedUnimplementedWorkspaceContentServiceServer))
}

// MockInWorkspaceServiceClient is a mock of InWorkspaceServiceClient interface.
type MockInWorkspaceServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockInWorkspaceServiceClientMockRecorder
}

// MockInWorkspaceServiceClientMockRecorder is the mock recorder for MockInWorkspaceServiceClient.
type MockInWorkspaceServiceClientMockRecorder struct {
	mock *MockInWorkspaceServiceClient
}

// NewMockInWorkspaceServiceClient creates a new mock instance.
func NewMockInWorkspaceServiceClient(ctrl *gomock.Controller) *MockInWorkspaceServiceClient {
	mock := &MockInWorkspaceServiceClient{ctrl: ctrl}
	mock.recorder = &MockInWorkspaceServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInWorkspaceServiceClient) EXPECT() *MockInWorkspaceServiceClientMockRecorder {
	return m.recorder
}

// EvacuateCGroup mocks base method.
func (m *MockInWorkspaceServiceClient) EvacuateCGroup(arg0 context.Context, arg1 *api.EvacuateCGroupRequest, arg2 ...grpc.CallOption) (*api.EvacuateCGroupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvacuateCGroup", varargs...)
	ret0, _ := ret[0].(*api.EvacuateCGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvacuateCGroup indicates an expected call of EvacuateCGroup.
func (mr *MockInWorkspaceServiceClientMockRecorder) EvacuateCGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvacuateCGroup", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).EvacuateCGroup), varargs...)
}

// MountProc mocks base method.
func (m *MockInWorkspaceServiceClient) MountProc(arg0 context.Context, arg1 *api.MountProcRequest, arg2 ...grpc.CallOption) (*api.MountProcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MountProc", varargs...)
	ret0, _ := ret[0].(*api.MountProcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MountProc indicates an expected call of MountProc.
func (mr *MockInWorkspaceServiceClientMockRecorder) MountProc(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountProc", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).MountProc), varargs...)
}

// MountSysfs mocks base method.
func (m *MockInWorkspaceServiceClient) MountSysfs(arg0 context.Context, arg1 *api.MountProcRequest, arg2 ...grpc.CallOption) (*api.MountProcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MountSysfs", varargs...)
	ret0, _ := ret[0].(*api.MountProcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MountSysfs indicates an expected call of MountSysfs.
func (mr *MockInWorkspaceServiceClientMockRecorder) MountSysfs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountSysfs", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).MountSysfs), varargs...)
}

// PrepareForUserNS mocks base method.
func (m *MockInWorkspaceServiceClient) PrepareForUserNS(arg0 context.Context, arg1 *api.PrepareForUserNSRequest, arg2 ...grpc.CallOption) (*api.PrepareForUserNSResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareForUserNS", varargs...)
	ret0, _ := ret[0].(*api.PrepareForUserNSResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareForUserNS indicates an expected call of PrepareForUserNS.
func (mr *MockInWorkspaceServiceClientMockRecorder) PrepareForUserNS(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForUserNS", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).PrepareForUserNS), varargs...)
}

// SetupPairVeths mocks base method.
func (m *MockInWorkspaceServiceClient) SetupPairVeths(arg0 context.Context, arg1 *api.SetupPairVethsRequest, arg2 ...grpc.CallOption) (*api.SetupPairVethsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetupPairVeths", varargs...)
	ret0, _ := ret[0].(*api.SetupPairVethsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupPairVeths indicates an expected call of SetupPairVeths.
func (mr *MockInWorkspaceServiceClientMockRecorder) SetupPairVeths(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupPairVeths", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).SetupPairVeths), varargs...)
}

// Teardown mocks base method.
func (m *MockInWorkspaceServiceClient) Teardown(arg0 context.Context, arg1 *api.TeardownRequest, arg2 ...grpc.CallOption) (*api.TeardownResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Teardown", varargs...)
	ret0, _ := ret[0].(*api.TeardownResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teardown indicates an expected call of Teardown.
func (mr *MockInWorkspaceServiceClientMockRecorder) Teardown(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).Teardown), varargs...)
}

// UmountProc mocks base method.
func (m *MockInWorkspaceServiceClient) UmountProc(arg0 context.Context, arg1 *api.UmountProcRequest, arg2 ...grpc.CallOption) (*api.UmountProcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UmountProc", varargs...)
	ret0, _ := ret[0].(*api.UmountProcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UmountProc indicates an expected call of UmountProc.
func (mr *MockInWorkspaceServiceClientMockRecorder) UmountProc(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UmountProc", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).UmountProc), varargs...)
}

// UmountSysfs mocks base method.
func (m *MockInWorkspaceServiceClient) UmountSysfs(arg0 context.Context, arg1 *api.UmountProcRequest, arg2 ...grpc.CallOption) (*api.UmountProcResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UmountSysfs", varargs...)
	ret0, _ := ret[0].(*api.UmountProcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UmountSysfs indicates an expected call of UmountSysfs.
func (mr *MockInWorkspaceServiceClientMockRecorder) UmountSysfs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UmountSysfs", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).UmountSysfs), varargs...)
}

// WriteIDMapping mocks base method.
func (m *MockInWorkspaceServiceClient) WriteIDMapping(arg0 context.Context, arg1 *api.WriteIDMappingRequest, arg2 ...grpc.CallOption) (*api.WriteIDMappingResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteIDMapping", varargs...)
	ret0, _ := ret[0].(*api.WriteIDMappingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteIDMapping indicates an expected call of WriteIDMapping.
func (mr *MockInWorkspaceServiceClientMockRecorder) WriteIDMapping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteIDMapping", reflect.TypeOf((*MockInWorkspaceServiceClient)(nil).WriteIDMapping), varargs...)
}
