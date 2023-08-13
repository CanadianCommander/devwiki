package template

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"html/template"
	"io"
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
	templates := append(tmplPathSlice(templateFiles), getDefaultComponentTmpls()...)
	renderer.AddFromFilesFuncs(name, defaultTemplateFuncs(), templates...)
}

// AddTemplateDef adds a template definition to the renderer
// The key difference between this and AddTemplate is that this method lets you specify a template definition to render
// ### params
// - renderer: the renderer to add the template to
// - name: the name of the new template
// - defName: the name of the template definition to render
// - templateFiles: the files that comprise the template
func AddTemplateDef(renderer multitemplate.Renderer, name string, defName string, templateFiles ...string) {

	templates := make([]string, 0, len(templateFiles))
	templates = append(templates, fmt.Sprintf("{{template \"%s\"}}", defName))
	for _, file := range templateFiles {
		templateData, err := readTemplate(tmplPath(file))
		if err != nil {
			log.Fatal().Msgf("Could not load template file: %s, during startup", file)
			panic(err)
		}
		templates = append(templates, templateData)
	}

	renderer.AddFromStringsFuncs(
		name,
		defaultTemplateFuncs(),
		templates...)
}

// ===============================================
// private
// ===============================================

// readTemplate loads the contents of a template from disk.
func readTemplate(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Error().Msgf("could not open template file: %s", path)
		return "", err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Error().Msgf("could not read template file: %s", path)
		return "", err
	}

	return string(bytes), nil
}

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
		"GetEditorComponentScripts": getEditorComponentScripts,
		"GetEditorComponentStyles":  getEditorComponentStyles,
		"GetThemeCss":               GetCurrentThemeCss,
		"GetImageInTheme":           GetImageInTheme,
		"NewUUID":                   uuid.NewString,
	}
}
