package entity

import (
	"errors"
	"time"
)

type Phone struct {
	merk         string
	model        string
	body         string
	display      string
	chipset      string
	os           string
	camera       string
	battery      int
	ram          int
	storage      int
	release_date time.Time
}

func NewPhone(
	merkSet string,
	modelSet string,
	bodySet string,
	displaySet string,
	chipsetSet string,
	osSet string,
	cameraSet string,
	batterySet int,
	ramSet int,
	storageSet int,
	releaseSet time.Time) (*Phone, error) {

	if merkSet == "" {
		return nil, errors.New("MERK NOT SET")
	}

	if modelSet == "" {
		return nil, errors.New("MODEL NOT SET")
	}

	if bodySet == "" {
		return nil, errors.New("BODY NOT SET")
	}

	if displaySet == "" {
		return nil, errors.New("DISPLAY NOT SET")
	}

	if chipsetSet == "" {
		return nil, errors.New("CHIPSET NOT SET")
	}

	if osSet == "" {
		return nil, errors.New("OS NOT SET")
	}

	if cameraSet == "" {
		return nil, errors.New("CAMERA NOT SET")
	}

	if batterySet == 0 {
		return nil, errors.New("BATTERY NOT SET")
	}

	if ramSet == 0 {
		return nil, errors.New("RAM NOT SET")
	}

	if storageSet == 0 {
		return nil, errors.New("STORAGE NOT SET")
	}

	return &Phone{
		merk:         merkSet,
		model:        modelSet,
		body:         bodySet,
		display:      displaySet,
		chipset:      chipsetSet,
		os:           osSet,
		camera:       cameraSet,
		battery:      batterySet,
		ram:          ramSet,
		storage:      storageSet,
		release_date: releaseSet,
	}, nil
}

func (ph *Phone) GetMerk() string {
	return ph.merk
}

func (ph *Phone) GetModel() string {
	return ph.model
}

func (ph *Phone) GetBody() string {
	return ph.body
}

func (ph *Phone) GetDisplay() string {
	return ph.display
}

func (ph *Phone) GetChipset() string {
	return ph.chipset
}

func (ph *Phone) GetOs() string {
	return ph.os
}

func (ph *Phone) GetCamera() string {
	return ph.camera
}

func (ph *Phone) GetBattery() int {
	return ph.battery
}

func (ph *Phone) GetRam() int {
	return ph.ram
}

func (ph *Phone) GetStorage() int {
	return ph.storage
}

func (ph *Phone) GetReleaseDate() time.Time {
	return ph.release_date
}
