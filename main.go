package main

import (
	"github.com/AntonVTR/probation_test/grapthql_fasthttp/testutil"

	handler "github.com/lab259/graphql-fasthttp-handler"
	"github.com/valyala/fasthttp"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &testutil.EmployerSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		h.ServeHTTP(ctx)
	})
}
