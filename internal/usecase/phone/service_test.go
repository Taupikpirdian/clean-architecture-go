package phone

import (
	"challenge/clean-architecture-go/internal/repository/mocks"
	"challenge/clean-architecture-go/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repoPhoneMock = new(mocks.PhoneRepositoryMock)
)

func TestGetPhone(t *testing.T) {
	/*
		ini ngarah ke internal > repository > mocks, nama func harus sama GetPhoneByDateRelease
	*/
	repoPhoneMock.On("GetPhoneByDateRelease", mock.Anything, mock.Anything).Return(testdata.GenerateListPhoneByEntity(), nil)

	useCase := NewPhoneUseCaseInteractor(repoPhoneMock)
	listPhone, errGetPhone := useCase.GetPhone()
	assert.Nil(t, errGetPhone)
	assert.NotNil(t, listPhone)
}
