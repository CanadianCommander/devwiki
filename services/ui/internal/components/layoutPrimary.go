package components

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ===============================================
// Types
// ===============================================

const LayoutPrimaryComponentName = "layout/primary"

type LayoutPrimaryComponent struct {
}

// ===============================================
// Public
// ===============================================

func (primary *LayoutPrimaryComponent) Name() string {
	return LayoutPrimaryComponentName
}

func (primary *LayoutPrimaryComponent) RegisterComponent(router *gin.RouterGroup, renderer multitemplate.Renderer) error {
	// Primary layout is only used by other components. It doesn't have any endpoints of its own.
	return nil
}

func (primary *LayoutPrimaryComponent) Render(c *gin.Context) {
	c.HTML(http.StatusOK, primary.Name(), primary.TemplateParams(c))
}

func (primary *LayoutPrimaryComponent) TemplateParams(c *gin.Context) gin.H {
	return gin.H{
		"leftNav": GlobalComponentRegistry.GetComponent(LeftNavComponentName).TemplateParams(c),
	}
}
