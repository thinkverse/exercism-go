package gigasecond

import "time"

// AddGigasecond returns the time t+gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
