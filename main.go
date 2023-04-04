package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create router

	router := gin.Default()

	// Test route

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// File upload route

	router.POST("/upload", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["files[]"]

		for i, file := range files {
			log.Println(fmt.Sprintf("File %d: %s", i+1, file.Filename))
		}

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d files uploaded", len(files))})
	})

	// Start server

	err := router.Run()

	if err != nil {
		log.Println(err)
	}
}
