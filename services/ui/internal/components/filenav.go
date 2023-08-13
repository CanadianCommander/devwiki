package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ===============================================
// Type
// ===============================================

type FileNavComponent struct {
}

// ===============================================
// Public
// ===============================================

func (fileNav *FileNavComponent) Name() string {
	return "components/filenav"
}

func (fileNav *FileNavComponent) RegisterComponent(router *gin.RouterGroup, renderer multitemplate.Renderer) error {
	router.GET("/filenav", fileNav.Render)
	template.AddTemplate(renderer, fileNav.Name(), "components/filenav/index")
	return nil
}

func (fileNav *FileNavComponent) Render(c *gin.Context) {
	c.HTML(http.StatusOK, "components/filenav", gin.H{})
}
