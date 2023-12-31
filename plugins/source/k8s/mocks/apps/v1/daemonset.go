// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/apps/v1 (interfaces: DaemonSetsGetter,DaemonSetInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/apps/v1"
	v12 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// MockDaemonSetsGetter is a mock of DaemonSetsGetter interface.
type MockDaemonSetsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetsGetterMockRecorder
}

// MockDaemonSetsGetterMockRecorder is the mock recorder for MockDaemonSetsGetter.
type MockDaemonSetsGetterMockRecorder struct {
	mock *MockDaemonSetsGetter
}

// NewMockDaemonSetsGetter creates a new mock instance.
func NewMockDaemonSetsGetter(ctrl *gomock.Controller) *MockDaemonSetsGetter {
	mock := &MockDaemonSetsGetter{ctrl: ctrl}
	mock.recorder = &MockDaemonSetsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetsGetter) EXPECT() *MockDaemonSetsGetterMockRecorder {
	return m.recorder
}

// DaemonSets mocks base method.
func (m *MockDaemonSetsGetter) DaemonSets(arg0 string) v12.DaemonSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DaemonSets", arg0)
	ret0, _ := ret[0].(v12.DaemonSetInterface)
	return ret0
}

// DaemonSets indicates an expected call of DaemonSets.
func (mr *MockDaemonSetsGetterMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaemonSets", reflect.TypeOf((*MockDaemonSetsGetter)(nil).DaemonSets), arg0)
}

// MockDaemonSetInterface is a mock of DaemonSetInterface interface.
type MockDaemonSetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonSetInterfaceMockRecorder
}

// MockDaemonSetInterfaceMockRecorder is the mock recorder for MockDaemonSetInterface.
type MockDaemonSetInterfaceMockRecorder struct {
	mock *MockDaemonSetInterface
}

// NewMockDaemonSetInterface creates a new mock instance.
func NewMockDaemonSetInterface(ctrl *gomock.Controller) *MockDaemonSetInterface {
	mock := &MockDaemonSetInterface{ctrl: ctrl}
	mock.recorder = &MockDaemonSetInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonSetInterface) EXPECT() *MockDaemonSetInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockDaemonSetInterface) Apply(arg0 context.Context, arg1 *v11.DaemonSetApplyConfiguration, arg2 v10.ApplyOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockDaemonSetInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockDaemonSetInterface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockDaemonSetInterface) ApplyStatus(arg0 context.Context, arg1 *v11.DaemonSetApplyConfiguration, arg2 v10.ApplyOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockDaemonSetInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockDaemonSetInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockDaemonSetInterface) Create(arg0 context.Context, arg1 *v1.DaemonSet, arg2 v10.CreateOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDaemonSetInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDaemonSetInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockDaemonSetInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDaemonSetInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDaemonSetInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockDaemonSetInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockDaemonSetInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockDaemonSetInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockDaemonSetInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDaemonSetInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDaemonSetInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockDaemonSetInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.DaemonSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.DaemonSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDaemonSetInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDaemonSetInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockDaemonSetInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockDaemonSetInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockDaemonSetInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockDaemonSetInterface) Update(arg0 context.Context, arg1 *v1.DaemonSet, arg2 v10.UpdateOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDaemonSetInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDaemonSetInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockDaemonSetInterface) UpdateStatus(arg0 context.Context, arg1 *v1.DaemonSet, arg2 v10.UpdateOptions) (*v1.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockDaemonSetInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockDaemonSetInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockDaemonSetInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockDaemonSetInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockDaemonSetInterface)(nil).Watch), arg0, arg1)
}
