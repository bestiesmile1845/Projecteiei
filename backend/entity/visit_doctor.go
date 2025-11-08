package entity

import "gorm.io/gorm"

type VisitDoctor struct {
	gorm.Model

	// FK -> Doctor (D_ID)
	DoctorID *uint   `valid:"required~กรุณาเลือกแพทย์"`
	Doctor   *Doctor `gorm:"references:ID" valid:"-"`

	Appointments []Appointment `gorm:"foreignKey:VisitDoctorID"`
}
