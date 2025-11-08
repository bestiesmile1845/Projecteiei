package entity

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model

	VisitDoctorID *uint 
	VisitDoctor   *VisitDoctor `gorm:"references:ID"`


	PregnantWomen []PregnantWoman `gorm:"foreignKey:AppointmentID"`
}
