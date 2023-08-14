package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Types
// ===============================================

const LeftNavComponentName = "components/leftnav"

type LeftNavComponent struct {
}

// ===============================================
// Public
// ===============================================

func (leftNav *LeftNavComponent) Name() string {
	return LeftNavComponentName
}

func (leftNav *LeftNavComponent) RegisterComponent(router *gin.RouterGroup, renderer multitemplate.Renderer) error {
	router.GET("/leftnav", leftNav.Render)
	template.AddTemplate(renderer, LeftNavComponentName, "components/leftnav/leftnav")
	return nil
}

func (leftNav *LeftNavComponent) Render(c *gin.Context) {
	c.HTML(200, LeftNavComponentName, leftNav.TemplateParams(c))
}

func (leftNav *LeftNavComponent) TemplateParams(c *gin.Context) gin.H {
	return gin.H{
		"Show":    c.Query("show") == "true",
		"fileNav": GlobalComponentRegistry.GetComponent(FileNavComponentName).TemplateParams(c),
	}
}
