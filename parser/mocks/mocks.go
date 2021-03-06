// Code generated by MockGen. DO NOT EDIT.
// Source: parser/parser.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCSVReader is a mock of CSVReader interface.
type MockCSVReader struct {
	ctrl     *gomock.Controller
	recorder *MockCSVReaderMockRecorder
}

// MockCSVReaderMockRecorder is the mock recorder for MockCSVReader.
type MockCSVReaderMockRecorder struct {
	mock *MockCSVReader
}

// NewMockCSVReader creates a new mock instance.
func NewMockCSVReader(ctrl *gomock.Controller) *MockCSVReader {
	mock := &MockCSVReader{ctrl: ctrl}
	mock.recorder = &MockCSVReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCSVReader) EXPECT() *MockCSVReaderMockRecorder {
	return m.recorder
}

// ReadAll mocks base method.
func (m *MockCSVReader) ReadAll() ([][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll")
	ret0, _ := ret[0].([][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockCSVReaderMockRecorder) ReadAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockCSVReader)(nil).ReadAll))
}
