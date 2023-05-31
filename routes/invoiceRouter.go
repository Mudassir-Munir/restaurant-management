package routes

import (
	"github.com/Mudassir-Munir/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/invoices", controllers.GetInvoices())
	incommingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incommingRoutes.GET("/invoices", controllers.CreateInvoice())
	incommingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}
