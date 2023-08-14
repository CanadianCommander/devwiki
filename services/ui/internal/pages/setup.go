package pages

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Public
// ===============================================

func InitPages(router *gin.RouterGroup, renderer multitemplate.Renderer) {
	initViewPage(router.Group("/"), renderer)
	initLandingPage(router.Group("/welcome"), renderer)
	initEditPage(router.Group("/edit"), renderer)
}
