package todo

type Task struct {
	ID          int
	Description string
	State       bool
}

func (t *Task) updateTask(newTask Task) {
	t.Description = newTask.Description
	t.State = newTask.State
}
