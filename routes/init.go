package routes

import (
	"net/http"
	http_handler "serverless/http"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, response.Success(map[string]any{
			"message": "pong",
		}))
	})

	r.POST("/login", http_handler.Login)
	// for nostr NIP-05
	// r.GET("/.well-known/nostr.json", controller.Cors, controller.NIP05)
}
