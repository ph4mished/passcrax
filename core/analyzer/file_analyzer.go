package analyzer

import (
    "fmt"
    "os"
	"strings"
	"bufio"
	"errors"
	"io/fs"
)


func FileAnalyze(hashFile string)string{
var result string
		    filename, err := os.Open(hashFile)
			if err != nil{
				fmt.Printf("%s[!] Error: Cannot open %s%s", bred, hashFile, rst)
			}
			if _, err = os.Stat(hashFile); errors.Is(err, fs.ErrNotExist){
				fmt.Printf("\n%s[!] Error: %s does not exist%s\n", bred, hashFile, rst)
				return ""
				}
			scanner := bufio.NewScanner(filename)
			for scanner.Scan(){
				hashLine := scanner.Text()
				hashLine = strings.TrimSpace(hashLine)
				fmt.Printf("\n%s", hashLine)
				if len(hashLine) != 0{
				result = PassAnalyze(hashLine)
				}
			
			}
			return result
		}

