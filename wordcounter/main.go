package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

//gives information as to what type of result is displayed
//result could be lines, words or bytes
var operationINFO string = "Lines"

func readFile(file string, lines bool, bytes bool) {
	text, err := getContent(file)
	if err != nil {
		fmt.Println("Error reading file %s: %v\n",file, err)
	}
	value := count(text, lines, bytes)
	fmt.Printf("Result for file %s: %s %d\n", file, operationINFO ,value)
}

func main() {
	//defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count Lines")
	bytes := flag.Bool("b", false, "Count Bytes")
	filename := flag.String("file", "", "File containing text")
	flag.Parse() //parsing the flags provided by the user

	//get the filenames from the cli
	filenames := flag.Args()

	//if filename or filenames are not provided then let switch our focus to capturing input from standard input
	if *filename == "" && len(filenames) == 0 {
		fmt.Println("No file or files provided, using standard input.............")
		readFile(*filename, *lines, *bytes)
	}

	if len(filenames) >= 1 {
		for _, file := range filenames{
			fmt.Printf("Processing file: %s\n", file)
			readFile(file, *lines, *bytes)
		}
		return
	}

	if *filename != ""{
		readFile(*filename, *lines, *bytes)
	}
	
}

// • Go back to the wordcounter project, Your First Command-Line Program
// in Go, on page 1, and update the wc tool to read data from files in
// addition to STDIN.
// • Update the wc tool to process multiple files.

func getContent(filename string) (io.Reader, error){

	if filename != "" {
		//open the file
		content, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		//return the file
		return bytes.NewReader(content), nil
	}else{
		//no file provided, make use of os.Stdin
		return os.Stdin, nil
	}
}

func count(r io.Reader, countLines bool, countBytes bool ) int {
	//scanner used to read text from a reader(such as file)
	scanner := bufio.NewScanner(r)

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
		operationINFO = "Words"
	}else if countBytes{
		scanner.Split(bufio.ScanBytes)
		operationINFO = "Bytes"
	}
	
	//count lines if we arent counting words and bytes
	result := 0
	for scanner.Scan() {
		result++
	}
	return result
}
