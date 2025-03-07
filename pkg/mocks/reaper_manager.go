// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/k8ssandra/k8ssandra-operator/apis/reaper/v1alpha1"

	v1beta1 "github.com/k8ssandra/cass-operator/apis/cassandra/v1beta1"
)

// ReaperManager is an autogenerated mock type for the Manager type
type ReaperManager struct {
	mock.Mock
}

// AddClusterToReaper provides a mock function with given fields: ctx, cassdc
func (_m *ReaperManager) AddClusterToReaper(ctx context.Context, cassdc *v1beta1.CassandraDatacenter) error {
	ret := _m.Called(ctx, cassdc)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.CassandraDatacenter) error); ok {
		r0 = rf(ctx, cassdc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields: _a0
func (_m *ReaperManager) Connect(_a0 *v1alpha1.Reaper) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Reaper) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyClusterIsConfigured provides a mock function with given fields: ctx, cassdc
func (_m *ReaperManager) VerifyClusterIsConfigured(ctx context.Context, cassdc *v1beta1.CassandraDatacenter) (bool, error) {
	ret := _m.Called(ctx, cassdc)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.CassandraDatacenter) bool); ok {
		r0 = rf(ctx, cassdc)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.CassandraDatacenter) error); ok {
		r1 = rf(ctx, cassdc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
