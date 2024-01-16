package main

import (
	"fmt"
	"log"
	"os"
	"rmzstartup/auth"
	"rmzstartup/handler"
	"rmzstartup/middleware"
	"rmzstartup/payment"
	"rmzstartup/repository"
	"rmzstartup/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env file")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect database: %s", err))
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get database instance: %s", err))
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(1000)
	db = conn
	return nil
}

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("Failed to initialize database : %v", err)
	}

	userRepository := repository.NewUserRepository(db)
	campaignRepository := repository.NewCampaignRepo(db)
	transactionRepository := repository.NewTransactionRepo(db)

	userService := service.NewUserService(userRepository)
	authService := auth.NewJWTService()
	campaignService := service.NewCampaignService(campaignRepository)
	paymentService := payment.NewPaymentService(transactionRepository, campaignRepository)
	transactionService := service.NewTransactionService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvalaible)
	api.POST("/avatars", middleware.AuthMiddleWare(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", middleware.AuthMiddleWare(authService, userService), userHandler.FetchUser)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleWare(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleWare(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleWare(authService, userService), campaignHandler.UploadImageCampaign)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleWare(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", middleware.AuthMiddleWare(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", middleware.AuthMiddleWare(authService, userService), transactionHandler.CreateTrasaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.Run()
}
