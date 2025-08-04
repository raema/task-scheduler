package main

import "testing"

func TestCreateTask(t *testing.T) {
	task := Task{
		Name:         "Task 1",
		Duration:     10,
		Dependencies: []string{},
	}

	if task.Name != "Task 1" {
		t.Errorf("Expected task name to be 'Task 1', got %s", task.Name)
	}

	if task.Duration != 10 {
		t.Errorf("Expected task duration to be 10, got %d", task.Duration)
	}

	if len(task.Dependencies) != 0 {
		t.Errorf("Expected task dependencies to be empty, got %v", task.Dependencies)
	}

}
