package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

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
	args := []string{"build",".","errors"}
	cmd := exec.Command("go",args...)
	cmd.Dir = proj
	if err := cmd.Run(); err != nil {
		// return fmt.Errorf("'go build' failed: %s", err)
		return &stepErr{step: "go build", msg: "go build failed", cause: err} //custom error handling technique
	}
	_, err := fmt.Fprintln(out, "Go Build: SUCCESS")
	return err
}
