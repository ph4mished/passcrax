package cracker

import "PassCrax/core/utils"

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PassCrack(targetHash string, hashtype string) string {

	var wordlist_dir = "Wordlist/"
	wordlist_files, err := filepath.Glob(filepath.Join(wordlist_dir, "*.txt"))
	if err != nil {
		fmt.Printf("\n%sError Scanning Wordlist Directory: %v %s\n", red, err, rst)
		return ""
	}
	if len(wordlist_files) == 0 {
		fmt.Printf("\n%sError: No Files Found In %s%s\n", red, wordlist_dir, rst)
		return ""
	}
	for _, filename := range wordlist_files {

		fmt.Printf("\n%sScanning File: %s...%s", bblu, filename, rst)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("\n%sError: File Cannot Be Opened!%s\n", red, rst)
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := scanner.Text()
			word = strings.TrimSpace(word)

			hash_type, err := utils.HashFormats(word, hashtype)
			if err != nil {
				fmt.Printf("\n%sError: %s%s", red, err, rst)
				file.Close()
				return ""
			}

			if hash_type == targetHash {
				fmt.Printf("\n%sPassword Found:%s %s%s%s\n", bgrn, rst, borng, word, rst)
				file.Close()
				return word
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("\n%sError Reading File: %v %s\n", red, err, rst)
		}
		file.Close()
	}
	fmt.Printf("\n%sPassword Not Found!%s\n", bred, rst)
	return ""
}
