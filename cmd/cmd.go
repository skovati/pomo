package cmd

import (
    "fmt"
    "flag"
    "time"
    "os"
)

var usage string = "pomo: a simple way to stay focused\n" +
    "usage: pomo [command] [flags]"

func Run() {
    // status command
    statusCmd := flag.NewFlagSet("status", flag.ExitOnError)
    statusVerbose := statusCmd.Bool("v", false, "enable verbose status output")
    // set command
    defDur, _ := time.ParseDuration("25m")
    defBreakDur, _ := time.ParseDuration("5m")
    setCmd := flag.NewFlagSet("set", flag.ExitOnError)
    dur := setCmd.Duration("d", defDur, "duration of pomodoro")
    breakDur := setCmd.Duration("b", defBreakDur, "duration of break")
    pomoNum := setCmd.Int("n", 3, "number of pomodoros until longer break")

    switch os.Args[1] {
    case "status":
        statusCmd.Parse(os.Args[2:])
        fmt.Println("status subcommand")
        if(*statusVerbose) {
            fmt.Println("verbose mode")
        }
    case "set":
        setCmd.Parse(os.Args[2:])
        fmt.Printf("set subcommand\n")
        fmt.Printf("duration: %s\n", *dur)
        fmt.Printf("break duration: %s\n", *breakDur)
        fmt.Printf("pomo num: %d\n", *pomoNum)
    default:
        fmt.Printf("%s\n", usage)
        statusCmd.Parse(os.Args[2:])
        setCmd.Parse(os.Args[2:])
        fmt.Printf("%s\n", "pomo status:")
        statusCmd.PrintDefaults()
        fmt.Printf("%s\n", "pomo set:")
        setCmd.PrintDefaults()
    }
}

func status() {
    
}

func set() {
    
}
