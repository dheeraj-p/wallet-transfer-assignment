package main

import (
	"log"
	"wallet-transfer-assignment/internal/api/routes"
	"wallet-transfer-assignment/internal/wallet"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	walletRepository := wallet.NewRepository()
	walletService := wallet.NewService(walletRepository)

	routes.RegisterHealthRoute(app)
	routes.RegisterWalletRoutes(app, walletService)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
