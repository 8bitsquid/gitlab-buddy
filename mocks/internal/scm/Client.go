// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package scm

import (
	mock "github.com/stretchr/testify/mock"
	scm "gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/scm"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateGroupRepo provides a mock function with given fields: _a0, _a1
func (_m *Client) CreateGroupRepo(_a0 interface{}, _a1 scm.IRepository) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, scm.IRepository) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllGroupRepos provides a mock function with given fields: _a0
func (_m *Client) GetAllGroupRepos(_a0 interface{}) []scm.IRepository {
	ret := _m.Called(_a0)

	var r0 []scm.IRepository
	if rf, ok := ret.Get(0).(func(interface{}) []scm.IRepository); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]scm.IRepository)
		}
	}

	return r0
}
