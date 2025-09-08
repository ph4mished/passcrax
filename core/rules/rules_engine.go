package rules

import (
	"strconv"
	"strings"
)

//there will be improvisations where the parser and rule engine will support more commands on a line
//this code need shortening 

/*others to add are:
r = reverse,
l = lowercase all, ...Done
u = uppercase all. ..Done
D<n> = Delete char at position n,
[ = Delete first char, ..Done
] = Delete last char,   ...Done
{ = rotate the word to the left,
} = rotate the word to the right,
t = toggle case of all letters (small t),
C = lowercase first, capitalize rest(capital c), ..done
p<n> = repeat the word n times,
f = reflect the word (append reversed of word)
: = no change (passthrough)
*/

//for this one a lexer, parser and ast is required else, things are gonna get complicated
//note that sometimes a command wont be the always be the first. else the script will assume the next command is a strings. such as 'c $s^sad'  after capitalization the rest are ignore.
//I'll come back to it

//there will be improvisations where the parser and rule engine will support more commands on a line

//this function wouldn't do any splitting any more. the splitting will be done by the lexer
//so it will be like
//Let's peak(firstPar, secondPar, word string)
func LeetSpeak(realRune, repRune rune, word string) string {
    return strings.ReplaceAll(word, string(realRune), string(repRune)
}

func Prepend(prepRune rune, word string) string {
	return string(prepRune)+word
}


func Append(appRune rune,  word string) string {
	return word+string(appRune)
}

func Toggle(ruleString string, word string) string {
    //I'll surely shorten this when I'm less busy
	split := strings.Split(ruleString, "")
	num := split[1]
	rrune := []rune(word)
	number, _ := strconv.Atoi(num)
	if number > 0 {
		return ""
	}
	if len(word) == 0 {
		return ""
	}
	letter := rrune[number]
	char := string(letter)
	if char == strings.ToLower(char) {
		charset := strings.ToUpper(char)
		changeWord := strings.Replace(word, char, charset, 1)
		return changeWord
	} else {
		altered := strings.ToLower(char)
		newWord := strings.Replace(word, char, altered, 1)
		return newWord
	}
}

func Capitalize(word string) string {
	return strings.Title(word)
}



func CapitalizeAll(word string) string{
   return strings.ToTitle(word)
   }
   
   
   
func LowerAll(word) string {
return strings.ToLower(word)
}


func LowerFirstCapAll(word string) string {
 low := strings.ToLower(word[:1])
 return low+strings.ToTitle(word[1:])
}
   
   

func Duplicate(dCount int, word string) string {
	rep := ruleString * 2
	return strings.Repeat(word, rep)
}

func Reverse(word string) string {
    n := len(word)
    reversed := make([]rune, n)
    for _, r := range word {
        n--
        reversed[n] = r
    }
    return string(reversed[n:])
}

func DeleteFirst(word string) string {
  return word[1:]
}

func DeleteLast(word string) string {
  return word[:len(word)-1]
}
