package main

import (
	"fmt"
	"log"
	"telmed_backend/config"
	"telmed_backend/db"
	"telmed_backend/migrations"
	"telmed_backend/routes"
	"telmed_backend/utils"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Hello world")
	utils.ImportEnv()
	config.LoadCfg()
	db.Connect()
	if config.MIGRATE {
		migrations.Migrate()
	}	

	app := fiber.New()
	app.Use(logger.New())
	
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
