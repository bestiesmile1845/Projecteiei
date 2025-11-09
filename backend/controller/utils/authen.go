package controller

import (
	"fmt"
	"net/http"

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
    User  interface{} `json:"user"` // สามารถเป็น entity.womanomer หรือ entity.docloyee
    Role  string      `json:"role"`
    Name  string      `json:"name"`
}

// POST /login
func Login(c *gin.Context) {
    var payload LoginPayload
    

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ค้นหา username หรือ email
    if err := entity.DB().
        Preload("UserRole").
        Where("username = ? OR email = ?", payload.Username, payload.Username).
        First(&signin).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(signin.Password), []byte(payload.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
        return
    }

    // สร้าง JWT token
    jwtWrapper := service.JwtWrapper{
        SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
        Issuer:          "AuthService",
        ExpirationHours: 24,
    }
    signedToken, err := jwtWrapper.GenerateToken(signin.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
        return
    }

    role := signin.UserRole.RoleName
    // DEBUG: แสดง role ที่ดึงมา
    fmt.Println("Role from UserRole:", role)

    var response LoginResponse

    switch role {
    case "woman":
        var woman entity.PregnantWoman
        if err := entity.DB().Where("signin_id = ?", signin.ID).First(&woman).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "customer data not found for this signin"})
            return
        }
        response = LoginResponse{
            Token: signedToken,
            ID:    woman.ID,
            User:  woman,
            Role:  role,
            Name:  woman.FullName,
        }

    case "doctor", "Admin":
        var doc entity.Doctor
        if err := entity.DB().Where("signin_id = ?", signin.ID).First(&doc).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "docloyee data not found for this signin"})
            return
        }
        response = LoginResponse{
            Token: signedToken,
            ID:    doc.ID,
            User:  doc,
            Role:  role,
            Name:  doc.FullName,
        }

    default:
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user role"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": response})
}