package entity

import "gorm.io/gorm"

type Husband struct {
	gorm.Model
	FullName    string
	Age         int
	CitizenID   string
	PhoneNumber string
	Email       string

	PregnantWomen []PregnantWoman `gorm:"foreignKey:HusbandID"`
}
