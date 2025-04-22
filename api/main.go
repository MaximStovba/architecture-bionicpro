package main

import (
	"log"

	"bionic.pro/reports-api/routes"
	"bionic.pro/reports-api/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация подключения к JWKS
	if err := utils.InitJWKS(); err != nil {
		log.Fatalf("Failed to initialize JWKS: %v", err)
	}

	server := gin.Default()

	// Настройка CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Регистрация маршрутов
	routes.RegisterRoutes(server)

	log.Println("Server running on :8000")
	server.Run(":8000")
}
