package http_handler

import (
	"net/http"
	"serverless/controller"
	"serverless/domain"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var payload domain.LoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, "invalid json data"))
		return
	}

	controller.Login(c, payload)
}
