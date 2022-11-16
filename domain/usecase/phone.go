package usecase

import "clean-architecture-go/domain/entity"

type PhoneUseCase interface {
	GetPhone() ([]*entity.Phone, error) // Get Data Last 3 Month
}
