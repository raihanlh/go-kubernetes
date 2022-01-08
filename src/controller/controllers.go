package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Route(app *fiber.App)
}

func RouteAll(app *fiber.App, controllers ...Controller) {
	for _, controller := range controllers {
		controller.Route(app)
	}
}
