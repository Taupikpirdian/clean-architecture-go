package repository

import (
	"clean-architecture-go/domain/entity"
	"time"
)

type PhoneRepository interface {
	GetPhoneByDateRelease(StartDate time.Time, EndDate time.Time) ([]*entity.Phone, error)
}
