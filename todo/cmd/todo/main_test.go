package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	todo "github.com/Micah-Shallom/modules"
)

var (
	binName  = "todo"
	fileName = "todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}
	fmt.Println("Running tests....")
	result := m.Run() //runs all the test in a package
	fmt.Println("Cleaning up....")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task number 1"
	dir, err := os.Getwd() //used to get the current working directory
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath,"-task", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf(" 1: %s\n", task)
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})
	
	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
		l :=  todo.List{}

		if err := l.Get(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if l[0].Done != true{
			t.Errorf("Expected task %s to be %t but got %t", l[0].Task, true, l[0].Done)
		}
	})

}
