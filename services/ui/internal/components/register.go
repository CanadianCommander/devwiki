package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/gocomp"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Static
// ===============================================

var components = []gocomp.Component{
	&LeftNavComponent{},
	&FileNavComponent{},
}

// ===============================================
// Public
// ===============================================

// RegisterComponents registers all components with the specified router
// ### params
// - router: the router to register the components with
// - renderer: the renderer to register the templates with
// ### returns
// - error: an error if one occured
func RegisterComponents(router *gin.RouterGroup, renderer multitemplate.Renderer) error {

	for _, component := range components {
		err := component.RegisterComponent(router, renderer)
		if err != nil {
			return err
		}
	}

	return nil
}
