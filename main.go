package main

import (
	"fmt"
	"log"
	"os"
	"rmzstartup/auth"
	"rmzstartup/handler"
	"rmzstartup/middleware"
	"rmzstartup/repository"
	"rmzstartup/service"

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
	userService := service.NewUserService(userRepository)
	authService := auth.NewJWTService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvalaible)
	api.POST("/avatars", middleware.AuthMiddleWare(authService, userService), userHandler.UploadAvatar)
	router.Run()
}

// func authMiddleWare(authService auth.JWTService, userService service.UserService) gin.HandlerFunc {
// 	func(c *gin.Context) {
// 		autHeader := c.GetHeader("Authorization")

// 		if !strings.Contains(autHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}

// 		tokenString := ""
// 		arrayToken := strings.Split(autHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		token, err = authService.ValidateToken(tokenString)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
// 			return
// 		}
// 		claim, ok := token.ClaimSet(jwt.MapClaims)
// 		if !ok || token.
// 	}
// }
