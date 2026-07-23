package internal

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDir = "files"

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dst := filepath.Join(uploadDir, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": file.Filename,
		"size":     file.Size,
		"path":     dst,
	})
}

func DownloadFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "File downloaded successfully"})
}

func DeleteFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "File delete successfully"})
}

func GetFileList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get File List successfully"})
}
