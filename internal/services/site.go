package services

import (
	"strings"

	"gorm.io/gorm/clause"

	"github.com/ehabterra/go-site/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoSiteService struct {
	db *gorm.DB
}

func NewGoSiteService(db *gorm.DB) *GoSiteService {

	return &GoSiteService{db}
}

func (s *GoSiteService) Create(site *models.GoSite) {

	site.SiteID = strings.ReplaceAll(uuid.New().String(), "-", "")

	// Create
	s.db.Create(site)
}

func (s *GoSiteService) Update(site *models.GoSite) {
	// Update
	s.db.Save(site)
}

func (s *GoSiteService) Delete(id string) {
	// Delete
	s.db.Transaction(func(tx *gorm.DB) error {
		tx.Select(clause.Associations).Delete(&models.GoSite{SiteID: id})

		return nil // commit
	})

}

func (s *GoSiteService) Get(id string, site *models.GoSite) {
	// Get first
	s.db.Preload(clause.Associations).Where(map[string]interface{}{"site_id": id}).First(site)
}

func (s *GoSiteService) GetAll(sites *[]*models.GoSite) {
	// Get all
	s.db.Preload(clause.Associations).Find(sites)
}
