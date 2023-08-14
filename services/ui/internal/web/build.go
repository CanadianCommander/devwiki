package web

import (
	"github.com/CanadianCommander/devwiki/services/lib/gocomp"
	"github.com/CanadianCommander/devwiki/services/ui/internal/components"
	"github.com/CanadianCommander/devwiki/services/ui/internal/pages"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// ===============================================
// Public
// ===============================================

// SetupPages setups the editor page routes
func SetupPages() *gin.Engine {
	router := gin.Default()
	renderer := multitemplate.NewRenderer()

	// Pages
	pages.InitPages(router.Group("/"), renderer)

	// Components
	componentGroup := router.Group("/components")

	_, err := gocomp.AutoLoadSimpleComponents(componentGroup.Group("/auto"), renderer, "./assets/components/auto/")
	if err != nil {
		slog.Error("Unexpected error loading auto components. Bootup failed!", err)
		panic(err)
	}

	err = components.GlobalComponentRegistry.RegisterComponents(componentGroup, renderer)
	if err != nil {
		slog.Error("Unexpected error registering components. Bootup failed!", err)
		panic(err)
	}

	loadStaticContent(router)

	router.HTMLRender = renderer
	return router
}

// ===============================================
// Private
// ===============================================

// loadStaticContent sets up static content paths for the editor
func loadStaticContent(router *gin.Engine) {

	router.StaticFile("/favicon.ico", "./assets/img/logo-dark.svg")
	router.Static("/assets", "./assets")
}
