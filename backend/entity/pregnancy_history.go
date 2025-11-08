package entity

import (
	"time"

	"gorm.io/gorm"
)

type PregnancyHistory struct {
	gorm.Model

	// FK -> PregnantWoman (P_ID)
	PregnantWomanID *uint          `gorm:"column:p_id" valid:"required~กรุณาเลือกมารดา"`
	PregnantWoman   *PregnantWoman `gorm:"references:ID" valid:"-"`

	Gravida                int
	DeliveryOrAbortionDate time.Time
	GestationAge           int
	BabyWeight             float64
	DeliveryPlace          string

	// FK -> Gender
	GenderID *uint   `valid:"-"`
	Gender   *Gender `gorm:"references:ID" valid:"-"`
}
