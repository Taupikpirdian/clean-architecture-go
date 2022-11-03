package response

import (
	"challenge/clean-architecture-go/domain/entity"
	"time"
)

type PhoneAttributs struct {
	Merk        string    `json:"merk"`
	Model       string    `json:"model"`
	Body        string    `json:"body"`
	Display     string    `json:"display"`
	Chipset     string    `json:"chipset"`
	Os          string    `json:"os"`
	Camera      string    `json:"camera"`
	Battery     int       `json:"battery"`
	Ram         int       `json:"ram"`
	Storage     int       `json:"storage"`
	ReleaseDate time.Time `json:"release_date"`
}

func NewPhoneResponse(dataPhone *entity.Phone) PhoneAttributs {
	return PhoneAttributs{
		Merk:        dataPhone.GetMerk(),
		Model:       dataPhone.GetModel(),
		Body:        dataPhone.GetBody(),
		Display:     dataPhone.GetDisplay(),
		Chipset:     dataPhone.GetChipset(),
		Os:          dataPhone.GetOs(),
		Camera:      dataPhone.GetCamera(),
		Battery:     dataPhone.GetBattery(),
		Ram:         dataPhone.GetRam(),
		Storage:     dataPhone.GetStorage(),
		ReleaseDate: dataPhone.GetReleaseDate(),
	}
}
