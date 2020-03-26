// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common (interfaces: ReloadSpacesState)

// Package common is a generated GoMock package.
package common

import (
	gomock "github.com/golang/mock/gomock"
	network "github.com/juju/juju/core/network"
	reflect "reflect"
)

// MockReloadSpacesState is a mock of ReloadSpacesState interface
type MockReloadSpacesState struct {
	ctrl     *gomock.Controller
	recorder *MockReloadSpacesStateMockRecorder
}

// MockReloadSpacesStateMockRecorder is the mock recorder for MockReloadSpacesState
type MockReloadSpacesStateMockRecorder struct {
	mock *MockReloadSpacesState
}

// NewMockReloadSpacesState creates a new mock instance
func NewMockReloadSpacesState(ctrl *gomock.Controller) *MockReloadSpacesState {
	mock := &MockReloadSpacesState{ctrl: ctrl}
	mock.recorder = &MockReloadSpacesStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReloadSpacesState) EXPECT() *MockReloadSpacesStateMockRecorder {
	return m.recorder
}

// SaveSpacesFromProvider mocks base method
func (m *MockReloadSpacesState) SaveSpacesFromProvider(arg0 []network.SpaceInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSpacesFromProvider", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSpacesFromProvider indicates an expected call of SaveSpacesFromProvider
func (mr *MockReloadSpacesStateMockRecorder) SaveSpacesFromProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSpacesFromProvider", reflect.TypeOf((*MockReloadSpacesState)(nil).SaveSpacesFromProvider), arg0)
}

// SaveSubnetsFromProvider mocks base method
func (m *MockReloadSpacesState) SaveSubnetsFromProvider(arg0 []network.SubnetInfo, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSubnetsFromProvider", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSubnetsFromProvider indicates an expected call of SaveSubnetsFromProvider
func (mr *MockReloadSpacesStateMockRecorder) SaveSubnetsFromProvider(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSubnetsFromProvider", reflect.TypeOf((*MockReloadSpacesState)(nil).SaveSubnetsFromProvider), arg0, arg1)
}
