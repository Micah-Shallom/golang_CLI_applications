package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type executer interface{
	execute() (string, error)
}


var sig = make(chan os.Signal, 1)
var errCh = make(chan error)
var done = make(chan struct{})

func main(){
	
	//relaying signals to the sig channel
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

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
		"Git Push: SUCCESS",
		proj,
		[]string{"push","origin"},
		10 * time.Second,
	)

	go func() {
		for _, s := range pipeline{
			msg, err := s.execute()
			if err != nil {
				errCh <- err
				return
			}
			_, err = fmt.Fprintln(out, msg)
			if err != nil {
				errCh<- err
				return
			}
		}
		close(done)
	}()

	for{
		select{
		case rec :=<-sig:
			signal.Stop(sig)
			return fmt.Errorf("%s: Exiting: %w", rec, ErrSignal)
		case err := <-errCh:
			return err
		case <- done:
			return nil
		}
	}
}
