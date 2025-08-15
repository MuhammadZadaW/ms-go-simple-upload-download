package main

import (
	"fmt"
	"github-port/ms-go-simple-upload-download/configs"
	"github-port/ms-go-simple-upload-download/controllers"
	"github-port/ms-go-simple-upload-download/dto"

	"github.com/gin-gonic/gin"
)

func main() {

	configs, err := configs.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	r.Use(func(ctx *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(500, dto.ErrorDto{
					Status:  500,
					Message: "Internal server error",
				})
			}
		}()

		ctx.Next()
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, dto.ErrorDto{
			Status:  404,
			Message: "Route not found",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, dto.SuccessDto{
			Status:  200,
			Message: "Server is running",
		})
	})

	ctrl := controllers.ProvideController(configs)
	r.POST("/upload", ctrl.Upload)
	r.POST("/multi-upload", ctrl.MultiUpload)
	r.GET("/list", ctrl.ListFiles)
	r.GET("/download/:filename", ctrl.Download)

	r.Run(fmt.Sprintf(":%d", configs.Server.Port))
}
