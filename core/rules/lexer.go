package rules
import (
    "strings"
)
func Lexer(ruleWord string, word string)(string){

        ruleWord = strings.TrimSpace(ruleWord)
    
        if len(ruleWord) == 0 {
         return ""
        }
    
        command := ruleWord[0]
       switch command {
       case 's':
       return LeetSpeak(ruleWord, word)
        
       case '$':
   return Append(ruleWord, word)

       case '#':
   return ""

       case '^':
    return Prepend(ruleWord, word)

       case 'd':
    return Duplicate(ruleWord, word)

       case 'c':
    return Capitalize(ruleWord, word)

       case 'T':
    return Toggle(ruleWord, word)

       case ':':
   return ""

       default:
   return ""

      }
          return ""
}

