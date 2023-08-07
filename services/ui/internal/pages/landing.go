package pages

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ===============================================
// Public
// ===============================================

func initLandingPage(router *gin.RouterGroup, renderer multitemplate.Renderer) {

	router.GET("/", welcomeIndex)
	AddTemplate(renderer, "welcome/index", "layout/primary", "page/welcome/index")

}

// ===============================================
// Endpoints
// ===============================================

// GET /welcome
func welcomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome/index", gin.H{})
}
