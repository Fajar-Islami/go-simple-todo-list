package http

import (
	"context"
	"fmt"
	"go-todo-list/internal/delivery/http/handler"
	"go-todo-list/internal/infrastructure/container"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func HTTPRouteInit(containerConf *container.Container) {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("", HealthCheck)

	api := app.Group("/api/v1") // /api

	agApi := api.Group("/activity-groups")
	handler.ActivitiesRoute(agApi, containerConf)

	todoApi := api.Group("/todo-items")
	handler.TodosRoute(todoApi, containerConf)

	// Start server
	port := fmt.Sprintf("%s:%d", containerConf.Apps.Host, containerConf.Apps.HttpPort)
	if containerConf.Apps.HttpPort == 0 {
		port = ":8090"
	}
	go func() {
		if err := app.Listen(port); err != nil && err != http.ErrServerClosed {
			app.Server().Logger.Printf("shutting down the server : ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		app.Server().Logger.Printf("shutting down the server :", err)
	}

}

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": "Server is up and running",
	})
}
