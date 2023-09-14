package main

import (
	"github.com/finaenno/sertifikasi-go.git/controllers/authenticationController"
	"github.com/finaenno/sertifikasi-go.git/controllers/penawaranSertifikasiController"
	"github.com/finaenno/sertifikasi-go.git/middlewares"
	"github.com/finaenno/sertifikasi-go.git/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	api := r.Group("/api")
	{

		api.POST("/register", authenticationController.Register)
		api.POST("/login", authenticationController.Login)

		protected := api.Group("protected").Use(middlewares.Authz())
		{
			protected.GET("/penawaranSertifikasi", penawaranSertifikasiController.Index)
			protected.GET("/penawaranSertifikasi/:id", penawaranSertifikasiController.Show)
			protected.POST("/penawaranSertifikasi", penawaranSertifikasiController.Create)
			protected.PUT("/penawaranSertifikasi/:id", penawaranSertifikasiController.Update)
			protected.DELETE("/penawaranSertifikasi", penawaranSertifikasiController.Delete)

		}
	}

}
