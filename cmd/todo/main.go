package main

import (
  "bufio"
  "flag"
	"fmt"

  "io:
	"os"

  "strings"

	"github.com/betology/todo"
)

// Default file name
var todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {

		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for the Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyleft 2020\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	// Parsing comman line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all task")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

  if os.Getenv("TODO_FILENAME") != "" {
    todoFileName = os.Getenv("TODO_FILENAME")
  }
	// Define an items list
	l := &todo.List{}

	// Use the Get command  to read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the provided flags
	switch {
	case *list:
		// List current to do items
    fmt.Print(l)

	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		// Add the task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		// Invalid flag privided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
