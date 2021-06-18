# pomo
Simple pomodoro CLI application for productivity and focus

## usage 
`pomo [command] [args]`

add a task to the pomo db
`pomo task add [title] [description]`
`pomo task add "Refactor main method" "Add comments, manage memory, add switch statements"`

list all tasks in db ie, tasks to do
`pomo task list`

remove a task from the db
`pomo task remove [id]`
`pomo task remove 2`

start a pomodoro session with all tasks in db
`pomo start all`

start a pomodoro sessions with select tasks
`pomo start [ids, comma seperated, in order]`
`pomo start 2,4,5,3`

## installation
make sure $GOPATH/bin is in $PATH, and then
`go get github.com/skovati/pomo`
`pomo ...`
