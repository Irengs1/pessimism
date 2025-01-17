// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/engine (interfaces: Manager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	core "github.com/base-org/pessimism/internal/core"
	invariant "github.com/base-org/pessimism/internal/engine/invariant"
	gomock "github.com/golang/mock/gomock"
)

// EngineManager is a mock of Manager interface.
type EngineManager struct {
	ctrl     *gomock.Controller
	recorder *EngineManagerMockRecorder
}

// EngineManagerMockRecorder is the mock recorder for EngineManager.
type EngineManagerMockRecorder struct {
	mock *EngineManager
}

// NewEngineManager creates a new mock instance.
func NewEngineManager(ctrl *gomock.Controller) *EngineManager {
	mock := &EngineManager{ctrl: ctrl}
	mock.recorder = &EngineManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *EngineManager) EXPECT() *EngineManagerMockRecorder {
	return m.recorder
}

// DeleteInvariantSession mocks base method.
func (m *EngineManager) DeleteInvariantSession(arg0 core.InvSessionUUID) (core.InvSessionUUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvariantSession", arg0)
	ret0, _ := ret[0].(core.InvSessionUUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInvariantSession indicates an expected call of DeleteInvariantSession.
func (mr *EngineManagerMockRecorder) DeleteInvariantSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvariantSession", reflect.TypeOf((*EngineManager)(nil).DeleteInvariantSession), arg0)
}

// DeployInvariantSession mocks base method.
func (m *EngineManager) DeployInvariantSession(arg0 *invariant.DeployConfig) (core.InvSessionUUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployInvariantSession", arg0)
	ret0, _ := ret[0].(core.InvSessionUUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeployInvariantSession indicates an expected call of DeployInvariantSession.
func (mr *EngineManagerMockRecorder) DeployInvariantSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployInvariantSession", reflect.TypeOf((*EngineManager)(nil).DeployInvariantSession), arg0)
}

// EventLoop mocks base method.
func (m *EngineManager) EventLoop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventLoop")
	ret0, _ := ret[0].(error)
	return ret0
}

// EventLoop indicates an expected call of EventLoop.
func (mr *EngineManagerMockRecorder) EventLoop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventLoop", reflect.TypeOf((*EngineManager)(nil).EventLoop))
}

// Shutdown mocks base method.
func (m *EngineManager) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *EngineManagerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*EngineManager)(nil).Shutdown))
}

// Transit mocks base method.
func (m *EngineManager) Transit() chan core.InvariantInput {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transit")
	ret0, _ := ret[0].(chan core.InvariantInput)
	return ret0
}

// Transit indicates an expected call of Transit.
func (mr *EngineManagerMockRecorder) Transit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transit", reflect.TypeOf((*EngineManager)(nil).Transit))
}
