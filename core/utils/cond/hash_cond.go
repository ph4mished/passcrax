package cond

//import "passcrax/core/utils"
import "passcrax/core/cracker"
import "passcrax/core/cracker/brute"
import "passcrax/core/file"

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	bcyn  = "\033[1;36m"
	bgrn  = "\033[1;32m"
	borng = "\033[1;38;5;208m"
	bblu  = "\033[1;34m"
	bred  = "\033[1;31m"
	bylw  = "\033[1;33m"
	grn   = "\033[32m"
	blu   = "\033[34m"
	ylw   = "\033[33m"
	red   = "\033[31m"
	orng  = "\033[38;5;208m"
	rst   = "\033[0m"
)

func HashConditions(targetHash string, hashtype string, mode string, charset string, dictDir string, startLen int, endLen int) string {
	if len(targetHash) == 0 {
		fmt.Printf("\n%s[!] Error: No hash set!%s %sUse %s'%sset hash <hashstring>%s'\n", bred, rst, bgrn, rst, bylw, rst)
	} else if len(hashtype) == 0 {
		fmt.Printf("\n%s[!] Error: No hash type set!%s %sUse %s '%sset hashtype <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
	} else if (len(charset) == 0) && (mode == "brute" || mode == "auto") {
		fmt.Printf("\n%s[!] Error: No brute characters set!%s %sUse %s'%sset charset <value>%s'\n%s [+] The value should be enclosed in%s %ssquare brakets%s %sfollowed by a%s %srange of alphabets or numbers or special chars%s %swhich are seperated by%s %shyphens%s.\n %s[+] Remember the spaces too%s\n %seg.%s %sset charset [a-g A-R 0-5 &*^#]%s\n", bred, rst, bgrn, rst, bylw, rst, bgrn, rst, bylw, rst, bgrn, rst, bylw, rst, bgrn, rst, bylw, rst, bgrn, rst, bgrn, rst, borng, rst)
	} else if len(mode) == 0 {
		fmt.Printf("\n%s[!] Error: No mode set!%s %sUse %s'%sset mode <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
	}
	if mode == "brute" && len(targetHash) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		if endLen == 0 && startLen == 0 {
			fmt.Printf("\n%s[!] Error: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode%s\n %sUse %s'%sset brute-range <min-max>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			return ""
		}
		brute_crx := brute.BruteGen(targetHash, hashtype, charset, startLen, endLen)
		if len(brute_crx) != 0 {
			fmt.Printf("\n\n%s[~] Password Found:%s %s%s%s\n", bgrn, rst, borng, brute_crx, rst)
		} else {
			fmt.Printf("\n\n%s[!] Password Not Found!%s\n", bred, rst)
		}
	} else if mode == "dict" && len(targetHash) != 0 && len(hashtype) != 0 {
		if len(dictDir) != 0 {
			letgo := cracker.PassCrack(dictDir, targetHash, hashtype, "")
			if len(letgo) != 0 {
				fmt.Printf("\n\n%s[~] Password Found:%s %s%s%s\n", bgrn, rst, borng, letgo, rst)
				return ""
			} else {
				fmt.Printf("\n\n%s[!] Password Not Found!%s\n", bred, rst)
			}
		} else {
			fmt.Printf("\n\n%s[!] No wordlist path defined!%s \n%s[~] Falling back to default wordlists%s\n", bylw, rst, bgrn, rst)
			nin := cracker.PassCrack("", targetHash, hashtype, "")
			if len(nin) != 0 {
				fmt.Printf("\n\n%s[~] Password Found:%s %s%s%s\n", bgrn, rst, borng, nin, rst)
				return ""
			} else {
				fmt.Printf("\n\n%s[!] Password Not Found!%s\n", bred, rst)
			}
		}
	} else if mode == "auto" && len(targetHash) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		var cracked_password string
		if len(dictDir) != 0 {
			cracked_password = cracker.PassCrack(dictDir, targetHash, hashtype, "")
			if len(cracked_password) != 0 {
				fmt.Printf("\n\n%s[~] Password Found:%s %s%s%s\n", bgrn, rst, borng, cracked_password, rst)
				return ""
			} else {
				fmt.Printf("\n\n%s[!] Password Not Found!%s\n", bred, rst)
			}
		} else {
			fmt.Printf("\n\n%s[!] No wordlist path defined!%s \n%s[~] Falling back to default wordlists%s\n", bylw, rst, bgrn, rst)
			cracked_password = cracker.PassCrack("", targetHash, hashtype, "")
		}
		if cracked_password != "" {
		}
		if startLen != 0 && endLen != 0 {
			fmt.Printf("\n%s[~] Switching To Bruteforce (%d-%dcharacters)...%s\n", bylw, startLen, endLen, rst)
		} else {
			fmt.Printf("\n\n%s[!] Password Not Found In Wordlist....%s\n\n%s[~] Switching To Bruteforce (1-7 characters)...%s\n", bylw, rst, bblu, rst)
			startLen = 1
			endLen = 7
		}
		result := brute.BruteGen(targetHash, hashtype, charset, startLen, endLen)
		if len(result) != 0 {
			fmt.Printf("\n\n%s[~] Password Found:%s %s%s%s\n", bgrn, rst, borng, result, rst)
			return ""
		} else {
			fmt.Printf("\n\n%s[!] Password Not Found!%s\n", bred, rst)
		}
		return ""
	}
	return ""
}

func FileConditions(hashFile string, hashtype string, mode string, charset string, dictDir string, startLen int, endLen int, outputFile string) string {
	if len(hashtype) == 0 {
		fmt.Printf("\n%s[!] Error: No hash type set!%s %sUse %s '%sset hashtype <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
	} else if len(charset) == 0 && (mode == "brute" || mode == "auto") {
		fmt.Printf("\n%s[!] Error: No brute characters set!%s %sUse %s'%sset charset <value>%s'\n %s[+] The value should be enclosed in%s %ssquare brakets%s %sfollowed by a%s %srange of alphabets or numbers or special chars%s %swhich are seperated by%s %shyphens%s\n %s[+] Remember the spaces too%s. \n %seg.%s %sset charset [a-g A-R 0-5 &*^#]%s\n", bred, rst, bgrn, rst, bylw, rst, bgrn, rst, bylw, rst, bgrn, rst, bylw, rst, bgrn, rst, bylw, rst, bgrn, rst, bgrn, rst, borng, rst)
	} else if mode == "brute" && len(hashFile) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		if endLen == 0 && startLen == 0 {
			fmt.Printf("\n%s[!] Error: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode%s\n %sUse %s'%sset brute-range <min-max>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			return ""
		}
		file.BruteFile(hashFile, hashtype, charset, outputFile, startLen, endLen)
		totalbrtLines, brtLines := file.FileRead(outputFile)
		var brtnum, allbrtnum int
		for brtnum, _ = range brtLines {
		}
		for allbrtnum, _ = range totalbrtLines {
		}

		brtnumber := brtnum + 1
		totalbrtnum := allbrtnum + 1
		brtnumtop := totalbrtnum - brtnumber

		fmt.Printf("\n\n%s[~] Cracking Completed: %s%s%d/%d%s%s passwords recovered!%s\n", bgrn, rst, borng, brtnumtop, totalbrtnum, rst, bgrn, rst)
		fmt.Printf("\n%s[~] Results copied to %s%s%s %ssuccessfully!%s\n", bgrn, bylw, outputFile, rst, bgrn, rst)
	} else if mode == "dict" && len(hashFile) != 0 && len(hashtype) != 0 {
		if len(dictDir) != 0 {
			file.DictFile(dictDir, hashFile, hashtype, outputFile)
			var dictnum, alldictnum int
			totaldictLines, readDictLines := file.FileRead(outputFile)
			for dictnum, _ = range readDictLines {
			}

			for alldictnum, _ = range totaldictLines {
			}
			dictnumber := dictnum + 1
			totaldictnum := alldictnum + 1
			dictnumtop := totaldictnum - dictnumber

			fmt.Printf("\n\n%s[~] Cracking Completed: %s%s%d/%d%s%s passwords recovered!%s\n", bgrn, rst, borng, dictnumtop, totaldictnum, rst, bgrn, rst)
			fmt.Printf("\n%s[~] Results copied to %s%s%s %ssuccessfully!%s\n", bgrn, bylw, outputFile, rst, bgrn, rst)
		} else {
			fmt.Printf("\n\n%s[!] No wordlist path defined!%s \n%s[~] Falling back to default wordlists%s\n", bylw, rst, bgrn, rst)
			file.DictFile("", hashFile, hashtype, outputFile)
			var dictnum, alldictnum int
			totaldictLines, readDictLines := file.FileRead(outputFile)
			for dictnum, _ = range readDictLines {
			}

			for alldictnum, _ = range totaldictLines {
			}
			dictnumber := dictnum + 1
			totaldictnum := alldictnum + 1
			dictnumtop := totaldictnum - dictnumber

			fmt.Printf("\n\n%s[~] Cracking Completed: %s%s%d/%d%s%s passwords recovered!%s\n", bgrn, rst, borng, dictnumtop, totaldictnum, rst, bgrn, rst)
			fmt.Printf("\n%s[~] Results copied to %s%s%s %ssuccessfully!%s\n", bgrn, bylw, outputFile, rst, bgrn, rst)
		}
	} else if mode == "auto" && len(hashFile) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		if len(dictDir) != 0 {
			cracked_hashline := file.DictFile(dictDir, hashFile, hashtype, outputFile)
			if cracked_hashline != "" {
			}
		} else {
			fmt.Printf("\n\n%s[!] No wordlist path defined!%s \n%s[~] Falling back to default wordlists%s\n", bylw, rst, bgrn, rst)
			cracked_hashline := file.DictFile("", hashFile, hashtype, outputFile)
			if cracked_hashline != "" {
			}
		}
		var num, allnum int
		var read_line string
		tmp_file := "/tmp.txt"
		dir := filepath.Dir(outputFile)
		addend := dir + tmp_file
		tmpFile, err := os.OpenFile(addend, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		defer tmpFile.Close()
		defer os.RemoveAll(addend)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		totalLines, readLines := file.FileRead(outputFile)
		for num, read_line = range readLines {
			tue := fmt.Sprintf("%s\n", read_line)
			tmpFile.Write([]byte(tue))
		}

		for allnum, _ = range totalLines {
		}
		number := num + 1
		totalnum := allnum + 1
		numtop := totalnum - number

		fmt.Printf("\n%s[~] %s%s%d/%d%s%s passwords recovered!%s\n", bgrn, rst, borng, numtop, totalnum, rst, bgrn, rst)
		fmt.Printf("\n%s[~] %s%s%d%s %shashes remain uncracked!%s\n", bgrn, rst, borng, number, rst, bgrn, rst)
		if startLen != 0 && endLen != 0 {
			fmt.Printf("\n%s[~] Switching To Bruteforce (%d-%dcharacters)...%s\n", bylw, startLen, endLen, rst)
		} else {
			fmt.Printf("\n%s[~] Switching To Bruteforce (1-7characters)...%s\n", bylw, rst)
			startLen = 1
			endLen = 7
		}

		file.BruteFile(addend, hashtype, charset, outputFile, startLen, endLen)
		fmt.Printf("\n\n%s[~] Results copied to %s%s%s %ssuccessfully!%s\n", bgrn, bylw, outputFile, rst, bgrn, rst)
		//}//else if mode != "dict" && mode != "brute" && mode != "auto"{
		//mode = ""
		//validMode := "dict, brute, auto"
		//fmt.Printf("\n%s[!] Inputted Mode Is Invalid!%s \n %sThese are the list of accepted modes to use for file hashes:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, validMode, rst)
	} else {
		return ""
	}
	return ""
}
