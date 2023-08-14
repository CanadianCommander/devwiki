package record

// ===============================================
// Types
// ===============================================

type FileType string

const (
	FileTypeUnknown   FileType = "application/octet-stream"
	FileTypeMarkdown           = "text/markdown"
	FileTypeDirectory          = "inode/directory"
)

type File struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Parent   *File    `json:"parent"`
	Children []*File  `json:"children"`
	Type     FileType `json:"type"`
}

// ===============================================
// Public
// ===============================================

// NewFile creates a new file record
func NewFile(name string, path string, parent *File, children []*File, fileType FileType) *File {
	return &File{
		Name:     name,
		Path:     path,
		Parent:   parent,
		Children: children,
		Type:     fileType,
	}
}

// ===============================================
// Getters
// ===============================================

func (file *File) IsDirectory() bool {
	return file.Type == FileTypeDirectory
}

func (file *File) HasChildren() bool {
	return file.Children != nil && len(file.Children) > 0
}

func (file *File) HasParent() bool {
	return file.Parent != nil
}
