/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/apiserver-network-proxy/proto/agent (interfaces: AgentService_ConnectServer)

// Package mock_agent is a generated GoMock package.
package mock_agent

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
	client "sigs.k8s.io/apiserver-network-proxy/konnectivity-client/proto/client"
)

// MockAgentService_ConnectServer is a mock of AgentService_ConnectServer interface.
type MockAgentService_ConnectServer struct {
	ctrl     *gomock.Controller
	recorder *MockAgentService_ConnectServerMockRecorder
}

// MockAgentService_ConnectServerMockRecorder is the mock recorder for MockAgentService_ConnectServer.
type MockAgentService_ConnectServerMockRecorder struct {
	mock *MockAgentService_ConnectServer
}

// NewMockAgentService_ConnectServer creates a new mock instance.
func NewMockAgentService_ConnectServer(ctrl *gomock.Controller) *MockAgentService_ConnectServer {
	mock := &MockAgentService_ConnectServer{ctrl: ctrl}
	mock.recorder = &MockAgentService_ConnectServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentService_ConnectServer) EXPECT() *MockAgentService_ConnectServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockAgentService_ConnectServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAgentService_ConnectServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockAgentService_ConnectServer) Recv() (*client.Packet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*client.Packet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockAgentService_ConnectServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockAgentService_ConnectServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAgentService_ConnectServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockAgentService_ConnectServer) Send(arg0 *client.Packet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockAgentService_ConnectServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockAgentService_ConnectServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockAgentService_ConnectServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockAgentService_ConnectServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAgentService_ConnectServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockAgentService_ConnectServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockAgentService_ConnectServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockAgentService_ConnectServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockAgentService_ConnectServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockAgentService_ConnectServer)(nil).SetTrailer), arg0)
}
