package handlers

import (
	"wallet-transfer-assignment/internal/api/dtos"
	"wallet-transfer-assignment/internal/wallet"

	"github.com/gofiber/fiber/v2"
)

func TransferHandler(service wallet.WalletService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dtos.TransferRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}

		err := service.TransferFunds(c.Context(), req)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		// Placeholder response for initial scaffolding
		return c.Status(200).JSON(fiber.Map{"status": "received", "idempotencyKey": req.IdempotencyKey})
	}
}
