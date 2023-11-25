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
	"github.com/joho/godotenv"
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

var defaultTemplate string = `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
	<title>{{ .Title }}</title>
	</head>
	<body>
	<h2>{{ .FileName}}</h2>
	{{ .Body }}
	</body>
	</html>
`


type content struct {
	Title string
	FileName string
	Body template.HTML
}

func main() {
	//get default template from dotenv file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file", err)
		return
	}

	//parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()

	//if user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, tFname string ,out io.Writer, skipPreview bool) error {
	//Read all the data from the input file and check for errors
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(input, tFname, filename) //responsible for converting markdown to html
	if err != nil {
		return err
	}
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

func parseContent(input []byte, tFname string, filename string) ([]byte, error) {

	if defaultTemplate == ""{
		defaultTemplate = os.Getenv("defaultTemplate")
	}

	//Parse the markdown file through blackfriday and bluemonday to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	//parse the content of the defaultTemplate const into a new template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	//if user provided alternate template file, replace template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
	}

	//instantiate the content type, adding the title and the body
	c := content{
		Title: "Markdown Preview Tool",
		FileName: filename,
		Body: template.HTML(body),
	}
	//Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// //Write html to bytes buffer
	// buffer.WriteString(header)
	// buffer.Write(body)
	// buffer.WriteString(footer)

	//execute the template with the content type
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
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
