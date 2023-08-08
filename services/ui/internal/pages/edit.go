package pages

import (
	"github.com/CanadianCommander/devwiki/services/ui/internal/template"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ===============================================
// Public
// ===============================================

// initEditorPage setups the editor page routes
func initEditorPage(router *gin.RouterGroup, renderer multitemplate.Renderer) {

	router.GET("/", editIndex)
	template.AddTemplate(renderer, "edit/index", "layout/primary", "page/edit/index")

	router.POST("/btn", btn)

}

// ===============================================
// Endpoints
// ===============================================

// GET /edit
func editIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "edit/index", gin.H{
		"LoadEditor": true,
	})
}

// POST /edit/btn
func btn(c *gin.Context) {
	c.String(http.StatusOK, "<wiki-editor></wiki-editor>")
}
