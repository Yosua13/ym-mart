package handler

import (
	"database/sql"
	"pretest-golang-tdi/repository"
	"pretest-golang-tdi/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CheckoutHandler(c *fiber.Ctx) error {
	// Ambil user_id dari claims JWT
	claims := util.GetUserClaims(c)
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	userID := claims.UserID

	order, err := repository.Checkout(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func GetUserOrdersHandler(c *fiber.Ctx) error {
	// Ambil user_id dari claims JWT
	claims := util.GetUserClaims(c)
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	userID := claims.UserID

	orders, err := repository.GetUserOrders(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

func GetOrderByIDHandler(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	}
	return c.JSON(order)
}

func PayForOrderHandler(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
	}

	err = repository.UpdateOrderStatus(orderID, "Selesai")
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Payment successful, order status updated to Selesai"})
}
