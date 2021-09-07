// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-game-api/internal/metrics (interfaces: Metrics)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// AddSuccessMultiCreateHeroesCounter mocks base method.
func (m *MockMetrics) AddSuccessMultiCreateHeroesCounter(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSuccessMultiCreateHeroesCounter", arg0)
}

// AddSuccessMultiCreateHeroesCounter indicates an expected call of AddSuccessMultiCreateHeroesCounter.
func (mr *MockMetricsMockRecorder) AddSuccessMultiCreateHeroesCounter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSuccessMultiCreateHeroesCounter", reflect.TypeOf((*MockMetrics)(nil).AddSuccessMultiCreateHeroesCounter), arg0)
}

// IncSuccessCreateHeroCounter mocks base method.
func (m *MockMetrics) IncSuccessCreateHeroCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncSuccessCreateHeroCounter")
}

// IncSuccessCreateHeroCounter indicates an expected call of IncSuccessCreateHeroCounter.
func (mr *MockMetricsMockRecorder) IncSuccessCreateHeroCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncSuccessCreateHeroCounter", reflect.TypeOf((*MockMetrics)(nil).IncSuccessCreateHeroCounter))
}

// IncSuccessDescribeHeroCounter mocks base method.
func (m *MockMetrics) IncSuccessDescribeHeroCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncSuccessDescribeHeroCounter")
}

// IncSuccessDescribeHeroCounter indicates an expected call of IncSuccessDescribeHeroCounter.
func (mr *MockMetricsMockRecorder) IncSuccessDescribeHeroCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncSuccessDescribeHeroCounter", reflect.TypeOf((*MockMetrics)(nil).IncSuccessDescribeHeroCounter))
}

// IncSuccessListHeroesCounter mocks base method.
func (m *MockMetrics) IncSuccessListHeroesCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncSuccessListHeroesCounter")
}

// IncSuccessListHeroesCounter indicates an expected call of IncSuccessListHeroesCounter.
func (mr *MockMetricsMockRecorder) IncSuccessListHeroesCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncSuccessListHeroesCounter", reflect.TypeOf((*MockMetrics)(nil).IncSuccessListHeroesCounter))
}

// IncSuccessRemoveHeroCounter mocks base method.
func (m *MockMetrics) IncSuccessRemoveHeroCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncSuccessRemoveHeroCounter")
}

// IncSuccessRemoveHeroCounter indicates an expected call of IncSuccessRemoveHeroCounter.
func (mr *MockMetricsMockRecorder) IncSuccessRemoveHeroCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncSuccessRemoveHeroCounter", reflect.TypeOf((*MockMetrics)(nil).IncSuccessRemoveHeroCounter))
}

// IncSuccessUpdateHeroCounter mocks base method.
func (m *MockMetrics) IncSuccessUpdateHeroCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncSuccessUpdateHeroCounter")
}

// IncSuccessUpdateHeroCounter indicates an expected call of IncSuccessUpdateHeroCounter.
func (mr *MockMetricsMockRecorder) IncSuccessUpdateHeroCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncSuccessUpdateHeroCounter", reflect.TypeOf((*MockMetrics)(nil).IncSuccessUpdateHeroCounter))
}