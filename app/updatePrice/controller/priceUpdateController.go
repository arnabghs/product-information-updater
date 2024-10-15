package priceUpdatecontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	updatePriceModel "product-information-updater/app/updatePrice/models"
	"product-information-updater/app/updatePrice/service"
)

type Handler interface {
	HandleSubmit(c *gin.Context) // rename
}

type handler struct {
	updateSvc priceUpdateService.Service
}

func NewPriceUpdateHandler(updateSvc priceUpdateService.Service) Handler {
	return &handler{
		updateSvc: updateSvc,
	}
}

func (h *handler) HandleSubmit(c *gin.Context) {
	var reqBody updatePriceModel.RequestBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.updateSvc.Process(reqBody)
	if err != nil {
		log.Printf("Failed to process request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
