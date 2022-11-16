package phone

import "clean-architecture-go/domain/usecase"

type PhoneHandlerInteractor struct {
	PhoneUseCase usecase.PhoneUseCase
}

func NewPhoneHandler(usecaseImplement usecase.PhoneUseCase) *PhoneHandlerInteractor {
	return &PhoneHandlerInteractor{PhoneUseCase: usecaseImplement}
}
