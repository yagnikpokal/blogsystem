// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/routes.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRoutes is a mock of Routes interface.
type MockRoutes struct {
	ctrl     *gomock.Controller
	recorder *MockRoutesMockRecorder
}

// MockRoutesMockRecorder is the mock recorder for MockRoutes.
type MockRoutesMockRecorder struct {
	mock *MockRoutes
}

// NewMockRoutes creates a new mock instance.
func NewMockRoutes(ctrl *gomock.Controller) *MockRoutes {
	mock := &MockRoutes{ctrl: ctrl}
	mock.recorder = &MockRoutesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutes) EXPECT() *MockRoutesMockRecorder {
	return m.recorder
}

// AllArticle mocks base method.
func (m *MockRoutes) AllArticle(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AllArticle", w, r)
}

// AllArticle indicates an expected call of AllArticle.
func (mr *MockRoutesMockRecorder) AllArticle(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllArticle", reflect.TypeOf((*MockRoutes)(nil).AllArticle), w, r)
}

// GetArticle mocks base method.
func (m *MockRoutes) GetArticle(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetArticle", w, r)
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockRoutesMockRecorder) GetArticle(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockRoutes)(nil).GetArticle), w, r)
}

// HealthCheck mocks base method.
func (m *MockRoutes) HealthCheck(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HealthCheck", w, r)
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockRoutesMockRecorder) HealthCheck(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockRoutes)(nil).HealthCheck), w, r)
}

// InsertArticle mocks base method.
func (m *MockRoutes) InsertArticle(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertArticle", w, r)
}

// InsertArticle indicates an expected call of InsertArticle.
func (mr *MockRoutesMockRecorder) InsertArticle(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertArticle", reflect.TypeOf((*MockRoutes)(nil).InsertArticle), w, r)
}
