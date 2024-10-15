package priceUpdatecontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	updatePriceModel "product-information-updater/app/updatePrice/models"
	"product-information-updater/app/updatePrice/service"
)

type Handler interface {
	UpdatePrice(c *gin.Context) // rename
}

type handler struct {
	updateSvc priceUpdateService.Service
}

func NewPriceUpdateHandler(updateSvc priceUpdateService.Service) Handler {
	return &handler{
		updateSvc: updateSvc,
	}
}

func (h *handler) UpdatePrice(ginCtx *gin.Context) {
	productID := ginCtx.Param("productID")
	if productID == "" {
		log.Printf("product id not found")
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "product id not found"})
		return
	}

	var reqBody updatePriceModel.RequestBody
	if err := ginCtx.ShouldBindJSON(&reqBody); err != nil {
		log.Printf("Failed to unmarshall request body: %v", err)
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.updateSvc.Process(ginCtx, productID, reqBody)
	if err != nil {
		log.Printf("Failed to process request: %v", err)
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
	}

	ginCtx.JSON(http.StatusOK, gin.H{"status": "success"})
}
