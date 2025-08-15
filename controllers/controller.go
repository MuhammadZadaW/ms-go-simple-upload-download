package controllers

import (
	"fmt"
	"github-port/ms-go-simple-upload-download/configs"
	"github-port/ms-go-simple-upload-download/dto"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ControllerStruct struct {
	config *configs.ConfigStruct
}

func ProvideController(config *configs.ConfigStruct) *ControllerStruct {
	return &ControllerStruct{
		config: config,
	}
}

func (c *ControllerStruct) Upload(ctx *gin.Context) {

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, dto.ErrorDto{
			Status:  400,
			Message: "File is required",
		})
		return
	}

	name := ctx.PostForm("name")
	if name != "" {
		ext := filepath.Ext(file.Filename)
		file.Filename = fmt.Sprintf("%s%s", name, ext)
	}

	err = ctx.SaveUploadedFile(file, c.config.Upload.DestinationPath+"/"+file.Filename)
	if err != nil {
		ctx.JSON(500, dto.ErrorDto{
			Status:  500,
			Message: "Failed to save file",
		})
		return
	}

	ctx.JSON(200, dto.SuccessDto{
		Status:  200,
		Message: "File uploaded successfully",
	})
}

func (c *ControllerStruct) ListFiles(ctx *gin.Context) {

	files, err := os.ReadDir(c.config.Download.SourcePath)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status":  500,
			"message": "Failed to list files",
		})
		return
	}

	var filesList []string
	for _, file := range files {
		filesList = append(filesList, file.Name())
	}

	ctx.JSON(200, dto.SuccessListDto{
		Status:  200,
		Message: "Files listed successfully",
		Data:    filesList,
	})
}

func (c *ControllerStruct) Download(ctx *gin.Context) {

	filename := ctx.Param("filename")

	if filename == "" {
		ctx.JSON(400, dto.ErrorDto{
			Status:  400,
			Message: "Filename is required",
		})
		return
	}

	_, err := os.Stat(c.config.Download.SourcePath + "/" + filename)
	if os.IsNotExist(err) {
		ctx.JSON(404, dto.ErrorDto{
			Status:  404,
			Message: "File not found",
		})
		return
	}

	ctx.FileAttachment(c.config.Download.SourcePath+"/"+filename, filename)
}
