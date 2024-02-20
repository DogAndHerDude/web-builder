package site

import "github.com/labstack/echo/v4"

type SiteHandlers struct {
	siteService ISiteService
}

func (h *SiteHandlers) CreateSite(c echo.Context) error {
	return nil
}

func RegisterHandlers(e *echo.Group, s ISiteService) {
	subGroup := e.Group("/site")
	h := &SiteHandlers{
		siteService: s,
	}

	subGroup.POST("/", h.CreateSite)
}
