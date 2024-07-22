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
	config               *Config
	nodeRendererFuncsTmp map[ast.NodeKind]NodeRendererFunc
	maxKind              int
	nodeRenderer         *nodeRendererFuncs
	nodeRendererFuncs    []NodeRendererFunc
	initSync             sync.Once
}

var _ renderer.Renderer = &docxRenderer{}

// AddOptions has no effect on this renderer
// The method is to satisfy goldmark's Renderer interface
func (r *docxRenderer) AddOptions(options ...renderer.Option) {

}

// Satisfies the NodeRendererFuncRegisterer interface
// used to add NodeRenderers
func (r *docxRenderer) Register(kind ast.NodeKind, v NodeRendererFunc) {
	r.nodeRendererFuncsTmp[kind] = v
	if int(kind) > r.maxKind {
		r.maxKind = int(kind)
	}
}

func (r *docxRenderer) Render(w io.Writer, source []byte, node ast.Node) error {
	r.initSync.Do(func() {
		r.nodeRenderer.RegisterFuncs(r)

		for _, nr := range r.nodeRendererFuncsTmp {
			r.nodeRendererFuncs = append(r.nodeRendererFuncs, nr)
		}
		r.nodeRendererFuncs = make([]NodeRendererFunc, r.maxKind+1)
		for kind, nr := range r.nodeRendererFuncsTmp {
			r.nodeRendererFuncs[kind] = nr
		}
	})
	doc := docx.NewA4()
	// para1 := doc.AddParagraph()
	writer := &Writer{
		Docx:   doc,
		states: newStates(),
	}
	// para1.AddText(string(source)).AddTab()
	err := ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		s := ast.WalkStatus(ast.WalkContinue)
		var err error
		f := r.nodeRendererFuncs[n.Kind()]
		if f != nil {
			s, err = f(writer, source, n, entering)
		}
		return s, err
	})

	if err != nil {
		return err
	}

	_, err = doc.WriteTo(w)
	if err != nil {
		return fmt.Errorf("failed to write document to writer %v", err)
	}

	return nil
}
