package controller

import (
	"net/http"

	"github.com/BevisDev/BevisBot/internal/app/config"
	"github.com/gin-gonic/gin"
)

type SupportController struct {
}

func NewSupportController() *SupportController {
	return &SupportController{}
}

func (s *SupportController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"profile": config.AppConfig.Server.Profile,
		"version": config.AppConfig.Server.Version,
	})
}
