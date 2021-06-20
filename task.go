package main

type Task struct {
    title string
    desc string
    points int
}

func NewTask(title string, desc string, points int) Task {
    ret := Task{
        title: title,
        desc: desc,
        points: points,
    }
    return ret
}
