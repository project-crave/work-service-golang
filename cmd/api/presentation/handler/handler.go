package api

import (
	"net/http"
	"crave/hub/cmd/work/cmd/api/presentation/controller"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctrl controller.IController
}

func NewHandler(ctrl controller.IController) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Craete(c *gin.Context) {

	

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
