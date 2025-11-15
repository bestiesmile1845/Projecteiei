package main

import (

	 "github.com/bestiesmile1845/Projecteiei/controller"
	// booking "github.com/Piyawat777/Final2/controller/Booking"
	// carousel "github.com/Piyawat777/Final2/controller/Carousel"
	// check "github.com/Piyawat777/Final2/controller/CheckInOut"
	// checkroom "github.com/Piyawat777/Final2/controller/Checkroom"
	// chk_payment "github.com/Piyawat777/Final2/controller/Chk_Payment"
	// customer "github.com/Piyawat777/Final2/controller/Customer"
	// hotel "github.com/Piyawat777/Final2/controller/Hotel"
	// employee "github.com/Piyawat777/Final2/controller/Manage_Employee"
	// payment "github.com/Piyawat777/Final2/controller/Payment"
	// repreq "github.com/Piyawat777/Final2/controller/RepReq"
	// reviewht "github.com/Piyawat777/Final2/controller/Review"
	// room "github.com/Piyawat777/Final2/controller/Room"
	// storage "github.com/Piyawat777/Final2/controller/Storage"
	// "github.com/bestiesmile1845/Projecteiei/middlewares"

	"github.com/bestiesmile1845/Projecteiei/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	
	r.POST("/login", controller.Login)


	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Username")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
