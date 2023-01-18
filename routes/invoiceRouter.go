package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/soumyaa03/pizzy-backend/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoice", controller.GetInvoices())
	incomingRoutes.GET("/invoice/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoice", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoice/:invoice_id", controller.UpdateInvoice())

}
