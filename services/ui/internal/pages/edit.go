package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ===============================================
// Public
// ===============================================

// initEditorPage setups the editor page routes
func initEditorPage(router *gin.RouterGroup) {

	router.GET("/", editIndex)
	router.POST("/btn", btn)
}

// ===============================================
// Endpoints
// ===============================================

// GET /
func editIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"Title": "Hello World",
	})
}

// POST /btn
func btn(c *gin.Context) {
	c.String(http.StatusOK, "<wiki-editor></wiki-editor>")
}
