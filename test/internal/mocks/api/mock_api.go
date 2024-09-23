// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\akawadia\Downloads\CryptoTracker\internal\api\api.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAPIClient is a mock of APIClient interface.
type MockAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockAPIClientMockRecorder
}

// MockAPIClientMockRecorder is the mock recorder for MockAPIClient.
type MockAPIClientMockRecorder struct {
	mock *MockAPIClient
}

// NewMockAPIClient creates a new mock instance.
func NewMockAPIClient(ctrl *gomock.Controller) *MockAPIClient {
	mock := &MockAPIClient{ctrl: ctrl}
	mock.recorder = &MockAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIClient) EXPECT() *MockAPIClientMockRecorder {
	return m.recorder
}

// GetAPIResponse mocks base method.
func (m *MockAPIClient) GetAPIResponse(endpoint string, params map[string]string) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIResponse", endpoint, params)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetAPIResponse indicates an expected call of GetAPIResponse.
func (mr *MockAPIClientMockRecorder) GetAPIResponse(endpoint, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIResponse", reflect.TypeOf((*MockAPIClient)(nil).GetAPIResponse), endpoint, params)
}
