package mocks

import mock "github.com/stretchr/testify/mock"

type URLSaver struct {
	mock.Mock
}

func (_m *URLSaver) SaveURL(urlToSave string) (int64, error) {
	ret := _m.Called(urlToSave)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(urlToSave)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(urlToSave)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(urlToSave)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewURLSaver interface {
	mock.TestingT
	Cleanup(func())
}

func NewURLSaver(t mockConstructorTestingTNewURLSaver) *URLSaver {
	mock := &URLSaver{}
	mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
