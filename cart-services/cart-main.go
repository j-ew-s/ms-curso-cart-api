package cartService

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/j-ew-s/ms-curso-cart-api/cart-services/api"
	sessionServices "github.com/j-ew-s/ms-curso-cart-api/session-services"
)

func SetRoutes(router *fasthttprouter.Router) {
	router.GET("/", sessionServices.AuthSessionValidator(api.GetCart))
	router.POST("/", sessionServices.AuthSessionValidator(api.AddToCart))
	router.DELETE("/:id", sessionServices.AuthSessionValidator(api.RemoveFromCart))
	router.DELETE("/", sessionServices.AuthSessionValidator(api.ClearCart))
}
