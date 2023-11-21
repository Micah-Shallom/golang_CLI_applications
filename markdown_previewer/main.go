package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
	<head>
	</head>
	<body>
	`
	footer = `
	</body>
	</html>
	`
)

func main() {
	//parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	//if user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer) error {
	//Read all the data from the input file and check for errors
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input) //responsible for converting markdown to html
	// outName := fmt.Sprintf("%s.html", filepath.Base(filename))

	//Create temporary file and check for errors
	temp, err := os.CreateTemp("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprintln(out, outName)
	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) []byte {
	//Parse the markdown file through blackfriday and bluemonday to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	//Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	//Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outFName string, data []byte) error {
	//write the bytes to the file
	return os.WriteFile(outFName, data, 0644)
}
