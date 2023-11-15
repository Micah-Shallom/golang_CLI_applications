package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	//scanner used to read text from a reader(such as file)
	scanner := bufio.NewScanner(r)

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}else if countBytes{
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

func main() {
	//defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count Lines")
	bytes := flag.Bool("b", true, "Count Bytes")
	flag.Parse() //parsing the flags provided by the user

	count := count(os.Stdin, *lines, *bytes)
	fmt.Println(count)
}