package main

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {

}

type Task struct {
	Name         string
	Duration     int
	Dependencies []string
}

func parse(input string) []Task {
	if input == "" {
		return []Task{}
	}

	var tasks []Task
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		fields := strings.Split(line, ",")
		name := strings.TrimSpace(fields[0])
		duration, err := strconv.Atoi(strings.TrimSpace(fields[1]))
		if err != nil {
			continue
		}

		var dependencies []string
		for i := 2; i < len(fields); i++ {
			dependency := strings.TrimSpace(fields[i])
			dependencies = append(dependencies, dependency)
		}

		tasks = append(tasks, Task{
			Name:         name,
			Duration:     duration,
			Dependencies: dependencies,
		})

	}
	return tasks

}

func expectedTime(tasks []Task) int {
	endTimeByTask := make(map[string]int)

	for len(endTimeByTask) < len(tasks) {

		for _, task := range tasks {
			if _, done := endTimeByTask[task.Name]; done {
				continue // task is done
			}

			ready := true
			startTime := 0

			for _, dependency := range task.Dependencies {
				if endTime, done := endTimeByTask[dependency]; !done {
					ready = false //dependency not done
					break
				} else if endTime > startTime {
					startTime = endTime
				}

			}

			if ready {
				endTimeByTask[task.Name] = startTime + task.Duration
			}
		}
	}

	totalTime := 0
	for _, endTime := range endTimeByTask {
		if endTime > totalTime {
			totalTime = endTime
		}
	}
	return totalTime
}

func run(tasks []Task) int {
	startTime := time.Now()
	done := make(map[string]chan struct{})

	// initialize channels
	for _, task := range tasks {
		done[task.Name] = make(chan struct{})
	}

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	// launch all tasks
	for _, task := range tasks {
		go func(t Task) {
			defer wg.Done()

			// Wait for dependencies
			for _, dependency := range t.Dependencies {
				<-done[dependency]
			}

			// Execute task
			time.Sleep(time.Duration(t.Duration) * time.Second)
			close(done[t.Name])
		}(task)
	}

	wg.Wait()
	return int(time.Since(startTime).Seconds())
}
