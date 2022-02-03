// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/plugins/kafka (interfaces: Consumer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	kafka "github.com/conduitio/conduit/pkg/plugins/kafka"
	gomock "github.com/golang/mock/gomock"
	kafka0 "github.com/segmentio/kafka-go"
)

// Consumer is a mock of Consumer interface.
type Consumer struct {
	ctrl     *gomock.Controller
	recorder *ConsumerMockRecorder
}

// ConsumerMockRecorder is the mock recorder for Consumer.
type ConsumerMockRecorder struct {
	mock *Consumer
}

// NewConsumer creates a new mock instance.
func NewConsumer(ctrl *gomock.Controller) *Consumer {
	mock := &Consumer{ctrl: ctrl}
	mock.recorder = &ConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Consumer) EXPECT() *ConsumerMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *Consumer) Ack() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *ConsumerMockRecorder) Ack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*Consumer)(nil).Ack))
}

// Close mocks base method.
func (m *Consumer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *ConsumerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Consumer)(nil).Close))
}

// Get mocks base method.
func (m *Consumer) Get(arg0 context.Context) (*kafka0.Message, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*kafka0.Message)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *ConsumerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Consumer)(nil).Get), arg0)
}

// StartFrom mocks base method.
func (m *Consumer) StartFrom(arg0 kafka.Config, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartFrom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartFrom indicates an expected call of StartFrom.
func (mr *ConsumerMockRecorder) StartFrom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFrom", reflect.TypeOf((*Consumer)(nil).StartFrom), arg0, arg1)
}
