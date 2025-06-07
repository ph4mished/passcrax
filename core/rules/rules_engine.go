package rules
import (
  "strconv"
    "strings"
)
//for this one a lexer, parser and ast is required else, things are gonna get complicated
//note that sometimes a command wont be the always be the first. else the script will assume the next command is a strings. such as 'c $s^sad'  after capitalization the rest are ignore.
//I'll come back to it

//there will be improvisations where the parser and rule engine will support more commands on a line

func LeetSpeak(ruleString string, word string)string{
     //if the string that follows is "s" it might be perceived as a command and it will be skipped
  //this isnt good
    var replace string
    split := strings.Split(ruleString, "")
    from := split[1]
    to := split[2]

    replace = strings.ReplaceAll(word, from, to)
    return replace
return replace
}

func Prepend(ruleString string, word string)string{
    var preppend string  
 split := strings.Split(ruleString, "^")
 if len(split) > 1 {
   join := strings.Join(split, "")
    reSplit := strings.Split(join, " ")
     reJoin := strings.Join(reSplit, "")
result := reJoin+word
return result
}else{
    split = strings.Split(ruleString, "")
        first := split[1]

        preppend = first+word
}
    return preppend
}


func Append(ruleString string, word string)string{
    
    var appnd string
    ruleString = strings.TrimSpace(ruleString)
   split := strings.Split(ruleString, "$")
   if len(split) > 1 {
     join := strings.Join(split, "")
     reSplit := strings.Split(join, " ")
     reJoin := strings.Join(reSplit, "")
  result := word+reJoin
  return result
  }else{

    split = strings.Split(ruleString, "")
        last := split[1]

        appnd = word+last
}
    return appnd
}


func Toggle(ruleString string, word string)string{
  split := strings.Split(ruleString, "")
  num := split[1]
rrune := []rune(word)
number, _ := strconv.Atoi(num)
 if number > 0{
    return ""
  }
  if len(word) == 0 {
  return ""
  }
letter := rrune[number]
char := string(letter)
if  char == strings.ToLower(char) {
charset := strings.ToUpper(char)
changeWord:= strings.Replace(word, char, charset, 1)
return changeWord
}else{
  altered := strings.ToLower(char)
  newWord := strings.Replace(word, char, altered, 1)
return newWord
}
}


func Capitalize(ruleString string, word string)string{
    var capital string
  capital = strings.Title(word)
    return capital
}

func Duplicate(ruleString string, word string)string{
    var repeat string
    strings.Split(ruleString, "")
    number := strings.Count(ruleString, "d")
    rep := number * 2
    repeat = strings.Repeat(word, rep)
    return repeat
}
