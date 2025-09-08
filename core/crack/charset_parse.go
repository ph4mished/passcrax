package crack

import (
        "strings"
        "unicode"
)


func ParseCharset(charsetStr string) []rune {
        var allChars []rune
    seen := make(map[rune]bool)
        charsetStr = strings.Trim(charsetStr, "[]")

    charset := []rune(charsetStr)
        for i, ch := range charset {
        var none rune = 0
           prevRune := none
            nextRune := none
            if i > 0 {
                prevRune = charset[i-1]
            }
            if i < len(charset)-1 {
                nextRune = charset[i+1]
            }

            if ch == '-' && prevRune != none && nextRune != none {
                //this for ranges, range parsing starts here
                
                        if (unicode.IsLetter(prevRune) && unicode.IsLetter(nextRune)) || (unicode.IsDigit(prevRune) && unicode.IsDigit(nextRune)) {
                            //valid rages
                                for r := prevRune; r <= nextRune; r++ {
                                        if !seen[r] {
                                                seen[r] = true
                                                allChars = append(allChars, r)
                                        }
                                }
                        }else {
                            //if not a valid range, assume it's a literal.
                                for _, uniSymbol := range charset {
                                        if !seen[uniSymbol] {
                                                seen[uniSymbol] = true
                                                allChars = append(allChars, uniSymbol)
                                        }
                                }
                        }

            } else {
                //else assume it's a literal
                                if !seen[ch] {
                                        seen[ch] = true
                                        allChars = append(allChars, ch)
                                }
            }
        }
        return allChars
}
