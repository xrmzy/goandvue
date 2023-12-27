package main

import (
	"fmt"
	"log"
	"rmzstartup/handler"
	"rmzstartup/repository"
	"rmzstartup/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() error {
	dsn := "host=localhost user=postgres password=1234 dbname=db_rmz port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect database: %s", err))
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to get Database instance: %s", err))
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

	userByEmail, err := userRepository.FindByEmail("testing1111@mail.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	if userByEmail.Id == uuid.Nil {
		fmt.Println("user not Found")
	} else {
		fmt.Println(userByEmail.Name)
	}

	fmt.Println(userByEmail.Name)

	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	router.Run()
}
