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

func TestParseLineNoDependencies(t *testing.T) {
	input := "Task 1, 10"
	tasks := parse(input)

	if tasks[0].Name != "Task 1" {
		t.Errorf("Expected task name 'Task 1', got %s", tasks[0].Name)
	}

	if tasks[0].Duration != 10 {
		t.Errorf("Expected task duration 10, got %d", tasks[0].Duration)
	}
}
func TestParseLineWithDependencies(t *testing.T) {
	input := "Task 2, 10, Task 1"
	tasks := parse(input)

	if tasks[0].Name != "Task 2" {
		t.Errorf("Expected task name 'Task 2', got %s", tasks[0].Name)
	}

	if tasks[0].Duration != 10 {
		t.Errorf("Expected task duration 10, got %d", tasks[0].Duration)
	}

	if tasks[0].Dependencies[0] != "Task 1" {
		t.Errorf("Expected dependency 'Task 1' got %s", tasks[0].Dependencies[0])
	}
}

func TestExpectedTime(t *testing.T) {
	tasks := []Task{
		{Name: "Task 1", Duration: 3, Dependencies: []string{}},
		{Name: "Task 2", Duration: 2, Dependencies: []string{"Task 1"}},
		{Name: "Task 3", Duration: 1, Dependencies: []string{"Task 1"}},
		{Name: "Task 4", Duration: 4, Dependencies: []string{"Task 2", "Task 3"}},
	}
	time := expectedTime(tasks)
	if time != 9 {
		t.Errorf("Expected 9, got %d", time)
	}
}
