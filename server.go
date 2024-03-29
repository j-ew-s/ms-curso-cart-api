package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/go-redis/redis"
	cartService "github.com/j-ew-s/ms-curso-cart-api/cart-services"
	"github.com/j-ew-s/ms-curso-cart-api/database"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	db := SetDataBase()
	client := cartService.SetCartSetvice(db)
	cartService.SetRoutes(router, client)
	fasthttp.ListenAndServe(":5300", CORS(router.Handler))
}

var (
	corsAllowHeaders     = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// CORS handles CORS
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}

func SetDataBase() *redis.Client {

	return database.SetupRedis()

}
