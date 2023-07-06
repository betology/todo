package todo_test

import (
    "io/ioutil"
    "os"
    "testing"
    "github.com/betology/todo"
)

func TestAdd( t *testing.T) {
    l := todo.List{}

    taskName := "New Task"
    l.Add(taskName)

    if l[0].Task != taskName {
        t.Errorf("Expedted %q, got %q instead.", taskName, l[0].Task)
    }
}

// TestComplete test the Complete method of the List type
func TestComplete(t *testing.T) {
    l := todo.List{}

    taskName := "New Task"
    l.Add(taskName)

    if l[0].Task != taskName {
        t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
    }
    if l[0].Done {
        t.Errorf("New task should not be completed.")
    }

    l.Complete(1)

    if !l[0].Done {
        t.Errorf("New task should be completed.")
    }
}

// TestDelete test the Delete method of the List type
func TestDelete(t *testing.T) {
    l := todo.List{}

    task := []string{
        "New Task 1",
        "New Task 2",
        "New Task 3",
    }

    for _, v := range task {
        l.Add(v)
    }

    if l[0].Task != task[0] {
        t.Errorf("Expected %q, got %q intead.", task[0], l[0].Task)
    }

    l.Delete(2)

    if len(l) != 2 {
        t.Errorf("Expected list lenght %d, got %d instead.", 2, len(l))
    }
    if l[1].Task != task[2] {
        t.Errorf("Expected %q, got %q insted.", task[2], l[1].Task)
    }
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
    l1 := todo.List{}
    l2 := todo.List{}
 
    taskName := "New Task"
    l1.Add (taskName)

    if l1[0].Task != taskName {
        t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
    }

    tf, err := ioutil.TempFile("", "")

    if err != nil {
        t.Fatalf("Error creatin temp file: %s", err)
    }
    defer os.Remove(tf.Name())

    if err := l1.Save(tf.Name()); err != nil {
        t.Fatalf("Error saving list to file: %s", err)
    }

    if err := l2.Get(tf.Name()); err != nil {
        t.Fatalf("Error getting list from file: %s", err)
    }

    if l1[0].Task != l2[0].Task {
        t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
    }
}
