package main

import (
	"log"
	"strings"
	"telmed_backend/config"
	"telmed_backend/db"
	"telmed_backend/migrations"
	"telmed_backend/routes"
	"telmed_backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	utils.ImportEnv()
	config.LoadCfg()
	db.Connect()
	if config.MIGRATE {
		migrations.Migrate()
	}

	db.InitServices()

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.HasPrefix(c.Path(), "api")
		},
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.MountRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
