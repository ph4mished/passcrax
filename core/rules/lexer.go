package rules

//import "rules/ruless"
import (
	"fmt"
	"strings"
	//  "bufio"
	//  "os"
)
//NEW LOGIC
//the Lexer that was be used for the charset parser will be used here. but it will need a counter to tell commands and characters apart. the counter restarts after each command 
//the counter knows 'r' command takes nothing, so it restarts to 0 and moves to the next. if at 0 and it meets 's', it means substitute, the counter reads to three and then restarts but moves on to the next until it meets EOF
//this new logic was made because the old one was space delimited, so it couldn't deal with chained commands 
//but with the new logic, if it sees a space it means the next logics that follows a for a new word. but if chained commands,then it means they all apply to one word

//OLD LOGICS
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
    /*others to add are: 
r = reverse,
l = lowercase all,
u = uppercase all
D<n> = Delete char at position n,
[ = Delete first char, 
] = Delete last char,
{ = rotate the word to the left,
} = rotate the word to the right,
t = toggle case of all letters (small t),
C = lowercase first, capitalize rest(capital c),
p<n> = repeat the word n times,
f = reflect the word (append reversed of word)
: = no change (passthrough)
*/
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
