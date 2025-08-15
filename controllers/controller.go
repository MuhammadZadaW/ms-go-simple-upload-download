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

	if file.Size > int64(c.config.Upload.LimitMB*1024*1024) {
		ctx.JSON(400, dto.ErrorDto{
			Status:  400,
			Message: "File size is too large",
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

func (c *ControllerStruct) MultiUpload(ctx *gin.Context) {

	result := make([]dto.StatusDto, 0)

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(400, dto.ErrorDto{
			Status:  400,
			Message: "Invalid form data",
		})
		return
	}

	files := form.File["files"]

	if len(files) == 0 {
		ctx.JSON(400, dto.ErrorDto{
			Status:  400,
			Message: "Files are required",
		})
		return
	}

	for _, file := range files {

		status := dto.StatusDto{
			Filename: file.Filename,
			Status:   200,
			Message:  "File uploaded successfully",
		}

		func() {
			if file.Size > int64(c.config.Upload.LimitMB*1024*1024) {
				status.Status = 400
				status.Message = "File size is too large"
				return
			}

			err = ctx.SaveUploadedFile(file, c.config.Upload.DestinationPath+"/"+file.Filename)
			if err != nil {
				status.Status = 500
				status.Message = "Failed to save file"
				return
			}
		}()

		result = append(result, status)
	}

	ctx.JSON(200, dto.MultiResponseDto{
		Status:  200,
		Message: "Files uploaded successfully",
		Data:    result,
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
