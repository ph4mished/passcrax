package analyzer

//import "passcrax/core/utils"
import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func FileAnalyze(hashFile string) string {
	/*
	   var numlines int
	   var hashline, result string
	   var same []string
	   //boolan := make(map[string]bool)
	   allLines := utils.FileLaunch(hashFile, os.O_RDWR, 0755)
	   //lets use the length of lines to keep them together
	   //put the hashes together by length and check their hashtypes
	   //this is to reduce the spamming in the hashid output

	   	   for numlines, hashline = range allLines{
	   	   part := len(hashline)
	   	   if len(hashline) == part {

	   	   //result = PassAnalyze(hashline)
	   	   //if !boolan[result] {
	   	  // if result == result{
	   	  // boolan[result] = true
	   	   fmt.Println("Hashline: ", len(hashline))
	   	   same = append(same, hashline)
	   	   }
	   	  // }
	   	   }
	   	   numlines = numlines + 1
	   	  // if len(hashline) != 0{
	   				//	fmt.Printf("There are %d lines after analysis\n", len(result))
	   					//}
	   					fmt.Printf("hashfile lines: %d\n", numlines)
	   					fmt.Println("these matched: ", same)
	   			        //fmt.Println(result)

	   					//this should printout the percentage of lines with the same hashtype
	*/
	var result string
	filename, err := os.Open(hashFile)
	defer filename.Close()
	if err != nil {
		fmt.Printf("%s[!] Error: Cannot open %s%s", bred, hashFile, rst)
	}
	if _, err = os.Stat(hashFile); errors.Is(err, fs.ErrNotExist) {
		fmt.Printf("\n%s[!] Error: %s does not exist%s\n", bred, hashFile, rst)
		return ""
	}
	scanner := bufio.NewScanner(filename)
	for scanner.Scan() {
		hashLine := scanner.Text()
		hashLine = strings.TrimSpace(hashLine)
		fmt.Printf("\n%s", hashLine)
		if len(hashLine) != 0 {
			result = PassAnalyze(hashLine)
		}

	}
	return result
}
