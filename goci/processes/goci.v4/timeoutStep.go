package main

import (
	"context"
	"os/exec"
	"time"
)

type timeoutStep struct {
	step
	timeout time.Duration
}

func newTimeoutStep(name, exe, message, proj string, args []string, timeout time.Duration) timeoutStep {
	s := timeoutStep{}
	s.step = newStep(name, exe, message, proj, args)
	s.timeout = timeout
	if s.timeout == 0 {
		s.timeout = 30 * time.Second
	}
	return s
} 

func (s timeoutStep) execute() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	//command context includes a context that terminates the code execution on timeout
	cmd := exec.CommandContext(ctx, s.exe, s.args...)
	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil { 
		if ctx.Err() == context.DeadlineExceeded {
			return "", &stepErr{
				step: s.name,
				msg: "failed timeout",
				cause: context.DeadlineExceeded,
			}
		}

		return "", &stepErr{
			step: s.name,
			msg: "failed to execute",
			cause: err,
		}
	}
	return s.message, nil
}