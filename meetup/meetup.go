package meetup

import (
	"time"
)

// WeekSchedule represents the schedule of the day
type WeekSchedule int8

// Various schedules supported
const (
	_                   = iota
	First  WeekSchedule = iota
	Second WeekSchedule = iota
	Third  WeekSchedule = iota
	Fourth WeekSchedule = iota
	Fifth  WeekSchedule = iota
	Teenth WeekSchedule = iota
	Last   WeekSchedule = iota
)

var scheduleStartDayMap = map[WeekSchedule]int{
	First:  1,
	Second: 8,
	Third:  15,
	Fourth: 22,
	Fifth:  29,
	Teenth: 13,
}

// Day returns the expected day given the schedule.
func Day(week WeekSchedule, weekDay time.Weekday, month time.Month, year int) int {
	var day int
	if week == Last {
		day = getDaysInMonth(month, year) - 6
	} else {
		day = scheduleStartDayMap[week]
	}
	for {
		ts := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if ts.Weekday() == weekDay {
			return day
		}
		day++
	}
}

func getDaysInMonth(month time.Month, year int) int {
	ts := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return ts.Day()
}
