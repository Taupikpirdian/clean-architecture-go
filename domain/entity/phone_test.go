package entity_test

import (
	"clean-architecture-go/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
positif casenya
*/
func TestNewPhone(t *testing.T) {
	phone, err := entity.NewPhone(
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
	assert.Nil(t, err)
	assert.Equal(t, "Samsung", phone.GetMerk())
}

/*
negatif casenya
*/
func TestValidasiErrorNewPhoneMerk(t *testing.T) {
	_, err := entity.NewPhone(
		"",
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

	assert.NotNil(t, err)
	assert.Equal(t, "MERK NOT SET", err.Error())
}
