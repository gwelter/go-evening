package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	start time.Time
	end   time.Time
}

func NewDateRange(start, end time.Time) (DateRange, error) {
	if start.IsZero() {
		return DateRange{}, errors.New("start cannot be nil")
	}
	if end.IsZero() {
		return DateRange{}, errors.New("end cannot be nil")
	}
	if end.Before(start) {
		return DateRange{}, errors.New("start cannot be after end")
	}
	return DateRange{
		start: start,
		end:   end,
	}, nil
}

func (d DateRange) Hours() float64 {
	return d.end.Sub(d.start).Hours()
}

func main() {
	lifetime := DateRange{
		start: time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		end:   time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(lifetime.Hours())

	travelInTime := DateRange{
		start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		end:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(travelInTime.Hours())
}
