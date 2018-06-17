// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/almamedia/fargate/elbv2 (interfaces: Client)

// Package client is a generated GoMock package.
package client

import (
	gomock "github.com/golang/mock/gomock"
	elbv2 "github.com/almamedia/fargate/elbv2"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateListener mocks base method
func (m *MockClient) CreateListener(arg0 elbv2.CreateListenerParameters) (string, error) {
	ret := m.ctrl.Call(m, "CreateListener", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateListener indicates an expected call of CreateListener
func (mr *MockClientMockRecorder) CreateListener(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateListener", reflect.TypeOf((*MockClient)(nil).CreateListener), arg0)
}

// CreateLoadBalancer mocks base method
func (m *MockClient) CreateLoadBalancer(arg0 elbv2.CreateLoadBalancerParameters) (string, error) {
	ret := m.ctrl.Call(m, "CreateLoadBalancer", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancer indicates an expected call of CreateLoadBalancer
func (mr *MockClientMockRecorder) CreateLoadBalancer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancer", reflect.TypeOf((*MockClient)(nil).CreateLoadBalancer), arg0)
}

// CreateTargetGroup mocks base method
func (m *MockClient) CreateTargetGroup(arg0 elbv2.CreateTargetGroupParameters) (string, error) {
	ret := m.ctrl.Call(m, "CreateTargetGroup", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTargetGroup indicates an expected call of CreateTargetGroup
func (mr *MockClientMockRecorder) CreateTargetGroup(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTargetGroup", reflect.TypeOf((*MockClient)(nil).CreateTargetGroup), arg0)
}

// DescribeListeners mocks base method
func (m *MockClient) DescribeListeners(arg0 string) (elbv2.Listeners, error) {
	ret := m.ctrl.Call(m, "DescribeListeners", arg0)
	ret0, _ := ret[0].(elbv2.Listeners)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeListeners indicates an expected call of DescribeListeners
func (mr *MockClientMockRecorder) DescribeListeners(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListeners", reflect.TypeOf((*MockClient)(nil).DescribeListeners), arg0)
}

// DescribeLoadBalancers mocks base method
func (m *MockClient) DescribeLoadBalancers() (elbv2.LoadBalancers, error) {
	ret := m.ctrl.Call(m, "DescribeLoadBalancers")
	ret0, _ := ret[0].(elbv2.LoadBalancers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancers indicates an expected call of DescribeLoadBalancers
func (mr *MockClientMockRecorder) DescribeLoadBalancers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockClient)(nil).DescribeLoadBalancers))
}

// DescribeLoadBalancersByName mocks base method
func (m *MockClient) DescribeLoadBalancersByName(arg0 []string) (elbv2.LoadBalancers, error) {
	ret := m.ctrl.Call(m, "DescribeLoadBalancersByName", arg0)
	ret0, _ := ret[0].(elbv2.LoadBalancers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancersByName indicates an expected call of DescribeLoadBalancersByName
func (mr *MockClientMockRecorder) DescribeLoadBalancersByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancersByName", reflect.TypeOf((*MockClient)(nil).DescribeLoadBalancersByName), arg0)
}
