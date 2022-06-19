package cartService

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/j-ew-s/ms-curso-cart-api/cart-services/api"
)

func SetRoutes(router *fasthttprouter.Router) {
	router.GET("/", api.GetCart)
	router.POST("/", api.AddToCart)
	router.DELETE("/", api.RemoveFromCart)
	router.DELETE("/", api.ClearCart)
}
