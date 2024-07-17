package web

import (
	"em/internal/db/models"
	"em/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type WebController struct {
	srv *service.TaskService
	app *fiber.App
	log *log.Logger
}

func CreateNewWebController(app *fiber.App, log *log.Logger, srv *service.TaskService) *WebController {
	return &WebController{
		srv: srv,
		app: app,
		log: log,
	}
}

// service's handlers
func (wc *WebController) RegisterRouters() {
	// http://127.0.0.1:1200/user/create
	wc.app.Get("/user/create", func(c *fiber.Ctx) error {
		var req models.UserModel
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.CreateUser(&req))
	})

	// http://127.0.0.1:1200/user/delete
	wc.app.Get("/user/delete", func(c *fiber.Ctx) error {
		var req service.IdRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.DeleteUser(req.ID))
	})

	// http://127.0.0.1:1200/user/update
	wc.app.Get("/user/update", func(c *fiber.Ctx) error {
		var req models.UserModel
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.UpdateUser(&req))
	})

	// http://127.0.0.1:1200/user/get-users
	wc.app.Get("/user/get-users", func(c *fiber.Ctx) error {
		var req *models.UserModel
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.GetUser(req))
	})

	// http://127.0.0.1:1200/user/get-time
	wc.app.Get("/user/get-time", func(c *fiber.Ctx) error {
		var req service.IdRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.GetSummaryTime(req.ID))
	})

	// http://127.0.0.1:1200/task/create
	wc.app.Get("/task/create", func(c *fiber.Ctx) error {
		var req models.TaskModel
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.CreateTask(&req))
	})

	// http://127.0.0.1:1200/task/start
	wc.app.Get("/task/start", func(c *fiber.Ctx) error {
		var req service.IdRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.StartTask(req.ID))
	})

	// http://127.0.0.1:1200/task/finish
	wc.app.Get("/task/finish", func(c *fiber.Ctx) error {
		var req service.IdRequest
		if err := c.BodyParser(&req); err != nil {
			wc.log.Fatal(err)
		}

		return c.JSON(wc.srv.FinishTask(req.ID))
	})

}
