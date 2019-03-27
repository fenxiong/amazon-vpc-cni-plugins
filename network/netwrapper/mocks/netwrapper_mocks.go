// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-plugins/network/netwrapper (interfaces: Net)

// Package mock_netwrapper is a generated GoMock package.
package mock_netwrapper

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNet is a mock of Net interface
type MockNet struct {
	ctrl     *gomock.Controller
	recorder *MockNetMockRecorder
}

// MockNetMockRecorder is the mock recorder for MockNet
type MockNetMockRecorder struct {
	mock *MockNet
}

// NewMockNet creates a new mock instance
func NewMockNet(ctrl *gomock.Controller) *MockNet {
	mock := &MockNet{ctrl: ctrl}
	mock.recorder = &MockNetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNet) EXPECT() *MockNetMockRecorder {
	return m.recorder
}

// InterfaceByName mocks base method
func (m *MockNet) InterfaceByName(arg0 string) (*net.Interface, error) {
	ret := m.ctrl.Call(m, "InterfaceByName", arg0)
	ret0, _ := ret[0].(*net.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InterfaceByName indicates an expected call of InterfaceByName
func (mr *MockNetMockRecorder) InterfaceByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InterfaceByName", reflect.TypeOf((*MockNet)(nil).InterfaceByName), arg0)
}

// Interfaces mocks base method
func (m *MockNet) Interfaces() ([]net.Interface, error) {
	ret := m.ctrl.Call(m, "Interfaces")
	ret0, _ := ret[0].([]net.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Interfaces indicates an expected call of Interfaces
func (mr *MockNetMockRecorder) Interfaces() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Interfaces", reflect.TypeOf((*MockNet)(nil).Interfaces))
}
