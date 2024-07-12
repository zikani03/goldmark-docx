package goldmarkdocx

import (
	"github.com/yuin/goldmark/renderer"
)

// New returns a Docx renderer
func New() renderer.Renderer {
	return &docxRenderer{
		config: DefaultConfig(),
		nr:     &nodeRendererFuncs{},
	}
}
