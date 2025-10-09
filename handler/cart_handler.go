package handler

import (
	"pretest-golang-tdi/repository"
	"pretest-golang-tdi/util"

	"github.com/gofiber/fiber/v2"
)

type AddToCartRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func AddToCartHandler(c *fiber.Ctx) error {
	claims := util.GetUserClaims(c)
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	userID := claims.UserID

	req := new(AddToCartRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if req.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Quantity must be positive"})
	}

	err := repository.AddItemToCart(userID, req.ProductID, req.Quantity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Item added to cart successfully"})
}

func GetUserCartHandler(c *fiber.Ctx) error {
	claims := util.GetUserClaims(c)
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	userID := claims.UserID

	cart, err := repository.GetUserCart(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Cart is empty or not found", "cart": nil})
	}

	totalPrice := 0.0
	for _, item := range cart.CartItems {
		totalPrice += item.Product.Price * float64(item.Quantity)
	}

	return c.JSON(fiber.Map{
		"cart":        cart,
		"total_price": totalPrice,
	})
}
