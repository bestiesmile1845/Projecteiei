package entity

import "gorm.io/gorm"

type VaccineType struct {
	gorm.Model
	Name         string
	Vaccinations []Vaccination `gorm:"foreignKey:VaccineTypeID"`
}
