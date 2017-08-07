package clock

import "fmt"

const testVersion = 4

type Clock struct {
	Hour int
	Minute int
}

func New(hours, minutes int) Clock {
	hours, minutes = convertTimes(hours, minutes)
	hours, minutes = formatClock(hours, minutes)
	return Clock{hours, minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	hours := c.Hour
	minutes += c.Minute
	hours, minutes = convertTimes(hours, minutes)
	c.Hour, c.Minute = formatClock(hours, minutes)
	return c
}

func convertTimes(hours, minutes int) (int, int) {
	if minutes >= 60 || minutes <= -60 {
		hours += minutes / 60
		minutes = minutes % 60
	}
	return hours, minutes
}

func formatClock(hours, minutes int) (int, int) {
	if minutes < 0 {
		hours -= 1
		minutes = 60 + minutes
	}
	if hours < 0 {
		hours = hours % 24
		hours = 24 + hours
	}
	if hours >= 24 {
		hours = hours % 24
	}
	return hours, minutes
}
