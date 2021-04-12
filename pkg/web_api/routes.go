package web_api

import (
	"github.com/ehabterra/go-site/internal/controllers/home"
	"github.com/gofiber/fiber/v2"
)

// ConfigureRoutes to add all routes
func (ser *Server) ConfigureRoutes(app *fiber.App) {
	app.Get("/", home.Index)

	app.Get("/site/:id", ser.SiteCtrl.GetSite)
	app.Get("/site", ser.SiteCtrl.GetSites)

	app.Post("/site", ser.SiteCtrl.CreateSite)
	app.Put("/site/:id", ser.SiteCtrl.UpdateSite)
	app.Delete("/site/:id", ser.SiteCtrl.DeleteSite)
}
