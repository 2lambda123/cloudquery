// Code generated by MockGen. DO NOT EDIT.
// Source: appflow.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	appflow "github.com/aws/aws-sdk-go-v2/service/appflow"
	gomock "github.com/golang/mock/gomock"
)

// MockAppflowClient is a mock of AppflowClient interface.
type MockAppflowClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppflowClientMockRecorder
}

// MockAppflowClientMockRecorder is the mock recorder for MockAppflowClient.
type MockAppflowClientMockRecorder struct {
	mock *MockAppflowClient
}

// NewMockAppflowClient creates a new mock instance.
func NewMockAppflowClient(ctrl *gomock.Controller) *MockAppflowClient {
	mock := &MockAppflowClient{ctrl: ctrl}
	mock.recorder = &MockAppflowClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppflowClient) EXPECT() *MockAppflowClientMockRecorder {
	return m.recorder
}

// DescribeConnector mocks base method.
func (m *MockAppflowClient) DescribeConnector(arg0 context.Context, arg1 *appflow.DescribeConnectorInput, arg2 ...func(*appflow.Options)) (*appflow.DescribeConnectorOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeConnector")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnector", varargs...)
	ret0, _ := ret[0].(*appflow.DescribeConnectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnector indicates an expected call of DescribeConnector.
func (mr *MockAppflowClientMockRecorder) DescribeConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnector", reflect.TypeOf((*MockAppflowClient)(nil).DescribeConnector), varargs...)
}

// DescribeConnectorEntity mocks base method.
func (m *MockAppflowClient) DescribeConnectorEntity(arg0 context.Context, arg1 *appflow.DescribeConnectorEntityInput, arg2 ...func(*appflow.Options)) (*appflow.DescribeConnectorEntityOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeConnectorEntity")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectorEntity", varargs...)
	ret0, _ := ret[0].(*appflow.DescribeConnectorEntityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectorEntity indicates an expected call of DescribeConnectorEntity.
func (mr *MockAppflowClientMockRecorder) DescribeConnectorEntity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectorEntity", reflect.TypeOf((*MockAppflowClient)(nil).DescribeConnectorEntity), varargs...)
}

// DescribeConnectorProfiles mocks base method.
func (m *MockAppflowClient) DescribeConnectorProfiles(arg0 context.Context, arg1 *appflow.DescribeConnectorProfilesInput, arg2 ...func(*appflow.Options)) (*appflow.DescribeConnectorProfilesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeConnectorProfiles")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectorProfiles", varargs...)
	ret0, _ := ret[0].(*appflow.DescribeConnectorProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectorProfiles indicates an expected call of DescribeConnectorProfiles.
func (mr *MockAppflowClientMockRecorder) DescribeConnectorProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectorProfiles", reflect.TypeOf((*MockAppflowClient)(nil).DescribeConnectorProfiles), varargs...)
}

// DescribeConnectors mocks base method.
func (m *MockAppflowClient) DescribeConnectors(arg0 context.Context, arg1 *appflow.DescribeConnectorsInput, arg2 ...func(*appflow.Options)) (*appflow.DescribeConnectorsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeConnectors")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConnectors", varargs...)
	ret0, _ := ret[0].(*appflow.DescribeConnectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConnectors indicates an expected call of DescribeConnectors.
func (mr *MockAppflowClientMockRecorder) DescribeConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConnectors", reflect.TypeOf((*MockAppflowClient)(nil).DescribeConnectors), varargs...)
}

// DescribeFlow mocks base method.
func (m *MockAppflowClient) DescribeFlow(arg0 context.Context, arg1 *appflow.DescribeFlowInput, arg2 ...func(*appflow.Options)) (*appflow.DescribeFlowOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFlow")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFlow", varargs...)
	ret0, _ := ret[0].(*appflow.DescribeFlowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFlow indicates an expected call of DescribeFlow.
func (mr *MockAppflowClientMockRecorder) DescribeFlow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFlow", reflect.TypeOf((*MockAppflowClient)(nil).DescribeFlow), varargs...)
}

// DescribeFlowExecutionRecords mocks base method.
func (m *MockAppflowClient) DescribeFlowExecutionRecords(arg0 context.Context, arg1 *appflow.DescribeFlowExecutionRecordsInput, arg2 ...func(*appflow.Options)) (*appflow.DescribeFlowExecutionRecordsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeFlowExecutionRecords")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFlowExecutionRecords", varargs...)
	ret0, _ := ret[0].(*appflow.DescribeFlowExecutionRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFlowExecutionRecords indicates an expected call of DescribeFlowExecutionRecords.
func (mr *MockAppflowClientMockRecorder) DescribeFlowExecutionRecords(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFlowExecutionRecords", reflect.TypeOf((*MockAppflowClient)(nil).DescribeFlowExecutionRecords), varargs...)
}

// ListConnectorEntities mocks base method.
func (m *MockAppflowClient) ListConnectorEntities(arg0 context.Context, arg1 *appflow.ListConnectorEntitiesInput, arg2 ...func(*appflow.Options)) (*appflow.ListConnectorEntitiesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListConnectorEntities")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConnectorEntities", varargs...)
	ret0, _ := ret[0].(*appflow.ListConnectorEntitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectorEntities indicates an expected call of ListConnectorEntities.
func (mr *MockAppflowClientMockRecorder) ListConnectorEntities(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectorEntities", reflect.TypeOf((*MockAppflowClient)(nil).ListConnectorEntities), varargs...)
}

// ListConnectors mocks base method.
func (m *MockAppflowClient) ListConnectors(arg0 context.Context, arg1 *appflow.ListConnectorsInput, arg2 ...func(*appflow.Options)) (*appflow.ListConnectorsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListConnectors")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConnectors", varargs...)
	ret0, _ := ret[0].(*appflow.ListConnectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectors indicates an expected call of ListConnectors.
func (mr *MockAppflowClientMockRecorder) ListConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectors", reflect.TypeOf((*MockAppflowClient)(nil).ListConnectors), varargs...)
}

// ListFlows mocks base method.
func (m *MockAppflowClient) ListFlows(arg0 context.Context, arg1 *appflow.ListFlowsInput, arg2 ...func(*appflow.Options)) (*appflow.ListFlowsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListFlows")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFlows", varargs...)
	ret0, _ := ret[0].(*appflow.ListFlowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFlows indicates an expected call of ListFlows.
func (mr *MockAppflowClientMockRecorder) ListFlows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFlows", reflect.TypeOf((*MockAppflowClient)(nil).ListFlows), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAppflowClient) ListTagsForResource(arg0 context.Context, arg1 *appflow.ListTagsForResourceInput, arg2 ...func(*appflow.Options)) (*appflow.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appflow.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*appflow.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAppflowClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppflowClient)(nil).ListTagsForResource), varargs...)
}
