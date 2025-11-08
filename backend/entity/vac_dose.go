package entity

import "gorm.io/gorm"

type VacDose struct {
	gorm.Model
	Field  string
	Method string

	Vaccinations []Vaccination `gorm:"foreignKey:VacDoseID"`
}
