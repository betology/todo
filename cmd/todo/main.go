package main 

import (
    "fmt"
    "os"
    "strings"

    "github.com/betology/todo"
)

// Hardcoding the file name
const todoFileName = ".todo.json"

func main ) {

    // Define an items list
    l :=&todo.List{}


    // Use the Get method to read to do items from file
    if err := l.Get(todoFileName); err != nill {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
