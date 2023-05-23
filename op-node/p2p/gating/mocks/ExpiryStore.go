// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"

	time "time"
)

// ExpiryStore is an autogenerated mock type for the ExpiryStore type
type ExpiryStore struct {
	mock.Mock
}

type ExpiryStore_Expecter struct {
	mock *mock.Mock
}

func (_m *ExpiryStore) EXPECT() *ExpiryStore_Expecter {
	return &ExpiryStore_Expecter{mock: &_m.Mock}
}

// IPBanExpiry provides a mock function with given fields: ip
func (_m *ExpiryStore) IPBanExpiry(ip net.IP) (time.Time, error) {
	ret := _m.Called(ip)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(net.IP) (time.Time, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(net.IP) time.Time); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(net.IP) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpiryStore_IPBanExpiry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPBanExpiry'
type ExpiryStore_IPBanExpiry_Call struct {
	*mock.Call
}

// IPBanExpiry is a helper method to define mock.On call
//   - ip net.IP
func (_e *ExpiryStore_Expecter) IPBanExpiry(ip interface{}) *ExpiryStore_IPBanExpiry_Call {
	return &ExpiryStore_IPBanExpiry_Call{Call: _e.mock.On("IPBanExpiry", ip)}
}

func (_c *ExpiryStore_IPBanExpiry_Call) Run(run func(ip net.IP)) *ExpiryStore_IPBanExpiry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP))
	})
	return _c
}

func (_c *ExpiryStore_IPBanExpiry_Call) Return(_a0 time.Time, _a1 error) *ExpiryStore_IPBanExpiry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExpiryStore_IPBanExpiry_Call) RunAndReturn(run func(net.IP) (time.Time, error)) *ExpiryStore_IPBanExpiry_Call {
	_c.Call.Return(run)
	return _c
}

// PeerBanExpiry provides a mock function with given fields: id
func (_m *ExpiryStore) PeerBanExpiry(id peer.ID) (time.Time, error) {
	ret := _m.Called(id)

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(peer.ID) (time.Time, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) time.Time); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(peer.ID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExpiryStore_PeerBanExpiry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PeerBanExpiry'
type ExpiryStore_PeerBanExpiry_Call struct {
	*mock.Call
}

// PeerBanExpiry is a helper method to define mock.On call
//   - id peer.ID
func (_e *ExpiryStore_Expecter) PeerBanExpiry(id interface{}) *ExpiryStore_PeerBanExpiry_Call {
	return &ExpiryStore_PeerBanExpiry_Call{Call: _e.mock.On("PeerBanExpiry", id)}
}

func (_c *ExpiryStore_PeerBanExpiry_Call) Run(run func(id peer.ID)) *ExpiryStore_PeerBanExpiry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(peer.ID))
	})
	return _c
}

func (_c *ExpiryStore_PeerBanExpiry_Call) Return(_a0 time.Time, _a1 error) *ExpiryStore_PeerBanExpiry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExpiryStore_PeerBanExpiry_Call) RunAndReturn(run func(peer.ID) (time.Time, error)) *ExpiryStore_PeerBanExpiry_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewExpiryStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewExpiryStore creates a new instance of ExpiryStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExpiryStore(t mockConstructorTestingTNewExpiryStore) *ExpiryStore {
	mock := &ExpiryStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
