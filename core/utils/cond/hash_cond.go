package cond

import "passcrax/core/cracker/file"
import "passcrax/core/cracker"

import (
"fmt"
"path/filepath"
"os"
)


const (
    bcyn = "\033[1;36m"
	bgrn = "\033[1;32m"
	bblu = "\033[1;34m"
	bred = "\033[1;31m"
	bylw = "\033[1;33m"
	grn = "\033[32m"
	blu = "\033[34m"
	ylw = "\033[33m"
	red = "\033[31m"
	orng = "\033[38;5;208m"
	rst = "\033[0m"
)

func HashConditions(targetHash string, hashtype string, mode string, startLen int, endLen int)string{
 if len(targetHash) == 0 {
				fmt.Printf("\n%s[!] Error: No hash set!%s %sUse %s'%sset hash <hashstring>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			} else if len(hashtype) == 0 {
				fmt.Printf("\n%s[!] Error: No hash type set!%s %sUse %s '%sset hashtype <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			}else if len(mode) == 0 {
			fmt.Printf("\n%s[!] Error: No mode set!%s %sUse %s'%sset mode <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			}
			  if mode == "brute" && len(targetHash) != 0 && len(hashtype) != 0 {
				if endLen == 0 && startLen == 0 {
					fmt.Printf("\n%s[!] Error: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode%s\n %sUse %s'%sset brute-range <min-max>%s'\n", bred, rst, bgrn, rst, bylw, rst)
				//	continue
				}
				cracker.BruteGen(targetHash, hashtype, startLen, endLen)
			}else if mode == "dict" && len(targetHash) != 0 && len(hashtype) != 0 { 
				cracker.PassCrack(targetHash, hashtype, "")
			} else if mode == "auto" && len(targetHash) != 0 && len(hashtype) != 0 {
				cracked_password := cracker.PassCrack(targetHash, hashtype, "")
				if cracked_password != "" {
				}
				fmt.Printf("\n%s[!] Password Not Found In Wordlist....%s\n%s[~] Switching To Bruteforce (1-7 characters)...%s\n", bylw, rst, bblu, rst)
				startLen = 1
				endLen = 7
				cracker.BruteGen(targetHash, hashtype, startLen, endLen)
			}else {
				return ""
			}
			return ""
			}
			
			

func FileConditions(hashFile string, hashtype string, mode string, startLen int, endLen int, outputFile string )string{
if len(hashtype) == 0 {
			 fmt.Printf("\n%s[!] Error: No hash type set!%s %sUse %s '%sset hashtype <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
				}else if mode == "brute" && len(hashFile) != 0 && len(hashtype) != 0 {
			  if endLen == 0 && startLen == 0 {
					fmt.Printf("\n%s[!] Error: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode%s\n %sUse %s'%sset brute-range <min-max>%s'\n", bred, rst, bgrn, rst, bylw, rst)
				return ""
				}
				//with the brutefile cracking, the ouput file should first be checked to see if its empty before it starts cracking
				//if it contains some hashes, then those were the output of the dict file. it should check that which wasn't cracked and crack it. this avoids wasting time cracking a hash whose result has already being found and continues with those which weren't cracked.
				file.BruteFile(hashFile, hashtype, outputFile, startLen, endLen)
				fmt.Printf("\n%s[~] Results copied to %s%s%s %ssuccessfully!%s\n", bgrn, bylw, outputFile, rst, bgrn, rst)
			} else if mode == "dict" && len(hashFile) != 0 && len(hashtype) != 0 {
			file.DictFile(hashFile, hashtype, outputFile)
			fmt.Printf("\n%s[~] Results copied to %s%s%s %ssuccessfully!%s\n", bgrn, bylw, outputFile, rst, bgrn, rst)
			}else if mode == "auto" && len(hashFile) != 0 && len(hashtype) != 0 {
				cracked_hashline := file.DictFile(hashFile, hashtype, outputFile)
				if cracked_hashline != "" {
				}
				fmt.Printf("\n%s[!] Password Not Found In Wordlist....%s\n%s[~] Switching To Bruteforce (1-7characters)...%s\n", bylw, rst, bblu, rst)
				startLen = 1
				endLen = 7
				tmp_file := "/tmp.txt"
				dir := filepath.Dir(outputFile)
				addend := dir+tmp_file
				tmpFile, err := os.OpenFile(addend, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
				if err != nil {
				fmt.Println("Error: ", err)
				}
				readLines := file.FileRead(outputFile)
				for _, read_line := range readLines{
				tue := fmt.Sprintf("%s\n", read_line)
				tmpFile.Write([]byte(tue))
				}
				file.BruteFile(addend, hashtype, outputFile, startLen, endLen)
		        fmt.Printf("\n%s[~] Results copied to %s%s%s %ssuccessfully!%s\n", bgrn, bylw, outputFile, rst, bgrn, rst)
			}else if mode != "dict" && mode != "brute" && mode != "auto"{
			mode = ""
			validMode := "dict, brute, auto"
			fmt.Printf("\n%s[!] Inputted Mode Is Invalid!%s \n %sThese are the list of accepted modes to use for file hashes:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, validMode, rst)
			 }else{
			return ""
			}
			return ""
		}
