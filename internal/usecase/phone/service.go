package phone

import (
	"challenge/clean-architecture-go/domain/entity"
	"time"
)

func (ph *PhoneUseCaseInteractor) GetPhone() ([]*entity.Phone, error) { // name method sesuai dengan usecase di domain
	var (
		reqDateNow    = time.Now()
		reqDate3Month = time.Now().AddDate(0, -3, 0)
	)

	phone, errGetPhone := ph.RepoPhone.GetPhoneByDateRelease(reqDateNow, reqDate3Month)
	if errGetPhone != nil {
		return nil, errGetPhone
	}

	return phone, nil
}
