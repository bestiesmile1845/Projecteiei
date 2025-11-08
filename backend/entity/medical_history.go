package entity

import "gorm.io/gorm"

type MedicalHistory struct {
	gorm.Model

	// FK -> PregnantWoman (P_ID)
	PregnantWomanID *uint          
	PregnantWoman   *PregnantWoman `gorm:"references:ID"`

	ChronicDiseases         string
	GeneticDiseases         string
	DrugAllergies           string
	FamilyHistoryHT         bool
	FamilyHistoryCongenital bool
	OtherFamilyHistory      string
}
