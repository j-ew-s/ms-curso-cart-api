package api

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/j-ew-s/ms-curso-cart-api/cart-services/models"
	"github.com/j-ew-s/ms-curso-cart-api/common"
	database "github.com/j-ew-s/ms-curso-cart-api/database"
	"github.com/valyala/fasthttp"
)

type CartApi struct {
	DBConn *redis.Client
}

func (c CartApi) GetCart(ctx *fasthttp.RequestCtx) {
	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	res := database.GetCart(c.DBConn, id)
	common.PrepareResponse(ctx, 200, res)
}

func (c CartApi) RemoveFromCart(ctx *fasthttp.RequestCtx) {
	fmt.Println("RemoveFromCart")

}

func (c CartApi) ClearCart(ctx *fasthttp.RequestCtx) {
	fmt.Println("ClearCart")

}

func (c CartApi) AddToCart(ctx *fasthttp.RequestCtx) {
	p := &models.Product{}

	err := json.Unmarshal(ctx.PostBody(), &p)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	currentCart := database.GetCart(c.DBConn, id)

	currentCart.Products = append(currentCart.Products, *p)

	database.AddToCart(c.DBConn, *currentCart, id)

	res := database.GetCart(c.DBConn, id)

	common.PrepareResponse(ctx, 201, res)

}
