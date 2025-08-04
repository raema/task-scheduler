# Task Scheduler

Task scheduler command line tool. Schedules and optionally runs a series of tasks in parallel,
according to a task list specification.

The schema for the task list is:

```sh
name, duration in seconds, dependencies (as a list of names)
...
```

## Build
```sh
go build 
```
## Run
```sh
./task-scheduler
```

## Test
```sh
go test
```