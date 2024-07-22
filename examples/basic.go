package main

import (
	"os"

	"github.com/yuin/goldmark"
	goldmarkdocx "github.com/zikani03/goldmark-docx"
)

func main() {

	md2docx := goldmark.New(
		goldmark.WithRenderer(
			goldmarkdocx.New(),
		),
	)

	source, err := os.ReadFile("../README.md")
	if err != nil {
		panic(err)
	}
	dest, err := os.Create("./test.docx")
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	err = md2docx.Convert(source, dest)
	if err != nil {
		panic(err)
	}
}
