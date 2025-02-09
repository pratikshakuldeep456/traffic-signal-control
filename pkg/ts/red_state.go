package ts

import "fmt"

type RedState struct{}

func (r *RedState) execute(ctx *TrafficSignalContext) {
	fmt.Println("Signal is RED")
	ctx.Wait(ctx.Duration.RedDuration)
	ctx.SetState(&GreenState{})
}
