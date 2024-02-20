package site

import (
	"time"

	"app/db"

	"github.com/google/uuid"
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

func (s *SiteService) GetSiteByID(id string) (*db.Site, error) {
	site := db.Site{}
	err := s.DB.Select(&site, `
    SELECT * FROM site
    OUTER LEFT JOIN page
    ON site.id = page.site_id
    WHERE id=$1 LIMIT 1
  `, id)
	if err != nil {
		return nil, err
	}

	return &site, nil
}

func (s *SiteService) CreateSite(userID string) (*db.Site, error) {
	siteID := uuid.NewString()
	_, err := s.DB.Exec(`
  INSERT INTO site (
    id = $1,
    title = $2,
    user_id = $3,
    created_at = $4,
  ) VALUES ($1, $2, $3, $4)
    `, siteID, "", userID, time.Now())
	if err != nil {
		return nil, err
	}

	_, pageErr := s.DB.Exec(`
  INSERT INTO page (
    id,
    title,
    nodes,
    dependancies,
    created_at,
  ) VALUES ($1, $2, $3, $4)
    `,
		uuid.NewString(),
		"Landing page",
		siteID,
		make([]byte, 0),
		make([]byte, 0),
		time.Now(),
	)
	if pageErr != nil {
		return nil, pageErr
	}

	site := &db.Site{}
	siteErr := db.DB.Get(site, "SELECT * FROM site OUTER LEFT JOIN page ON site.id = page.site_id")
	if siteErr != nil {
		return nil, siteErr
	}

	return site, nil
}

func (s *SiteService) PublishSite(id string) {
}

func New(db *sqlx.DB) *SiteService {
	return &SiteService{}
}
