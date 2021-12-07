// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jakub-dzon/k4e-operator/internal/repository/edgedeployment (interfaces: Repository)

// Package edgedeployment is a generated GoMock package.
package edgedeployment

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/jakub-dzon/k4e-operator/api/v1alpha1"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// ListForEdgeDevice mocks base method.
func (m *MockRepository) ListForEdgeDevice(arg0 context.Context, arg1, arg2 string) ([]v1alpha1.EdgeDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForEdgeDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].([]v1alpha1.EdgeDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForEdgeDevice indicates an expected call of ListForEdgeDevice.
func (mr *MockRepositoryMockRecorder) ListForEdgeDevice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForEdgeDevice", reflect.TypeOf((*MockRepository)(nil).ListForEdgeDevice), arg0, arg1, arg2)
}

// Patch mocks base method.
func (m *MockRepository) Patch(arg0 context.Context, arg1, arg2 *v1alpha1.EdgeDeployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch.
func (mr *MockRepositoryMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRepository)(nil).Patch), arg0, arg1, arg2)
}

// Read mocks base method.
func (m *MockRepository) Read(arg0 context.Context, arg1, arg2 string) (*v1alpha1.EdgeDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.EdgeDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRepositoryMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRepository)(nil).Read), arg0, arg1, arg2)
}

// RemoveFinalizer mocks base method.
func (m *MockRepository) RemoveFinalizer(arg0 context.Context, arg1 *v1alpha1.EdgeDeployment, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFinalizer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFinalizer indicates an expected call of RemoveFinalizer.
func (mr *MockRepositoryMockRecorder) RemoveFinalizer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFinalizer", reflect.TypeOf((*MockRepository)(nil).RemoveFinalizer), arg0, arg1, arg2)
}
