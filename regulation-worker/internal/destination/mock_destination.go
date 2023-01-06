// Code generated by MockGen. DO NOT EDIT.
// Source: destination.go

// Package destination is a generated GoMock package.
package destination

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	backendconfig "github.com/rudderlabs/rudder-server/config/backend-config"
	pubsub "github.com/rudderlabs/rudder-server/utils/pubsub"
)

// MockdestMiddleware is a mock of destMiddleware interface.
type MockdestMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockdestMiddlewareMockRecorder
}

// MockdestMiddlewareMockRecorder is the mock recorder for MockdestMiddleware.
type MockdestMiddlewareMockRecorder struct {
	mock *MockdestMiddleware
}

// NewMockdestMiddleware creates a new mock instance.
func NewMockdestMiddleware(ctrl *gomock.Controller) *MockdestMiddleware {
	mock := &MockdestMiddleware{ctrl: ctrl}
	mock.recorder = &MockdestMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdestMiddleware) EXPECT() *MockdestMiddlewareMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockdestMiddleware) Subscribe(ctx context.Context, topic backendconfig.Topic) pubsub.DataChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, topic)
	ret0, _ := ret[0].(pubsub.DataChannel)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockdestMiddlewareMockRecorder) Subscribe(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockdestMiddleware)(nil).Subscribe), ctx, topic)
}
