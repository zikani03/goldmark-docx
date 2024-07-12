package goldmarkdocx

import "image/color"

type listType int

const (
	ListTypeNumbered listType = iota
	ListTypeBulleted
	ListTypeNone
)

type aligment int

const (
	AlignLeft aligment = iota
	AlignRight
	AlignMiddle
	AlignJustify
)

type Style struct {
	Title     string
	Font      string
	FontSize  float32
	Spacing   float32
	FgColor   color.Color
	BgColor   color.Color
	Alignment aligment
}

type Styles struct {
	Title  *Style
	Normal *Style
	H1     *Style
	H2     *Style
	H3     *Style
	H4     *Style
	H5     *Style
	Code   *Style
	Header *Style
	Footer *Style

	LinkColor color.Color
}

func DefaultStyles() Styles {
	return Styles{
		// TODO: add the actual default styles..
	}
}
