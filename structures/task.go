package structures

type Task struct {
	ID          int
	Description string
	State       bool
}

func (t *Task) UpdateTask(newTask Task) {
	t.Description = newTask.Description
	t.State = newTask.State
}

func (t *Task) EqualsTo(compTask Task) bool {
	if t.Description == compTask.Description && t.State == compTask.State && t.ID == compTask.ID {
		return true
	}
	return false
}
