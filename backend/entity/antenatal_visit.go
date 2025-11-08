package entity

import (
	"time"

	"gorm.io/gorm"
)

type AntenatalVisit struct {
	gorm.Model

	// FK -> Pregnancy
	PregnancyID *uint      `valid:"required~กรุณาเลือกครรภ์"`
	Pregnancy   *Pregnancy `gorm:"references:ID" valid:"-"`

	VisitDate        time.Time
	GestationalAge   int
	Weight           float64
	BloodPressure    string
	HeightFundus     float64
	FetalHeartSound  string
	FetalMovement    string
	UrineProtein     string
	UrineSugar       string
	Swelling         string
	MedicalDiagnosis string
}
