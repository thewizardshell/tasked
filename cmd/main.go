package main

import (
	"database/sql"
	"log"
	"tasked/internal/config"
	"tasked/internal/handler"
	"tasked/internal/middleware"
	"tasked/internal/repository"
	"tasked/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	_ "tasked/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Tasked API
// @version 1.0
// @description API para gestión de tareas
// @host localhost:8080
// @BasePath /
func main() {
	godotenv.Load()
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(middleware.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	router.Run(":" + cfg.Port)
}
