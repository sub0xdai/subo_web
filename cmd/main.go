package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/sub0x/subo_web/internal/handlers"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "SUBO Portfolio",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Static files
	app.Static("/static", "./web/static", fiber.Static{
		ByteRange: true,
		Browse: false,
		MaxAge: 3600,
	})

	// Routes
	app.Get("/", handlers.HandleHome)
	app.Get("/projects", handlers.HandleProjects)
	app.Get("/blog", handlers.HandleBlog)
	app.Get("/about", handlers.HandleAbout)
	
	// API Routes
	app.Post("/api/chat", handlers.HandleChatPost)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
