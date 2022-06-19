package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func GetCart(ctx *fasthttp.RequestCtx) {
	fmt.Println("GetCart")
}

func RemoveFromCart(ctx *fasthttp.RequestCtx) {
	fmt.Println("ClearCart")

}

func ClearCart(ctx *fasthttp.RequestCtx) {
	fmt.Println("ClearCart")

}

func AddToCart(ctx *fasthttp.RequestCtx) {
	fmt.Println("ClearCart")

}
