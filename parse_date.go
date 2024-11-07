package main

import (
	"fmt"
	"time"
)

func ParseDate(input string) (time.Time, error) {

	// It'll parse in multiple formats
	formats := []string{"2006-01-02", "02-01-2006", "01/02/2006"} // define all input formats for date
	for _, format := range formats {
		date, err := time.Parse(format, input)
		if err == nil {
			return date, err
		}
	}

	return time.Time{}, fmt.Errorf("Invalid date format. Failed to parse date\nPlease use YYYY-MM-DD, DD-MM-YYYY, or MM/DD/YYYY")
}
