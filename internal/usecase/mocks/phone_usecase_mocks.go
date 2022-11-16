package mocks

import (
	"clean-architecture-go/domain/entity"

	"github.com/stretchr/testify/mock"
)

type PhoneUseCaseMocks struct {
	mock.Mock
}

/*
implement dari usecase si domain, jadi nama method harus sama seperti interface si domain > usecase > phone > GetPhone()
*/
func (ph *PhoneUseCaseMocks) GetPhone() ([]*entity.Phone, error) {
	args := ph.Called()
	return args.Get(0).([]*entity.Phone), args.Error(1)
}
