package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/template/django"

	"meta-go/db"
	"meta-go/router"
	"meta-go/util/middleware"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName: "meta-go",
		Views:   django.New("./template", ".html"),
	})

	app.Static("/static", "./static")

	app.Use(
		middleware.Gate(),
		encryptcookie.New(encryptcookie.Config{
			Key: "vtDs1VZzSZ+flREhrrybKFBY8j1K8g0NyvRmv9+8MRA=",
		}),
		favicon.New(favicon.Config{
			File: "./static/favicon.ico",
		}),
		middleware.UserAuth(),
	)

	db.Init()
	router.Register(app)

	start(app)
}

func start(app *fiber.App) {
	if err := app.Listen(":9000"); err != nil {
		log.Panic(err)
	}
}

func gracefulStart(app *fiber.App) {
	go func() {
		if err := app.Listen(":9000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	_ = <-c
	log.Println("Gracefully shutting down...")

	_ = app.Shutdown()

	log.Println("Shutdown successfully.")
}
