package cracker

import (
	"passcrax/core/rules"
	"passcrax/core/utils"
	//"passcrax/core/file"
)

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
	"unicode"
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

var hash_type string
var err error

func PassCrack(dict_dir string, targetHash string, hashtype string, rulefile string) string {
	var wordlist_dir string
	if len(dict_dir) != 0 {
		wordlist_dir = dict_dir
	} else {
		wordlist_dir = "Wordlists/"
	}
	wordlist_files, err := filepath.Glob(filepath.Join(wordlist_dir, "*.txt"))
	if err != nil {
		fmt.Printf("\n%s[!] Error Scanning Wordlist Directory: %v %s\n", bred, err, rst)
		return ""
	}
	if len(wordlist_files) == 0 {
		fmt.Printf("\n%s[!] Error: No Files Found In %s%s\n", bred, wordlist_dir, rst)
		return ""
	}
	for fileNum, filename := range wordlist_files {
		startTime := time.Now()

		results := utils.FileLaunch(filename, 0, 0644)
		for _, word := range results {
			hash_type, err = utils.HashFormats(word, hashtype)
			if err != nil {
				fmt.Printf("\n%s[!] Error: %s%s", bred, err, rst)
				return ""
			}
			for _, hashChar := range targetHash {
				if unicode.IsUpper(hashChar) {
					altHashtype := strings.ToUpper(hash_type)
					if altHashtype == targetHash {
						return word
					}
				}
				if hash_type == targetHash {
					return word
				}
			}

			if rulefile != "" {
				rule_words := rules.FindRuleWord(rulefile, word)
				for _, rule_word := range rule_words {
					hash_type, err := utils.HashFormats(rule_word, hashtype)
					if err != nil {
						fmt.Printf("\n%s[!] Error: %s%s", bred, err, rst)
						return ""
					}

					if hash_type == targetHash {
						return rule_word
					}
				}
			}

		}
		currentFileNum := fileNum + 1
		utils.PrintProgress(currentFileNum, len(wordlist_files), startTime)
	}
	return ""
}
