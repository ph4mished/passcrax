package utils

import (
//	"fmt"
	"regexp"
	"strings"
)

func levenshtein(a, b string) int {
	la := len(a)
	lb := len(b)

	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
	}

	for i := 0; i <= la; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			dp[i][j] = min(
                //delete
				dp[i-1][j]+1,
                //insert
				dp[i][j-1]+1,
                //replace or subtitute
				dp[i-1][j-1]+cost,
			)
		}
	}

	return dp[la][lb]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

type Command struct {
	Name  string         
	Regex *regexp.Regexp
}

var commands = []Command{
	{"set hash", regexp.MustCompile(`(?i)^set\s+hash\s+(.+)$`)},
	{"set hashtype", regexp.MustCompile(`(?i)^set\s+hashtype\s+(.+)$`)},
	{"identify", regexp.MustCompile(`(?i)^identify\s+(.+)$`)},
	{"set mode", regexp.MustCompile(`(?i)^set\s+mode\s+(.+)$`)},
	{"set charset", regexp.MustCompile(`(?i)^set\s+charset\s+(.+)$`)},
	{"set brute-range", regexp.MustCompile(`^(?i)set\s+brute-range\s+\d+\s*-\s*\d+$`)},
	{"load rulefile", regexp.MustCompile(`^(?i)load\s+rulefile\s+(.+)$`)},
	{"drop rulefile", regexp.MustCompile(`^(?i)drop\s+rulefile\s*$`)},
	{"load dictdir", regexp.MustCompile(`^(?i)load\s+dictdir\s+(.+)$`)},
	{"load hashfile", regexp.MustCompile(`^(?i)load\s+hashfile\s+(.+)$`)},
	{"set outputfile", regexp.MustCompile(`^(?i)set\s+outputfile\s+(.+)$`)},
}

func DidYouMean(input string) string {
	best := ""
	bestDist := 999
	for _, cand := range commands {
		d := levenshtein(strings.ToLower(input), strings.ToLower(cand.Name))
		//fmt.Printf("Distance between %-12q and %-12q = %d\n", input, cand.Name, d)
		if d < bestDist {
			bestDist = d
			best = cand.Name
		//	return best
		}
	}
	return best
}
