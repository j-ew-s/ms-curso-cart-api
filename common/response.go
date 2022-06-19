package common

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func PrepareResponse(ctx *fasthttp.RequestCtx, code int, response interface{}) {

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	ctx.Response.SetStatusCode(code)

	start := time.Now()

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), response)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}
