// Code generated by MockGen. DO NOT EDIT.
// Source: downtimes_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	datadogV1 "github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	gomock "github.com/golang/mock/gomock"
)

// MockDowntimesAPIClient is a mock of DowntimesAPIClient interface.
type MockDowntimesAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockDowntimesAPIClientMockRecorder
}

// MockDowntimesAPIClientMockRecorder is the mock recorder for MockDowntimesAPIClient.
type MockDowntimesAPIClientMockRecorder struct {
	mock *MockDowntimesAPIClient
}

// NewMockDowntimesAPIClient creates a new mock instance.
func NewMockDowntimesAPIClient(ctrl *gomock.Controller) *MockDowntimesAPIClient {
	mock := &MockDowntimesAPIClient{ctrl: ctrl}
	mock.recorder = &MockDowntimesAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDowntimesAPIClient) EXPECT() *MockDowntimesAPIClientMockRecorder {
	return m.recorder
}

// ListDowntimes mocks base method.
func (m *MockDowntimesAPIClient) ListDowntimes(arg0 context.Context, arg1 ...datadogV1.ListDowntimesOptionalParameters) ([]datadogV1.Downtime, *http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDowntimes", varargs...)
	ret0, _ := ret[0].([]datadogV1.Downtime)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDowntimes indicates an expected call of ListDowntimes.
func (mr *MockDowntimesAPIClientMockRecorder) ListDowntimes(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDowntimes", reflect.TypeOf((*MockDowntimesAPIClient)(nil).ListDowntimes), varargs...)
}

// ListMonitorDowntimes mocks base method.
func (m *MockDowntimesAPIClient) ListMonitorDowntimes(arg0 context.Context, arg1 int64) ([]datadogV1.Downtime, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMonitorDowntimes", arg0, arg1)
	ret0, _ := ret[0].([]datadogV1.Downtime)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMonitorDowntimes indicates an expected call of ListMonitorDowntimes.
func (mr *MockDowntimesAPIClientMockRecorder) ListMonitorDowntimes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMonitorDowntimes", reflect.TypeOf((*MockDowntimesAPIClient)(nil).ListMonitorDowntimes), arg0, arg1)
}
