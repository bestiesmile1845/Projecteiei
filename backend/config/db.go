package config

import (
	"fmt"

	"github.com/bestiesmile1845/Projecteiei/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("Mother.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	// AutoMigrate will create the tables and columns if they do not exist
	db.AutoMigrate(
		&entity.PregnantWoman{},
		&entity.Doctor{},
	)
	// GenderMale := entity.Genders{Gender: "Male"}
	// GenderFemale := entity.Genders{Gender: "Female"}

	// db.FirstOrCreate(&GenderMale, &entity.Genders{Gender: "Male"})
	// db.FirstOrCreate(&GenderFemale, &entity.Genders{Gender: "Female"})

	// hashedPasswordAd, _ := HashPassword("123456")
	// Admin := entity.Admin{
	// 	Username: "Admin",
	// 	Password: hashedPasswordAd,
	// 	Email: "Admin@gmail.com",
	// 	Firstname: "Thawan",
	// 	Lastname:  "Banda",
	// 	GenderID: 2,
	// }

	hashedPasswordAd, _ := HashPassword("123456")
	Admin := entity.Doctor{
		Username:    "Doctor",
		Password:    hashedPasswordAd,
		Email:       "Doctor@gmail.com",
		FullName:    "Doctor D",
		PhoneNumber: "0655765587",
	}

	hashedPassword, _ := HashPassword("789012")
	Member := entity.PregnantWoman{
		Username:    "woman",
		Password:    hashedPassword,
		Email:       "woman@gmail.com",
		FullName:    "Thawamhathai Bandasak",
		PhoneNumber: "0655765586",
		HN:          "HN123456789",
		CitizenID:   "11125634569",
	}

	// StartDate, _ := time.Parse("2006-01-02 15:04:05", "2024-08-31 14:30:00")
	// EndDate, _ := time.Parse("2006-01-02 15:04:05", "2024-08-31 14:30:00")
	// Class := &entity.Class{
	// 	ClassName: "Hatha Yoga",
	// 	Deets:  "Introduction to yoga for strength & flexibility",
	// 	StartDate: StartDate,
	// 	EndDate:  EndDate,
	// 	TrainerID: 1,
	// 	ClassPic: "aa",
	// 	ParticNum: 30,
	// 	ClassTypeID: 1,
	// 	AdminID: 1,
	// }

	// db.FirstOrCreate(&Admin, entity.Admin{Email: "PsAdmin@gmail.com"})
	// db.FirstOrCreate(&Member, entity.Member{Email: "Ps@gmail.com"})

	db.FirstOrCreate(&Admin, &entity.Doctor{
		Username:    "Doctor",
		Password:    hashedPasswordAd,
		Email:       "Doctor@gmail.com",
		FullName:    "Doctor D",
		PhoneNumber: "0655765587",
	})
	db.FirstOrCreate(&Member, &entity.PregnantWoman{
		Username:    "woman",
		Password:    hashedPassword,
		Email:       "woman@gmail.com",
		FullName:    "Thawamhathai Bandasak",
		PhoneNumber: "0655765586",
		HN:          "HN123456789",
		CitizenID:   "11125634569",
	})

}
