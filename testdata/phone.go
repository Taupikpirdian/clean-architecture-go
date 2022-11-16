package testdata

import (
	"clean-architecture-go/domain/entity"
	"time"
)

func GenerateListPhoneByEntity() []*entity.Phone {
	phonelist := make([]*entity.Phone, 0)
	phone, _ := entity.NewPhone(
		"Samsung",
		"A50",
		"Plastic",
		"15",
		"Snap Dragon 553",
		"Android 10",
		"720p",
		5400,
		8,
		128,
		time.Now())
	phonelist = append(phonelist, phone)

	return phonelist
}
