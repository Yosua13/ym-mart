package router

import (
	"pretest-golang-tdi/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Rute untuk Toko
	api.Post("/stores", handler.CreateStoreHandler)
	api.Get("/stores", handler.GetAllStoresHandler)
	api.Get("/stores/:id", handler.GetStoreByIDHandler)

	// Rute untuk Produk
	api.Post("/products", handler.CreateProductHandler)
	api.Get("/products", handler.GetAllProductsHandler)
	api.Get("/products/:id", handler.GetProductByIDHandler)

	// Rute untuk Keranjang Belanja
	api.Post("/cart/items", handler.AddToCartHandler)
	api.Get("/cart/user/:user_id", handler.GetUserCartHandler)

	// Rute untuk Transaksi
	api.Post("/checkout", handler.CheckoutHandler)
	api.Get("/transactions/user/:user_id", handler.GetUserOrdersHandler)
	api.Get("/transactions/:id", handler.GetOrderByIDHandler)
	api.Post("/transactions/:id/pay", handler.PayForOrderHandler)
}
