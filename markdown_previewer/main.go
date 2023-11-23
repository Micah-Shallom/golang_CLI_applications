package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// const (
// 	header = `<!DOCTYPE html>
// 	<html>
// 	<head>
// 	</head>
// 	<body>
// 	`
// 	footer = `
// 	</body>
// 	</html>
// 	`
// )

const (
	defaultTemplate = `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
	<title>{{ .Title }}</title>
	</head>
	<body>
	{{ .Body }}
	</body>
	</html>
`
)

type content struct {
	Title string
	Body template.HTML
}

func main() {
	//parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()

	//if user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer, skipPreview bool) error {
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
	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}
	if skipPreview {
		return nil
	}

	defer os.Remove(outName)
	return preview(outName)
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

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	//define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "cmd.exe"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	//append filename to parameter slice
	cParams = append(cParams, fname)

	//locate executable in path
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	//open the file using default program
	err = exec.Command(cPath, cParams...).Run()

	//Give the browser some time to open the file before deleting it
	time.Sleep(2 * time.Second)
	return err
}
