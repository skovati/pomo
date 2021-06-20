package main

import (
    "testing"
    "time"
    "io"
    "os"
    "strings"
    "bytes"
)

func TestLongTimer(t *testing.T) {
    // redirect stdout
    // caputure original stdout
    origStdout := os.Stdout
    r, w, err := os.Pipe()
    if err != nil {
        t.Errorf("Stdout redirection failed: %v", err)
    }
    // set stdout to capture stream
    os.Stdout = w

    capChan := make(chan string)

    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        capChan <- buf.String()
    }()

    // get current time, run timer, and check time since
    before := time.Now()
    expected, _ := time.ParseDuration("3s")
    countdown("", expected)
    diff := time.Now().Sub(before)
    if diff < time.Second*3 {
        t.Errorf("Countdown timer did not last long enough, got %s, expected %s", diff.String(), expected.String())
    }

    // close temp stdout
    err = w.Close()
    if err != nil {
        t.Errorf("Error closing temp stdout: %v", err)
    }

    // read in captured stdout
    captured := <-capChan
    if err != nil {
        t.Errorf("Error reading temp stdout: %v", err)
    }

    // restore stdout
    os.Stdout = origStdout

    // check for correct timer output
    lines := strings.Split(string(captured), "\n")
    expectedFirstLine := "00:00:02" + ansiClear
    if lines[1] != expectedFirstLine {
        t.Errorf("Incorrect timer output, got %s, expected, %s", lines[1], expectedFirstLine)
    }
}

func TestShortTimer(t *testing.T) {
    // redirect stdout, caputure original stdout
    origStdout := os.Stdout
    _, w, err := os.Pipe()
    if err != nil {
        t.Errorf("Stdout redirection failed: %v", err)
    }
    // set new stdout
    os.Stdout = w

    // text countdown time
    before := time.Now()
    expected, _ := time.ParseDuration("0.1s")
    countdown("", expected)
    diff := time.Now().Sub(before)
    // and check to make sure duration is as expected
    if diff < time.Second/10 {
        t.Errorf("Countdown timer did not last long enough, got %s, expected %s", diff.String(), expected.String())
    }

    // close redir stdout
    err = w.Close()
    if err != nil {
        t.Errorf("Error closing temp stdout: %v", err)
    }

    // restore original stdout
    os.Stdout = origStdout
}
