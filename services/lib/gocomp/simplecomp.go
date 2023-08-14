package gocomp

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// ===============================================
// Type
// ===============================================

type SimpleComponent struct {
	name      string
	templates []string
}

// ===============================================
// Public
// ===============================================

// NewSimpleComponent creates a new simple component
func NewSimpleComponent(name string, templates []string) SimpleComponent {
	return SimpleComponent{
		name:      name,
		templates: templates,
	}
}

// Name returns the name of the component
func (simpleComp *SimpleComponent) Name() string {
	return simpleComp.name
}

// RegisterComponent registers the component
func (simpleComp *SimpleComponent) RegisterComponent(router *gin.RouterGroup, renderer multitemplate.Renderer) error {
	renderer.AddFromFilesFuncs(simpleComp.name, template.DefaultTemplateFuncs(), simpleComp.templates...)
	router.GET(simpleComp.name, simpleComp.Render)
	return nil
}

// Render simply renders the template using all query parameters as variables
func (simpleComp *SimpleComponent) Render(c *gin.Context) {
	c.HTML(200, simpleComp.name, simpleComp.TemplateParams(c))
}

// TemplateParams returns the template parameters to use when rendering this component
func (simpleComp *SimpleComponent) TemplateParams(c *gin.Context) gin.H {
	queryParams := c.Request.URL.Query()
	processedParams := make(map[string]interface{})
	for key, param := range queryParams {
		if len(param) != 1 {
			slog.Warn(
				"Query parameter has multiple values. This is not supported by simple components",
				"key", key,
				"values", param)
		} else {
			processedParams[key] = param[0]
		}
	}

	return processedParams
}
