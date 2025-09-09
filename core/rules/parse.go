package rules

import (
	"strings"
)


//this parser accepts the token and does the mangling 
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
t = toggle case of all letters (small t), ...Done
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
func LeetSpeak(i *int, ruleRune []rune, word string) string {
     
    newWord := strings.ReplaceAll(word, string(ruleRune[*i+1]), string(ruleRune[*i+2]))
    *i += 2
    return newWord
}

func Prepend(i *int, ruleRune []rune, word string)  string {
	newWord := string(ruleRune[*i+1])+word
       *i += 1
    return newWord
}


func Append(i *int, ruleRune []rune, word string)  string {
	newWord :=  word+string(ruleRune[*i+1])
     *i += 1
    return newWord
}

func ToggleAll(i *int, ruleRune []rune, word string)  string {
   var  allRunes []rune
    for _, char := range word {
        if char >= 'a' && char <= 'z' {
            allRunes = append(allRunes, char-32)
        } else if char >= 'A' && char <= 'Z' {
            allRunes = append(allRunes, char+32)
        } else {
            allRunes = append(allRunes, char)
        }
    }
    return string(allRunes)
}

func Capitalize(i *int, ruleRune []rune, word string)  string {
	return strings.Title(word)
    
}



func CapitalizeAll(i *int, ruleRune []rune, word string)  string{
   return strings.ToTitle(word)
   }
   
   
   
func LowerAll(i *int, ruleRune []rune, word string)  string {
return strings.ToLower(word)
}


func LowerFirstCapAll(i *int, ruleRune []rune, word string)  string {
 low := strings.ToLower(word[:1])
 return low+strings.ToTitle(word[1:])
}
   
   

func Duplicate(i *int, ruleRune []rune, word string)  string {
    start := *i
        for *i < len(ruleRune) && ruleRune[*i] == 'd' {
            *i++
        }
        dCount := *i - start 
        rep := dCount * 2
        *i-- //dcrement to avoid skipping the char pointed that's not 'd'
	return strings.Repeat(word, rep)
	
}

func Reverse(i *int, ruleRune []rune, word string)  string {
    n := len(word)
    reversed := make([]rune, n)
    for _, r := range word {
        n--
        reversed[n] = r
    }
    return string(reversed[n:])
}

func DeleteFirst(i *int, ruleRune []rune, word string)  string {
  return word[1:]
}

func DeleteLast(i *int, ruleRune []rune, word string)  string {
  return word[:len(word)-1]
}
