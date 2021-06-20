package main

import (
    "fmt"
    "time"
)

const sec = time.Second
const maxPomoCounter = 4
const header = "=== pomodoro timer, a simple way to stay focused ===\n\n"

func PomoLoop(pomoDuration time.Duration, breakDuration time.Duration, ) {
    numTasks := 8
    pomoCounter := 1

    // loop over first 
    for i := 0; i < numTasks; i++ {
        countdown(fmt.Sprintf("%sPomodoro #%d/%d\nRemaining time:", header, i+1, numTasks), pomoDuration)
        if pomoCounter == 4 {
            countdown(fmt.Sprintf("%sRemaining long break time:", header), breakDuration*6)
        } else {
            countdown(fmt.Sprintf("%sRemaining short break time:", header), breakDuration)
        }
        pomoCounter++
    }
}

func countdown(message string, duration time.Duration) {
    timer := time.NewTimer(duration)
    ticker := time.NewTicker(sec)
    // loop until the time
    for {
        select {
        // if the ticker ticks, a second has passed, so we decrement duration left
        case <-ticker.C:
            duration -= time.Duration(sec)
            // clear screen
            fmt.Print("\033[H\033[2J")

            // print message and time
            fmt.Println(message)
            fmt.Print(format(duration))
        // if the timer channel is nonempty 
        case <-timer.C:
            return
        }
    }
}

func format(d time.Duration) string {
    // round duration to nearest second so we can do integer arithmetic
    rd := d.Round(time.Second)
    // calc hours
    hours := rd / time.Hour
    // and subtract hours from dur
    rd -= hours * time.Hour
    // same for mins
    mins := rd / time.Minute
    rd -= mins * time.Minute
    // calc secs
    secs := rd / time.Second

    if hours > 1 {
        return fmt.Sprintf("%02d:%02d", mins, secs)
    } else {
        return fmt.Sprintf("%02d:%02d:%02d", hours, mins, secs)
    }
}
