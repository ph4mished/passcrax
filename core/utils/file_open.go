package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func FileLaunch(filename string, commands int, permission os.FileMode) []string {
	var lineWords []string
	file, err := os.OpenFile(filename, commands, permission)

	if err != nil {
		fmt.Printf("\n%s[!] Error: %s%s", bred, err, rst)
	}
	defer file.Close()
	if _, err = os.Stat(filename); errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("\n%s[!] Error: %s does not exist%s\n", bred, filename, rst)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			lineWords = append(lineWords, line)
		}
	}
	return lineWords
}
