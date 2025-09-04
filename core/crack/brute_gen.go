package crack

import "passcrax/core/utils"
import "github.com/fatih/color"

import (
	//	"fmt"
	"time"
)

var (
	bgrn  = color.New(color.FgGreen, color.Bold)
	bred  = color.New(color.FgRed, color.Bold)
	borng = color.New(color.FgHiYellow, color.Bold)
	bblu  = color.New(color.FgBlue, color.Bold)
	bcyn  = color.New(color.FgCyan, color.Bold)
	bylw  = color.New(color.FgYellow, color.Bold)
	grn   = color.New(color.FgGreen)
	red   = color.New(color.FgRed)
	//   orng = color.New(color.FgHiYellow)
	blu = color.New(color.FgBlue)
	cyn = color.New(color.FgCyan)
	ylw = color.New(color.FgYellow)
)

/*const (
	borng = "\033[1;38;5;208m"
	bgrn  = "\033[1;32m"
	bblu  = "\033[1;34m"
	bred  = "\033[1;31m"
	bylw  = "\033[1;33m"
	grn   = "\033[32m"
	blu   = "\033[34m"
	ylw   = "\033[33m"
	red   = "\033[31m"
	orng  = "\033[38;5;208m"
	rst   = "\033[0m"
)*/

// this function looks messy. it needs cleaning  but not that much cos it's still readable
func BruteGen(targetHash string, hashtype string, charSet string, startLen int, endLen int) string {
	var charset = ParseCharset(charSet)
	//brute range wont only support ranges. it will also support single int
	if startLen < 1 || endLen < 1 || startLen > endLen {
		bred.Printf("\n[!] Error: Invalid Length Parameters: Minimum length cannot be '0' !")
		return ""
	}
	if startLen == endLen {
		bylw.Printf("\n\n\n[~] Brute-Forcing Length %d ...", startLen)
	} else {
		bylw.Printf("\n\n\n[~] Brute-Forcing From Length %d To %d...", startLen, endLen)
	}
	startTime := time.Now()

	for length := startLen; length <= endLen; length++ {
		bblu.Printf("\n\n[+] Trying Length: %d\n", length)

		total := 1
		for i := 0; i < length; i++ {
			total *= len(charset)
		}

		for i := 0; i <= total; i++ {
			utils.PrintProgress(i, total, startTime)
			text := make([]rune, length)
			n := i
			for j := length - 1; j >= 0; j-- {
				text[j] = charset[n%len(charset)]
				n /= len(charset)
			}

			word := string(text)
			result :=  WordCrack(targetHash, word, hashtype)
if result != ""{
	                              return result
	                          }
			/*	hash_type, worked := utils.HashFormats(word, hashtype)
				           if !worked {
				               return ""
				               }
							if hash_type == targetHash {
								return word
							}*/
		}
	}
	return ""
}
