package main

import (
	"net/http"

	"github.com/pluralsight/inventoryservice/database"
	"github.com/pluralsight/inventoryservice/product"
	"github.com/pluralsight/inventoryservice/receipt"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	receipt.SetupRoutes(apiBasePath)
	http.ListenAndServe("localhost:5000", nil)
}
