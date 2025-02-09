package ts

import "fmt"

type GreenState struct{}

func (r *GreenState) execute(ctx *TrafficSignalContext) {
	fmt.Println("Signal is GREEN")
	ctx.Wait(ctx.Duration.GreenDuration)
	ctx.SetState(&YellowState{})

}
