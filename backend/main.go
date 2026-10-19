package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	deliveryhttp "siakup-fs-test-starter/module/announcement/delivery/http"
	"siakup-fs-test-starter/module/announcement/repository/memory"
	"siakup-fs-test-starter/module/announcement/usecase"
)

func main() {
	repo := memory.NewRepository()
	uc := usecase.NewUseCase(repo)
	handler := deliveryhttp.NewAnnouncementHandler(uc)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Content-Type",
	}))
	handler.RegisterRoutes(app)

	log.Println("backend listening on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
