package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Parse a formatted time like hours:minutes:seconds,milliseconds
// Ex) 00:00:00,000
// Source: https://github.com/konifar/go-srt/blob/5d3adcc53519032da72bc4334cf2b37447df6b3a/scanner.go#L132
func parseTime(input string) (time.Duration, error) {
	regex := regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2}),(\d{3})`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) < 4 {
		return time.Duration(0), fmt.Errorf("invalid time format:%s", input)
	}

	hour, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Duration(0), err
	}
	minute, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Duration(0), err
	}
	second, err := strconv.Atoi(matches[3])
	if err != nil {
		return time.Duration(0), err
	}
	millisecond, err := strconv.Atoi(matches[4])
	if err != nil {
		return time.Duration(0), err
	}

	return time.Duration(time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute + time.Duration(second)*time.Second + time.Duration(millisecond)*time.Millisecond), nil
}
