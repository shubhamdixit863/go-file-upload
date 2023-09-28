package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goFileService/internal/application"
	"net/http"
	"path/filepath"
)

type Handler struct {
	FileUplaodService application.FileUploadService
}

func (h Handler) UploadFile(c *gin.Context) {
	/// Our handler to upload files

	// Getting the file and uploading it to the local directory

	file, _ := c.FormFile("file")
	fileName := filepath.Base(file.Filename)
	dest := "./uploads/" + fileName

	fmt.Println()
	err := c.SaveUploadedFile(file, dest)
	absPath, err := filepath.Abs(dest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Call the file upload method of AWS
	filePath, err := h.FileUplaodService.Upload(absPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"filePath": filePath,
	})

}
