package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	// replace with your own docs folder, usually "github.com/sing3demons/reponame/docs"
	_ "swagger/recipes/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/", hello)
	app.Get("/accounts/:id", ShowAccount)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}

// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Param id path int true "Account ID"
// @Router /accounts/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	str := c.Params("id")
	id, err := strconv.Atoi(str)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	return c.JSON(fiber.Map{
		"id": id,
	})
}

type Account struct {
	id int64
}

type HTTPError struct {
	status  string
	message string
}
