// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "doggos/repo/model"

	mock "github.com/stretchr/testify/mock"
)

// DoggoRepository is an autogenerated mock type for the DoggoRepository type
type DoggoRepository struct {
	mock.Mock
}

// GetDoggos provides a mock function with given fields: page, limit, breedID
func (_m *DoggoRepository) GetDoggos(page int, limit int, breedID string) ([]model.DoggoDto, error) {
	ret := _m.Called(page, limit, breedID)

	var r0 []model.DoggoDto
	if rf, ok := ret.Get(0).(func(int, int, string) []model.DoggoDto); ok {
		r0 = rf(page, limit, breedID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.DoggoDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(page, limit, breedID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}