package entity

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	FullName    string
	PhoneNumber string
	Email       string

	VisitDoctors []VisitDoctor `gorm:"foreignKey:DoctorID"`
}
