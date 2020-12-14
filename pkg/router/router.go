package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orlowskilp/api-go-aws-ecs/pkg/sys"
)

// GetKernelMethod handles GET for /kernel
func GetKernelMethod(ctx *gin.Context) {
	ctx.String(http.StatusOK, sys.GetKernelInfo())
}

// GetHostnameMethod handles GET for /hostname
func GetHostnameMethod(ctx *gin.Context) {
	ctx.String(http.StatusOK, sys.GetHostname())
}

// SetupRouter registers API endpoints and returns a router instance
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hostname", GetHostnameMethod)
	router.GET("/kernel", GetKernelMethod)

	return router
}
