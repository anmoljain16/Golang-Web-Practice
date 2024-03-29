package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Fullname struct {
	FirstName string
	LastName  string
}

func getAnmol(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(&fname)
}

var fname Fullname

func postAnmol(ctx *fiber.Ctx) error {
	body := new(Fullname)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadGateway).SendString(err.Error())
		return err
	}
	fname = Fullname{
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	return ctx.Status(fiber.StatusOK).JSON(fname)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	appgroup := app.Group("/anmol")
	appgroup.Get("", getAnmol)
	appgroup.Post("", postAnmol)
	app.Listen(":80")
}
