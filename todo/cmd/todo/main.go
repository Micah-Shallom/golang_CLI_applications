package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	todo "github.com/Micah-Shallom/golang_CLI_applications/todo"
)

func main() {
	//Parsing command line flags
	// task := flag.String("task", "", "Task to be included into the todo list")
	add := flag.Bool("add", false, "Add task to the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("del", 0, "Item to be deleted")
	verbose := flag.Bool("v", false, "Show more item information")
	show := flag.Bool("show-incompleted", false, "Show incompleted tasks")
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Micah Shallom\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyrigt 2023\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information")
		flag.PrintDefaults()
	}

	var todoFileName = ".todo.json"

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	//define items list
	l := &todo.List{}

	//use the GET command to read to-do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//decide what to do based on the provided flags
	switch {
	case *list:
		formatted := ""
		for k, t := range *l {
			prefix := "  "
			if *show {
				if t.Done {
					continue
				}
			} else {
				if t.Done {
					prefix = "X "
				}
			}
			if *verbose {
				date, time := formatDateTime(t.CreatedAt)
				formatted += fmt.Sprintf("%s%d: %s  %s  %s\n", prefix, k+1, t.Task, date, time)
			} else {
				formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
			}
		}
		fmt.Println(formatted)
	// 	fmt.Print(l) //leverages the springer interface

	// case *list && *verbose:
	// 	// if *verbose {
	// 	// 	for _, item := range *l {

	// 	// }
	// 	fmt.Print(l)

	// // case *task != "":
	// // 	l.Add(*task)

	// // 	if err := l.Save(todoFileName); err != nil {
	// // 		fmt.Fprintln(os.Stderr, err)
	// // 		os.Exit(1)
	// 	}

	case *add:
		// when any arguments (excluding flags) are provided, they will be used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		//save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *complete > 0:
		//complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		//save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *delete > 0:
		//delete the given item
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}

}

//getTask function decides where to get the decsripton for a new task from: argument or STDIN

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan() //scans the input

	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil //returns scanned text and error
}


func formatDateTime(l time.Time) (string, string) {
	//Extract and print date component
	year, month, day := l.Date()
	date := fmt.Sprintf("%d-%02d-%02d", year, month, day)

	// Extract and print the time components
	hour, minute, second := l.Clock()
	time := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	return date, time
}
