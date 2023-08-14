package pages

import (
	"github.com/CanadianCommander/devwiki/services/ui/internal/components"
	"github.com/gin-gonic/gin"
	"maps"
)

// ===============================================
// Public
// ===============================================

// GetPrimaryLayoutArgs returns the base template arguments that you must include when using the primary layout
func GetPrimaryLayoutArgs(c *gin.Context) gin.H {
	args := gin.H{}

	maps.Copy(args, components.GlobalComponentRegistry.GetComponent(components.LayoutPrimaryComponentName).TemplateParams(c))
	return args
}
