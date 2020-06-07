package http

import (
	"net/http"

	"github.com/dm/chat-x-back/auth"
	"github.com/dm/chat-x-back/models"
	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	useCase auth.UseCase
}

// NewHandler ...
func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// SignUp ...
func (h *Handler) SignUp(c *gin.Context) {
	data := new(models.User)

	if err := c.BindJSON(data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), data.FirstName, data.LastName, data.EMail, data.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
