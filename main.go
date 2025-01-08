package main

import (
	"donation-api/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador Gin
	router := gin.Default()

	// Configura o CORS
	router.Use(cors.Default())

	// Cria o controlador de doações
	donationController := controller.NewController()

	// Rotas da API
	router.GET("/api/doacoes", donationController.GetDonations)
	router.POST("/api/doacoes", donationController.CreateDonation)
	router.GET("/api/doacoes/:id", donationController.GetDonation)
	router.PUT("/api/doacoes/:id", donationController.UpdateDonation)
	router.DELETE("/api/doacoes/:id", donationController.DeleteDonation)

	// Rota para servir os arquivos estáticos
	router.Static("/static", "./frontend/static")
	router.StaticFile("/", "./frontend/static/index.html")

	router.Run(":8080")
}
