package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	server = gin.Default()
}

func main() {

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	// fmt.Println("In main.go")

	// log.Fatal(server.Run(":" + config.ServerPort))
}
