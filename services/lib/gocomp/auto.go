package gocomp

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
	"path"
	"strings"
)

// ===============================================
// Public
// ===============================================

// AutoLoadSimpleComponents automatically loads all components as simple components in the specified directory
// and registers them with the router
// ### params
// - router: the router to register the components with
// - renderer: the renderer to register the templates with
// - directory: the directory to load the components from
// ### returns
// - error: an error if one occured
// - []SimpleComponent: the list of simple components that were loaded
func AutoLoadSimpleComponents(
	router *gin.RouterGroup,
	renderer multitemplate.Renderer,
	directory string) ([]SimpleComponent, error) {
	return AutoLoadSimpleComponentsUsingResolver(
		router,
		renderer,
		directory,
		discoverComponentsUsingDefaultStrategy)
}

// AutoLoadSimpleComponentsUsingResolver automatically loads all components as simple components in the specified directory
// and registers them with the router
// ### params
// - router: the router to register the components with
// - renderer: the renderer to register the templates with
// - directory: the directory to load the components from
// - resolver: a function that takes a directory and returns a list of components to register
// ### returns
// - error: an error if one occured
// - []SimpleComponent: the list of simple components that were loaded
func AutoLoadSimpleComponentsUsingResolver(
	router *gin.RouterGroup,
	renderer multitemplate.Renderer,
	directory string,
	resolver func(directory string) (map[string][]string, error)) ([]SimpleComponent, error) {

	components, err := resolver(directory)
	if err != nil {
		slog.Error("Unexpected error resolving components", err)
		return nil, err
	}

	var simpleComponents []SimpleComponent
	for name, templates := range components {
		componentName := path.Base(name)
		newComponent := NewSimpleComponent(componentName, templates)
		simpleComponents = append(simpleComponents, newComponent)

		// register
		err = newComponent.RegisterComponent(router, renderer)
		if err != nil {
			slog.Error("Unexpected error registering component", err)
			return nil, err
		}
	}

	return simpleComponents, nil
}

// ===============================================
// Private
// ===============================================

func discoverComponentsUsingDefaultStrategy(directory string) (map[string][]string, error) {
	templates, err := getAllTemplates(directory)
	if err != nil {
		return nil, err
	}

	var components = make(map[string][]string)
	for _, template := range templates {
		if path.Base(template) == "index.gohtml" {
			components[path.Dir(template)] = templates
		}
	}

	return components, nil
}

func getAllTemplates(directory string) ([]string, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var components []string
	for _, entry := range entries {
		if entry.IsDir() {
			subDirectoryTemplates, err := getAllTemplates(path.Join(directory, entry.Name()))
			if err != nil {
				return nil, err
			}

			components = append(components, subDirectoryTemplates...)
		} else if strings.Contains(entry.Name(), ".gohtml") {
			components = append(components, path.Join(directory, entry.Name()))
		}
	}

	return components, nil
}
