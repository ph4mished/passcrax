package file

import (
	"bufio"
	//	"fmt"
	"errors"
	"io/fs"
	"os"
	"strings"
)

import "github.com/fatih/color"

var (
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
)

// this function is RAM expensive
func FileRead(hashFile string) ([]string, []string) {
	var allFirst, allLast []string
	filename, err := os.Open(hashFile)
	if err != nil {
		bred.Printf("[!] Error: Cannot open %s", hashFile)
	}
	if _, err = os.Stat(hashFile); errors.Is(err, fs.ErrNotExist) {
		bred.Printf("\n[!] Error: %s does not exist\n", hashFile)
	}
	scanner := bufio.NewScanner(filename)
	for scanner.Scan() {
		hashLine := strings.TrimSpace(scanner.Text())

		if !strings.HasPrefix(hashLine, "#") {
			if strings.Contains(hashLine, ":") {
				split := strings.Split(hashLine, ":")
				first := split[0]
				last := split[1]
				allLast = append(allLast, last)
				if len(last) == 0 {
					allFirst = append(allFirst, first)
				}

			}
		}
	}
	return allLast, allFirst
}
