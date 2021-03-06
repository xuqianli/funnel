package mocks

import config "tes/config"
import ga4gh_task_ref "tes/server/proto"

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// StartWorker provides a mock function with given fields: project, zone, id, conf
func (_m *Client) StartWorker(project string, zone string, id string, conf config.Worker) error {
	ret := _m.Called(project, zone, id, conf)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, config.Worker) error); ok {
		r0 = rf(project, zone, id, conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Template provides a mock function with given fields: project, zone, id
func (_m *Client) Template(project string, zone string, id string) (*ga4gh_task_ref.Resources, error) {
	ret := _m.Called(project, zone, id)

	var r0 *ga4gh_task_ref.Resources
	if rf, ok := ret.Get(0).(func(string, string, string) *ga4gh_task_ref.Resources); ok {
		r0 = rf(project, zone, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ga4gh_task_ref.Resources)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(project, zone, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
