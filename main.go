package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	file := flag.String("file", "", "task file")
	validate := flag.Bool("validate", false, "validate the task list")
	run := flag.Bool("run", false, "run the task list")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	input, err := readTaskFile(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading task file: %s\n", err)
		os.Exit(0)
	}

	tasks := parse(input)

	if *validate {
		time := expectedTime(tasks)
		fmt.Printf("Expected time: %d\n", time)
	}

	if *run {
		expectedTime := expectedTime(tasks)
		runTime := runTasks(tasks)
		difference := runTime - expectedTime
		fmt.Printf("Run time (sec): %d\n", runTime)
		fmt.Printf("Expected time (sec): %d\n", expectedTime)
		fmt.Printf("Time difference (sec): %d\n", difference)

	}

}

type Task struct {
	Name         string
	Duration     int
	Dependencies []string
}

func readTaskFile(filename string) (string, error) {
	if filename != "" {
		data, err := os.ReadFile(filename)
		if err != nil {
			return "", fmt.Errorf("error reading from file: %s", err)
		}
		return string(data), nil
	} else {
		return "", nil
	}
}

func parse(input string) []Task {
	if input == "" {
		return []Task{}
	}

	var tasks []Task
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // skip empty lines
		}

		fields := strings.Split(line, ",")
		if len(fields) < 2 {
			fmt.Fprintf(os.Stderr, "skipping malformed line: %s\n", line)
			os.Exit(1)
		}

		name := strings.TrimSpace(fields[0])
		duration, err := strconv.Atoi(strings.TrimSpace(fields[1]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid duration in line: %s\n", line)
			os.Exit(1)
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

func runTasks(tasks []Task) int {
	// * save start time at beginning of runTasks
	// * create a map of channel structs for tracking signals
	// * create a channel in the map for each task
	// * create WaitGroup to track goroutine completion
	// * launch each task in parallel
	// * tasks wait for each dependency to signal done via channel before running
	// * WaitGroup waits for all goroutines to be Done
	// * return actual runtime since start time of runTasks

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
