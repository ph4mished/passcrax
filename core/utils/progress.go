package utils

import (
	"fmt"
	"time"
)

func formatDuration(timeDur time.Duration) string {
	secs := int(timeDur.Seconds())

	mins := secs / 60
	secs = secs % 60

	hours := mins / 60
	mins = mins % 60

	days := hours / 24
	hours = hours % 24

	if days > 0 {
		return fmt.Sprintf("%dd %02dh:%02d:%02ds", days, hours, mins, secs)
	}
	return fmt.Sprintf("%02dh:%02d:%02ds", hours, mins, secs)
}

func PrintProgress(cracked, total int, startTime time.Time) {
	//var totalTime time.Time
	progress := float64(cracked) / float64(total)
	percent := progress * 100

	barLength := 10
	filledLength := int(progress * float64(barLength))
	bar := ""
	for i := 0; i < barLength; i++ {
		if i < filledLength {
			bar = "\u2588" + bar
		} else {
			bar = bar + "_"
		}
	}

	timeElapsed := time.Since(startTime).Seconds()
	elapsed := time.Since(startTime)
	velocity := float64(cracked) / timeElapsed
	var eta string
	if cracked > 0 {
		remaining := time.Duration(float64(elapsed) * (float64(total)/float64(cracked) - 1))
		if remaining < 0 {
			eta = "ETA: \u221E"
		} else {
			eta = fmt.Sprintf("ETA: %s", formatDuration(remaining))
		}
	} else {
		eta = "ETA: --"
	}
	fmt.Printf("\r%s[%s%s%s%s%s] %.2f%% %s|%s %s%.2fH/s %s%s|%s %s%s%s", bgrn, rst, bblu, bar, rst, bgrn, percent, bcyn, rst, bgrn, velocity, rst, bcyn, rst, bgrn, eta, rst)
}
