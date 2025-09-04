package cond

import (
	"fmt"
	"os"
//	"path/filepath"

	"passcrax/core/crack"
	"passcrax/core/file"

	"github.com/fatih/color"
)

var (
	bgrn = color.New(color.FgGreen, color.Bold)
	bred = color.New(color.FgRed, color.Bold)
	bblu = color.New(color.FgBlue, color.Bold)
	bcyn = color.New(color.FgCyan, color.Bold)
	bylw = color.New(color.FgYellow, color.Bold)
	grn  = color.New(color.FgGreen)
	red  = color.New(color.FgRed)
	blu  = color.New(color.FgBlue)
	cyn  = color.New(color.FgCyan)
	ylw  = color.New(color.FgYellow)
)

// This looks so much messy and unmaintainable
func HashConditions(targetHash string, hashtype string, mode string, charset string, dictDir string, startLen int, endLen int) string {
	if len(targetHash) == 0 {
		bred.Print("\n[!] Error: No hash set!")
		bgrn.Print(" Use ")
		bylw.Print("'set hash <hashstring>'")
	} else if len(hashtype) == 0 {
		bred.Print("\n[!] Error: No hash type set!")
		bgrn.Print(" Use ")
		bylw.Print("'set hashtype <value>'")
	} else if (len(charset) == 0) && (mode == "brute" || mode == "auto") {
		bred.Print("\n[!] Error: No brute characters set")
		bgrn.Print(" Use ")
		bylw.Print("'set charset <value>'")
		bgrn.Print("\n [+] The value should be enclosed in ")
		bylw.Print("square brackets")
		bgrn.Print(" followed by a ")
		bylw.Print("range of alphabets or numbers or special chars")
		bgrn.Print(" which are separated by ")
		bylw.Print("hyphens")
		bblu.Print("\n[+] Remember the spaces too. \n eg. ")
		bylw.Println("set charset [a-g A-R 0-5 &*^#]")
	} else if len(mode) == 0 {
		bred.Print("\n[!] Error: No mode set!")
		bgrn.Print(" Use ")
		bylw.Print("'set mode <value>'")
	}
	if mode == "brute" && len(targetHash) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		if endLen == 0 && startLen == 0 {
			bred.Println("\n[!] Error: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode")
			bgrn.Print(" Use ")
			bylw.Println("'set brute-range <min-max>'")
			return ""
		}
		brute_crx := crack.BruteGen(targetHash, hashtype, charset, startLen, endLen)
		if len(brute_crx) != 0 {
			bgrn.Println("\n\n[~] Password Found:")
			bblu.Println(brute_crx)
		} else {
			bred.Println("\n\n[!] Password Not Found!")
		}
	} else if mode == "dict" && len(targetHash) != 0 && len(hashtype) != 0 {
		if len(dictDir) != 0 {
			bgrn.Print("\n[~] Utilizing ")
			bylw.Printf("%s", dictDir)
			bgrn.Println(" for wordlist cracking")
			letgo := crack.PassCrack(dictDir, targetHash, hashtype, "")
			if len(letgo) != 0 {
				bgrn.Println("\n\n[~] Password Found:")
				bylw.Println(letgo)
				return ""
			} else {
				bred.Println("\n\n[!] Password Not Found!")
			}
		} else {
			bylw.Println("\n\n[!] No wordlist path defined!")
			bgrn.Println("[~] Falling back to default wordlists")
			nin := crack.PassCrack("", targetHash, hashtype, "")
			if len(nin) != 0 {
				bgrn.Println("\n\n[~] Password Found:")
				bylw.Println(nin)
				return ""
			} else {
				bred.Println("\n\n[!] Password Not Found!")
			}
		}
	} else if mode == "auto" && len(targetHash) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		var cracked_password string
		if len(dictDir) != 0 {
			bgrn.Print("\n[~] Utilizing ")
			bylw.Printf("%s", dictDir)
			bgrn.Println(" for wordlist cracking")
			cracked_password = crack.PassCrack(dictDir, targetHash, hashtype, "")
			if len(cracked_password) != 0 {
				bylw.Println(cracked_password)
				return ""
			} else {
				bred.Println("\n\n[!] Password Not Found!")
			}
		} else {
			bylw.Println("\n\n[!] No wordlist path defined!")
			bgrn.Println("[~] Falling back to default wordlists")
			cracked_password = crack.PassCrack("", targetHash, hashtype, "")
		}
		if cracked_password != "" {
		}
		if startLen != 0 && endLen != 0 {
			bylw.Println("\n\n[!] Password Not Found In Wordlist....")
			bylw.Printf("\n[~] Switching To Bruteforce (%d-%d characters)...\n", startLen, endLen)
		} else {
			bylw.Println("\n\n[!] Password Not Found In Wordlist....")
			bylw.Println("\n[~] Switching To Bruteforce (1-7 characters)...")
			startLen = 1
			endLen = 7
		}
		result := crack.BruteGen(targetHash, hashtype, charset, startLen, endLen)
		if len(result) != 0 {
			bgrn.Println("\n\n[~] Password Found:")
			bylw.Println(result)
			return ""
		} else {
			bred.Println("\n\n[!] Password Not Found!")
		}
		return ""
	}
	return ""
}

func FileConditions(hashFile string, hashtype string, mode string, charset string, dictDir string, startLen int, endLen int, outputFile string) string {
	if len(hashtype) == 0 {
		bred.Print("\n[!] Error: No hash type set!")
		bgrn.Print(" Use ")
		bylw.Print("'set hashtype <value>'")
	} else if len(charset) == 0 && (mode == "brute" || mode == "auto") {
		bred.Print("\n[!] Error: No brute characters set")
		bgrn.Print(" Use ")
		bylw.Print("'set charset <value>'")
		bgrn.Print("\n [+] The value should be enclosed in ")
		bylw.Print("square brackets")
		bgrn.Print(" followed by a ")
		bylw.Print("range of alphabets or numbers or special chars")
		bgrn.Print(" which are separated by ")
		bylw.Print("hyphens")
		bblu.Print("\n[+] Remember the spaces too. \n eg. ")
		bylw.Println("set charset [a-g A-R 0-5 &*^#]")
	} else if mode == "brute" && len(hashFile) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		if endLen == 0 && startLen == 0 {
			bred.Println("\n[!] Error: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode")
			bgrn.Print(" Use ")
			bylw.Print("'set brute-range <min-max>'\n")
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

		bgrn.Printf("\n[~] Cracking Completed: ")
		bblu.Printf("%d/%d", brtnumtop, totalbrtnum)
		bgrn.Printf("  passwords recovered", brtnumtop)

		bgrn.Print("\n[~] Results copied to ")
		bylw.Print(outputFile)
		bgrn.Print(" successfully!")
	} else if mode == "dict" && len(hashFile) != 0 && len(hashtype) != 0 {
		if len(dictDir) != 0 {
			bgrn.Print("\n[~] Utilizing ")
			bylw.Print(dictDir)
			bgrn.Println(" for wordlist cracking")
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

			bgrn.Printf("\n\n[~] Cracking Completed: ")
			bblu.Printf("%d/%d", dictnumtop, totaldictnum)
			bgrn.Print(" passwords recovered", )

			bgrn.Print("\n[~] Results copied to ")
			bylw.Print(outputFile)
			bgrn.Print(" successfully!")
		} else {
			bylw.Println("\n\n[!] No wordlist path defined!")
			bgrn.Println("[~] Falling back to default wordlists")
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
			bgrn.Printf("\n\n[~] Cracking Completed: ")
			bblu.Printf("%d/%d", dictnumtop, totaldictnum)
			bgrn.Print(" passwords recovered")

			bgrn.Print("\n[~] Results copied to ")
			bylw.Print(outputFile)
			bgrn.Print(" successfully!")
		}
	} else if mode == "auto" && len(hashFile) != 0 && len(hashtype) != 0 && len(charset) != 0 {
		if len(dictDir) != 0 {
			bgrn.Print("\n[~] Utilizing ")
			bblu.Print(dictDir)
			bgrn.Println(" for wordlist cracking")
			file.DictFile(dictDir, hashFile, hashtype, outputFile)
		} else {
			bylw.Println("\n\n[!] No wordlist path defined!")
			file.DictFile("", hashFile, hashtype, outputFile)
		}
		var num, allnum int
		var read_line string
		tmp_file := "tmp/passcrax_tmp.txt"
	//	dir := filepath.Dir(outputFile)
	//	addend := dir + tmp_file
		tmpFile, err := os.OpenFile(tmp_file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		defer tmpFile.Close()
		defer os.RemoveAll(tmp_file)
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

		bgrn.Printf("\n\n[~] ")
		bblu.Printf("%d/%d", numtop, totalnum)
		bgrn.Print(" passwords recovered")
		bgrn.Print("\n[~] ")
		bblu.Print(number)
		bgrn.Print(" hashes remain uncracked!")
		if startLen != 0 && endLen != 0 {
			bgrn.Println("\n[~] Bruteforce Range Defined")
			bylw.Printf("\n[~] Switching To Bruteforce (%d-%d characters)...\n", startLen, endLen)
		} else {
			bylw.Println("\n[~] Switching To Bruteforce (1-7 characters)...")
			startLen = 1
			endLen = 7
		}

		file.BruteFile(tmp_file, hashtype, charset, outputFile, startLen, endLen)
		bgrn.Print("\n[~] Results copied to ")
		bylw.Printf("%s", outputFile)
		bgrn.Print(" successfully!")
	} else {
		return ""
	}
	return ""
}
