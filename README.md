# Task Scheduler

Task scheduler command line tool. Schedules and optionally runs a series of tasks in parallel,
according to a task list specification.

The schema for the task list file is:

```sh
name, duration in seconds, dependencies (as a list of names)
...
```

## Build
```sh
go build 
```
## Usage
```sh
./task-scheduler
Usage of ./task-scheduler:
  -file string
        task file
  -run
        run the task list
  -validate
        validate the task list
```

## Example Task File
```sh
Task 1, 1 
Task 2, 2, Task 1
Task 3, 2, Task 2
```

## Example Run
```sh
./task-scheduler -file examples/taskfile -validate
Expected time: 5

./task-scheduler -file examples/taskfile -run     
Run time: 5
Expected time: 5
Difference: 0
```
## Test
```sh
go test
```