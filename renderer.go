package goldmarkdocx

import (
	"fmt"
	"io"
	"sync"

	"github.com/fumiama/go-docx"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

type docxRenderer struct {
	config   *Config
	nr       *nodeRendererFuncs
	initSync sync.Once
}

var _ renderer.Renderer = &docxRenderer{}

// AddOptions has no effect on this renderer
// The method is to satisfy goldmark's Renderer interface
func (r *docxRenderer) AddOptions(options ...renderer.Option) {

}

// Satisfies the NodeRendererFuncRegisterer interface
// used to add NodeRenderers
func (r *docxRenderer) Register(kind ast.NodeKind, v NodeRendererFunc) {
	// r.nodeRendererFuncsTmp[kind] = v
	// if int(kind) > r.maxKind {
	// 	r.maxKind = int(kind)
	// }
}

func (r *docxRenderer) Render(w io.Writer, source []byte, node ast.Node) error {
	r.initSync.Do(func() {

	})
	doc := docx.NewA4()
	para1 := doc.AddParagraph()
	para1.AddText(string(source)).AddTab()

	_, err := doc.WriteTo(w)
	if err != nil {
		return fmt.Errorf("failed to write document to writer %v", err)
	}

	return nil
}
