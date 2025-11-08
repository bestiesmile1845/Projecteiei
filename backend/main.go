package main

import (

	// "github.com/Piyawat777/Final2/controller"
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
	// "github.com/Piyawat777/Final2/middlewares"

	"github.com/bestiesmile1845/Projecteiei/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// api := r.Group("")
	// {
	// 	protected := api.Use(middlewares.Authorizes())
	// 	{

	// 		protected.GET("/Employees", employee.ListEmployees)
	// 		protected.GET("/Employee/:id", employee.GetEmployee)
	// 		protected.GET("Employees/officer/:id", employee.ListEmployeeByID)
	// 		protected.POST("/Employees", employee.CreateEmployee)
	// 		protected.PATCH("/Employees", employee.UpdateEmployee)
	// 		protected.DELETE("/Employees/:id", employee.DeleteEmployee)
	// 		protected.PATCH("/Employees/:id/password", employee.ResetEmployeePassword)

	// 		//=================================================== Booking Routes
	// 		protected.GET("/bookings", booking.ListBookings)
	// 		protected.GET("/booking/:id", booking.GetBooking)
	// 		protected.GET("/bookings/user/:id", booking.ListBookingsByUID)
	// 		protected.POST("/bookings", booking.CreateBooking)
	// 		protected.PATCH("/bookings/cancel/:id", booking.CancelByBookingNumber)
	// 		protected.PATCH("/bookings/:id", booking.UpdateBooking)
	// 		protected.DELETE("/bookings/:id", booking.DeleteBooking)
	// 		protected.DELETE("/bookings/customer/:id", booking.DeleteBookingByCID)
	// 		protected.GET("/bookingbydate", booking.ListBookingsBydate)
	// 		protected.GET("/bookingtotalgroupbydate", booking.ListBookingsTotalbyCID)
	// 		protected.GET("/rooms/available", booking.ListAvailableRooms)
	// 		//=================================================== Booking Routes

	// 		//=================================================== Check Payment Routes
	// 		protected.GET("/chk_payments", chk_payment.ListCHK_Payments)
	// 		protected.GET("/chk_payment/:id", chk_payment.GetCHK_Payment)
	// 		protected.POST("/chk_payments", chk_payment.CreateCHK_Payment)
	// 		protected.PATCH("/chk_payments/:id", chk_payment.UpdateCHK_Payment)
	// 		protected.DELETE("/chk_payments/:id", chk_payment.DeleteCHK_Payment)
	// 		// ---Status---
	// 		protected.GET("/chk_payment/statuses", chk_payment.ListStatuses)
	// 		protected.GET("/chk_payment/status/:id", chk_payment.GetStatus)
	// 		protected.POST("/chk_payment/statuses", chk_payment.CreateStatus)
	// 		protected.PATCH("/chk_payment/statuses", chk_payment.UpdateStatus)
	// 		protected.DELETE("/chk_payment/statuses/:id", chk_payment.DeleteStatus)
	// 		//=================================================== Check Payment Routes

	// 		//==================================================Customer Routes
	// 		protected.GET("/customers", customer.ListCustomers)
	// 		protected.GET("/customer/:id", customer.GetCustomerByID)
	// 		r.POST("/customers", customer.CreateCustomer)
	// 		protected.PATCH("/customer/:id", customer.UpdateCustomer)
	// 		protected.PATCH("/admin/customer/:id", customer.UpdateCustomerByAdmin)
	// 		protected.DELETE("/customers/:id", customer.DeleteCustomer)
	// 		protected.PATCH("/customer/:id/change-password", customer.ChangePassword)

	// 		//Gender
	// 		r.GET("/customers/genders", customer.ListGender)
	// 		protected.GET("/customer/genders/:id", customer.GetGender)
	// 		protected.POST("/customers/genders", customer.CreateGender)
	// 		protected.PATCH("/customers/genders", customer.UpdateGender)
	// 		protected.DELETE("/customers/genders/:id", customer.DeleteGender)

	// 		//==================================================Customer Routes

	// 		//========================= checkInOut routes
	// 		//status
	// 		protected.GET("/checkinoutstatus/:id", check.GetCheckInOutStatus)
	// 		protected.GET("/checkinoutstatuses", check.ListCheckInOutStatuses)
	// 		protected.POST("/checkinoutstatus", check.CreateCheckInOutStatus)
	// 		protected.PATCH("/checkinoutstatus", check.UpdateCheckInOutStatus)
	// 		protected.DELETE("/checkinoutstatus/:id", check.DeleteCheckInOutStatus)
	// 		//main
	// 		protected.GET("/checkinout/:id", check.GetCheckInOut)
	// 		protected.GET("/checkinouts", check.ListCheckInOuts)
	// 		protected.POST("/checkinout", check.CreateCheckInOut)
	// 		protected.PATCH("/checkin", check.UpdateCheckIn)
	// 		protected.PATCH("/checkout", check.UpdateCheckOut)
	// 		//protected.PATCH("/checkinout", check.UpdateCheckInOut)
	// 		protected.DELETE("/checkinout/:id", check.DeleteCheckInOut)
	// 		protected.PATCH("/checkinout/:id", check.CheckOut)

	// 		//========================= repreq routes
	// 		//type
	// 		protected.GET("/repairtype/:id", repreq.GetRepairType)
	// 		protected.GET("/repairtypes", repreq.ListRepairTypes)
	// 		protected.POST("/repairtype", repreq.CreateRepairType)
	// 		protected.PATCH("/repairtype", repreq.UpdateRepairType)
	// 		protected.DELETE("/repairtype/:id", repreq.DeleteRepairType)
	// 		//main
	// 		protected.GET("/repairreq/:id", repreq.GetRepairReqByCid)
	// 		protected.GET("/repairreqs", repreq.ListRepairReqs)
	// 		protected.GET("/rooms/customer/:id", repreq.GetListRoomByCID)
	// 		protected.POST("/repairreq", repreq.CreateRepairReq)
	// 		protected.PATCH("/repairreq", repreq.UpdateRepairReq)
	// 		protected.DELETE("/repairreq/:id", repreq.DeleteRepairReq)

	// 		//==================================================Room Routes
	// 		protected.GET("/rooms", room.ListRooms)
	// 		protected.GET("/room/:id", room.GetRoom)
	// 		protected.POST("/rooms", room.CreateRoom)
	// 		protected.PUT("/rooms", room.UpdateRoom)
	// 		protected.DELETE("/rooms/:id", room.DeleteRoom)

	// 		protected.POST("/roomtypes", room.CreateRoomType)
	// 		protected.GET("/roomtypes", room.ListRoomTypes)
	// 		protected.GET("/roomtypes/:id", room.GetRoomType)
	// 		protected.PATCH("/roomtypes", room.UpdateRoomType)
	// 		protected.DELETE("/roomtypes/:id", room.DeleteRoomType)
	// 		protected.POST("/roomtypes/:id/images", room.UploadRoomTypeImages)
	// 		protected.DELETE("/roomtype-images/:imageID", room.DeleteRoomTypeImage)
	// 		protected.DELETE("/roomtype-images/bulk-delete", room.DeleteRoomTypeImagesBulk)

	// 		protected.GET("/room_zones", room.ListRoomZones)
	// 		protected.GET("/room_zone/:id", room.GetRoomZone)
	// 		protected.POST("/room_zones", room.CreateRoomZone)
	// 		protected.PUT("/room_zones", room.UpdateRoomZone)
	// 		protected.DELETE("/room_zones/:id", room.DeleteRoomZone)

	// 		protected.GET("/states", room.ListStates)
	// 		protected.GET("/state/:id", room.GetState)
	// 		protected.POST("/states", room.CreateState)
	// 		protected.PUT("/states", room.UpdateState)
	// 		protected.DELETE("/states/:id", room.DeleteState)
	// 		//===================================================Room

	// 		// ======================================= PAYMENT
	// 		protected.GET("/payments", payment.ListPayments)
	// 		protected.GET("/payment/:id", payment.GetPayment)
	// 		protected.GET("/payment/customer/:id", payment.ListPaymentByUID)
	// 		protected.POST("/payment", payment.CreatePayment)
	// 		protected.PATCH("/payments", payment.UpdatePayment)

	// 		protected.GET("/paymentmethods", payment.ListPaymentMethods)
	// 		protected.GET("/methods/paymet/:id", payment.ListMethodsByPID)
	// 		protected.GET("/method/:id", payment.GetMethod)
	// 		protected.GET("/places", payment.ListPlaces)
	// 		protected.GET("/priceroom/customer/:id", payment.PriceRoomCID)

	// 		// ======================================= PAYMENT

	// 		//----------review----------------------
	// 		// Review Routes
	// 		r.GET("/Reviews", reviewht.ListReviews)
	// 		protected.GET("/Review/:id", reviewht.GetReview)
	// 		protected.POST("/Reviews", reviewht.CreateReview)
	// 		protected.PATCH("/Reviews", reviewht.UpdateReview)
	// 		protected.DELETE("/Reviews/:id", reviewht.DeleteReview)

	// 		// Systemwork Routes
	// 		protected.GET("/Systemworks", reviewht.ListSystemworks)

	// 		//==================================================Storage Routes
	// 		protected.GET("/storages", storage.ListStorages)
	// 		protected.GET("/storage/:id", storage.GetStorage)
	// 		protected.POST("/storages", storage.CreateStorage)
	// 		protected.PUT("/storages", storage.UpdateStorage)
	// 		protected.DELETE("/storages/:id", storage.DeleteStorage)

	// 		protected.GET("/products", storage.ListProducts)
	// 		protected.GET("/product/:id", storage.GetProduct)
	// 		protected.POST("/products", storage.CreateProduct)
	// 		protected.PATCH("/products", storage.UpdateProduct)
	// 		protected.DELETE("/products/:id", storage.DeleteProduct)

	// 		protected.GET("/product_types", storage.ListProductTypes)
	// 		protected.GET("/product_type/:id", storage.GetProductType)
	// 		protected.POST("/product_types", storage.CreateProductType)
	// 		protected.PATCH("/pProduct_types", storage.UpdateProductType)
	// 		protected.DELETE("/pProduct_types/:id", storage.DeleteProductType)

	// 		//===================================================Storage
	// 		//==================================================Checkroom Routes
	// 		protected.GET("/checkrooms", checkroom.ListCheckroom)
	// 		protected.GET("/checkroom/:id", checkroom.GetCheckroom)
	// 		protected.POST("/checkrooms", checkroom.CreateCheckroom)
	// 		protected.PATCH("/checkroomsupdate/:id", checkroom.UpdateCheckroom)
	// 		protected.DELETE("/checkrooms/:id", checkroom.DeleteCheckroom)
	// 		//Gender
	// 		protected.GET("/damages", checkroom.ListDamage)
	// 		protected.GET("/checkrooms/damages/:id", checkroom.GetDamage)
	// 		protected.POST("/checkrooms/damages", checkroom.CreateDamage)
	// 		protected.PATCH("/checkrooms/damages", checkroom.UpdateDamage)
	// 		protected.DELETE("/checkrooms/damages/:id", checkroom.DeleteDamage)
	// 		//StatusCR
	// 		protected.GET("/status", checkroom.ListStatus)
	// 		protected.GET("/checkrooms/statuscrs/:id", checkroom.GetStatus)
	// 		protected.POST("/checkrooms/status", checkroom.CreateStatus)
	// 		protected.PATCH("/checkrooms/status", checkroom.UpdateStatus)
	// 		protected.DELETE("/checkrooms/status/:id", checkroom.DeleteStatus)

	// 		//==================================================uploadImage
	// 		protected.POST("/carousels", carousel.CreateCarousel)       // สำหรับเพิ่มใหม่
	// 		protected.PATCH("/carousels/:id", carousel.UpdateCarousel)  // สำหรับแก้ไข
	// 		protected.DELETE("/carousels/:id", carousel.DeleteCarousel) // สำหรับลบ

	// 		//==================================================uploadImage
	// 		protected.PATCH("/settings", hotel.UpdateSetting)

	// 	}
	// }
	// r.GET("/listroomtypes", room.ListRoomTypes)
	// r.POST("/login", controller.Login)
	// r.POST("/check-email", customer.CheckEmail)
	// r.POST("/request-otp", customer.RequestOTPHandler)
	// r.POST("/verify-otp", customer.VerifyOTPHandler)
	// r.POST("/reset-password", customer.ResetCustomerPassword)
	// r.GET("/settings", hotel.GetSetting)
	// r.GET("/carousels", carousel.ListCarousels)
	// r.GET("/carousels/:id", carousel.GetCarousel) // ย้ายมาที่นี่: สามารถดูได้โดยไม่ต้องล็อกอิน

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
