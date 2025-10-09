package handler

import (
	"pretest-golang-tdi/model"
	"pretest-golang-tdi/repository"
	"pretest-golang-tdi/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateStoreHandler(c *fiber.Ctx) error {
	store := new(model.Store)
	if err := c.BodyParser(store); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Ambil user_id dari claims JWT yang sudah divalidasi oleh middleware
	claims := util.GetUserClaims(c)
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	userID := claims.UserID

	// Panggil repository dengan userID
	if err := repository.CreateStore(store, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(store)
}

func GetStoreByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid ID"})
	}

	store, err := repository.GetStoreByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(store)
}

func GetAllStoresHandler(c *fiber.Ctx) error {
	search := c.Query("search")

	stores, err := repository.GetAllStores(search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(stores)
}
