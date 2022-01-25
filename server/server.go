package server

import (
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gotrack/pkg"
)

func Run(dc pkg.DependencyContainer) error {
	server := fiber.New(
		fiber.Config{
			ErrorHandler: errorMiddleware,
		},
	)
	server.Use(recover.New())

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)
	go func() {
		<-quitChan
		_ = server.Shutdown()
		_ = dc.Close()
	}()
	return server.Listen(dc.Config.Server.GetAddress())
}
