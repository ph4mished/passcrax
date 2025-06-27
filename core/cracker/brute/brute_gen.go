package brute

import "passcrax/core/utils"

import (
	"fmt"
	"time"
)

const (
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
)

//[]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.,;!?:'*#€@_&-+()/✓[]}{><∆§×÷=°^¢$¥£~|•√π`")

func BruteGen(targetHash string, hashtype string, charSet string, startLen int, endLen int) string {
	var charset = ParseCharset(charSet)
	if startLen < 1 || endLen < 1 || startLen > endLen {
		fmt.Printf("\n%s[!] Error: Invalid Length Parameters: Minimum length cannot be '0' !%s", bred, rst)
		return ""
	}

	fmt.Printf("\n\n\n%s[~] Brute-Forcing From Length %d To %d...%s", bylw, startLen, endLen, rst)
	startTime := time.Now()

	for length := startLen; length <= endLen; length++ {
		fmt.Printf("\n\n%s[+] Trying Length: %d%s\n", bblu, length, rst)

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
			hash_type, err := utils.HashFormats(word, hashtype)
			if err != nil {
				fmt.Printf("\n%s[!] Error: %s%s",
					bred, err, rst)
				return ""
			}

			if hash_type == targetHash {
				return word
			}
		}
	}
	return ""
}
