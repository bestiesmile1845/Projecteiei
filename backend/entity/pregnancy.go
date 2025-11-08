package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pregnancy struct {
	gorm.Model

	// FK -> PregnantWoman (P_ID)
	PregnantWomanID *uint          `gorm:"column:p_id" valid:"required~กรุณาเลือกมารดา"`
	PregnantWoman   *PregnantWoman `gorm:"references:ID" valid:"-"`

	PregnancyNo        int
	LMP                time.Time // Last Menstrual Period
	EDC                time.Time // Expected Date of Confinement
	PrePregnancyWeight float64
	Height             float64
	PrePregnancyBMI    float64

	AntenatalVisits []AntenatalVisit `gorm:"foreignKey:PregnancyID"`
	LabResults      []LabResult      `gorm:"foreignKey:PregnancyID"`
	FetalKickCounts []FetalKickCount `gorm:"foreignKey:PregnancyID"`
}
