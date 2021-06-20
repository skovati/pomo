package main

import (
    "testing"
    "time"
)

func TestLongTimer(t *testing.T) {
    before := time.Now()
    expected, _ := time.ParseDuration("5s")
    countdown("", expected)
    diff := time.Now().Sub(before)
    if diff < time.Second*5 {
        t.Errorf("Countdown timer did not last long enough, got %s, expected %s", diff.String(), expected.String())
    }
}

func TestShortTimer(t *testing.T) {
    before := time.Now()
    expected, _ := time.ParseDuration("0.1s")
    countdown("", expected)
    diff := time.Now().Sub(before)
    if diff < time.Second/10 {
        t.Errorf("Countdown timer did not last long enough, got %s, expected %s", diff.String(), expected.String())
    }
}
