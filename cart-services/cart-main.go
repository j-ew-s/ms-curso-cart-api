package cartService

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/go-redis/redis"
	authServices "github.com/j-ew-s/ms-curso-cart-api/auth-services"
	"github.com/j-ew-s/ms-curso-cart-api/cart-services/api"
)

type CartService struct {
	*api.CartApi
}

func SetCartSetvice(db *redis.Client) *api.CartApi {

	cartServices := &api.CartApi{
		DBConn: db,
	}

	return cartServices
}

func SetRoutes(router *fasthttprouter.Router, client *api.CartApi) {
	router.GET("/:id", authServices.AuthSessionValidator(client.GetCart))
	router.POST("/:id", authServices.AuthSessionValidator(client.AddToCart))
	router.DELETE("/:id", authServices.AuthSessionValidator(client.RemoveFromCart))
	router.DELETE("/", authServices.AuthSessionValidator(client.ClearCart))
}
