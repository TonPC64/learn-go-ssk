package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release")
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello")
	})
	router.Run(":5001")
}
