package api

import (
	v1 "github.com/MuhammadyusufAdhamov/booking/api/v1"
	"github.com/MuhammadyusufAdhamov/booking/config"
	"github.com/MuhammadyusufAdhamov/booking/storage"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/MuhammadyusufAdhamov/booking/api/docs"
)

type RouterOptions struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      localhost:8000
// @BasePath  /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:      opt.Cfg,
		Storage:  opt.Storage,
		InMemory: opt.InMemory,
	})

	router.Static("/media", "./media")

	apiV1 := router.Group("/v1")

	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.GET("/users", handlerV1.GetAllUsers)
	apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	apiV1.DELETE("/users/:id", handlerV1.DeleteUser)

	apiV1.GET("/owners/:id", handlerV1.GetOwner)
	apiV1.POST("/owners", handlerV1.CreateOwner)
	apiV1.GET("/owners", handlerV1.GetAllOwners)
	apiV1.PUT("/owners/:id", handlerV1.UpdateOwner)
	apiV1.DELETE("/owners/:id", handlerV1.DeleteOwner)

	apiV1.GET("/hotels/:id", handlerV1.GetHotel)
	apiV1.POST("/hotels", handlerV1.CreateHotel)
	apiV1.GET("/hotels", handlerV1.GetAllHotels)
	apiV1.PUT("/hotels/:id", handlerV1.UpdateHotel)
	apiV1.DELETE("/hotels/:id", handlerV1.DeleteHotel)

	apiV1.GET("/rooms/:id", handlerV1.GetRoom)
	apiV1.POST("/rooms", handlerV1.CreateRoom)
	apiV1.GET("/rooms", handlerV1.GetAllRooms)
	apiV1.PUT("/rooms/:id", handlerV1.UpdateRoom)
	apiV1.DELETE("/rooms/:id", handlerV1.DeleteRoom)

	apiV1.GET("/bookings/:id", handlerV1.GetBooking)
	apiV1.POST("/bookings", handlerV1.CreateBooking)
	apiV1.GET("/bookings", handlerV1.GetAllBookings)
	apiV1.PUT("/bookings/:id", handlerV1.UpdateBooking)
	apiV1.DELETE("/bookings/:id", handlerV1.DeleteBooking)

	apiV1.POST("/auth/register", handlerV1.Register)
	apiV1.POST("/auth/verify", handlerV1.Verify)
	apiV1.POST("/auth/login", handlerV1.Login)
	apiV1.POST("/auth/forgot-password", handlerV1.ForgotPassword)
	apiV1.POST("/auth/verify-forgot-password", handlerV1.VerifyForgotPassword)
	apiV1.POST("/auth/update-password", handlerV1.AuthMiddleware, handlerV1.UpdatePassword)

	apiV1.POST("/file-upload", handlerV1.UploadFile)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
