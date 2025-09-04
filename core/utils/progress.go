package utils

import (
	"fmt"
	"time"

//	"github.com/fatih/color"
)

/*var (
	bgrn = color.New(color.FgGreen, color.Bold)
	bred = color.New(color.FgRed, color.Bold)
	bblu = color.New(color.FgBlue, color.Bold)
	bcyn = color.New(color.FgCyan, color.Bold)
	bylw = color.New(color.FgYellow, color.Bold)
	grn  = color.New(color.FgGreen)
	red  = color.New(color.FgRed)
	blu  = color.New(color.FgBlue)
	cyn  = color.New(color.FgCyan)
	ylw  = color.New(color.FgYellow)
)*/

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
	progress := float64(cracked) / float64(total)
	percent := progress * 100

	barLength := 10
	filledLength := int(progress * float64(barLength))
	bar := ""
	for i := 0; i < barLength; i++ {
		if i < filledLength {
			bar += "\u2588"
		} else {
			bar += "_"
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

	bgrn.Print("\r[")
	bblu.Print(bar)
	bgrn.Print("]")
	bgrn.Printf(" %.2f%%", percent)
	bcyn.Print(" | ")
	bgrn.Printf("%.2fH/s", velocity)
	bcyn.Print(" | ")
	bgrn.Print(eta)
}

// Optional helper functions for consistent color usage throughout in this tool.. will be of  use later
func PrintError(message string) {
	bred.Printf("ERROR: %s\n", message)
}

func PrintWarning(message string) {
	bylw.Printf("WARNING: %s\n", message)
}

func PrintSuccess(message string) {
	bgrn.Printf("SUCCESS: %s\n", message)
}

func PrintInfo(message string) {
	bblu.Printf("INFO: %s\n", message)
}

func PrintDebug(message string) {
	cyn.Printf("DEBUG: %s\n", message)
}
