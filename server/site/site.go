package site

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type SiteHandler interface {
	GetSiteByID(id string)

	CreateSite()

	PublishSite(id string)
}

type SiteService struct {
	DB *sqlx.DB
}

func (s *SiteService) GetSiteByID(id string) {
}

func (s *SiteService) CreateSite() {
	log.Println("Not implemented!")
}

func (s *SiteService) PublishSite(id string) {
}

func New(db *sqlx.DB) *SiteService {
	return &SiteService{}
}
