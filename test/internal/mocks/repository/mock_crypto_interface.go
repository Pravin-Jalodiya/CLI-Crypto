// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\akawadia\Downloads\CryptoTracker\internal\repositories\crypto_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	models "cryptotracker/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCryptoRepository is a mock of CryptoRepository interface.
type MockCryptoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoRepositoryMockRecorder
}

// MockCryptoRepositoryMockRecorder is the mock recorder for MockCryptoRepository.
type MockCryptoRepositoryMockRecorder struct {
	mock *MockCryptoRepository
}

// NewMockCryptoRepository creates a new mock instance.
func NewMockCryptoRepository(ctrl *gomock.Controller) *MockCryptoRepository {
	mock := &MockCryptoRepository{ctrl: ctrl}
	mock.recorder = &MockCryptoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoRepository) EXPECT() *MockCryptoRepositoryMockRecorder {
	return m.recorder
}

// DisplayTopCryptocurrencies mocks base method.
func (m *MockCryptoRepository) DisplayTopCryptocurrencies(count int) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisplayTopCryptocurrencies", count)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisplayTopCryptocurrencies indicates an expected call of DisplayTopCryptocurrencies.
func (mr *MockCryptoRepositoryMockRecorder) DisplayTopCryptocurrencies(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplayTopCryptocurrencies", reflect.TypeOf((*MockCryptoRepository)(nil).DisplayTopCryptocurrencies), count)
}

// SearchCryptocurrency mocks base method.
func (m *MockCryptoRepository) SearchCryptocurrency(user *models.User, cryptoSymbol string) (float64, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCryptocurrency", user, cryptoSymbol)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SearchCryptocurrency indicates an expected call of SearchCryptocurrency.
func (mr *MockCryptoRepositoryMockRecorder) SearchCryptocurrency(user, cryptoSymbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCryptocurrency", reflect.TypeOf((*MockCryptoRepository)(nil).SearchCryptocurrency), user, cryptoSymbol)
}

// SetPriceAlert mocks base method.
func (m *MockCryptoRepository) SetPriceAlert(user *models.User, symbol string, targetPrice float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPriceAlert", user, symbol, targetPrice)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetPriceAlert indicates an expected call of SetPriceAlert.
func (mr *MockCryptoRepositoryMockRecorder) SetPriceAlert(user, symbol, targetPrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPriceAlert", reflect.TypeOf((*MockCryptoRepository)(nil).SetPriceAlert), user, symbol, targetPrice)
}
