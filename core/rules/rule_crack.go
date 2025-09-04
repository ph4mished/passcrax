package rules

import (
	//	"bufio"
	//	"fmt"
	//	"os"
	"strings"
)

func FindRuleWord(ruleString, word string) string {
	mangledWord := Lexer(strings.TrimSpace(ruleString), word)
	//is empty check really of relevance??
	//if len(mangledWord) != 0 {
	//return mangledWord
	//	}
	//}
	return mangledWord
}
