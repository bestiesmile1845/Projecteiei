package entity

import "gorm.io/gorm"

type CheckResult struct {
	gorm.Model
	Name       string
	LabResults []LabResult `gorm:"foreignKey:DCPResultID"`
}
