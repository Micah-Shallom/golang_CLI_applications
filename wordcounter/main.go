package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)


func main() {
	//defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count Lines")
	bytes := flag.Bool("b", false, "Count Bytes")
	filename := flag.String("file", "", "File containing text")
	flag.Parse() //parsing the flags provided by the user
	text, err := getContent(*filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	value := count(text, *lines, *bytes)
	fmt.Println(value)
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
	}else if countBytes{
		scanner.Split(bufio.ScanBytes)
	}

	result := 0
	for scanner.Scan() {
		result++
	}
	return result
}
