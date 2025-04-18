package analyzer

import (
	"fmt"
	"regexp"
	"strings"
)

var HASH_PATTERNS = map[string]*regexp.Regexp{
	"MD5":        regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"SHA-1":      regexp.MustCompile(`^(?i)[a-f0-9]{40}$`),
	"SHA-224":    regexp.MustCompile(`^(?i)[a-f0-9]{56}$`),
	"SHA-256":    regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),
	"SHA-384":    regexp.MustCompile(`^(?i)[a-f0-9]{96}$`),
	"SHA-512":    regexp.MustCompile(`^(?i)[a-f0-9]{128}$`),
	"NTLM":       regexp.MustCompile(`^(?i)[A-F0-9]{32}$`),
	"LM Hash":    regexp.MustCompile(`^(?i)[A-F.0-9]{32}$`),
	"MySQL v3+":  regexp.MustCompile(`^(?i)[a-f0-9]{16}$`),
	"MySQL v5+":  regexp.MustCompile(`^(?i)\*[A-F0-9]{40}$`),
	"bcrypt":     regexp.MustCompile(`^(?i)\$2[ayb]\$.{56}$`),
	"Argon2":     regexp.MustCompile(`^(?i)\$argon2[a-z]+\$.+`),
	"DES (Unix)": regexp.MustCompile(`^(?i).{13}$`),
}

const (
	grn = "\033[32m"
	blu = "\033[34m"
	ylw = "\033[33m"
	red = "\033[31m"
	rst = "\033[0m"
)

func findHash(hash string) string {

	if len(hash) == 0 {
		return fmt.Sprintf("%sInput Must Not Be Empty!%s\n", red, rst)
	}

	hash = strings.TrimSpace(hash)

	var matches []string
	for name, regex := range HASH_PATTERNS {
		if regex.MatchString(hash) {
			matches = append(matches, name)
		}
	}
	if len(matches) > 0 {
		return strings.Join(matches, " || ")
	}
	return fmt.Sprintf("%sUnknown Hash Format!%s\n", red, rst)
}

func PassAnalyze(hash string) {
	result := findHash(hash)
	fmt.Printf("%s\nHash Type:%s %s%s%s\n", grn, rst, ylw, result, rst)
}
