// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/krobus00/krokit (interfaces: OpensearchClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	strings "strings"

	gomock "github.com/golang/mock/gomock"
	kit "github.com/krobus00/krokit"
	opensearchapi "github.com/opensearch-project/opensearch-go/opensearchapi"
)

// MockOpensearchClient is a mock of OpensearchClient interface.
type MockOpensearchClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpensearchClientMockRecorder
}

// MockOpensearchClientMockRecorder is the mock recorder for MockOpensearchClient.
type MockOpensearchClientMockRecorder struct {
	mock *MockOpensearchClient
}

// NewMockOpensearchClient creates a new mock instance.
func NewMockOpensearchClient(ctrl *gomock.Controller) *MockOpensearchClient {
	mock := &MockOpensearchClient{ctrl: ctrl}
	mock.recorder = &MockOpensearchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpensearchClient) EXPECT() *MockOpensearchClientMockRecorder {
	return m.recorder
}

// CreateIndices mocks base method.
func (m *MockOpensearchClient) CreateIndices(arg0 context.Context, arg1 string, arg2 *strings.Reader) (*opensearchapi.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndices", arg0, arg1, arg2)
	ret0, _ := ret[0].(*opensearchapi.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIndices indicates an expected call of CreateIndices.
func (mr *MockOpensearchClientMockRecorder) CreateIndices(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndices", reflect.TypeOf((*MockOpensearchClient)(nil).CreateIndices), arg0, arg1, arg2)
}

// Index mocks base method.
func (m *MockOpensearchClient) Index(arg0 context.Context, arg1 string, arg2 kit.IndexModel) (*opensearchapi.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", arg0, arg1, arg2)
	ret0, _ := ret[0].(*opensearchapi.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Index indicates an expected call of Index.
func (mr *MockOpensearchClientMockRecorder) Index(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockOpensearchClient)(nil).Index), arg0, arg1, arg2)
}

// PutIndicesMapping mocks base method.
func (m *MockOpensearchClient) PutIndicesMapping(arg0 context.Context, arg1 []string, arg2 *strings.Reader) (*opensearchapi.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutIndicesMapping", arg0, arg1, arg2)
	ret0, _ := ret[0].(*opensearchapi.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutIndicesMapping indicates an expected call of PutIndicesMapping.
func (mr *MockOpensearchClientMockRecorder) PutIndicesMapping(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutIndicesMapping", reflect.TypeOf((*MockOpensearchClient)(nil).PutIndicesMapping), arg0, arg1, arg2)
}

// Search mocks base method.
func (m *MockOpensearchClient) Search(arg0 context.Context, arg1 []string, arg2 *strings.Reader) (*opensearchapi.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1, arg2)
	ret0, _ := ret[0].(*opensearchapi.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockOpensearchClientMockRecorder) Search(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockOpensearchClient)(nil).Search), arg0, arg1, arg2)
}
