package goldmarkdocx

import (
	"github.com/fumiama/go-docx"
)

type Writer struct {
	states *states
	Docx   *docx.Docx
}
