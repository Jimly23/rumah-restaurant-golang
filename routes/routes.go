package routes

import (
	"rumah-restaurant/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST(
		"/promo",
		handlers.CreatePromo,
	)

	router.GET(
		"/promo",
		handlers.GetAllPromo,
	)

	router.GET(
		"/promo/:id",
		handlers.GetPromoByID,
	)

	router.PUT(
		"/promo/:id",
		handlers.UpdatePromo,
	)

	router.DELETE(
		"/promo/:id",
		handlers.DeletePromo,
	)
}
