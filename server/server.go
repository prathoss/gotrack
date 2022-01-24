package server

import (
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"gotrack/pkg"
)

func Run(dc pkg.DependencyContainer) error {
	server := fiber.New()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)
	go func() {
		<-quitChan
		_ = server.Shutdown()
	}()
	if err := server.Listen(dc.Config.Server.GetAddress()); err != nil {
		return err
	}
	return nil
}
