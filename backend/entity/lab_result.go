package entity

import (
	"time"

	"gorm.io/gorm"
)

type LabResult struct {
	gorm.Model

	// FK -> Pregnancy
	PregnancyID *uint      `valid:"required~กรุณาเลือกครรภ์"`
	Pregnancy   *Pregnancy `gorm:"references:ID" valid:"-"`

	TestDate     time.Time
	Hct          float64 // Hematocrit
	Hb           float64 // Hemoglobin
	HbTyping     string
	OtherRemarks string

	// FK -> ผลตรวจ
	DCPResultID     *uint        `valid:"-"`
	DCPResult       *CheckResult `gorm:"foreignKey:DCPResultID" valid:"-"`
	AntiHIVResultID *uint        `valid:"-"`
	AntiHIVResult   *CheckResult `gorm:"foreignKey:AntiHIVResultID" valid:"-"`
}
