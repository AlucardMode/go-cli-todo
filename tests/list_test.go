package structures_tests

import (
	"testing"

	"github.com/AlucardMode/go-cli-todo/structures"
)

func TestAddTask(t *testing.T) {
	list := structures.List{Title: "Test List"}
	list.AddTask("Complete the Project")

	if len(list.Tasks) != 1 {
		t.Errorf("expected 1 test item but got %d", len(list.Tasks))
	}
	if list.Tasks[0].Description != "Complete the Project" {
		t.Errorf("Description does not match expected value: %s", list.Tasks[0].Description)
	}
}
