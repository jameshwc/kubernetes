/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by mockery v2.53.3. DO NOT EDIT.

package eviction

import mock "github.com/stretchr/testify/mock"

// MockNotifierFactory is an autogenerated mock type for the NotifierFactory type
type MockNotifierFactory struct {
	mock.Mock
}

type MockNotifierFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNotifierFactory) EXPECT() *MockNotifierFactory_Expecter {
	return &MockNotifierFactory_Expecter{mock: &_m.Mock}
}

// NewCgroupNotifier provides a mock function with given fields: path, attribute, threshold
func (_m *MockNotifierFactory) NewCgroupNotifier(path string, attribute string, threshold int64) (CgroupNotifier, error) {
	ret := _m.Called(path, attribute, threshold)

	if len(ret) == 0 {
		panic("no return value specified for NewCgroupNotifier")
	}

	var r0 CgroupNotifier
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64) (CgroupNotifier, error)); ok {
		return rf(path, attribute, threshold)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64) CgroupNotifier); ok {
		r0 = rf(path, attribute, threshold)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(CgroupNotifier)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int64) error); ok {
		r1 = rf(path, attribute, threshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotifierFactory_NewCgroupNotifier_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewCgroupNotifier'
type MockNotifierFactory_NewCgroupNotifier_Call struct {
	*mock.Call
}

// NewCgroupNotifier is a helper method to define mock.On call
//   - path string
//   - attribute string
//   - threshold int64
func (_e *MockNotifierFactory_Expecter) NewCgroupNotifier(path interface{}, attribute interface{}, threshold interface{}) *MockNotifierFactory_NewCgroupNotifier_Call {
	return &MockNotifierFactory_NewCgroupNotifier_Call{Call: _e.mock.On("NewCgroupNotifier", path, attribute, threshold)}
}

func (_c *MockNotifierFactory_NewCgroupNotifier_Call) Run(run func(path string, attribute string, threshold int64)) *MockNotifierFactory_NewCgroupNotifier_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockNotifierFactory_NewCgroupNotifier_Call) Return(_a0 CgroupNotifier, _a1 error) *MockNotifierFactory_NewCgroupNotifier_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotifierFactory_NewCgroupNotifier_Call) RunAndReturn(run func(string, string, int64) (CgroupNotifier, error)) *MockNotifierFactory_NewCgroupNotifier_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNotifierFactory creates a new instance of MockNotifierFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNotifierFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNotifierFactory {
	mock := &MockNotifierFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
