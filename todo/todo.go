package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task      		string
	Done      		bool
	CreatedAt 		time.Time
	CompletedAt 	time.Time
}

type Stringer interface{
	String() string
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}


func (l *List) Complete(i int) error{
	ls := *l
	
	if i <=0 || i > len(ls){
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error{
	ls := *l
	
	if i <=0 || i > len(ls){
		return fmt.Errorf("item %d does not exist, nothing to delete", i)
	}
	
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *List) Save(filename string) error{
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, data, 0644)
}

func (l *List) Get(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist){
			return nil
		}
		return err
	}
		return json.Unmarshal(data, l)
}

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := "  "
		if t.Done{
			prefix = "X "
		}
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}

