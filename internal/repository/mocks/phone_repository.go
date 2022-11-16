package mocks

import (
	"clean-architecture-go/domain/entity"
	"time"

	"github.com/stretchr/testify/mock"
)

type PhoneRepositoryMock struct {
	mock.Mock
}

func (ar *PhoneRepositoryMock) GetPhoneByDateRelease(StartDate time.Time, EndDate time.Time) ([]*entity.Phone, error) {
	args := ar.Called(StartDate, EndDate)
	return args.Get(0).([]*entity.Phone), args.Error(1)
}
