package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKernelMethod handles GET for /kernel
func GetKernelMethod(ctx *gin.Context) {
	ctx.String(http.StatusOK, GetKernelInfo())
}

// GetHostnameMethod handles GET for /hostname
func GetHostnameMethod(ctx *gin.Context) {
	ctx.String(http.StatusOK, GetHostname())
}

// SetupRouter registers API endpoints and returns a router instance
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hostname", GetHostnameMethod)
	router.GET("/kernel", GetKernelMethod)

	return router
}
