package brute

import (
	"strings"
	"unicode"
)

func ParseCharset(charset string) []rune {
	charset = strings.Trim(charset, "[]")
	tokens := strings.Split(charset, " ")
	seen := make(map[rune]bool)

	var allChars []rune
	for _, token := range tokens {
		if strings.Contains(token, "-") && len([]rune(token)) == 3 {
			runes := []rune(token)
			startRune, endRune := runes[0], runes[2]

			if (unicode.IsLetter(startRune) && unicode.IsLetter(endRune)) || (unicode.IsDigit(startRune) && unicode.IsDigit(endRune)) {
				for r := startRune; r <= endRune; r++ {
					if !seen[r] {
						seen[r] = true
						allChars = append(allChars, r)
					}
				}
			} else {
				for _, uniSymbol := range token {
					if !seen[uniSymbol] {
						seen[uniSymbol] = true
						allChars = append(allChars, uniSymbol)
					}
				}
			}

		} else {
			for _, symbol := range token {
				if !seen[symbol] {
					seen[symbol] = true
					allChars = append(allChars, symbol)
				}
			}
		}
	}
	return allChars
}
