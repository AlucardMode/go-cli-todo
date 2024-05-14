package structures

// import (
// 	"fmt"
// )

type List struct {
	Title string
	Tasks []Task
}

// // Constructor
// func NewList(title string) *List {
// 	return &List{
// 		Title: title,
// 		Tasks: []Task{},
// 	}
// }

// Functions
func (l *List) AddTask(description string) {
	l.Tasks = append(l.Tasks, Task{ID: len(l.Tasks) + 1, Description: description, State: false})
}

// // Searches for task in list with id. if id is found, uses updateTask method from Task Struct
// func (l *List) UpdateTask(id int, newTask Task) {
// 	_, err := l.FindTask(id)
// 	if err != nil {
// 		syserr := fmt.Errorf("Task %d not found in list", id)
// 		panic(syserr)
// 	}
// 	l.Tasks[id].updateTask(newTask)

// }

// // helpers
// func (l *List) FindTask(id int) (bool, error) {
// 	for _, task := range l.Tasks {
// 		if task.ID == id {
// 			return true, nil
// 		}
// 	}
// 	return false, fmt.Errorf("ERROR, Item %d not found in slice", id)
// }
