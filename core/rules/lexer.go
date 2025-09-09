package rules


var handleTokens = map[rune]func(i *int, ruleRune []rune, word string) string {
    's': LeetSpeak,
    '^': Prepend,
    '$': Append,
    'c': Capitalize,
    'C': LowerFirstCapAll,
    //'T': TogglePosition,
    't': ToggleAll,
    'd': Duplicate,
    'r': Reverse,
    'l': LowerAll,
    'u': CapitalizeAll,
    '[': DeleteFirst,
    ']': DeleteLast,
   // 'D': DeletePosition,
    
   // '{': RotateLeft,
   // '}': RotateRight,
}



func ParseRules(ruleStr, word string) string {
    
    ruleRune := []rune(ruleStr)
    original := word
    result := ""
  
    for i := 0; i < len(ruleRune); i++ {
        
        //space should reset word
        if ruleRune[i] == ' '{
           result += word + " " 
           word = original
           continue 
           }


  if  command, exists := handleTokens[ruleRune[i]]; exists{
         
          word = command(&i, ruleRune, word)
}
}
result += word
return result
}



/*func main() {
    rule := "dsa@r ss&"
    word := "pass"
    fmt.Println(ParseRules(rule, word))
}*/
