# pomo
Simple pomodoro CLI application for productivity and focus

## usage 
```sh
pomo [command] [args]
```

### add a task to the pomo db
```sh
# pomo task add [title] [description]
pomo task add "Refactor main method" "Add comments, manage memory, add switch statements"
```

### list all tasks in db ie, tasks to do
```sh
pomo task list
```

### remove a task from the db
```sh
pomo task remove [id]
pomo task remove 2
```

### start a pomodoro session with all tasks in db
```sh
pomo start all
```

### start a pomodoro sessions with select tasks
```sh
pomo start [ids, comma seperated, in order]
pomo start 2,4,5,3
```

## installation
make sure $GOPATH/bin is in $PATH, and then
```bash
go get github.com/skovati/pomo`

pomo ...
```
