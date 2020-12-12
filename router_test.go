package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetKernelMethod_NoParamsReqd_ReturnStatus200(t *testing.T) {
	// Assemble
	gin.SetMode(gin.TestMode)
	httpResponse := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpResponse)
	ginContext.Request = httptest.NewRequest("GET", "http://localhost/kernel", nil)

	// Act
	GetKernelMethod(ginContext)
	actualResponseBody, _ := ioutil.ReadAll(httpResponse.Body)

	// Assert
	assert.Equal(t, 200, httpResponse.Code, "Expected status code 200")
	// Expected value should be hard-coded
	assert.Equal(t, GetKernelInfo(), string(actualResponseBody))
}

func TestGetHostnameMethod_NoParamsReqd_ReturnStatus200(t *testing.T) {
	// Assemble
	gin.SetMode(gin.TestMode)
	httpResponse := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpResponse)
	ginContext.Request = httptest.NewRequest("GET", "http://localhost/hostname", nil)

	// Act
	GetHostnameMethod(ginContext)
	actualResponseBody, _ := ioutil.ReadAll(httpResponse.Body)

	// Assert
	assert.Equal(t, 200, httpResponse.Code, "Expected status code 200")
	// Expected value should be hard-coded
	assert.Equal(t, GetHostname(), string(actualResponseBody))
}
