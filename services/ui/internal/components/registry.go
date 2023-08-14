package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/gocomp"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// ===============================================
// Static
// ===============================================

type ComponentRegistry struct {
	components []gocomp.Component
	// used to speedup subsequent component lookups
	componentIndex map[string]gocomp.Component
}

var GlobalComponentRegistry = ComponentRegistry{
	componentIndex: make(map[string]gocomp.Component),
	components: []gocomp.Component{
		&LayoutPrimaryComponent{},
		&LeftNavComponent{},
		&FileNavComponent{},
	},
}

// ===============================================
// Public
// ===============================================

// RegisterComponents registers all components with the specified router
// ### params
// - router: the router to register the components with
// - renderer: the renderer to register the templates with
// ### returns
// - error: an error if one occurred
func (registry *ComponentRegistry) RegisterComponents(router *gin.RouterGroup, renderer multitemplate.Renderer) error {

	for _, component := range registry.components {
		err := component.RegisterComponent(router, renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetComponent returns the component with the specified name from the registry
func (registry *ComponentRegistry) GetComponent(name string) gocomp.Component {
	if registry.componentIndex[name] != nil {
		return registry.componentIndex[name]
	} else {
		for _, component := range registry.components {
			if component.Name() == name {
				registry.componentIndex[name] = component
				return component
			}
		}
	}

	return nil
}
