package main

import (
	"net/http"

	"github.com/pluralsight/inventoryservice/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe("localhost:5000", nil)
}
