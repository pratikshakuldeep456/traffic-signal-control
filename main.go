package main

import (
	"pratikshakuldeep456/traffic-signal-control/pkg/ts"
	"time"
)

func main() {

	durations := &ts.SignalDuration{
		RedDuration:    5 * time.Second,
		YellowDuration: 5 * time.Second,
		GreenDuration:  5 * time.Second,
	}

	signal := ts.NewTrafficSignal(*durations)

	go func() {
		time.Sleep(5 * time.Second)
		signal.TriggerEmergency()
	}()
	signal.Start()

}
