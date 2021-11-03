package timer

import (
    "time"
)

type Timer struct {
    title       string          // title of current session
    start       time.Time       // start time of timer
    dur         time.Duration   // duration of each pomodoro
    breakDur    time.Duration   // duration of the break time
    num         int             // number of sessions until task complete
}

func NewTimer(title string, dur time.Duration, breakDur time.Duration, num int) *Timer {
    t := Timer{}
    t.title = title
    t.dur = dur
    t.breakDur = breakDur
    t.num = num
    t.start = time.Now()
    return &t
}

func (t Timer) GetCount() time.Duration {
    return time.Now().Sub(t.start)
}

func (t Timer) ToString() string {
    return t.GetCount().String() + ": " + t.title
}
