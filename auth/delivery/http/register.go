package http

import (
	"github.com/dm/chat-x-back/auth"
	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/signUp", h.SignUp)
		// authEndpoints.POST("/signIn", h.SignIn)
	}
}
