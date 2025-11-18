package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bestiesmile1845/Projecteiei/config"
	"github.com/bestiesmile1845/Projecteiei/entity"  // ตรวจสอบเส้นทางให้ถูกต้อง
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
	Token string      `json:"token"`
	ID    uint        `json:"id"`
	User  interface{} `json:"user"` // สามารถเป็น entity.PregnantWoman หรือ entity.Doctor
	Role  string      `json:"role"` // Role ที่กำหนดโดยตรง
	Name  string      `json:"name"`
}

// --- Helper struct to hold common login data ---
type UserInterface struct {
	ID       uint
	Username string
	Password string
}

// POST /login
func Login(c *gin.Context) {
	var payload LoginPayload
	var user UserInterface
	var found bool = false
	var role string // **กำหนด Role โดยตรง**
	var userID uint
	var userName string
	var userDetails interface{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB() // ใช้ config.DB() แทน

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
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
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
		ID:    userID,
		User:  userDetails,
		Role:  role, // ใช้ role ที่กำหนดโดยตรง
		Name:  userName,
	}

	// ส่ง response กลับในรูปแบบ {"data": response}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

// GET /me
func GetMe(c *gin.Context) {
	// ดึง Authorization header
	clientToken := c.Request.Header.Get("Authorization")
	if clientToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
		return
	}

	// แยก Bearer token
	extractedToken := strings.Split(clientToken, "Bearer ")
	if len(extractedToken) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Format of Authorization Token"})
		return
	}
	clientToken = strings.TrimSpace(extractedToken[1])

	// Validate token
	jwtWrapper := service.JwtWrapper{
		SecretKey: "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:    "AuthService",
	}

	claims, err := jwtWrapper.ValidateToken(clientToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	// ใช้ username จาก token claims เพื่อค้นหาผู้ใช้
	username := claims.Username
	db := config.DB()

	var found bool = false
	var role string
	var userID uint
	var userName string
	var userDetails interface{}

	// ค้นหาใน Doctor
	var doc entity.Doctor
	if err := db.Where("username = ?", username).First(&doc).Error; err == nil {
		userDetails = doc
		role = "doctor"
		userID = doc.ID
		userName = doc.FullName
		found = true
	}

	// ถ้าไม่พบใน Doctor ค้นหาใน PregnantWoman
	if !found {
		var woman entity.PregnantWoman
		if err := db.Where("username = ?", username).First(&woman).Error; err == nil {
			userDetails = woman
			role = "pregnant"
			userID = woman.ID
			userName = woman.FullName
			found = true
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// ส่งข้อมูลผู้ใช้กลับ
	response := gin.H{
		"id":   userID,
		"role": role,
		"name": userName,
		"user": userDetails,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
