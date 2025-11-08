package entity

import "gorm.io/gorm"

type BloodType struct {
	gorm.Model
	Name   string
	Childs []Child `gorm:"foreignKey:BloodTypeID"`
}
