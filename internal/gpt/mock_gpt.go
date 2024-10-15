// Code generated by MockGen. DO NOT EDIT.
// Source: internal/gpt/gpt.go

// Package gpt is a generated GoMock package.
package gpt

import (
	context "context"
	reflect "reflect"

  "go.uber.org/mock/gomock"
)

// MockGPTModelClient is a mock of GPTModelClient interface.
type MockGPTModelClient struct {
	ctrl     *gomock.Controller
	recorder *MockGPTModelClientMockRecorder
}

// MockGPTModelClientMockRecorder is the mock recorder for MockGPTModelClient.
type MockGPTModelClientMockRecorder struct {
	mock *MockGPTModelClient
}

// NewMockGPTModelClient creates a new mock instance.
func NewMockGPTModelClient(ctrl *gomock.Controller) *MockGPTModelClient {
	mock := &MockGPTModelClient{ctrl: ctrl}
	mock.recorder = &MockGPTModelClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGPTModelClient) EXPECT() *MockGPTModelClientMockRecorder {
	return m.recorder
}

// FetchChangelog mocks base method.
func (m *MockGPTModelClient) FetchChangelog(ctx context.Context, commits string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchChangelog", ctx, commits)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchChangelog indicates an expected call of FetchChangelog.
func (mr *MockGPTModelClientMockRecorder) FetchChangelog(ctx, commits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchChangelog", reflect.TypeOf((*MockGPTModelClient)(nil).FetchChangelog), ctx, commits)
}

// FetchCommitMessage mocks base method.
func (m *MockGPTModelClient) FetchCommitMessage(ctx context.Context, diff string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCommitMessage", ctx, diff)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCommitMessage indicates an expected call of FetchCommitMessage.
func (mr *MockGPTModelClientMockRecorder) FetchCommitMessage(ctx, diff interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCommitMessage", reflect.TypeOf((*MockGPTModelClient)(nil).FetchCommitMessage), ctx, diff)
}
