package gigasecond

import (
	"time"
	"math"
)

const testVersion = 4
var gigasecond time.Duration = time.Duration(math.Pow(10, 9)) * time.Second

func AddGigasecond(initialTime time.Time) time.Time {
	finalTime := initialTime.Add(gigasecond)
	return finalTime
}
