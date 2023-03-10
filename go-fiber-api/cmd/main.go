package main

import (
	"log"

	"github.com/Amarendrapratihari/go-fiber-api/pkg/books"
	//"github.com/Amarendrapratihari/go-fiber-api/pkg/common/config"
	config "github.com/Amarendrapratihari/go-fiber-api/pkg/common/config/envs"
	"github.com/Amarendrapratihari/go-fiber-api/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
