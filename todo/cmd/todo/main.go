package main

import (
	"flag"
	"fmt"
	"os"
	todo "github.com/Micah-Shallom/modules"
)

func main() {
	//Parsing command line flags
	task := flag.String("task","","Task to be included into the todo list")
	list := flag.Bool("list",false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	flag.Usage =  func ()  {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Micah Shallom\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyrigt 2023\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information")
		flag.PrintDefaults()
	}


	const todoFileName = "todo.json"

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
		// for _, item := range *l{
		// 	if !item.Done{
		// 		fmt.Printf("%s\n",item.Task)
		// 	}
		// }
		fmt.Print(l)

	case *task != "":
		l.Add(*task)

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

	default:
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}


}
