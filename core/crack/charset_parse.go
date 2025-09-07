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
                
                        if (unicode.IsLetter(prevRune) && unicode.IsLetter(nextRune)) || (unicode.IsDigit(prevRune) && unicode.IsDigit(nextRune)) {
                                for r := prevRune; r <= nextRune; r++ {
                                        if !seen[r] {
                                                seen[r] = true
                                                allChars = append(allChars, r)
                                        }
                                }
                        }else {
                                for _, uniSymbol := range charset {
                                        if !seen[uniSymbol] {
                                                seen[uniSymbol] = true
                                                allChars = append(allChars, uniSymbol)
                                        }
                                }
                        }

            } else {
                                if !seen[ch] {
                                        seen[ch] = true
                                        allChars = append(allChars, ch)
                                }
            }
        }
        return allChars
}
