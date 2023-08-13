package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Types
// ===============================================

type LeftNavComponent struct {
}

// ===============================================
// Public
// ===============================================

func (leftNav *LeftNavComponent) Name() string {
	return "components/leftnav"
}

func (leftNav *LeftNavComponent) RegisterComponent(router *gin.RouterGroup, renderer multitemplate.Renderer) error {
	router.GET("/leftnav", leftNav.Render)
	template.AddTemplate(renderer, "components/leftnav", "components/leftnav/leftnav")
	return nil
}

func (leftNav *LeftNavComponent) Render(c *gin.Context) {
	c.HTML(200, "components/leftnav", gin.H{
		"Show": c.Query("show") == "true",
		"Filenav": gin.H{
			"Overlay": c.Query("show") == "true",
		},
	})
}
