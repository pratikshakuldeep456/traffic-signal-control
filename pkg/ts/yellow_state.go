package ts

import "fmt"

type YellowState struct{}

func (r *YellowState) execute(ctx *TrafficSignalContext) {
	fmt.Println("Signal is Yellow")
	ctx.Wait(ctx.Duration.YellowDuration)
	ctx.SetState(&RedState{})

}
