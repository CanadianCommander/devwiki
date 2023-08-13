package pages

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Public
// ===============================================

// initViewPage setups the view page routes
func initViewPage(router *gin.RouterGroup, renderer multitemplate.Renderer) {
	router.GET("/", viewIndex)
	template.AddTemplate(renderer, "view/index", "layout/primary", "page/view/index")
}

// ===============================================
// Endpoints
// ===============================================

// GET /view
func viewIndex(c *gin.Context) {
	c.HTML(200, "view/index", gin.H{})
}
