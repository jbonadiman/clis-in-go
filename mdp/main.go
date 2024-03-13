package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

const (
	header = `<!DOCTYPE html>
<html>

<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Markdown Preview Tool</title>
</head>

<body>
`

	footer = `</body>

</html>
`
)

func main() {
	filename := flag.String("file", "", "markdown file to preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outName)

	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) []byte {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(input, &buf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	body := bluemonday.UGCPolicy().SanitizeBytes(buf.Bytes())

	buf.Reset()
	buf.WriteString(header)
	buf.Write(body)
	buf.WriteString(footer)

	return buf.Bytes()
}

func saveHTML(outFname string, data []byte) error {
	return os.WriteFile(outFname, data, 0644)
}
