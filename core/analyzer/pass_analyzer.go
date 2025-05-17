package analyzer

import (
	"fmt"
	"regexp"
	"strings"
)

//Contribution Sake: Hash algorithms should be grouped by family order to avoid redundancy (repetition) of algorithms and make it easier to find related hashes.


//Problem: A better and robust approach should be addressed because shows output in unordered list. most likely hashed should be shown at the top and least likely, at the bottom
var unknown string
var HASH_PATTERNS = map[string]*regexp.Regexp{

	//BLAKE Family
	"BLAKE2b":    regexp.MustCompile(`^[a-f0-9]{128}$`),
	"BLAKE2s":    regexp.MustCompile(`^[a-f0-9]{64}$`),
	"BLAKE3-128": regexp.MustCompile(`^[a-f0-9]{32}$`),
	"BLAKE3-256": regexp.MustCompile(`^[a-f0-9]{64}$`),
	"BLAKE3-512": regexp.MustCompile(`^[a-f0-9]{128}$`),

	//Checksums
	"Adler-32":    regexp.MustCompile(`^(?i)[a-f0-9]{8}$`),
	"CRC-16":       regexp.MustCompile(`^(?i)[a-f0-9]{4}$`),
	"CRC-16-CCITT": regexp.MustCompile(`^(?i)[a-f0-9]{4}$`),
	"CRC-24":       regexp.MustCompile(`^(?i)[a-f0-9]{6}$`),
	"CRC-32":       regexp.MustCompile(`^(?i)[a-f0-9]{8}$`),
	"CRC-32B":      regexp.MustCompile(`^(?i)[a-f0-9]{8}$`),
	"CRC-64":       regexp.MustCompile(`^(?i)[a-f0-9]{16}$`),
	"CRC-96-ZIP":   regexp.MustCompile(`^(?i)[a-f0-9]{24}$`),

	//Database Hashes
	"MySQL v3+": regexp.MustCompile(`^[A-F0-9]{16}$`),
	"MySQL v5+": regexp.MustCompile(`^(?i)\*[A-F0-9]{40}$`),

	//EDON-R Family
	"EDON-R-256": regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),
	"EDON-R-512": regexp.MustCompile(`^(?i)[a-f0-9]{128}$`),

	//FNV Family
	"FNV-1a-32":  regexp.MustCompile(`^(?i)[a-f0-9]{8}$`),
	"FNV-1a-64":  regexp.MustCompile(`^(?i)[a-f0-9]{16}$`),
	"FNV-1a-128": regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"FNV-1-32":   regexp.MustCompile(`^(?i)[a-f0-9]{8}$`),
	"FNV-1-64":   regexp.MustCompile(`^(?i)[a-f0-9]{16}$`),
	"FNV-1-128":  regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),

	//Gost Family
	"Gost":           regexp.MustCompile(`^[a-f0-9]{64}$`),
	"Gost-CryptoPro": regexp.MustCompile(`^[a-f0-9]{64}$`),

	//Haval Family
	"Haval-128": regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"Haval-160": regexp.MustCompile(`^(?i)[a-f0-9]{40}$`),
	"Haval-192": regexp.MustCompile(`^(?i)[a-f0-9]{48}$`),
	"Haval-224": regexp.MustCompile(`^(?i)[a-f0-9]{56}$`),
	"Haval-256": regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),

	//MD Family
	"MD2":                     regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"MD4":                     regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"MD5":                     regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"Half MD5":                regexp.MustCompile(`^(?i)[a-f0-9]{16}$`),
	"Double MD5":              regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"MD5-Crypt":               regexp.MustCompile(`^\$1\$[a-zA-Z0-9./]{8}\$[a-zA-Z0-9./]{22}$`),
	"Cisco-IOS-MD5":           regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"FreeBSD-MD5":             regexp.MustCompile(`^\$1\$[a-zA-Z0-9./]{8}\$[a-zA-Z0-9./]{22}$`),
	"MD5-APR":                 regexp.MustCompile(`^\$apr1\$[a-zA-Z0-9./]{8}\$[a-zA-Z0-9./]{22}$`),
	"Apache-MD5":              regexp.MustCompile(`^\$apr1\$[a-zA-Z0-9./]{8}\$[a-zA-Z0-9./]{22}$`),

	//Murmur Family
	"Murmur3-32":  regexp.MustCompile(`^(?i)[a-f0-9]{8}$`),
	"Murmur3-128": regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),

	//Password Hashes
	"bcrypt":     regexp.MustCompile(`^\$2[abxy]\$\d{2}\$[./A-Za-z0-9]{53}$`),
	"Argon2":     regexp.MustCompile(`^\$argon2[a-z]+\$.+`),
	"DES (Unix)": regexp.MustCompile(`^[./A-Za-z0-9]{13}$`),

	//RIPEMD Family
	"RIPEMD-128": regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"RIPEMD-160": regexp.MustCompile(`(?i)^[a-f0-9]{40}$`),
	"RIPEMD-256": regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),
	"RIPEMD-320": regexp.MustCompile(`^(?i)[a-f0-9]{80}$`),

	//SHA Family
	"SHA-1":        regexp.MustCompile(`^[a-f0-9]{40}$`),
	"SHA-1-Base64": regexp.MustCompile(`^[a-zA-Z0-9+/]{28}={0,2}$`),
	"SHA-224":      regexp.MustCompile(`^[a-f0-9]{56}$`),
	"SHA-256":      regexp.MustCompile(`^[a-f0-9]{64}$`),
	"SHA-384":      regexp.MustCompile(`^[a-f0-9]{96}$`),
	"SHA-512":      regexp.MustCompile(`^[a-f0-9]{128}$`),
	"SHA3-224":     regexp.MustCompile(`^[a-f0-9]{56}$`),
	"SHA3-256":     regexp.MustCompile(`^[a-f0-9]{64}$`),
	"SHA3-384":     regexp.MustCompile(`^[a-f0-9]{96}$`),
	"SHA3-512":     regexp.MustCompile(`^[a-f0-9]{128}$`),

	//SNEFRU Family
	"SNEFRU-128": regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"SNEFRU-256": regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),

	//Skein Family
	"Skein256-128":   regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"Skein256-256":   regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),
	"Skein512-128":   regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"Skein512-256":   regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),
	"Skein512-512":   regexp.MustCompile(`^(?i)[a-f0-9]{128}$`),
	"Skein512-1024":  regexp.MustCompile(`^(?i)[a-f0-9]{256}$`),
	"Skein1024-408":  regexp.MustCompile(`^(?i)[a-f0-9]{102}$`),
	"Skein1024-1024": regexp.MustCompile(`^(?i)[a-f0-9]{256}$`),

	//Tiger Family
	"Tiger-128":   regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"Tiger-160":   regexp.MustCompile(`^(?i)[a-f0-9]{40}$`),
	"Tiger-192":   regexp.MustCompile(`^(?i)[a-f0-9]{48}$`),
	"Tiger-128,3": regexp.MustCompile(`^(?i)[a-f0-9]{32}$`),
	"Tiger-160,3": regexp.MustCompile(`^(?i)[a-f0-9]{40}$`),
	"Tiger-192,3": regexp.MustCompile(`^(?i)[a-f0-9]{48}$`),

	//Windows Hashes
	"NTLM": regexp.MustCompile(`^[A-F0-9]{32}$`),
	"LM":   regexp.MustCompile(`^[A-F0-9]{32}$`),

	//Whirlpool Family
	"Whirlpool":     regexp.MustCompile(`^(?i)[a-f0-9]{128}$`),
	"Whirlpool-1":   regexp.MustCompile(`^(?i)[a-f0-9]{128}$`),
	"Whirlpool-2":   regexp.MustCompile(`^(?i)[a-f0-9]{128}$`),
	"Whirlpool-224": regexp.MustCompile(`^(?i)[a-f0-9]{56}$`),
	"Whirlpool-256": regexp.MustCompile(`^(?i)[a-f0-9]{64}$`),
	"Whirlpool-384": regexp.MustCompile(`^(?i)[a-f0-9]{96}$`),

	"FCS-16": regexp.MustCompile(`^(?i)[a-f0-9]{4}$`),
	"FCS-32": regexp.MustCompile(`^(?i)[a-f0-9]{8}$`),
}

const (
    bred = "\033[1;31m"
	grn = "\033[32m"
	blu = "\033[34m"
	ylw = "\033[33m"
	red = "\033[31m"
	bcyn = "\033[1;36m"
	orng = "\033[38;5;208m"
	rst = "\033[0m"
)

func FindHash(targetHash string) string {

	targetHash = strings.TrimSpace(targetHash)

	var matches []string
	for name, regex := range HASH_PATTERNS {
		if regex.MatchString(targetHash) {
			matches = append(matches, name)
		}
	}
	if len(matches) > 0 {
		return strings.Join(matches, "\n[+] ")
	}
	unknown = fmt.Sprintf("%sUnknown Hash Format!%s\n", red, rst)
	return unknown
}

func PassAnalyze(targetHash string){
	result := FindHash(targetHash)
	if result == unknown {
	fmt.Printf("\n%s[+]Unknown Hash Format!%s\n", bred, rst)
	}else if strings.ContainsAny(result, " [+] ") {
		fmt.Printf("%s\nPOSSIBLE HASH TYPES%s %s\n[+] %s%s\n", bcyn, rst, orng, result, rst)
	} else {
		fmt.Printf("%s\nHASH TYPE:%s %s%s%s\n", bcyn, rst, orng, result, rst)
	}
}
