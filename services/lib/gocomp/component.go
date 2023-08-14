package gocomp

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Type
// ===============================================

type Component interface {
	// Name returns the name of the component
	Name() string

	// RegisterComponent registers the component
	RegisterComponent(router *gin.RouterGroup, renderer multitemplate.Renderer) error

	// Render this component to the gin context (write response)
	Render(c *gin.Context)

	// TemplateParams returns the template parameters to use when rendering this component
	TemplateParams(c *gin.Context) gin.H
}
