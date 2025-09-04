package file

import (
	"bufio"
	//	"fmt"
	"errors"
	"io/fs"
	"os"
	"strings"
)

func FileCount(hashFile string) int {
	totalLines := 0
	filename, err := os.Open(hashFile)
	if err != nil {
		bred.Printf("[!] Error: Cannot open %s", hashFile)
	}
	if _, err = os.Stat(hashFile); errors.Is(err, fs.ErrNotExist) {
		bred.Printf("\n[!] Error: %s does not exist%s\n", hashFile)
	}
	scanner := bufio.NewScanner(filename)
	for scanner.Scan() {
		hashLine := strings.TrimSpace(scanner.Text())
		if hashLine != "" {
			totalLines++
		}
	}
	return totalLines
}
