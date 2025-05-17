package main

//above all the code will be kept simple and concise
import (
	"PassCrax/core/analyzer"
	"PassCrax/core/utils"
	"PassCrax/core/cracker" 
	"PassCrax/core/cracker/file"
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
	"time"
	"io/fs"
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
	reRange       = regexp.MustCompile(`^(?i)set\srange\s\d+\s*-\s*\d+$`)
	reSetInput    = regexp.MustCompile(`^(?i)set\s+hashfile\s+(.+)$`)
    reSetOutput    = regexp.MustCompile(`^(?i)set\s+outputfile\s+(.+)$`)
    )

var targetHash, hashtype, mode, hashFile, outputFile string
	var startLen, endLen int
	var usingHashFile bool
	
	func ifEmpty(if_full, if_empty string) string {
	if if_full == "" {
		return if_empty
	}
	return if_full
}

func Status(){
	fmt.Printf("\n%sCURRENT SETTINGS%s", bcyn, rst)
	fmt.Printf("\n%sHash%s: %s%s%s", grn, rst, ylw, ifEmpty(targetHash, "Not Set"), rst)
	fmt.Printf("\n%sHash Type%s: %s%s%s", grn, rst, ylw, ifEmpty(hashtype, "Not Set"), rst)
	fmt.Printf("\n%sMode%s: %s%s%s", grn, rst, ylw, ifEmpty(mode, "Not Set"), rst)
	fmt.Printf("\n%sBrute Min Length%s: %s%d %s", grn, rst, ylw, startLen, rst)
	fmt.Printf("\n%sBrute Max Length%s: %s%d%s\n", grn, rst, ylw, endLen, rst)
}

func FileStatus(){
    fmt.Printf("\n%sCURRENT FILE HASH SETTINGS%s", bcyn, rst)
	fmt.Printf("\n%sHash File%s: %s%s%s", grn, rst, ylw, ifEmpty(hashFile, "Not Set"), rst)
	fmt.Printf("\n%sOutput File%s: %s%s%s", grn, rst, ylw, ifEmpty(outputFile, "Not Set"), rst)
	fmt.Printf("\n%sHash Type%s: %s%s%s", grn, rst, ylw, ifEmpty(hashtype, "Not Set"), rst)
	fmt.Printf("\n%sMode%s: %s%s%s", grn, rst, ylw, ifEmpty(mode, "Not Set"), rst)
	fmt.Printf("\n%sBrute Min Length%s: %s%d %s", grn, rst, ylw, startLen, rst)
	fmt.Printf("\n%sBrute Max Length%s: %s%d%s\n", grn, rst, ylw, endLen, rst)
}



func main(){
	
	scanner := bufio.NewScanner(os.Stdin)

	// Ctrl+C handling for user to exit
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		fmt.Printf("\n%sProgram Terminated!%s\n", bred, rst)
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
			fmt.Printf("\n%sProgram Terminated!%s\n", bred, rst)
			fmt.Println("\n")
			return
		
		case input == "help":
			utils.Help()
		
		//Problem: It's impossible to switch back from filestatus to status. One has to close the program and restart before they can be able to see hash status. It should be more flexible.
        case input == "status":
			if usingHashFile {
			//targetHash = ""
			FileStatus()
			//}else if hashFile != "" && targetHash != ""{
		//	hashFile = ""
		//	Status()
			}else{
			Status()
			}
			
			
		case reSetHash.MatchString(input):
            usingHashFile = false
			targetHash = strings.TrimSpace(reSetHash.FindStringSubmatch(input)[1])
			fmt.Printf("\n%sHash set to:%s %s%s%s\n", grn, rst, bylw, targetHash, rst)


		case reSetHashtype.MatchString(input):
			hashtype = strings.TrimSpace(reSetHashtype.FindStringSubmatch(input)[1])
			if hashtype == analyzer.CheckValidHashType(hashtype){
			fmt.Printf("\n%sHash Type set to:%s %s%s%s\n", grn, rst, bylw, hashtype, rst)
			}else{
			fmt.Printf("\n%sInputted Hashtype Is Invalid!%s \n %sThese are the list of valid inputs to use:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, analyzer.CheckValidHashType(hashtype), rst)
			  hashtype = ""
			}
//unknown hashes should not be accepted cos they cant be cracked here. if hashid failed to identify, what shows passcrax will be able to crack
//Idea: hashes that were shown as unknown will not be cracked. all hashes will be reverified to see if it results in unknown. 
// lets not trouble hashid
		case reHashid.MatchString(input):
		//because of input getting accepted before they are verified, 
		//former inputs are lost even if current ones are defined unknown
			targetHash = strings.TrimSpace(reHashid.FindStringSubmatch(input)[1])
			analyzer.PassAnalyze(targetHash)
			fmt.Printf("\n%sHash set to:%s %s%s%s\n", grn, rst, bylw, targetHash, rst)
			
		case reSetMode.MatchString(input):
			mode = strings.TrimSpace(reSetMode.FindStringSubmatch(input)[1])
			if mode == analyzer.CheckValidMode(mode){
			fmt.Printf("\n%sMode set to:%s %s%s%s\n", grn, rst, bylw, mode, rst)
			}else{
			fmt.Printf("\n%sInputted Mode Is Invalid!%s \n %sThese are the list of accepted modes to use:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, analyzer.CheckValidMode(mode), rst)
			  mode = ""
			  }
			

		case reRange.MatchString(input):
		downcase := strings.ToLower(input)
     	trimStart := strings.TrimPrefix(downcase, "set range")
		num := strings.Split(trimStart, "-")
		startNum := strings.TrimSpace(num[0])
	    endNum := strings.TrimSpace(num[1])
	    
		
	value, err := strconv.Atoi(startNum)
	if err != nil{
	fmt.Println("Error: ", err)
	}
	startLen = value
	
	
	val, err := strconv.Atoi(endNum)
	if err != nil{
	fmt.Println("Error: ", err)
	}
	endLen = val
	
	
	if startLen >= endLen {
	fmt.Printf("\n%sError: Minimum length cannot be greater than Maximum length%s\n", bred, rst)
	startLen = 0
	endLen = 0
	    }else{   
	fmt.Printf("\n%sBrute Min Length set to:%s %s%d%s\n", grn, rst, bylw, startLen, rst)
	fmt.Printf("\n%sBrute Max Length set to:%s %s%d%s\n", grn, rst, bylw, endLen, rst)
	}
	
       case reSetInput.MatchString(input):
           usingHashFile = true
	   hashFile = strings.TrimSpace(reSetInput.FindStringSubmatch(input)[1])
	   if _, err := os.Stat(hashFile); errors.Is(err, fs.ErrNotExist){
	   fmt.Printf("\n%sError: %s does not exist%s", bred, hashFile, rst)
	   fmt.Printf("\n%sCross check if you didn't make a typo in writing the file name%s\n", bylw, rst)
	   hashFile = ""
	   continue
	   }
	   fmt.Printf("\n%sHash file set to:%s %s%s%s\n", grn, rst, bylw, hashFile, rst)
	   if strings.Contains(hashFile, "/"){
	   split := strings.Split(hashFile, ".")
	   begin := split[0]
	   end := split[1]
	   outputFile = fmt.Sprintf("%s_cracked.%s", begin, end)
	   fmt.Printf("\n%sOutput file set to:%s %s%s%s\n", grn, rst, bylw, outputFile, rst)
	  }else{
	   outputFile = fmt.Sprintf("cracked_%s", hashFile)
	   fmt.Printf("\n%sOutput file set to:%s %s%s%s\n", grn, rst, bylw, outputFile, rst)
	   }
	  
	   case reSetOutput.MatchString(input):
	    outputFile = strings.TrimSpace(reSetOutput.FindStringSubmatch(input)[1])
	   fmt.Printf("\n%sOutput file set to:%s %s%s%s\n", grn, rst, bylw, outputFile, rst)
	   
	   
	  
		case input == "run":
		if len(hashFile) != 0 && len(outputFile)  != 0{
		targetHash = ""
		
		if len(hashtype) == 0 {
				fmt.Printf("\n%sError: No hash type set!%s %sUse %s'%sset hashtype <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			 }else if mode == "brute" {
				if endLen == 0 && startLen == 0 {
					fmt.Printf("\n%sError: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode%s\n %sUse %s'%sset range <min-max>%s'\n", bred, rst, bgrn, rst, bylw, rst)
					continue
				}
				file.BruteFile(hashFile, hashtype, outputFile, startLen, endLen)
				fmt.Printf("\n%sResults has being copied to %s %s%s%s\n", bgrn, rst, bylw, outputFile, rst)
			} else if mode == "dict" {
			file.DictFile(hashFile, hashtype, outputFile)
			fmt.Printf("\n%sResults has being copied to %s %s%s%s\n", bgrn, rst, bylw, outputFile, rst)
			//} //else if mode == "auto" {
			    //cracked_password := cracker.PassCrack(targetHash, hashtype)
				//if cracked_password != "" {
				//	continue
				//}
			//	fmt.Printf("\n%sPassword Not Found In Wordlist....%s\n%sSwitching To Bruteforce (1-7 characters)...%s\n", ylw, rst, bblu, rst)
				//time.Sleep(2 * time.Second)
				//startLen = 1
			//	endLen = 7
				//cracker.BruteGen(targetHash, hashtype, startLen, endLen)
			}else if mode != "dict" && mode != "brute"{
			mode = ""
			validMode := "dict, brute"
			fmt.Printf("\n%sInputted Mode Is Invalid!%s \n %sThese are the list of accepted modes to use:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, validMode, rst)
			}else{
			fmt.Println("mAd")
			}
			//'START' is here to make a jump over 'set hash error' 
			goto START
		}
		 if len(targetHash) == 0 {
				fmt.Printf("\n%sError: No hash set!%s %sUse %s'%sset hash <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			} else if len(hashtype) == 0 {
				fmt.Printf("\n%sError: No hash type set!%s %sUse %s '%sset hashtype <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			}else if len(mode) == 0 {
			fmt.Printf("\n%sError: No mode set!%s %sUse %s'%sset mode <value>%s'\n", bred, rst, bgrn, rst, bylw, rst)
			}else if mode == "brute" {
				if endLen == 0 && startLen == 0 {
					fmt.Printf("\n%sError: Minimum Length and Maximum Length Cannot Be Empty In BruteForce Mode%s\n %sUse %s'%sset range <min-max>%s'\n", bred, rst, bgrn, rst, bylw, rst)
					continue
				}
				cracker.BruteGen(targetHash, hashtype, startLen, endLen)
			} else if mode == "dict" {
				cracker.PassCrack(targetHash, hashtype)
			} else if mode == "auto" {
				cracked_password := cracker.PassCrack(targetHash, hashtype)
				if cracked_password != "" {
					continue
				}
				fmt.Printf("\n%sPassword Not Found In Wordlist....%s\n%sSwitching To Bruteforce (1-7 characters)...%s\n", ylw, rst, bblu, rst)
				time.Sleep(2 * time.Second)
				startLen = 1
				endLen = 7
				cracker.BruteGen(targetHash, hashtype, startLen, endLen)
				
				//this was repeated for run because some words bypass the mode valid check. this is only some strings to hold it for now
			}else if mode != "dict" && mode != "brute" && mode != "auto" {
			mode = ""
			validMode := "dict, brute, auto"
			fmt.Printf("\n%sInputted Mode Is Invalid!%s \n %sThese are the list of accepted modes to use:%s%s %s%s\n", bred, rst, bgrn, rst, bylw, validMode, rst)
			}else {
				continue
			}
			
		default:
			fmt.Printf("\n%sUnknown command:%s %sType%s '%shelp%s' %sfor available commands.%s\n", bred, rst, grn, rst, bylw, rst, grn, rst)
		}
	}
}
