package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds a gigasecond to time
func AddGigasecond(t time.Time) time.Time {
	gs, _ := time.ParseDuration("1000000000s")
	return t.Add(gs)
}
