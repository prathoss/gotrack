package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gotrack/pkg"
	"gotrack/pkg/applogger"
)

func errorMiddleware(ctx *fiber.Ctx, err error) error {
	var enf pkg.ErrorNotFound
	if errors.As(err, &enf) {
		return ctx.Status(http.StatusNotFound).JSON(enf)
	}

	var eid pkg.ErrorInvalidData
	if errors.As(err, &eid) {
		return ctx.Status(http.StatusBadRequest).JSON(eid)
	}

	var eu pkg.ErrorUnauthorized
	if errors.As(err, &eu) {
		return ctx.Status(http.StatusUnauthorized).JSON(eu)
	}

	applogger.Error(
		"Internal server error",
		err,
		zap.String(
			"url",
			fmt.Sprintf("%s%s",
				string(ctx.Request().URI().Path()),
				string(ctx.Request().URI().QueryString()),
			),
		),
	)
	return ctx.SendStatus(http.StatusInternalServerError)
}
