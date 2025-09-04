package analyzer

import (
//	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	s_bblu = color.New(color.FgBlue, color.Bold).SprintFunc()
	s_bvol = color.New(color.FgMagenta, color.Bold).SprintFunc()
	s_bylw = color.New(color.FgYellow, color.Bold).SprintFunc()
	bgrn   = color.New(color.FgGreen, color.Bold)
	bred   = color.New(color.FgRed, color.Bold)
	borng  = color.New(color.FgHiYellow, color.Bold)
	bblu   = color.New(color.FgBlue, color.Bold)
	bcyn   = color.New(color.FgCyan, color.Bold)
	bylw   = color.New(color.FgYellow, color.Bold)
	grn    = color.New(color.FgGreen)
	red    = color.New(color.FgRed)
	orng   = color.New(color.FgHiYellow)
	blu    = color.New(color.FgBlue)
	cyn    = color.New(color.FgCyan)
	ylw    = color.New(color.FgYellow)
)

func PassAnalyze(targetHash string) string {
	targetHash = strings.TrimSpace(targetHash)

	var matches []string
	for regex, HashTypes := range HASH_DATABASE {
		if regex.MatchString(targetHash) {
			for _, hashtype := range HashTypes {
				name := "\n" + s_bblu("  "+hashtype.Name) + s_bvol("\n    PassCrax Format: ") + s_bylw(hashtype.PassCrax)
				matches = append(matches, name)
			}
		}
	}
	if len(matches) > 0 {
		return strings.Join(matches, "\n ")
	} else {
		return bred.Sprint("[!] Unknown Hash Format!\n")
	}
}
