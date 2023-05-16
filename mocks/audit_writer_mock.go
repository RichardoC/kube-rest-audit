// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/RichardoC/kube-audit-rest/internal/audit_writer (interfaces: AuditWritter)

// Package mymock is a generated GoMock package.
package mymock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuditWritter is a mock of AuditWritter interface.
type MockAuditWritter struct {
	ctrl     *gomock.Controller
	recorder *MockAuditWritterMockRecorder
}

// MockAuditWritterMockRecorder is the mock recorder for MockAuditWritter.
type MockAuditWritterMockRecorder struct {
	mock *MockAuditWritter
}

// NewMockAuditWritter creates a new mock instance.
func NewMockAuditWritter(ctrl *gomock.Controller) *MockAuditWritter {
	mock := &MockAuditWritter{ctrl: ctrl}
	mock.recorder = &MockAuditWritterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditWritter) EXPECT() *MockAuditWritterMockRecorder {
	return m.recorder
}

// LogEvent mocks base method.
func (m *MockAuditWritter) LogEvent(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogEvent", arg0)
}

// LogEvent indicates an expected call of LogEvent.
func (mr *MockAuditWritterMockRecorder) LogEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogEvent", reflect.TypeOf((*MockAuditWritter)(nil).LogEvent), arg0)
}

// Sync mocks base method.
func (m *MockAuditWritter) Sync() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sync")
}

// Sync indicates an expected call of Sync.
func (mr *MockAuditWritterMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockAuditWritter)(nil).Sync))
}
