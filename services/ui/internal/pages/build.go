package pages

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

// ===============================================
// Public
// ===============================================

// SetupPages setups the editor page routes
func SetupPages() *gin.Engine {
	router := gin.Default()

	registerTemplateFunctions(router)
	initEditorPage(router.Group("/editor"))
	loadStaticContent(router)

	return router
}

// loadStaticContent sets up static content paths for the editor
func loadStaticContent(router *gin.Engine) {

	router.StaticFile("/favicon.ico", "./assets/img/favicon.webp")
	router.LoadHTMLGlob("./assets/html/**/*")
	router.Static("/assets", "./assets")
}

// registerTemplateFunctions for use during template rendering
func registerTemplateFunctions(router *gin.Engine) {

	router.SetFuncMap(template.FuncMap{
		"getEditorComponentScripts": getEditorComponentScripts,
		"getEditorComponentStyles":  getEditorComponentStyles,
	})
}
