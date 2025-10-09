package router

import (
	"pretest-golang-tdi/handler"
	"pretest-golang-tdi/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", handler.RegisterHandler)
	api.Post("/login", handler.LoginHandler)

	auth := api.Group("", middleware.Protected())

	// --- Rute Khusus Pembeli ---
	auth.Post("/cart/items", middleware.Authorize("pembeli"), handler.AddToCartHandler)
	auth.Get("/cart", middleware.Authorize("pembeli"), handler.GetUserCartHandler)
	auth.Post("/checkout", middleware.Authorize("pembeli"), handler.CheckoutHandler)
	auth.Get("/transactions", middleware.Authorize("pembeli"), handler.GetUserOrdersHandler)
	// Catatan: GetOrderByID bisa juga diakses oleh penjual untuk melihat detail order tokonya
	auth.Get("/transactions/:id", middleware.Authorize("pembeli"), handler.GetOrderByIDHandler)
	auth.Post("/transactions/:id/pay", middleware.Authorize("pembeli"), handler.PayForOrderHandler)

	// --- Rute Khusus Penjual ---
	auth.Post("/stores", middleware.Authorize("penjual"), handler.CreateStoreHandler)
	auth.Post("/products", middleware.Authorize("penjual"), handler.CreateProductHandler)
	// Tambahkan rute lain untuk penjual di sini...

	// --- Rute Khusus Manajer ---
	// auth.Get("/management/all-transactions", middleware.Authorize("manager"), handler.GetAllTransactionsHandler) // Contoh rute baru untuk ke depannya
	// Tambahkan rute lain untuk manajer di sini...

	// --- Rute yang Bisa Diakses oleh Beberapa Peran (Shared Routes) ---
	auth.Get("/products", middleware.Authorize("pembeli", "penjual", "manager"), handler.GetAllProductsHandler)
	auth.Get("/products/:id", middleware.Authorize("pembeli", "penjual", "manager"), handler.GetProductByIDHandler)
	auth.Get("/stores", middleware.Authorize("pembeli", "penjual", "manager"), handler.GetAllStoresHandler)
	auth.Get("/stores/:id", middleware.Authorize("pembeli", "penjual", "manager"), handler.GetStoreByIDHandler)
	auth.Get("/transactions/:id", middleware.Authorize("pembeli", "manager"), handler.GetOrderByIDHandler)
}
