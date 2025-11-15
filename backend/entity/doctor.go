package entity

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	FullName    string
	PhoneNumber string
	Email       string
	
	// **ฟิลด์ที่เพิ่มเข้ามาสำหรับการล็อกอิน**
	Username 	string 	`gorm:"uniqueIndex"`
	Password 	string

	VisitDoctors []VisitDoctor `gorm:"foreignKey:DoctorID"`
}