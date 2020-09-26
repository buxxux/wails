package runtime

import (
	"strings"

	"github.com/wailsapp/wails/lib/interfaces"
)

// Dialog exposes an interface to native dialogs
type Dialog struct {
	renderer interfaces.Renderer
}

// NewDialog creates a new Dialog struct
func NewDialog(renderer interfaces.Renderer) *Dialog {
	return &Dialog{
		renderer: renderer,
	}
}

// SelectFile prompts the user to select a file
func (r *Dialog) SelectFile(params ...string) string {
	title := "Select File"
	filter := ""
	if len(params) > 0 {
		title = params[0]
	}
	if len(params) > 1 {
		filter = strings.Replace(params[1], " ", "", -1)
	}
	return r.renderer.SelectFile(title, filter)
}

// OpenDialog prompts the user to select a directory
func (r *Dialog) OpenDialog(params ...string) []string {
	title := "Select File"
	filter := ""
	if len(params) > 0 {
		title = params[0]
	}
	if len(params) > 1 {
		filter = strings.Replace(params[1], " ", "", -1)
	}
	return r.renderer.OpenDialog(title, filter)
}
