// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/policy/v1 (interfaces: PodDisruptionBudgetsGetter,PodDisruptionBudgetInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/policy/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/policy/v1"
	v12 "k8s.io/client-go/kubernetes/typed/policy/v1"
)

// MockPodDisruptionBudgetsGetter is a mock of PodDisruptionBudgetsGetter interface.
type MockPodDisruptionBudgetsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPodDisruptionBudgetsGetterMockRecorder
}

// MockPodDisruptionBudgetsGetterMockRecorder is the mock recorder for MockPodDisruptionBudgetsGetter.
type MockPodDisruptionBudgetsGetterMockRecorder struct {
	mock *MockPodDisruptionBudgetsGetter
}

// NewMockPodDisruptionBudgetsGetter creates a new mock instance.
func NewMockPodDisruptionBudgetsGetter(ctrl *gomock.Controller) *MockPodDisruptionBudgetsGetter {
	mock := &MockPodDisruptionBudgetsGetter{ctrl: ctrl}
	mock.recorder = &MockPodDisruptionBudgetsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodDisruptionBudgetsGetter) EXPECT() *MockPodDisruptionBudgetsGetterMockRecorder {
	return m.recorder
}

// PodDisruptionBudgets mocks base method.
func (m *MockPodDisruptionBudgetsGetter) PodDisruptionBudgets(arg0 string) v12.PodDisruptionBudgetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodDisruptionBudgets", arg0)
	ret0, _ := ret[0].(v12.PodDisruptionBudgetInterface)
	return ret0
}

// PodDisruptionBudgets indicates an expected call of PodDisruptionBudgets.
func (mr *MockPodDisruptionBudgetsGetterMockRecorder) PodDisruptionBudgets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodDisruptionBudgets", reflect.TypeOf((*MockPodDisruptionBudgetsGetter)(nil).PodDisruptionBudgets), arg0)
}

// MockPodDisruptionBudgetInterface is a mock of PodDisruptionBudgetInterface interface.
type MockPodDisruptionBudgetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPodDisruptionBudgetInterfaceMockRecorder
}

// MockPodDisruptionBudgetInterfaceMockRecorder is the mock recorder for MockPodDisruptionBudgetInterface.
type MockPodDisruptionBudgetInterfaceMockRecorder struct {
	mock *MockPodDisruptionBudgetInterface
}

// NewMockPodDisruptionBudgetInterface creates a new mock instance.
func NewMockPodDisruptionBudgetInterface(ctrl *gomock.Controller) *MockPodDisruptionBudgetInterface {
	mock := &MockPodDisruptionBudgetInterface{ctrl: ctrl}
	mock.recorder = &MockPodDisruptionBudgetInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodDisruptionBudgetInterface) EXPECT() *MockPodDisruptionBudgetInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockPodDisruptionBudgetInterface) Apply(arg0 context.Context, arg1 *v11.PodDisruptionBudgetApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PodDisruptionBudget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodDisruptionBudget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockPodDisruptionBudgetInterface) ApplyStatus(arg0 context.Context, arg1 *v11.PodDisruptionBudgetApplyConfiguration, arg2 v10.ApplyOptions) (*v1.PodDisruptionBudget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodDisruptionBudget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockPodDisruptionBudgetInterface) Create(arg0 context.Context, arg1 *v1.PodDisruptionBudget, arg2 v10.CreateOptions) (*v1.PodDisruptionBudget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodDisruptionBudget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockPodDisruptionBudgetInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockPodDisruptionBudgetInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockPodDisruptionBudgetInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.PodDisruptionBudget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodDisruptionBudget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockPodDisruptionBudgetInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.PodDisruptionBudgetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.PodDisruptionBudgetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockPodDisruptionBudgetInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.PodDisruptionBudget, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.PodDisruptionBudget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockPodDisruptionBudgetInterface) Update(arg0 context.Context, arg1 *v1.PodDisruptionBudget, arg2 v10.UpdateOptions) (*v1.PodDisruptionBudget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodDisruptionBudget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockPodDisruptionBudgetInterface) UpdateStatus(arg0 context.Context, arg1 *v1.PodDisruptionBudget, arg2 v10.UpdateOptions) (*v1.PodDisruptionBudget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.PodDisruptionBudget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockPodDisruptionBudgetInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockPodDisruptionBudgetInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockPodDisruptionBudgetInterface)(nil).Watch), arg0, arg1)
}
