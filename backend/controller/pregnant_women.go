package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bestiesmile1845/Projecteiei/entity"
	"egithub.com/bestiesmile1845/Projecteiei/config"

)

// POST /users
func CreatePregnantWoman(c *gin.Context) {
	var PregnantWoman entity.PregnantWoman

	// bind เข้าตัวแปร PregnantWoman
	if err := c.ShouldBindJSON(&PregnantWoman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// ตรวจสอบว่า username ซ้ำกันหรือไม่
	var existingPregnantWoman entity.PregnantWoman
	if err := db.Where("username = ?", PregnantWoman.Username).First(&existingPregnantWoman).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists in PregnantWoman"})
		return
	}
	// ตรวจสอบว่า username ซ้ำกันหรือไม่ใน table doctor
	var existingDoctor entity.Doctor
	if err := db.Where("username = ?", PregnantWoman.Username).First(&existingDoctor).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists in doctor"})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashedPassword, _ := config.HashPassword(PregnantWoman.Password)

	// สร้าง PregnantWoman
	m := entity.PregnantWoman{
		FullName: PregnantWoman.FullName,
		Email:     PregnantWoman.Email,
		Password:  hashedPassword,
		Username: PregnantWoman.Username,
		PhoneNumber: PregnantWoman.PhoneNumber,
		Age: PregnantWoman.Age,
	}

	// บันทึก
	if err := db.Create(&m).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": m})
}


// GET /PregnantWoman/:id
func GetPregnantWoman(c *gin.Context) {
	ID := c.Param("id")
	var PregnantWoman entity.PregnantWoman

	db := config.DB()
	results := db.Preload("Gender").First(&PregnantWoman, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if PregnantWoman.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, PregnantWoman)
}
func GetUsername(c *gin.Context) {
	Username := c.Param("username")
	var PregnantWoman entity.PregnantWoman

	db := config.DB()
	results := db.Preload("Gender").First(&PregnantWoman, Username)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if PregnantWoman.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, PregnantWoman)
}

func GetPassword(c *gin.Context) {
	Password := c.Param("password")
	var PregnantWoman entity.PregnantWoman

	db := config.DB()
	results := db.Preload("Gender").First(&PregnantWoman, Password)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if PregnantWoman.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, PregnantWoman)
}

// GET /users
func ListPregnantWomans(c *gin.Context) {

	var users []entity.PregnantWoman

	db := config.DB()
	results := db.Preload("Gender").Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// DELETE /users/:id
func DeletePregnantWoman(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM PregnantWomans WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}

// PATCH /users
func UpdatePregnantWoman(c *gin.Context) {
	var PregnantWoman entity.PregnantWoman

	PregnantWomanID := c.Param("PregnantWomanid")

	db := config.DB()
	result := db.First(&PregnantWoman, PregnantWomanID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&PregnantWoman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&PregnantWoman)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	var existingPregnantWoman entity.PregnantWoman
	if err := db.Where("username = ?", PregnantWoman.Username).First(&existingPregnantWoman).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

func CountPregnantWomans(c *gin.Context) {
	var count int64
	db := config.DB()
	if err := db.Model(&entity.PregnantWoman{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}