package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type executionStep struct {
	step
}

func newExecutionStep(name, exe, message, proj string, args []string) executionStep {
	s := executionStep{}
	s.step = newStep(name, exe, message, proj, args)
	return s
}

func (s executionStep) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj
	var out bytes.Buffer
	cmd.Stdout = &out
	
	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step: s.name,
			msg: "failed to execute",
			cause: err,
		}
	}

	if out.Len() > 0 {
		return "", &stepErr{
			step: s.name,
			msg: fmt.Sprintf("invalid format: %s", out.String()),
			cause: nil,
		}
	}
	return s.message, nil
}