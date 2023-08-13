package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Public
// ===============================================

// InitLeftNav initializes the leftnav component
func InitLeftNav(router *gin.RouterGroup, renderer multitemplate.Renderer) {
	router.GET("/leftnav", leftnav)
	template.AddTemplate(renderer, "components/leftnav", "components/leftnav/leftnav")
}

// ===============================================
// Endpoints
// ===============================================

// GET components/leftnav
func leftnav(c *gin.Context) {
	c.HTML(200, "components/leftnav", gin.H{
		"Show": c.Query("show") == "true",
		"Filenav": gin.H{
			"Overlay": c.Query("show") == "true",
		},
	})
}
