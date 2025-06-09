package rules

import (
 "fmt"
 "os"
 "bufio"
 "strings"
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
