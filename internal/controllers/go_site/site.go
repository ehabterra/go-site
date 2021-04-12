package go_site

import (
	"errors"
	"log"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/ehabterra/go-site/internal/models"

	"github.com/gofiber/fiber/v2"
)

type SiteService interface {
	Create(site *models.GoSite)
	Update(site *models.GoSite)
	Delete(id string)
	Get(id string, site *models.GoSite)
	GetAll(sites *[]*models.GoSite)
}

type SiteCtrl struct {
	siteService SiteService
}

func NewSiteCtrl(siteService SiteService) *SiteCtrl {

	return &SiteCtrl{siteService: siteService}
}

func (s *SiteCtrl) GetSites(c *fiber.Ctx) error {

	var sites []*models.GoSite
	s.siteService.GetAll(&sites)
	err := c.JSON(sites)
	if err != nil {
		err = s.error(c, err)
		return err
	}

	err = c.SendStatus(fasthttp.StatusOK)
	log.Println(err)
	return nil
}

func (s *SiteCtrl) GetSite(c *fiber.Ctx) error {
	id, err := s.getID(c)
	if err != nil {
		return err
	}

	site := new(models.GoSite)
	s.siteService.Get(id, site)

	if site.SiteID == "" {
		err := errors.New("no data")
		err = s.error(c, err)
		return err
	}

	err = c.JSON(site)
	if err != nil {
		err = s.error(c, err)
		return err
	}

	err = c.SendStatus(fasthttp.StatusOK)
	log.Println(err)

	return nil
}

func (s *SiteCtrl) CreateSite(c *fiber.Ctx) error {
	// Instantiate new Product struct
	site := new(models.GoSite)
	//  Parse body into product struct
	if err := c.BodyParser(site); err != nil {
		log.Println(err)
		err = s.error(c, err)
		return err
	}

	s.siteService.Create(site)
	err := c.JSON(site)
	if err != nil {
		err = s.error(c, err)
		return err
	}

	err = c.SendStatus(fasthttp.StatusOK)
	log.Println(err)

	return nil
}

func (s *SiteCtrl) UpdateSite(c *fiber.Ctx) error {
	id, err := s.getID(c)
	if err != nil {
		return err
	}

	// Instantiate new Product struct
	site := new(models.GoSite)
	//  Parse body into product struct
	if err := c.BodyParser(site); err != nil {
		log.Println(err)
		err = s.error(c, err)
		return err
	}

	site.SiteID = id

	s.siteService.Update(site)
	err = c.JSON(site)
	if err != nil {
		err = s.error(c, err)
		return err
	}

	err = c.SendStatus(fasthttp.StatusOK)
	log.Println(err)

	return nil
}

func (s *SiteCtrl) DeleteSite(c *fiber.Ctx) error {
	id, err := s.getID(c)
	if err != nil {
		return err
	}

	s.siteService.Delete(id)

	err = c.SendStatus(fasthttp.StatusOK)
	log.Println(err)

	return nil
}

func (s *SiteCtrl) error(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(&fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}

func (s *SiteCtrl) getID(c *fiber.Ctx) (string, error) {
	id := c.Params("id")

	if len(strings.TrimSpace(id)) == 0 {
		err := errors.New("id is empty")
		err = s.error(c, err)
		return "", err
	}
	return id, nil
}
