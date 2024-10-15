package router

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"product-information-updater/app/updatePrice/controller"
	"product-information-updater/app/updatePrice/repository"
	"product-information-updater/app/updatePrice/service"
)

func InitializeRouter(snsTopicARN string, snsSession *sns.SNS, mongoCollection *mongo.Collection) *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	productUpdateInfoRepo := priceUpdateRepository.NewProductUpdateInfoRepo(mongoCollection)
	priceUpdateSvc := priceUpdateService.NewPriceUpdateService(productUpdateInfoRepo, snsTopicARN, snsSession)
	priceUpdateHandler := priceUpdatecontroller.NewPriceUpdateHandler(priceUpdateSvc)

	router.POST("/submit", priceUpdateHandler.HandleSubmit)

	return router
}
