package file

import (
    "fmt"
    "os"
	"strings"
	"bufio"
	"errors"
	"io/fs"
)





func FileRead(outFile string)[]string{
var allFirst []string
		    filename, err := os.Open(outFile)
			if err != nil{
				fmt.Printf("%s[!] Error: Cannot open %s%s", bred, outFile, rst)
			}
			if _, err = os.Stat(outFile); errors.Is(err, fs.ErrNotExist){
				fmt.Printf("\n%s[!] Error: %s does not exist%s\n", bred, outFile, rst)
                 
				}
			scanner := bufio.NewScanner(filename)
			for scanner.Scan(){
				hashLine := scanner.Text()
				hashLine = strings.TrimSpace(hashLine)
				if strings.Contains(hashLine, "#") {
				continue
				}
				if strings.Contains(hashLine, ":"){
				split := strings.Split(hashLine, ":")
				first := split[0]
				last := split[1]
				if len(last) == 0 {
			    allFirst = append(allFirst, first)
				}
				}
				}
			return allFirst
			}
