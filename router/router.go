package router

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"product-information-updater/app/dbUtils"
	"product-information-updater/app/queueUtils"
	"product-information-updater/app/updatePrice/controller"
	"product-information-updater/app/updatePrice/repository"
	"product-information-updater/app/updatePrice/service"
)

func InitializeRouter(snsTopicARN string, snsSession *sns.SNS, mongoCollection *mongo.Collection) *gin.Engine {
	router := gin.Default()

	snsSess := queueUtils.NewSNSSession(snsSession)
	mongoColl := dbUtils.NewMongoColl(mongoCollection)
	productUpdateInfoRepo := priceUpdateRepository.NewProductUpdateInfoRepo(mongoColl)
	priceUpdateSvc := priceUpdateService.NewPriceUpdateService(productUpdateInfoRepo, snsTopicARN, snsSess)
	priceUpdateHandler := priceUpdatecontroller.NewPriceUpdateHandler(priceUpdateSvc)

	router.POST("api/v1/products/:productID", priceUpdateHandler.UpdatePrice)

	return router
}
