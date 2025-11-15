package controller

import (
	"fmt"
	"net/http"

	"github.com/bestiesmile1845/Projecteiei/entity" // ตรวจสอบเส้นทางให้ถูกต้อง
	"github.com/bestiesmile1845/Projecteiei/service" // ตรวจสอบเส้นทางให้ถูกต้อง
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload รับค่าจาก frontend
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse สำหรับตอบกลับ frontend
type LoginResponse struct {
	Token string 	 `json:"token"`
	ID 	  uint 	 `json:"id"`
	User  interface{} `json:"user"` // สามารถเป็น entity.PregnantWoman หรือ entity.Doctor
	Role  string 	 `json:"role"` // Role ที่กำหนดโดยตรง
	Name  string 	 `json:"name"`
}

// --- Helper struct to hold common login data ---
type UserInterface struct {
	ID 		 uint 
	Username string
	Password string 
}

// POST /login
func Login(c *gin.Context) {
	var payload LoginPayload
	var user UserInterface 
	var found bool = false
	var role string 	  // **กำหนด Role โดยตรง**
	var userID uint
	var userName string
	var userDetails interface{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	db := entity.DB() // สมมติว่ามีฟังก์ชัน DB() ที่คืนค่า GORM DB

	// 1. ลองค้นหาในตาราง Doctor
	var doc entity.Doctor
	if err := db.
		Where("username = ? OR email = ?", payload.Username, payload.Username).
		First(&doc).Error; err == nil {
		
		user = UserInterface{ID: doc.ID, Username: doc.Username, Password: doc.Password}
		userDetails = doc
		role = "doctor" // **กำหนด Role เป็น doctor**
		found = true
	}

	// 2. ถ้าไม่พบใน Doctor ลองค้นหาในตาราง PregnantWoman
	if !found {
		var woman entity.PregnantWoman
		if err := db.
			Where("username = ? OR email = ?", payload.Username, payload.Username).
			First(&woman).Error; err == nil {
			
			user = UserInterface{ID: woman.ID, Username: woman.Username, Password: woman.Password}
			userDetails = woman
			role = "pregnant" // **กำหนด Role เป็น pregnant**
			found = true
		}
	}

	// ถ้าไม่พบผู้ใช้ในทั้งสองตาราง
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	// ตรวจสอบรหัสผ่าน
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	
	// สร้าง JWT token
	jwtWrapper := service.JwtWrapper{
		SecretKey: 		 "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer: 		 "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
		return
	}

	// Set final response details based on the user object (doc or woman)
	switch details := userDetails.(type) {
	case entity.PregnantWoman:
		userID = details.ID
		userName = details.FullName
	case entity.Doctor:
		userID = details.ID
		userName = details.FullName
	default:
		// ไม่ควรเกิดขึ้น
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown user type"})
		return
	}

	// DEBUG: แสดง role ที่ดึงมา
	fmt.Println("Role determined:", role)


	response := LoginResponse{
		Token: signedToken,
		ID: 	userID,
		User: 	userDetails,
		Role: 	role, // ใช้ role ที่กำหนดโดยตรง
		Name: 	userName,
	}

	// ส่ง response กลับในรูปแบบ {"data": response}
	c.JSON(http.StatusOK, gin.H{"data": response})
}