package rules

import(
//"passcrax/core/utils"
)
import (
 "fmt"
 "os"
 "bufio"
 "strings"
 //"path/filepath"
 )
 
 const (
    borng = "\033[1;38;5;208m"
    bgrn = "\033[1;32m"
	bblu = "\033[1;34m"
	bred = "\033[1;31m"
	bylw = "\033[1;33m"
	grn = "\033[32m"
	blu = "\033[34m"
	ylw = "\033[33m"
	red = "\033[31m"
	orng = "\033[38;5;208m"
	rst = "\033[0m"
)


/*func DictRule(targetHash string, hashtype string, ruleFile string)string{
var hash_type string
var err error
//the hardcoded wordlist makes it brittle
var word string
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
			word = scanner.Text()
			word = strings.TrimSpace(word)
            rule_words := FindRuleWord(ruleFile, word)
            for _, rule_word := range rule_words{
            fmt.Println(rule_word)
            hash_type, err = utils.HashFormats(rule_word, hashtype)
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

		if err := scanner.Err(); err != nil {
			fmt.Printf("\n%sError Reading File: %v %s\n", red, err, rst)
			}
		file.Close()
}
fmt.Printf("\n%s[!] Password Not Found!%s\n", bred, rst)
	return ""
}*/



func FindRuleWord(ruleFile string, word string)[]string{
var ruleWords []string
  filename, err := os.OpenFile(ruleFile, os.O_RDWR, 0644)
  if err != nil {
  fmt.Printf("\n%s[!] Error: %s%s", bred, err, rst)
  }
  defer filename.Close()

  scanner := bufio.NewScanner(filename)
  for scanner.Scan(){
      rule := scanner.Text()
      rule = strings.TrimSpace(rule)
        rule_word := Lexer(rule, word)
        if len(rule_word) != 0{
        ruleWords = append(ruleWords, rule_word)
        }
}
  return ruleWords
}
