package main

import (
    "fmt"
    "time"
    "os"
    // "flag"
)

const usage = `
    durations entered in "XhYmZs" format
    pomo [focus duration] [break duration]

`

func main() {
    if len(os.Args) < 2 {
        fmt.Print(usage)
        os.Exit(1)
    }
    switch os.Args[1] {
        case "task":
            taskHandler(os.Args[1:])
    }
    pomoDuration, err := time.ParseDuration(os.Args[1])
    if err != nil {
        fmt.Print(usage)
        os.Exit(1)
    }
    breakDuration, err := time.ParseDuration(os.Args[2])
    if err != nil {
        fmt.Print(usage)
        os.Exit(1)
    }
    PomoLoop(pomoDuration, breakDuration)
    fmt.Println("\nTimer complete")
}

func taskHandler(args []string) {
    switch args[1] {
    case "add":
        // addTask(args[1:])
    case "remove":
        // removeTask(args[1:])
    }
}
