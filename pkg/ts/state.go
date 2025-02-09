package ts

import (
	"fmt"
	"time"
)

type TrafficSignalState interface {
	execute(*TrafficSignalContext)
}

type TrafficSignalContext struct {
	CurrentState TrafficSignalState
	Duration     SignalDuration
	Emergency    chan bool
}

func NewTrafficSignal(durations SignalDuration) *TrafficSignalContext {
	return &TrafficSignalContext{
		CurrentState: &RedState{},
		Duration:     durations,
		Emergency:    make(chan bool),
	}
}

func (ctx *TrafficSignalContext) SetState(state TrafficSignalState) {
	ctx.CurrentState = state

}

func (ctx *TrafficSignalContext) Start() {
	for {
		ctx.CurrentState.execute(ctx)
	}
}

func (ctx *TrafficSignalContext) Wait(duration time.Duration) {
	timer := time.NewTimer(duration)
	fmt.Println("timer", timer.C)
	select {
	case <-timer.C:
		return
	case <-ctx.Emergency:
		fmt.Println("Emergency detected! Switching to RED signal.")
		timer.Stop()
		ctx.SetState(&RedState{})
		ctx.Wait(ctx.Duration.RedDuration)
	}

}

func (ctx *TrafficSignalContext) TriggerEmergency() {
	ctx.Emergency <- true

}
