package web_api

import (
	"fmt"
	"log"

	"github.com/ehabterra/go-site/internal/services"
	"gorm.io/gorm"

	"github.com/ehabterra/go-site/internal/controllers/go_site"

	"github.com/ehabterra/go-site/internal/environment"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	SiteCtrl *go_site.SiteCtrl
}

func NewServer(db *gorm.DB) *Server {
	siteService := services.NewGoSiteService(db)
	siteCtrl := go_site.NewSiteCtrl(siteService)

	return &Server{siteCtrl}
}

func (ser *Server) configure() *fiber.App {

	// fiber
	app := fiber.New()

	// add routes
	ser.ConfigureRoutes(app)

	return app
}

func (ser *Server) Serve() {
	serverPort := environment.GetEnv("SERVER_PORT", "3000")
	serverHost := environment.GetEnv("SERVER_HOST", "")

	app := ser.configure()

	err := app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))
	if err != nil {
		log.Fatal(err)
	}
}
