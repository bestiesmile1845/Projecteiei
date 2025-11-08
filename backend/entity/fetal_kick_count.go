package entity

import (
	"time"

	"gorm.io/gorm"
)

type FetalKickCount struct {
	gorm.Model

	// FK -> Pregnancy
	PregnancyID *uint      `valid:"required~กรุณาเลือกครรภ์"`
	Pregnancy   *Pregnancy `gorm:"references:ID" valid:"-"`

	CountDate        time.Time
	KickCountMorning int
	KickCountLunch   int
	KickCountEvening int
}
