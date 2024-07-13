package goldmarkdocx

import (
	"context"
	"io/fs"
	"os"

	"github.com/fumiama/go-docx"
)

type Config struct {
	Context context.Context

	Docx    *docx.Docx
	ImageFS fs.FS
	Styles  *Styles
}

func DefaultConfig() *Config {
	return &Config{
		Context: context.Background(),
		Docx:    docx.NewA4(),
		ImageFS: os.DirFS("."),
		Styles:  DefaultStyles(),
	}
}
