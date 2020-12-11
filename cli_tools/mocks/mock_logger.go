// Code generated by MockGen. DO NOT EDIT.
// Source: logger.go

// Package mocks is a generated GoMock package.
package mocks

import (
	pb "github.com/GoogleCloudPlatform/compute-image-tools/proto/go/pb"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// User mocks base method
func (m *MockLogger) User(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "User", message)
}

// User indicates an expected call of User
func (mr *MockLoggerMockRecorder) User(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockLogger)(nil).User), message)
}

// Debug mocks base method
func (m *MockLogger) Debug(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Debug", message)
}

// Debug indicates an expected call of Debug
func (mr *MockLoggerMockRecorder) Debug(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), message)
}

// Trace mocks base method
func (m *MockLogger) Trace(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Trace", message)
}

// Trace indicates an expected call of Trace
func (mr *MockLoggerMockRecorder) Trace(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trace", reflect.TypeOf((*MockLogger)(nil).Trace), message)
}

// Metric mocks base method
func (m *MockLogger) Metric(metric *pb.OutputInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Metric", metric)
}

// Metric indicates an expected call of Metric
func (mr *MockLoggerMockRecorder) Metric(metric interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metric", reflect.TypeOf((*MockLogger)(nil).Metric), metric)
}

// MockOutputInfoReader is a mock of OutputInfoReader interface
type MockOutputInfoReader struct {
	ctrl     *gomock.Controller
	recorder *MockOutputInfoReaderMockRecorder
}

// MockOutputInfoReaderMockRecorder is the mock recorder for MockOutputInfoReader
type MockOutputInfoReaderMockRecorder struct {
	mock *MockOutputInfoReader
}

// NewMockOutputInfoReader creates a new mock instance
func NewMockOutputInfoReader(ctrl *gomock.Controller) *MockOutputInfoReader {
	mock := &MockOutputInfoReader{ctrl: ctrl}
	mock.recorder = &MockOutputInfoReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOutputInfoReader) EXPECT() *MockOutputInfoReaderMockRecorder {
	return m.recorder
}

// read mocks base method
func (m *MockOutputInfoReader) read() *pb.OutputInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "read")
	ret0, _ := ret[0].(*pb.OutputInfo)
	return ret0
}

// read indicates an expected call of read
func (mr *MockOutputInfoReaderMockRecorder) read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "read", reflect.TypeOf((*MockOutputInfoReader)(nil).read))
}

// MockToolLogger is a mock of ToolLogger interface
type MockToolLogger struct {
	ctrl     *gomock.Controller
	recorder *MockToolLoggerMockRecorder
}

// MockToolLoggerMockRecorder is the mock recorder for MockToolLogger
type MockToolLoggerMockRecorder struct {
	mock *MockToolLogger
}

// NewMockToolLogger creates a new mock instance
func NewMockToolLogger(ctrl *gomock.Controller) *MockToolLogger {
	mock := &MockToolLogger{ctrl: ctrl}
	mock.recorder = &MockToolLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockToolLogger) EXPECT() *MockToolLoggerMockRecorder {
	return m.recorder
}

// User mocks base method
func (m *MockToolLogger) User(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "User", message)
}

// User indicates an expected call of User
func (mr *MockToolLoggerMockRecorder) User(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockToolLogger)(nil).User), message)
}

// Debug mocks base method
func (m *MockToolLogger) Debug(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Debug", message)
}

// Debug indicates an expected call of Debug
func (mr *MockToolLoggerMockRecorder) Debug(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockToolLogger)(nil).Debug), message)
}

// Trace mocks base method
func (m *MockToolLogger) Trace(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Trace", message)
}

// Trace indicates an expected call of Trace
func (mr *MockToolLoggerMockRecorder) Trace(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trace", reflect.TypeOf((*MockToolLogger)(nil).Trace), message)
}

// Metric mocks base method
func (m *MockToolLogger) Metric(metric *pb.OutputInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Metric", metric)
}

// Metric indicates an expected call of Metric
func (mr *MockToolLoggerMockRecorder) Metric(metric interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metric", reflect.TypeOf((*MockToolLogger)(nil).Metric), metric)
}

// read mocks base method
func (m *MockToolLogger) read() *pb.OutputInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "read")
	ret0, _ := ret[0].(*pb.OutputInfo)
	return ret0
}

// read indicates an expected call of read
func (mr *MockToolLoggerMockRecorder) read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "read", reflect.TypeOf((*MockToolLogger)(nil).read))
}
