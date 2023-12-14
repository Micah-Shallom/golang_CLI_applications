package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type executer interface{
	execute() (string, error)
}

func main(){
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(proj string, out io.Writer) error {
	if proj == ""{
		return fmt.Errorf("project directory is required: %w", ErrValidation) //named error handling technique
	}
	pipeline := make([]executer, 3)
	pipeline[0] = newExecutionStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		proj,
		[]string{"build",".","errors"},
	)
	// pipeline[1] = newExecutionStep(
	// 	"go test",
	// 	"go",
	// 	"Go Test: SUCCESS",
	// 	proj,
	// 	[]string{"test", "-v"},
	// )
	pipeline[1] = newExecutionStep(
		"go fmt",
		"gofmt",
		"Gofmt: SUCCESS",
		proj,
		[]string{"-l","."},
	)
	pipeline[2] = newTimeoutStep(
		"git push",
		"git",
		"Git Push: SUCESS",
		proj,
		[]string{"push","origin","master"},
		10 * time.Second,
	)
	
	for _, s := range pipeline{
		msg, err := s.execute()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
