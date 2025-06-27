package file

import (
	"passcrax/core/cracker"
	"passcrax/core/cracker/brute"
	"passcrax/core/utils"
	//"passcrax/core/rules"
)

import (
	"fmt"
	"os"
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

func DictFile(dict_dir string, hashFile string, hashtype string, outputFile string) string {
	filename := utils.FileLaunch(hashFile, os.O_RDWR, 0755)
	file, _ := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	header := fmt.Sprintf("\n\n\n###---These Are Results From %s---\n### Time: %v\n\n", hashFile, time.Now())
	file.Write([]byte(header))

	for _, targetHash := range filename {
		if true {
			var word_dir string
			if len(dict_dir) != 0 {
				word_dir = dict_dir
			} else {
				word_dir = ""
			}
			now := cracker.PassCrack(word_dir, targetHash, hashtype, "")
			results := fmt.Sprintf("%s : %s\n", targetHash, now)
			file.Write([]byte(results))
		}
	}
	return ""
}

func BruteFile(hashFile string, hashtype string, charset string, outputFile string, min int, max int) {
	filename := utils.FileLaunch(hashFile, os.O_RDWR, 0644)
	file, _ := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	header := fmt.Sprintf("\n\n\n###---These Are Results From %s---\n### Time: %v\n\n", hashFile, time.Now())
	file.Write([]byte(header))

	for _, targetHash := range filename {
		if true {
			non := brute.BruteGen(targetHash, hashtype, charset, min, max)
			bond := fmt.Sprintf("%s : %s\n", targetHash, non)
			file.Write([]byte(bond))
		}
	}
}
