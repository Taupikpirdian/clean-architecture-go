package phone

import "challenge/clean-architecture-go/domain/repository"

type PhoneUseCaseInteractor struct {
	RepoPhone repository.PhoneRepository
}

func NewPhoneUseCaseInteractor(repoImplement repository.PhoneRepository) *PhoneUseCaseInteractor {
	return &PhoneUseCaseInteractor{RepoPhone: repoImplement}
}
