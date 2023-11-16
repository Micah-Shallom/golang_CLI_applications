package todo

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {

	l := List{}

	taskName := "Go to the gym"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q,got %q. \n", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := List{}

	taskName := "Go to school"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Error("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Error("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := List{}

	strArr := []string{
		"New String1",
		"New String2",
		"New String3",
	}

	for _, v := range strArr {
		l.Add(v)
	}

	if l[0].Task != strArr[0] {
		t.Errorf("Expected %q, got %q instead.\n", strArr[0], l[0].Task)
	}
	
	l.Delete(2)
	
	if len(l) != 2 {
		t.Errorf("Expected length is %d, got %d instead\n", 2, len(l))
	}
	
	if l[1].Task != strArr[2] {
		t.Errorf("Expected %q, got %q instead.\n", strArr[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T){
	l1 := List{}
	l2 := List{}

	taskName := "Go to church"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead. \n", taskName, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("Error creating temporary file %s", err )
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Errorf("Error Saving list to file: %s\n", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Errorf("Error Getting list from file: %s\n", err)
	}

	if l1[0].Task != l2[0].Task{
		t.Errorf("Task %q should match %q.\n", l1[0].Task, l2[0].Task)
	}
}