package main

import (
	"database/sql"
	"log"
	"tasked/internal/config"
	"tasked/internal/handler"
	"tasked/internal/middleware"
	"tasked/internal/repository"
	"tasked/internal/services"

	"github.com/gin-contrib/cors"
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

	//Users
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//Task
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:4173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	router.Use(middleware.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	router.GET("/tasks/:id", taskHandler.GetTask)
	router.GET("/users/:id/tasks", taskHandler.ListTasksByUser)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.PATCH("/tasks/:id/status", taskHandler.UpdateStatus)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)
	router.POST("/tasks", taskHandler.CreateTask)

	router.Run(":" + cfg.Port)
}
