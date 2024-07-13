package goldmarkdocx

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

// New returns a Docx renderer
func New() renderer.Renderer {

	nrs := make(map[ast.NodeKind]NodeRendererFunc)
	return &docxRenderer{
		config:               DefaultConfig(),
		nodeRenderer:         &nodeRendererFuncs{},
		nodeRendererFuncs:    nil,
		nodeRendererFuncsTmp: nrs,
	}
}
