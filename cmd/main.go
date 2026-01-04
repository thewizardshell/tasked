package main

import (
	"database/sql"
	"log"
	"tasked/internal/auth"
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

// @title           Tasked API
// @version         1.0
// @description     API para gestión de tareas con JWT
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Escribe "Bearer" seguido de un espacio y el JWT token.
func main() {
	godotenv.Load()
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tokenManager := auth.NewTokenManager(cfg.JWTSecret, cfg.JWTExpiryHrs)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, tokenManager)

	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	router.Use(middleware.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/login", userHandler.Login)
	router.POST("/users", userHandler.CreateUser)

	authMiddleware := middleware.AuthRequired(tokenManager)

	router.GET("/users/:id", authMiddleware, userHandler.GetUser)
	router.PUT("/users/:id", authMiddleware, userHandler.UpdateUser)
	router.DELETE("/users/:id", authMiddleware, userHandler.DeleteUser)

	router.GET("/tasks/:id", authMiddleware, taskHandler.GetTask)
	router.GET("/users/:id/tasks", authMiddleware, taskHandler.ListTasksByUser)
	router.PUT("/tasks/:id", authMiddleware, taskHandler.UpdateTask)
	router.PATCH("/tasks/:id/status", authMiddleware, taskHandler.UpdateStatus)
	router.DELETE("/tasks/:id", authMiddleware, taskHandler.DeleteTask)
	router.POST("/tasks", authMiddleware, taskHandler.CreateTask)

	router.Run(":" + cfg.Port)
}
