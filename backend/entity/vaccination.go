package entity

import (
	"time"

	"gorm.io/gorm"
)

type Vaccination struct {
	gorm.Model

	// FK -> PregnantWoman (P_ID)
	PregnantWomanID *uint          `gorm:"column:p_id" valid:"required~กรุณาเลือกมารดา"`
	PregnantWoman   *PregnantWoman `gorm:"references:ID" valid:"-"`

	IsPreviouslyVaccinated bool
	PreviousDoses          int
	LastPreviousDateYear   *time.Time
	Dose1DateDuringPreg    *time.Time
	Dose2DateDuringPreg    *time.Time
	Remarks                string

	// FK -> VaccineType
	VaccineTypeID *uint        `valid:"-"`
	VaccineType   *VaccineType `gorm:"references:ID" valid:"-"`

	// FK -> VacDose (VD_ID)
	VacDoseID *uint    `gorm:"column:vd_id" valid:"-"`
	VacDose   *VacDose `gorm:"references:ID" valid:"-"`
}
