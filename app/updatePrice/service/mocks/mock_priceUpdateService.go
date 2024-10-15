// Code generated by MockGen. DO NOT EDIT.
// Source: priceUpdateService.go

// Package mock_priceUpdateService is a generated GoMock package.
package mock_priceUpdateService

import (
	updatePriceModel "product-information-updater/app/updatePrice/models"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockService) Process(ginCtx *gin.Context, productID string, request updatePriceModel.RequestBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", ginCtx, productID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *MockServiceMockRecorder) Process(ginCtx, productID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockService)(nil).Process), ginCtx, productID, request)
}

// PublishToSNS mocks base method.
func (m *MockService) PublishToSNS(message updatePriceModel.ProductEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishToSNS", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishToSNS indicates an expected call of PublishToSNS.
func (mr *MockServiceMockRecorder) PublishToSNS(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishToSNS", reflect.TypeOf((*MockService)(nil).PublishToSNS), message)
}

// SaveToDb mocks base method.
func (m *MockService) SaveToDb(request updatePriceModel.ProductEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveToDb", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveToDb indicates an expected call of SaveToDb.
func (mr *MockServiceMockRecorder) SaveToDb(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveToDb", reflect.TypeOf((*MockService)(nil).SaveToDb), request)
}
