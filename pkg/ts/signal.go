package ts

import "time"

type Signaltype string

const (
	Red    Signaltype = "RED"
	Yellow Signaltype = "YELLOW"
	Green  Signaltype = "GREEN"
)

type SignalDuration struct {
	RedDuration    time.Duration
	YellowDuration time.Duration
	GreenDuration  time.Duration
}
