package record

// ===============================================
// Types
// ===============================================

type Workspace struct {
	Name string `json:"name"`

	// root of the file tree for this workspace.
	Root *File `json:"root"`
}

// ===============================================
// Public
// ===============================================

// NewWorkspace creates a new workspace
func NewWorkspace(name string, root *File) *Workspace {
	return &Workspace{
		Name: name,
		Root: root,
	}
}
