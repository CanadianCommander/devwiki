package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ===============================================
// Public
// ===============================================

func InitFileNav(router *gin.RouterGroup, renderer multitemplate.Renderer) {
	router.GET("/filenav", filenav)
	template.AddTemplate(renderer, "components/filenav", "components/filenav/index")
}

// ===============================================
// Endpoints
// ===============================================

// GET components/filenav
func filenav(c *gin.Context) {
	c.HTML(http.StatusOK, "components/filenav", gin.H{})
}
