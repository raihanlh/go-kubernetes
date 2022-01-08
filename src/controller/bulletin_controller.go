package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/raihanlh/go-kubernetes/src/entity"
	"github.com/raihanlh/go-kubernetes/src/service"
)

type BulletinController struct {
	bulletinService service.BulletinService
}

func NewBulletinController(bulletinService service.BulletinService) Controller {
	return &BulletinController{bulletinService}
}

func (b *BulletinController) Route(app *fiber.App) {
	app.Get("/board", b.GetAll)
	app.Post("/board", b.Insert)
	app.Get("/", func(context *fiber.Ctx) error {
		return context.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"data":    "success",
		})
	})
}

func (b *BulletinController) GetAll(context *fiber.Ctx) error {
	results, err := b.bulletinService.GetAll()
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    results,
	})
}

func (b *BulletinController) Insert(context *fiber.Ctx) error {
	var bulletin entity.Bulletin
	err := context.BodyParser(&bulletin)
	bulletin.CreatedAt = time.Now()
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	err = b.bulletinService.InsertOne(bulletin)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	fmt.Println(bulletin)
	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    bulletin,
	})
}
