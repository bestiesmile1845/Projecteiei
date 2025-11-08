package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("Mother.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate ตาม Entity ทั้งหมดที่เราใช้คุยกัน
	err = database.AutoMigrate(

		// -------- ระบบสมุดแม่และเด็กอิเล็กทรอนิกส์ --------
		// Master / Lookup
		&BloodType{},
		&CheckResult{},

		// บุคคล
		&Doctor{},
		&Husband{},
		&PregnantWoman{},
		&Child{},

		// การพบแพทย์ / นัดหมาย
		&VisitDoctor{},
		&Appointment{},

		// ประวัติการเจ็บป่วย
		&MedicalHistory{},

		// วัคซีน
		&VaccineType{},
		&VacDose{},
		&Vaccination{},

		// การตั้งครรภ์ และฝากครรภ์
		&Pregnancy{},
		&AntenatalVisit{},
		&FetalKickCount{},

		// ประวัติการตั้งครรภ์ที่ผ่านมา
		&PregnancyHistory{},

		// ผลแลป
		&LabResult{},
	)

	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	db = database

	// ใส่ข้อมูลเริ่มต้น (ถ้ามี)
	// SetupIntoDatabase(db)
}
