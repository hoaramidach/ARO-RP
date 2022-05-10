// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/compute (interfaces: DisksClient,ResourceSkusClient,VirtualMachinesClient,UsageClient,VirtualMachineScaleSetVMsClient,VirtualMachineScaleSetsClient,DiskEncryptionSetsClient)

// Package mock_compute is a generated GoMock package.
package mock_compute

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	gomock "github.com/golang/mock/gomock"
)

// MockDisksClient is a mock of DisksClient interface.
type MockDisksClient struct {
	ctrl     *gomock.Controller
	recorder *MockDisksClientMockRecorder
}

// MockDisksClientMockRecorder is the mock recorder for MockDisksClient.
type MockDisksClientMockRecorder struct {
	mock *MockDisksClient
}

// NewMockDisksClient creates a new mock instance.
func NewMockDisksClient(ctrl *gomock.Controller) *MockDisksClient {
	mock := &MockDisksClient{ctrl: ctrl}
	mock.recorder = &MockDisksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDisksClient) EXPECT() *MockDisksClientMockRecorder {
	return m.recorder
}

// DeleteAndWait mocks base method.
func (m *MockDisksClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait.
func (mr *MockDisksClientMockRecorder) DeleteAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockDisksClient)(nil).DeleteAndWait), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockDisksClient) Get(arg0 context.Context, arg1, arg2 string) (compute.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDisksClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDisksClient)(nil).Get), arg0, arg1, arg2)
}

// MockResourceSkusClient is a mock of ResourceSkusClient interface.
type MockResourceSkusClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourceSkusClientMockRecorder
}

// MockResourceSkusClientMockRecorder is the mock recorder for MockResourceSkusClient.
type MockResourceSkusClientMockRecorder struct {
	mock *MockResourceSkusClient
}

// NewMockResourceSkusClient creates a new mock instance.
func NewMockResourceSkusClient(ctrl *gomock.Controller) *MockResourceSkusClient {
	mock := &MockResourceSkusClient{ctrl: ctrl}
	mock.recorder = &MockResourceSkusClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceSkusClient) EXPECT() *MockResourceSkusClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockResourceSkusClient) List(arg0 context.Context, arg1 string) ([]compute.ResourceSku, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]compute.ResourceSku)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockResourceSkusClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceSkusClient)(nil).List), arg0, arg1)
}

// MockVirtualMachinesClient is a mock of VirtualMachinesClient interface.
type MockVirtualMachinesClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachinesClientMockRecorder
}

// MockVirtualMachinesClientMockRecorder is the mock recorder for MockVirtualMachinesClient.
type MockVirtualMachinesClientMockRecorder struct {
	mock *MockVirtualMachinesClient
}

// NewMockVirtualMachinesClient creates a new mock instance.
func NewMockVirtualMachinesClient(ctrl *gomock.Controller) *MockVirtualMachinesClient {
	mock := &MockVirtualMachinesClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachinesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachinesClient) EXPECT() *MockVirtualMachinesClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateAndWait mocks base method.
func (m *MockVirtualMachinesClient) CreateOrUpdateAndWait(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAndWait", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateAndWait indicates an expected call of CreateOrUpdateAndWait.
func (mr *MockVirtualMachinesClientMockRecorder) CreateOrUpdateAndWait(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAndWait", reflect.TypeOf((*MockVirtualMachinesClient)(nil).CreateOrUpdateAndWait), arg0, arg1, arg2, arg3)
}

// DeleteAndWait mocks base method.
func (m *MockVirtualMachinesClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string, arg3 *bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait.
func (mr *MockVirtualMachinesClientMockRecorder) DeleteAndWait(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockVirtualMachinesClient)(nil).DeleteAndWait), arg0, arg1, arg2, arg3)
}

// Get mocks base method.
func (m *MockVirtualMachinesClient) Get(arg0 context.Context, arg1, arg2 string, arg3 compute.InstanceViewTypes) (compute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVirtualMachinesClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualMachinesClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method.
func (m *MockVirtualMachinesClient) List(arg0 context.Context, arg1 string) ([]compute.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]compute.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVirtualMachinesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachinesClient)(nil).List), arg0, arg1)
}

// RedeployAndWait mocks base method.
func (m *MockVirtualMachinesClient) RedeployAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeployAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RedeployAndWait indicates an expected call of RedeployAndWait.
func (mr *MockVirtualMachinesClientMockRecorder) RedeployAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeployAndWait", reflect.TypeOf((*MockVirtualMachinesClient)(nil).RedeployAndWait), arg0, arg1, arg2)
}

// StartAndWait mocks base method.
func (m *MockVirtualMachinesClient) StartAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartAndWait indicates an expected call of StartAndWait.
func (mr *MockVirtualMachinesClientMockRecorder) StartAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAndWait", reflect.TypeOf((*MockVirtualMachinesClient)(nil).StartAndWait), arg0, arg1, arg2)
}

// StopAndWait mocks base method.
func (m *MockVirtualMachinesClient) StopAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopAndWait indicates an expected call of StopAndWait.
func (mr *MockVirtualMachinesClientMockRecorder) StopAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAndWait", reflect.TypeOf((*MockVirtualMachinesClient)(nil).StopAndWait), arg0, arg1, arg2)
}

// MockUsageClient is a mock of UsageClient interface.
type MockUsageClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsageClientMockRecorder
}

// MockUsageClientMockRecorder is the mock recorder for MockUsageClient.
type MockUsageClientMockRecorder struct {
	mock *MockUsageClient
}

// NewMockUsageClient creates a new mock instance.
func NewMockUsageClient(ctrl *gomock.Controller) *MockUsageClient {
	mock := &MockUsageClient{ctrl: ctrl}
	mock.recorder = &MockUsageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsageClient) EXPECT() *MockUsageClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockUsageClient) List(arg0 context.Context, arg1 string) ([]compute.Usage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]compute.Usage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUsageClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsageClient)(nil).List), arg0, arg1)
}

// MockVirtualMachineScaleSetVMsClient is a mock of VirtualMachineScaleSetVMsClient interface.
type MockVirtualMachineScaleSetVMsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetVMsClientMockRecorder
}

// MockVirtualMachineScaleSetVMsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetVMsClient.
type MockVirtualMachineScaleSetVMsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetVMsClient
}

// NewMockVirtualMachineScaleSetVMsClient creates a new mock instance.
func NewMockVirtualMachineScaleSetVMsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetVMsClient {
	mock := &MockVirtualMachineScaleSetVMsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetVMsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachineScaleSetVMsClient) EXPECT() *MockVirtualMachineScaleSetVMsClientMockRecorder {
	return m.recorder
}

// GetInstanceView mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) GetInstanceView(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineScaleSetVMInstanceView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceView", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetVMInstanceView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceView indicates an expected call of GetInstanceView.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) GetInstanceView(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceView", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).GetInstanceView), arg0, arg1, arg2, arg3)
}

// List mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) List(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) ([]compute.VirtualMachineScaleSetVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSetVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RunCommandAndWait mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) RunCommandAndWait(arg0 context.Context, arg1, arg2, arg3 string, arg4 compute.RunCommandInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCommandAndWait", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCommandAndWait indicates an expected call of RunCommandAndWait.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) RunCommandAndWait(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommandAndWait", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).RunCommandAndWait), arg0, arg1, arg2, arg3, arg4)
}

// MockVirtualMachineScaleSetsClient is a mock of VirtualMachineScaleSetsClient interface.
type MockVirtualMachineScaleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetsClientMockRecorder
}

// MockVirtualMachineScaleSetsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetsClient.
type MockVirtualMachineScaleSetsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetsClient
}

// NewMockVirtualMachineScaleSetsClient creates a new mock instance.
func NewMockVirtualMachineScaleSetsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetsClient {
	mock := &MockVirtualMachineScaleSetsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachineScaleSetsClient) EXPECT() *MockVirtualMachineScaleSetsClientMockRecorder {
	return m.recorder
}

// DeleteAndWait mocks base method.
func (m *MockVirtualMachineScaleSetsClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) DeleteAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).DeleteAndWait), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockVirtualMachineScaleSetsClient) List(arg0 context.Context, arg1 string) ([]compute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).List), arg0, arg1)
}

// MockDiskEncryptionSetsClient is a mock of DiskEncryptionSetsClient interface.
type MockDiskEncryptionSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDiskEncryptionSetsClientMockRecorder
}

// MockDiskEncryptionSetsClientMockRecorder is the mock recorder for MockDiskEncryptionSetsClient.
type MockDiskEncryptionSetsClientMockRecorder struct {
	mock *MockDiskEncryptionSetsClient
}

// NewMockDiskEncryptionSetsClient creates a new mock instance.
func NewMockDiskEncryptionSetsClient(ctrl *gomock.Controller) *MockDiskEncryptionSetsClient {
	mock := &MockDiskEncryptionSetsClient{ctrl: ctrl}
	mock.recorder = &MockDiskEncryptionSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiskEncryptionSetsClient) EXPECT() *MockDiskEncryptionSetsClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDiskEncryptionSetsClient) Get(arg0 context.Context, arg1, arg2 string) (compute.DiskEncryptionSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.DiskEncryptionSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDiskEncryptionSetsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDiskEncryptionSetsClient)(nil).Get), arg0, arg1, arg2)
}
