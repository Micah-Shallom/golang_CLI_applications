package main

import (
	"bytes"
	"errors"
	"os/exec"
	"syscall"
	"testing"
)

func TestRun(t *testing.T) {
	_, err := exec.LookPath("git")
	if err != nil {
		t.Skip("git not installed. Skipping test.")
	}
	var testCases = []struct {
		name     string
		proj     string
		out      string
		expErr   error
		setupGit bool
	}{
		{name: "success", proj: "./testdata/tool/", out: "Go Build: SUCCESS\nGofmt: SUCCESS\nGit Push: SUCCESS\n", expErr: nil, setupGit: true},
		{name: "fail", proj: "./testdata/toolErr", out: "", expErr: &stepErr{step: "go build"}, setupGit: false},
		{name: "failFormat", proj: "./testdata/toolFmtErr/", out: "", expErr: &stepErr{step: "go fmt"}, setupGit: false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setupGit {
				setupGit(t, tc.proj)
			}

			var out bytes.Buffer
			err := run(tc.proj, &out)
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error: %q. Got %q", tc.expErr, err)
					return
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error: %q. Got %q.\n", tc.expErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %q\n", err)
			}
			if out.String() != tc.out {
				t.Errorf("Expected output: %q. Got %q\n", tc.out, out.String())
			}

		})
	}
}

func setupGit(t *testing.T, proj string) {
	t.Helper()

	//use the LookPath() function to enquire if git is an installed package
	gitExec, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}

	var gitCMDList = []struct {
		name string
		args []string
		dir  string
		env  []string
	}{
		{"staging", []string{"add"}, "./testdata/tool", nil},
		{"commit", []string{"commit", "-m", "update: testCommit"}, "./testdata/tool", nil},
	}

	for _, g := range gitCMDList {
		g.args = append(g.args, g.dir)
		gitCmd := exec.Command(gitExec, g.args...)
		if err := gitCmd.Run(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestRunKill(t *testing.T) {
	var testCases = []struct {
		name   string
		proj   string
		sig    syscall.Signal
		expErr error
	}{
		{"SIGINT", "./testdata/tool", syscall.SIGINT, ErrSignal},
		{"SIGTERM", "./testdata/tool", syscall.SIGTERM, ErrSignal},
		{"SIGQUIT", "./testdata/tool", syscall.SIGQUIT, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			
		})
	}
}
