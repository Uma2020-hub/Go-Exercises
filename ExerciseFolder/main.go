package main

import (
	"time"
)

//Health ..
type Health interface {
	Exercise(tm time.Duration)
	Steps(n int)
}

//Watch ..
type Watch struct {
	ExerciseDuration time.Duration
	StepCount        int
	Calories         int
}

//Exercise ...
func (w *Watch) Exercise(t time.Duration) {
	w.Calories = w.Calories + int(t.Minutes())
}

//Steps ...
func (w *Watch) Steps(n int) {
	if n == 1000 {
		w.Calories = w.Calories + 1
	}

}

func main() {

}
