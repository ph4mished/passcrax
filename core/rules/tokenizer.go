package main

import "fmt"

func ParseRules(ruleStr, word string) {
    //all 'return' here in this script are only for testing, it will be removed. 
    ruleRune := []rune(ruleStr)
    fmt.Println("All: ", string(ruleRune))

    for i := 0; i < len(ruleRune); i++ {

//for prepending 
    if ruleRune[i] == '^'{
    if i+1 >= len(ruleRune) {
                    fmt.Printf("Error: '%s' command incomplete\n", string(ruleRune[i]))
                    break
                }
                Prepend(ruleRune[i+1], word)
                i += 1
}

//for appending
//for single paraméter based commands
    if ruleRune[i] == '$'{
                // Check if we have enough characters for parameters
                if i+1 >= len(ruleRune) {
                    fmt.Printf("Error: '%s' command incomplete\n", string(ruleRune[i]))
                    break
                }
                Append(ruleRune[i+1], word)

                // Skip ahead by 1 to avoid processing parameters as separate commands
                i += 1
            }



//for two parameter based commands
//leet 
        if ruleRune[i] == 's' {
            // Check if we have enough characters for parameters
            if i+2 >= len(ruleRune) {
                fmt.Println("Error: '%s' command incomplete", string(ruleRune[i]))
                break
            }

            // Process the complete 's' command with its parameters
          LeetSpeak(ruleRune[i+1], ruleRune[i+2], word)

            // Skip ahead by 2 to avoid processing parameters as separate commands
            i += 2
        }
        //single non parameter needed commands will be
        //defined here
       
    //d should be read to see until no d is seen anymore.
    //eg. ddddf
    //it should stop at f and return all d as rulestring
     /*  if ruleRune[i] == 'd' && ruleRune[i+1] != 'd' {
        Duplicate(
    }*/
    
    if ruleRune[i] == 'd' {
        // Found start of 'd' run - find how long it is
        start := i
        for i < len(ruleRune) && ruleRune[i] == 'd' {
            i++ // Keep moving while we see 'd'
        }
        // Now i points to first non-'d' character
        dCount := i - start // Number of consecutive 'd's
        
        // Process the duplicate rule with dCount
        Duplicate(dCount, word)
    } else {
        i++ // Move to next character for other rules
    }
    
    if ruleRune[i] == 'c' {
    Capitalize(word)
}

        if ruleRune[i] == '[' {
        DeleteFirst(word)
    }
    
        if ruleRune[i] == ']' {
        DeleteLastWord
    } 
    
            if ruleRune[i] == 'u' {
           CapitalizeAll(word)
           }
           
            if ruleRune[i] == 'l' {
            LowerAll(word)
        }
        
        if ruleRune[i] == 'C' {
            LowerFirstCapAll(word)    
            }
            
            
            if ruleRune[i] == 't' {
            Reversed(word)    
            }

        if ruleRune[i] == 't' || ruleRune[i] == '{' || ruleRune[i] == '}' || ruleRune[i] == 'r'{
            
                fmt.Printf("\nFound %s comand\n", string(ruleRune[i]))
        }

    }
}

func main() {
    rule := "  s d@se4d t r d"
    word := "pass"
    ParseRules(rule, word)
}package main

import "fmt"

func ParseRules(ruleStr, word string) {
    ruleRune := []rune(ruleStr)
    fmt.Println("All: ", string(ruleRune))

    for i := 0; i < len(ruleRune); i++ {

    if ruleRune[i] == '^' || ruleRune[i] == '$'{
                // Check if we have enough characters for parameters
                if i+1 >= len(ruleRune) {
                    fmt.Printf("Error: '%s' command incomplete\n", string(ruleRune[i]))
                    break
                }

                // Process the complete 's' command with its parameters
                fmt.Printf("Found '%s' command:\n", ruleRune[i])
                fmt.Println("  First parameter:", string(ruleRune[i+1]))

                // Skip ahead by i to avoid processing parameters as separate commands
                i += 1
            }


        if ruleRune[i] == 's' {
            // Check if we have enough characters for parameters
            if i+2 >= len(ruleRune) {
                fmt.Println("Error: '%s' command incomplete", string(ruleRune[i])
                break
            }

            // Process the complete 's' command with its parameters
            fmt.Println("Found 's' command:")
            fmt.Println("  s:", string(ruleRune[i]))
            fmt.Println("  First parameter:", string(ruleRune[i+1]))
            fmt.Println("  Second parameter:", string(ruleRune[i+2]))

            // Skip ahead by 2 to avoid processing parameters as separate commands
            i += 2
        }
        //single non parameter needed commands will be
        //defined here
        if ruleRune[i] == 'd' || ruleRune[i] == 'c' || ruleRune[i] == 't' || ruleRune[i] == '{' || ruleRune[i] == '}' || ruleRune[i] == '[' || ruleRune[i] == ']' || ruleRune[i] == 'u' || ruleRune[i] == 'l' || ruleRune[i] == 'r'{
                //d needs no char
                fmt.Printf("\nFound %s comand\n", string(ruleRune[i]))
        }

        //some functions need only one parameter
    }
}

func main() {
    rule := "  s d@se4d t r d"
    word := "pass"
    ParseRules(rule, word)
}