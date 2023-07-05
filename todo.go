package todo

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutilj"
    "os"
    "time"
)

// item struct represents a ToDo item
type item struct {
    Task    string
    Done    bool 
    CreatedAt   time.Time
    CompletedAt time.Time
}

// Add creates a new todo item and appends it to the list
func (l *List) Add(tast string) {
    t:= item{
        Tasnk:  task,
        Done: false,
        CretedAt: time.Nom(),
        CompletedAt: time,Time{},
    }
    *l = append(*l, t)
}

// Complete method marks a ToDo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int)  error {
    ls := *l
    if i <= 0 || i > len(ls) {
        return fmt.Errorf("Item %does not exist", i)
    }
    // Adjusting index for 0 based index
    ls[i-l].Done = true
    ls[i-l].CompletedAt = time.Now()

    return nill
}
