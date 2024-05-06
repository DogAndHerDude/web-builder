package site_handlers

import (
	"net/http"

<<<<<<< HEAD:server/internal/app/site/site_handlers/site_handlers.go
	"github.com/DogAndHerDude/web-builder/internal/app/site/site_service"
=======
>>>>>>> 973db0e (chore(project): udpate imports and add build targets):server/internal/app/site/handlers.go
	"github.com/DogAndHerDude/web-builder/internal/pkg/jwt_utils"
	custom_middleware "github.com/DogAndHerDude/web-builder/middleware"

	"github.com/labstack/echo/v4"
)

type SiteHandlers struct {
	siteService site_service.ISiteService
}

func (h *SiteHandlers) CreateSite(c echo.Context) error {
	claims, ok := c.Get("user").(jwt_utils.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	site, err := h.siteService.CreateSite(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unauthorized")
	}

	c.JSON(http.StatusCreated, struct {
		data struct {
			ID string `json:"id"`
		}
	}{
		data: struct {
			ID string `json:"id"`
		}{
			ID: site.ID,
		},
	})

	return nil
}

func (h *SiteHandlers) UpdateSite(c echo.Context) error {
	payload := &site_service.UpdateSiteDetailsPayload{}
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return err
	}

	// Handle the rest of the garbage

	return nil
}

func (h *SiteHandlers) CreatePage(c echo.Context) error {
	return nil
}

func (h *SiteHandlers) UpdatePage(c echo.Context) error {
	return nil
}

func (h *SiteHandlers) UpdatePageNodes(c echo.Context) error {
	payload := &site_service.UpdatePageNodesPayload{}
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(payload); err != nil {
		return err
	}

	return nil
}

func RegisterHandlers(e *echo.Group, s site_service.ISiteService) {
	subGroup := e.Group("/site")
	h := &SiteHandlers{
		siteService: s,
	}

	subGroup.Use(custom_middleware.NewAuthorizeMiddleware())

	subGroup.POST("/", h.CreateSite)
	subGroup.PATCH("/", h.UpdateSite)
	subGroup.POST("/page", h.CreatePage)
	subGroup.PATCH("/page", h.UpdatePage)
	subGroup.PATCH("/page/nodes", h.UpdatePageNodes)
}
