package main

import (
	"runtime"
	"time"
)

func reset1(t *time.Timer, d time.Duration) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}
	t.Reset(d)
}

func reset2(t *time.Timer, d time.Duration) {
	if !t.Stop() && len(t.C) > 0 {
		<-t.C
	}
	t.Reset(d)
}

func runTimer(fn func(t *time.Timer, d time.Duration)) {
	tmr := time.NewTimer(0)
	fn(tmr, time.Minute)
	select {
	case <-tmr.C:
		panic("unexpected firing of Timer")
	default:
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	for {
		runTimer(reset1)
		runTimer(reset2)
	}
}
