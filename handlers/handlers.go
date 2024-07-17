package handlers

import (
	"multipleParam_git/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	GetUniversalInfoHandlers2(c *gin.Context)
}

type handlerAdapter struct {
	s services.ServicePort
}

func NewHanerhandlerAdapter(s services.ServicePort) HandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) GetUniversalInfoHandlers2(c *gin.Context) {
	// catalog := c.Query("catalog")
	// if catalog == "" {
	catalog := c.QueryArray("catalog")
	if len(catalog) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Missing catalog query parameter"})
		return
	}
	dataResponse, err := h.s.GetUniversalInfoServices2(catalog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "dataResponse": dataResponse})
}
