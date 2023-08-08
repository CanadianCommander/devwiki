package template

import (
	"fmt"
	"os"
)

// ===============================================
// Private
// ===============================================

// getEditorComponentHash returns the hash of the Angular editor component.
// this hash should be appended to the editor component script name.
func getEditorComponentHash() string {
	return os.Getenv("EDITOR_COMPONENT_HASH")
}

// getEditorComponentScripts returns a list of scripts that need to be loaded in order to
// use the editor component
func getEditorComponentScripts() []string {
	return []string{
		fmt.Sprintf("/component/editor/runtime%s.js", getEditorComponentHash()),
		fmt.Sprintf("/component/editor/polyfills%s.js", getEditorComponentHash()),
		fmt.Sprintf("/component/editor/vendor%s.js", getEditorComponentHash()),
		fmt.Sprintf("/component/editor/main%s.js", getEditorComponentHash()),
	}
}

// getEditorComponentStyles returns a list of styles that need to be loaded in order to
// use the editor component
func getEditorComponentStyles() []string {
	return []string{
		fmt.Sprintf("/component/editor/styles%s.css", getEditorComponentHash()),
	}
}
