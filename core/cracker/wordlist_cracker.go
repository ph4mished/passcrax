package cracker

import ("passcrax/core/utils"
        "passcrax/core/rules"
)

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
var hash_type string
var err error



func PassCrack(targetHash string, hashtype string, rulefile string) string {
//the hardcoded wordlist makes it brittle

	var wordlist_dir = "Wordlists/"
	wordlist_files, err := filepath.Glob(filepath.Join(wordlist_dir, "*.txt"))
	if err != nil {
		fmt.Printf("\n%s[!] Error Scanning Wordlist Directory: %v %s\n", red, err, rst)
		return ""
	}
	if len(wordlist_files) == 0 {
		fmt.Printf("\n%s[!] Error: No Files Found In %s%s\n", red, wordlist_dir, rst)
		return ""
	}
	for _, filename := range wordlist_files {

		fmt.Printf("\n%s[~] Scanning File: %s...%s", bblu, filename, rst)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("\n%s[!] Error: File Cannot Be Opened!%s\n", red, rst)
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := scanner.Text()
			word = strings.TrimSpace(word)

			if rulefile != ""{
				rule_words := rules.FindRuleWord(rulefile, word)
				   for _, rule_word := range rule_words{
				   hash_type, err := utils.HashFormats(rule_word, hashtype)
			if err != nil {
				fmt.Printf("\n%s[!] Error: %s%s", red, err, rst)
				       file.Close()
				return ""
			}
			if hash_type == targetHash {
				fmt.Printf("\n%s[~] Password Found:%s %s%s%s\n", bgrn, rst, borng, rule_word, rst)
				file.Close()
				return rule_word
			}
			}
			}
			
				hash_type, err = utils.HashFormats(word, hashtype)
			if err != nil {
				fmt.Printf("\n%s[!] Error: %s%s", red, err, rst)
                file.Close()
				return ""
			}

			if hash_type == targetHash {
				fmt.Printf("\n%s[~] Password Found:%s %s%s%s\n", bgrn, rst, borng, word, rst)
				file.Close()
				return word
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("\n%sError Reading File: %v %s\n", red, err, rst)
		}
		file.Close()
	}
	fmt.Printf("\n%s[!] Password Not Found!%s\n", bred, rst)
	return ""
}
