package main

//above all the code will be kept simple and concise
import (
	"passcrax/core/analyzer"
	"passcrax/core/utils"
    "passcrax/core/utils/cond"
	"passcrax/core/cracker" 
	"passcrax/core/cracker/file"
)

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"io/fs"
	"path/filepath"
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



var (
	reSetHash     = regexp.MustCompile(`(?i)^set\s+hash\s+(.+)$`)
	reSetHashtype = regexp.MustCompile(`(?i)^set\s+hashtype\s+(.+)$`)
	reHashid      = regexp.MustCompile(`(?i)^hashid\s+(.+)$`)
	reSetMode     = regexp.MustCompile(`(?i)^set\s+mode\s+(.+)$`)
	reBruteRange       = regexp.MustCompile(`^(?i)set\sbrute-range\s\d+\s*-\s*\d+$`)
	reSetRules   = regexp.MustCompile(`^(?i)load\s+rulefile\s+(.+)$`)
	reUnsetRules   = regexp.MustCompile(`^(?i)drop\s+rulefile\s*$`)
	reSetInput    = regexp.MustCompile(`^(?i)load\s+hashfile\s+(.+)$`)
    reSetOutput    = regexp.MustCompile(`^(?i)set\s+outputfile\s+(.+)$`)
    )

var targetHash, hashtype, mode, hashFile, ruleFile, outputFile string
	var startLen, endLen int
	var usingHashFile bool
	
	func ifEmpty(if_full, if_null string) string {
	if if_full == "" {
		return if_null
	}
	return if_full
}

func Status(){
	fmt.Printf("\n%sCURRENT SETTINGS%s", bcyn, rst)
	fmt.Printf("\n%sHash%s: %s%s%s", grn, rst, ylw, ifEmpty(targetHash, "Not Set"), rst)
	fmt.Printf("\n%sHash Type%s: %s%s%s", grn, rst, ylw, ifEmpty(hashtype, "Not Set"), rst)
	fmt.Printf("\n%sRule File%s: %s%s%s", grn, rst, ylw, ifEmpty(ruleFile, "Not Set {Optional}"), rst)
	fmt.Printf("\n%sMode%s: %s%s%s", grn, rst, ylw, ifEmpty(mode, "Not Set\n"), rst)
	if mode == "brute" {
	fmt.Printf("\n%sBrute Min Length%s: %s%d%s", grn, rst, ylw, startLen, rst)
	fmt.Printf("\n%sBrute Max Length%s: %s%d%s\n", grn, rst, ylw, endLen, rst)
	}
}

func FileStatus(){
    fmt.Printf("\n%sCURRENT FILE HASH SETTINGS%s", bcyn, rst)
	fmt.Printf("\n%sHash File%s: %s%s%s", grn, rst, ylw, ifEmpty(hashFile, "Not Set"), rst)
	fmt.Printf("\n%sOutput File%s: %s%s%s", grn, rst, ylw, ifEmpty(outputFile, "Not Set"), rst)
	fmt.Printf("\n%sHash Type%s: %s%s%s", grn, rst, ylw, ifEmpty(hashtype, "Not Set"), rst)
	fmt.Printf("\n%sMode%s: %s%s%s", grn, rst, ylw, ifEmpty(mode, "Not Set\n"), rst)
	if mode == "brute" {
	fmt.Printf("\n%sBrute Min Length%s: %s%d %s", grn, rst, ylw, startLen, rst)
	fmt.Printf("\n%sBrute Max Length%s: %s%d%s\n", grn, rst, ylw, endLen, rst)
	}
}



func main(){
 
	
	scanner := bufio.NewScanner(os.Stdin)

	// Ctrl+C handling for user to exit
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		fmt.Printf("\n   %sProgram Terminated!%s\n", bred, rst)
		os.Exit(0)
	}()
	
	utils.Banner()
	Status()

	for {
START:
		fmt.Print("\n> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch {
		case input == "exit":
			fmt.Printf("\n   %sProgram Terminated!%s\n", bred, rst)
			fmt.Println("\n")
			return
		
		case input == "help":
			utils.Help()
		
        case input == "status":
			if usingHashFile {
			FileStatus()
			}else{
			Status()
			}
			
			
		case reSetHash.MatchString(input):
		hashFile = ""
            usingHashFile = false
			targetHash = strings.TrimSpace(reSetHash.FindStringSubmatch(input)[1])
			fmt.Printf("\n%s[~] Hash set to:%s %s%s%s\n", grn, rst, bylw, targetHash, rst)


		case reSetHashtype.MatchString(input):
			hashtype = strings.TrimSpace(reSetHashtype.FindStringSubmatch(input)[1])
			hashtype = strings.ToLower(hashtype)
			if hashtype == analyzer.CheckValidHashType(hashtype){
			fmt.Printf("\n%s[~] Hash Type set to:%s %s%s%s\n", grn, rst, bylw, hashtype, rst)
			}else{
			fmt.Printf("\n%s[!] Hashtype Value Is Invalid Or Unsupported!%s \n %s[~] These are the list of supported inputs to use:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, analyzer.CheckValidHashType(hashtype), rst)
			fmt.Printf("\n%s[~] Use '%shashid <hashstring>%s%s' if you don't know what hashtype to use%s\n", bgrn, bylw, rst, bgrn, rst)
			  hashtype = ""
			}

		case reHashid.MatchString(input):
			targetHash = strings.TrimSpace(reHashid.FindStringSubmatch(input)[1])
			//hashfiles may not always be 'txt'. 
			if filepath.Ext(targetHash) ==".txt" {
			usingHashFile = true
		     hashFile  = targetHash
			if len(analyzer.FileAnalyze(hashFile)) == 0{
         	fmt.Printf("\n%s[!] Error: %s is an empty file%s", bred, hashFile, rst)
         	hashFile = ""
			outputFile = ""
			continue
			}else{
			fmt.Printf("\n%s[[[ %sEnd Of File of '%s%s%s%s'%s %s]]]%s\n", bblu, bcyn, bgrn, hashFile, rst, bcyn, rst, bblu, rst )
			fmt.Printf("\n%s[~] Hash File set to:%s %s%s%s\n", grn, rst, bylw, hashFile, rst)
			}
	 fileDir := filepath.Dir(hashFile)
	 fileName := filepath.Base(hashFile)
	dot := strings.Index(fileName, ".")
	var name, ext string
if dot != -1 {
	name = fileName[:dot]
	ext = fileName[dot:]
	PRACK:
	alterName := file.IterName()
	   alteredFile := fmt.Sprintf("%s_%s%s", name, alterName, ext)
	   outputFile = filepath.Join(fileDir, alteredFile)
	   
	    _, err := os.Stat(outputFile)
	   if err == nil {
	   goto 
	   PRACK
	   }
	   fmt.Printf("\n%s[~] Output file set to:%s %s%s%s\n", grn, rst, bylw, outputFile , rst)
	 }
		}else{
			usingHashFile = false
			analyzer.PassAnalyze(targetHash)
			fmt.Printf("\n%s[~] Hash set to:%s %s%s%s\n", grn, rst, bylw, targetHash, rst)
			}
			
			
		case reSetMode.MatchString(input):
			mode = strings.TrimSpace(reSetMode.FindStringSubmatch(input)[1])
			 mode = strings.ToLower(mode)
			if mode == analyzer.CheckValidMode(mode){
			fmt.Printf("\n%s[~] Mode set to:%s %s%s%s\n", grn, rst, bylw, mode, rst)
			}else{
			fmt.Printf("\n%s[!] Mode Value Is Invalid!%s \n %s[~] These are the list of accepted modes to use:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, analyzer.CheckValidMode(mode), rst)
			  mode = ""
			  }
			

		case reBruteRange.MatchString(input):
		downcase := strings.ToLower(input)
     	trimStart := strings.TrimPrefix(downcase, "set brute-range")
		num := strings.Split(trimStart, "-")
		startNum := strings.TrimSpace(num[0])
	    endNum := strings.TrimSpace(num[1])
	    
	value, err := strconv.Atoi(startNum)
	if err != nil{
	fmt.Printf("%s[!] Error: %s%s", bred, err, rst)
	}
	startLen = value
	
	val, err := strconv.Atoi(endNum)
	if err != nil{
	fmt.Printf("%s[!] Error: %s%s", bred, err, rst)
	}
	endLen = val
	
	if startLen >= endLen {
	fmt.Printf("\n%s[!] Error: Minimum length cannot be greater than Maximum length%s\n", bred, rst)
	startLen = 0
	endLen = 0
	    }else{   
	fmt.Printf("\n%s[~] Brute Min Length set to:%s %s%d%s\n", grn, rst, bylw, startLen, rst)
	fmt.Printf("\n%s[~] Brute Max Length set to:%s %s%d%s\n", grn, rst, bylw, endLen, rst)
	}
	
       case reSetInput.MatchString(input):
           usingHashFile = true
           targetHash = ""
	   hashFile = strings.TrimSpace(reSetInput.FindStringSubmatch(input)[1])
	     _, err := os.Stat(hashFile)
	    if err != nil{
	   fmt.Printf("\n%s[!] Error: %s does not exist%s", bred, hashFile, rst)
	   fmt.Printf("\n%s[+] Cross check if it's not a file path typographical error%s\n", bylw, rst)
	   hashFile = ""
	   continue
	   }
	   fmt.Printf("\n%s[~] Hash file set to:%s %s%s%s\n", grn, rst, bylw, hashFile, rst)
	    fileDir := filepath.Dir(hashFile)
	 fileName := filepath.Base(hashFile)
	dot := strings.Index(fileName, ".")
	var name, ext string
	if dot != -1 {
	name = fileName[:dot]
	ext = fileName[dot:]
	PACK:
	alterName := file.IterName()
	   alteredFile := fmt.Sprintf("%s_%s%s", name, alterName, ext)
	   outputFile = filepath.Join(fileDir, alteredFile)
	   
	    _, err := os.Stat(outputFile)
	   if err == nil {
	   goto PACK
	   }
	   fmt.Printf("\n%s[~] Output file set to:%s %s%s%s\n", grn, rst, bylw, outputFile , rst)
	 }
	 
	   case reSetOutput.MatchString(input):
	   usingHashFile = true
	    outputFile = strings.TrimSpace(reSetOutput.FindStringSubmatch(input)[1])
	   fmt.Printf("\n%s[~] Output file set to:%s %s%s%s\n", grn, rst, bylw, outputFile, rst)
	   
	   case reSetRules.MatchString(input):
	   ruleFile = strings.TrimSpace(reSetRules.FindStringSubmatch(input)[1])
	   if _, err := os.Stat(ruleFile); errors.Is(err, fs.ErrNotExist){
	   fmt.Printf("\n%s[!] Error: %s does not exist%s", bred, ruleFile, rst)
	   fmt.Printf("\n%s[+] Cross check if it's not a file path typographical error%s\n", bylw, rst)
	   ruleFile = ""
	   continue
	   }
	    fmt.Printf("\n%s[~] Rule file set to:%s %s%s%s\n", grn, rst, bylw, ruleFile, rst)


	case reUnsetRules.MatchString(input):
		if len(ruleFile) == 0 {
			fmt.Printf("%s[!] Error:%sNo rule file loaded!%s\n", bred, bylw, rst)
		}else{
		ruleFile = ""
		fmt.Printf("\n%s[~] Rule file dropped successfully!%s\n", bgrn, rst)
		}
	   
	   
	  
		case input == "run":
			//hash file cracking
		if len(hashFile) != 0 && len(outputFile) != 0 {
		cond.FileConditions(hashFile, hashtype, mode, startLen, endLen, outputFile)
		}else if len(ruleFile) != 0 && len(hashFile) == 0 && usingHashFile == false{
		//rule file 
			if len(targetHash) == 0 {
			fmt.Printf("\n%s[!] Error: No hash set!%s %sUse %s'%sset hash <hashstring>%s'\n", bred, rst, bgrn, rst, bylw, rst)
		} else if len(hashtype) == 0 {
			fmt.Printf("\n%s[!] Error: No hash type set!%s %sUse %s '%sset hashtype <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
		}else if len(mode) == 0 {
		fmt.Printf("\n%s[!] Error: No mode set!%s %sUse %s'%sset mode <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
		}else if mode == "dict" {
			cracker.PassCrack(targetHash, hashtype, ruleFile)
		 }else if mode == "brute" {
		    cracker.BruteGen(targetHash, hashtype, startLen, endLen)
		 goto START
		}
		}else if len(targetHash) != 0{
		cond.HashConditions(targetHash, hashtype, mode, startLen, endLen)
		}else {
				continue
			}
			
			
		default:
			fmt.Printf("\n%s[!] Unknown command:%s %sType%s '%shelp%s' %sfor available commands.%s\n", bred, rst, grn, rst, bylw, rst, grn, rst)
		}
	}
}
