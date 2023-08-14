package components

import (
	"github.com/CanadianCommander/devwiki/services/lib/template"
	"github.com/CanadianCommander/devwiki/services/wiki/lib/record"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ===============================================
// Type
// ===============================================

const FileNavComponentName = "components/filenav"

type FileNavComponent struct {
}

// ===============================================
// Public
// ===============================================

func (fileNav *FileNavComponent) Name() string {
	return FileNavComponentName
}

func (fileNav *FileNavComponent) RegisterComponent(router *gin.RouterGroup, renderer multitemplate.Renderer) error {
	router.GET("/filenav", fileNav.Render)
	template.AddTemplate(renderer, FileNavComponentName, "components/filenav/index")
	return nil
}

func (fileNav *FileNavComponent) Render(c *gin.Context) {
	c.HTML(http.StatusOK, FileNavComponentName, fileNav.TemplateParams(c))
}

// TemplateParams returns the template parameters to use when rendering this component
func (fileNav *FileNavComponent) TemplateParams(c *gin.Context) gin.H {

	// TMP setup dummy data
	root := record.NewFile("root", "/", nil, []*record.File{
		record.NewFile("foo", "/foo", nil, nil, record.FileTypeMarkdown),
		record.NewFile("howto", "/howto", nil, nil, record.FileTypeMarkdown),
		record.NewFile("notes", "/notes", nil, []*record.File{
			record.NewFile("bar", "/howto/bar", nil, nil, record.FileTypeMarkdown),
			record.NewFile("bang", "/howto/bar", nil, nil, record.FileTypeMarkdown),
		}, record.FileTypeDirectory),
	}, record.FileTypeDirectory)

	return gin.H{
		"editWorkspace": c.Query("editWorkspace") == "true",
		"workspace":     record.NewWorkspace("My First Workspace", root),
	}
}
