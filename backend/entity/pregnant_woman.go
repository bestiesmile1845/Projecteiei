package entity

import "gorm.io/gorm"

type PregnantWoman struct {
	gorm.Model
	FullName    string `json:"full_name"`
	Age         int    `json:"age"`
	HN          string `json:"hn"`
	CitizenID   string `json:"citizen_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`

	// **ฟิลด์ที่เพิ่มเข้ามาสำหรับการล็อกอิน**
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"password"`

	// FK -> Appointment (A_ID)
	AppointmentID *uint        `gorm:"column:a_id" valid:"-"`
	Appointment   *Appointment `gorm:"references:ID" valid:"-"`

	// FK -> Husband
	HusbandID *uint    `valid:"-"`
	Husband   *Husband `gorm:"references:ID" valid:"-"`

	MedicalHistories   []MedicalHistory   `gorm:"foreignKey:PregnantWomanID"`
	Vaccinations       []Vaccination      `gorm:"foreignKey:PregnantWomanID"`
	Pregnancies        []Pregnancy        `gorm:"foreignKey:PregnantWomanID"`
	PregnancyHistories []PregnancyHistory `gorm:"foreignKey:PregnantWomanID"`
}
