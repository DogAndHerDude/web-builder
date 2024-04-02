package site_service

import (
	"time"

	"github.com/DogAndHerDude/web-builder/builder"
	"github.com/DogAndHerDude/web-builder/db"
	"github.com/DogAndHerDude/web-builder/publisher"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UpdateSiteDetailsPayload struct {
	ID          string `json:"id" validate:"required,uuid"`
	Title       string `json:"title" validate:"omitnil,min=2,max=40"`
	Description string `json:"description" validate:"omitnil,min=10,max=40"`
}

type CreatePagePayload struct {
	SiteID string `json:"siteId" validate:"required,uuid"`
}

type UpdatePageDetailsPayload struct {
	Title string `json:"title" validate:"omitnil,min=2,max=40"`
}

// TODO: Font families should be saved somehwere in utils or something
type UpdateSiteTemplatePayload struct {
	Pallete    []string `json:"pallete" validate:"omitnil,len=3,dive,hexcolor"` // Need to validate is as a tuple of #HEX
	FontFamily string   `json:"FontFmaily" validate:"omitnil"`                  // Needs to validate based on a given set of a vailable font families
}

type HTMLNodePayload struct {
	Tag              string             `json:"tag" validate:"required,html"`
	TextContent      string             `json:"textContent" validate:"omitnil,omitempty,html_encoded"`
	Dependency       string             `json:"dependency" validate:"omitnil,omitempty"` // Need to validate from all possible references to valid urls to dependencies
	Attributes       map[string]string  `json:"attributes" validate:"omitnil,omitempty"`
	ComponentID      string             `json:"componentId" validate:"omitnil,omitempty"`
	ComponentVersion string             `json:"componentVersion" validate:"omitnil,omitempty"`
	ClassList        []string           `json:"classList" validate:"required,omitnil,dive,html_encoded"`
	Children         []*HTMLNodePayload `json:"children" validate:"required"`
}

type UpdatePageNodesPayload struct {
	SiteID string             `json:"siteId" validate:"required,uuid"`
	PageID string             `json:"pageId" validate:"required,uuid"`
	Nodes  []*HTMLNodePayload `json:"nodes" validate:"required"` // Need to nest validate
}

type ISiteService interface {
	GetSiteByID(ID string) (*db.Site, error)
	CreateSite(userID string) (*db.Site, error)
	UpdateSite(payload UpdateSiteDetailsPayload) ([]string, error)
	// Probably move it into its own module? What's the difference between SiteTemplate and Template? Should they be unrelated?
	// UpdateSiteTemplate(siteID string) ([]string, error)
	CreatePage(siteID string) error
	UpdatePage(siteID string, pageID string) error
}

type SiteService struct {
	DB        *sqlx.DB
	builder   builder.SiteBuilder
	publisher publisher.Publisher
}

// ======= SITE START =======

func (s *SiteService) GetSiteByID(ID string) (*db.Site, error) {
	site := db.Site{}
	err := s.DB.Select(&site, `
    SELECT * FROM site
    OUTER LEFT JOIN page
    ON site.id = page.site_id
    WHERE id=$1 LIMIT 1
  `, ID)
	if err != nil {
		return nil, err
	}

	return &site, nil
}

func (s *SiteService) CreateSite(userID string) (*db.Site, error) {
	siteID := uuid.NewString()
	_, err := s.DB.Exec(`
  INSERT INTO "public"."Site" (
    id = $1,
    title = $2,
    user_id = $3,
    created_at = $4,
  ) VALUES ($1, $2, $3, $4)
    `, siteID, "Untitled Site", userID, time.Now())
	if err != nil {
		return nil, err
	}

	_, pageErr := s.DB.Exec(`
  INSERT INTO "public"."Page" (
    id = $1,
    title = $2,
    nodes = $3,
    site_id = $4
    created_at = $5,
  ) VALUES ($1, $2, $3, $4, $5)
    `,
		uuid.NewString(),
		"So",
		make([]byte, 0),
		siteID,
		time.Now(),
	)
	if pageErr != nil {
		return nil, pageErr
	}

	site := &db.Site{}
	siteErr := s.DB.Get(site, `
    SELECT *
    FROM "public"."Site"
    OUTER LEFT JOIN page ON "Site".id = "Page".site_id
    WHERE id="$1"`,
		siteID,
	)
	if siteErr != nil {
		return nil, siteErr
	}

	return site, nil
}

func (s *SiteService) UpdateSite(payload UpdateSiteDetailsPayload) ([]string, error) {
	// Itterate through fields and create an update query

	return []string{}, nil
}

func (s *SiteService) CreatePage(siteID string) error {
	return nil
}

func (s *SiteService) UpdatePage(siteID string, pageID string) error {
	return nil
}

func New(db *sqlx.DB, builder builder.SiteBuilder, publisher publisher.Publisher) *SiteService {
	return &SiteService{
		DB:        db,
		builder:   builder,
		publisher: publisher,
	}
}
