package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"path/filepath"
	"errors"
	"io/fs"
)

func FileRead(hashFile string) ([]string, []string) {
	var allFirst, allLast []string
	filename, err := os.Open(hashFile)
	if err != nil {
		fmt.Printf("%s[!] Error: Cannot open %s%s", bred, hashFile, rst)
	}
	if _, err = os.Stat(hashFile); errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("\n%s[!] Error: %s does not exist%s\n", bred, hashFile, rst)
	}
	scanner := bufio.NewScanner(filename)
	for scanner.Scan() {
		hashLine := scanner.Text()
		hashLine = strings.TrimSpace(hashLine)
		if strings.Contains(hashLine, "#") {
			continue
		}
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
	return allLast, allFirst
}
