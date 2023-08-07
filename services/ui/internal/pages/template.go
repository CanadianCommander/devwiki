package pages

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/rs/zerolog/log"
	"html/template"
	"os"
	"path"
	"strings"
)

// ===============================================
// Types
// ===============================================

const (
	TemplateBasePath = "assets"
)

// ===============================================
// Public
// ===============================================

// AddTemplate adds a template to the renderer
// ### params
// - renderer: the renderer to add the template to
// - name: the name of the new template
// - templateFiles: the files that comprise the template
func AddTemplate(renderer multitemplate.Renderer, name string, templateFiles ...string) {
	renderer.AddFromFilesFuncs(name, defaultTemplateFuncs(), append(tmplPathSlice(templateFiles), getDefaultComponentTmpls()...)...)
}

// ===============================================
// private
// ===============================================

// DefaultComponentTemplates holds a list of default component templates.
// Used to cache the default component templates.
var defaultComponentTemplates []string = nil

// tmplPath resolves the path of the template with given name
func tmplPath(name string) string {
	return fmt.Sprintf("%s/%s.gohtml", TemplateBasePath, strings.ReplaceAll(name, ".gohtml", ""))
}

// tmplPathSlice resolves the paths of the templates with given names
func tmplPathSlice(templates []string) []string {
	var templatePaths []string
	for _, template := range templates {
		templatePaths = append(templatePaths, tmplPath(template))
	}
	return templatePaths
}

// discoverTemplates discovers templates in the given directory
func discoverTemplates(directory string) ([]string, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var templates []string
	for _, entry := range entries {
		if entry.IsDir() {
			subDirectoryTemplates, err := discoverTemplates(path.Join(directory, entry.Name()))
			if err != nil {
				return nil, err
			}

			templates = append(templates, subDirectoryTemplates...)
		} else if strings.Contains(entry.Name(), ".gohtml") {
			templates = append(templates, path.Join(directory, entry.Name()))
		}
	}

	return templates, nil
}

// getDefaultComponentTmpls gets the default component templates
// these are added to every page for ease of use
func getDefaultComponentTmpls() []string {

	if defaultComponentTemplates == nil {
		components, err := discoverTemplates(path.Join(TemplateBasePath, "components"))
		if err != nil {
			log.Err(err).Msg("Error during component template discovery")
			panic(err)
		}
		defaultComponentTemplates = components
	}

	return defaultComponentTemplates
}

// defaultTemplateFuncs gets the default template functions
func defaultTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"getEditorComponentScripts": getEditorComponentScripts,
		"getEditorComponentStyles":  getEditorComponentStyles,
	}
}
