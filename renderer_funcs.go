package goldmarkdocx

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
)

type NodeRenderer interface {
	RegisterFuncs(NodeRendererFuncRegisterer)
}

type NodeRendererFuncRegisterer interface {
	Register(ast.NodeKind, NodeRendererFunc)
}

type NodeRendererFunc func(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error)

type nodeRendererFuncs struct{}

func (r *nodeRendererFuncs) RegisterFuncs(reg NodeRendererFuncRegisterer) {
	reg.Register(ast.KindDocument, r.renderDocument)
	reg.Register(ast.KindHeading, r.renderHeading)
	reg.Register(ast.KindList, r.renderList)
	reg.Register(ast.KindListItem, r.renderListItem)
	reg.Register(ast.KindText, r.renderText)
	reg.Register(ast.KindString, r.renderText)
	reg.Register(ast.KindEmphasis, r.renderEmphasis)
	reg.Register(ast.KindTextBlock, r.renderTextBlock)
	reg.Register(ast.KindLink, r.renderLink)
	reg.Register(ast.KindImage, r.renderImage)
}

func (r *nodeRendererFuncs) renderDocument(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	// Nothing to do here, continue
	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderHeading(w *Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	// n := node.(*ast.Heading)
	// segment := n.Text
	// w.Docx.AddParagraph().AddText(string(segment.Value(source))).Size("16")

	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderText(w *Writer, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Text)
	segment := n.Segment
	w.Docx.AddParagraph().AddText(string(segment.Value(source)))
	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderEmphasis(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderTextBlock(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderList(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderListItem(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderLink(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkContinue, nil
}

func (r *nodeRendererFuncs) renderImage(w *Writer, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return ast.WalkStop, fmt.Errorf("not implemented!")
}
