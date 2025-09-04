package rules

//import "rules/ruless"
import (
	"fmt"
	"strings"
	//  "bufio"
	//  "os"
)

// i'm still figuring out a way to loop over the ones which should be ignore whereby they will leave no spaces in the output. i mean those such as comments.
func Lexer(ruleWord string, word string) string {
	//my strategy
	//this lexer takes a line of a rule file and tokenizes it
	//it then divides each input per its command eg. substitute requires a length of three 'sa@'
	//it then sends these segregated data to the engine to be cracked
	//meaning this file plays the role of a parser and tokenizer

	ruleWord = strings.TrimSpace(ruleWord)
	//to ignore empty lines
	if len(ruleWord) == 0 {
		return ""
	}
	//pick up the word and look at the first letter. ie. the command
	command := ruleWord[0]
	switch command {
	case 's':
		return LeetSpeak(ruleWord, word)

	case '$':
		return Append(ruleWord, word)

	case '#':
		//  continue
		return ""

	case '^':
		return Prepend(ruleWord, word)

	case 'd':
		return Duplicate(ruleWord, word)

	case 'c':
		return Capitalize(word)

	case 'T':
		return Toggle(ruleWord, word)

	case ':':
		return ""

	default:
		return ""

	}
	fmt.Print("Found Nothing")
	return ""
}
