package crack

import (
	"passcrax/core/rules"
	"passcrax/core/utils"
	// "passcrax/core/rules"
)

//import "github.com/fatih/color"

import (
	"bufio"
	//	"fmt"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

/*const (
	bcyn  = cyan.Add(color.Bold)
	borng = orange.Add(color.Bold)
	bgrn  = green.Add(color.Bold)
	bblu  = blue.Add(color.Bold)
	bred  = red.Add(color.Bold)
	bylw  = yellow.Add(color.Bold)
    )*/

var hash_type string
var err error

//for external wordlist, it will check if the wordlist is utf-8 encoded

// this one will no more accept files because accepting files make it more complicated if rules are used and also adding more. cracking methods likd hybrid names this function more verbose including it's parameters.
// instead it will accept strings (wordlist string). this will be able to accept altered (rule altered/ hybrid altered) wordlist strings so rules and hybrid will be externally handle.
// This will ensure this function does one thing and does it well
// this function looks messy. it needs cleaning
// func DictDirLaunch(dict_dir string, targetHash string, hashtype string, ruleFile string) string {
func PassCrack(dict_dir string, targetHash string, hashtype string, ruleFile string) string {
	var wordlist_dir string
	if len(dict_dir) != 0 {
		wordlist_dir = dict_dir
	} else {
		wordlist_dir = "Wordlists/"
	}
	wordlist_files, err := filepath.Glob(filepath.Join(wordlist_dir, "*.txt"))
	if err != nil {
		bred.Printf("\n[!] Error Scanning Wordlist Directory: %v\n", err)
		return ""
	}
	if len(wordlist_files) == 0 {
		bred.Printf("\n[!] Error: No Files Found In %s\n", wordlist_dir)
		return ""
	}
	for fileNum, filename := range wordlist_files {
		startTime := time.Now()

		//the use of this function should be aborted because it loads all the wordlist contents into memory. it's prone to OOM if the file is larger than the ram (eg. rockyou.txt)
		//	results := utils.FileLaunch(filename, 0, 0644)
		file, err := os.Open(filename)

		if err != nil {
			bred.Printf("\n[!] Error: %v", err)
		}
		defer file.Close()
		if _, err = os.Stat(filename); errors.Is(err, fs.ErrNotExist) {
			bred.Printf("\n[!] Error: %s does not exist\n", filename)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := strings.TrimSpace(scanner.Text())
			if word != "" {
				result := WordCrack(targetHash, word, hashtype)
if result != ""{
	                              return result
	                           }
	                           }

			//now rules can be easily handled here
			if ruleFile != "" {
				file, err := os.Open(ruleFile)

				if err != nil {
					bred.Printf("\n[!] Error: %v", err)
				}
				defer file.Close()
				if _, err = os.Stat(ruleFile); errors.Is(err, fs.ErrNotExist) {
					bred.Printf("\n[!] Error: %s does not exist\n", ruleFile)
				}
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					ruleLine := strings.TrimSpace(scanner.Text())
					if ruleLine != "" {
						mangledWord := rules.FindRuleWord(ruleLine, word)
						result := WordCrack(targetHash, mangledWord, hashtype)
						if result != ""{
							return result
						}
					}
				}
			}

		}
		//lemme check if this will still work
		currentFileNum := fileNum + 1
		utils.PrintProgress(currentFileNum, len(wordlist_files), startTime)
	}
	return ""
}

// this function is plug and play. it makes wordlist mangling with rules and hybrid simple.
// this function could be used by bruteforce too
func WordCrack(targetHash, word, hashtype string) string {
	hash_type, worked := utils.HashFormats(word, hashtype)
	if !worked {
		return ""
	}
	for _, hashChar := range targetHash {
		if unicode.IsUpper(hashChar) {
			if strings.ToUpper(hash_type) == targetHash {
				return word
			}
		}
		if hash_type == targetHash {
			return word
		}
	}
	return ""
}
