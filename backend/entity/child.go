package entity

import (
	"time"

	"gorm.io/gorm"
)

type Child struct {
	gorm.Model
	FullName  string
	HN        int
	DOB       time.Time
	CitizenID string
	Address   string

	// FK -> BloodType (Blood_ID)
	BloodTypeID *uint      `gorm:"column:blood_id" valid:"-"`
	BloodType   *BloodType `gorm:"references:ID" valid:"-"`

}
