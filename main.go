package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"passcrax/core/analyzer"
	"passcrax/core/crack"
	"passcrax/core/file"
	"passcrax/core/utils"
	"passcrax/core/utils/cond"
	"passcrax/core/utils/help"
)

var (
	bgrn  = color.New(color.FgGreen, color.Bold).SprintFunc()
	bred  = color.New(color.FgRed, color.Bold).SprintFunc()
	bblu  = color.New(color.FgBlue, color.Bold).SprintFunc()
	bcyn  = color.New(color.FgCyan, color.Bold).SprintFunc()
	bylw  = color.New(color.FgYellow, color.Bold).SprintFunc() // Fixed: added .SprintFunc()
	borng = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	orng  = color.New(color.FgHiYellow).SprintFunc()
	grn   = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
	blu   = color.New(color.FgBlue).SprintFunc()
	cyn   = color.New(color.FgCyan).SprintFunc()
	ylw   = color.New(color.FgYellow).SprintFunc()
)

func Status() {
	fmt.Println(bcyn("\nCURRENT HASH SETTINGS"))
	fmt.Println(grn("Hash"), ":", ylw(ifEmpty(targetHash, "Not Set")))
	fmt.Println(grn("Hash Type"), ":", ylw(ifEmpty(hashtype, "Not Set")))
	fmt.Println(grn("Rule File"), ":", ylw(ifEmpty(ruleFile, "Not Set {Optional}")))
	fmt.Println(grn("Mode"), ":", ylw(ifEmpty(mode, "Not Set")))
	if mode == "brute" || mode == "auto" {
		fmt.Println(grn("Brute Charset"), ":", ylw(ifEmpty(charset, "Not Set")))
		fmt.Println(grn("Brute Min Length"), ":", ylw(ifEmpty(strconv.Itoa(startLen), "Not Set")))
		fmt.Println(grn("Brute Max Length"), ":", ylw(ifEmpty(strconv.Itoa(endLen), "Not Set")))
	}
	if mode == "dict" || mode == "auto" {
		fmt.Println(grn("Dict Path"), ":", ylw(ifEmpty(dictDir, "Not Set {Optional}")))
	}
}

func FileStatus() {
	fmt.Println(bcyn("\nCURRENT FILE HASH SETTINGS"))
	fmt.Println(grn("Hash File"), ":", ylw(ifEmpty(hashFile, "Not Set")))
	fmt.Println(grn("Output File"), ":", ylw(ifEmpty(outputFile, "Not Set")))
	fmt.Println(grn("Hash Type"), ":", ylw(ifEmpty(hashtype, "Not Set")))
	fmt.Println(grn("Mode"), ":", ylw(ifEmpty(mode, "Not Set")))

	if mode == "brute" || mode == "auto" {
		fmt.Println(grn("Brute Charset"), ":", ylw(ifEmpty(charset, "Not Set")))
		fmt.Println(grn("Brute Min Length"), ":", ylw(ifEmpty(strconv.Itoa(startLen), "Not Set")))
		fmt.Println(grn("Brute Max Length"), ":", ylw(ifEmpty(strconv.Itoa(endLen), "Not Set")))
	}
	if mode == "dict" || mode == "auto" {
		fmt.Println(grn("Dict Path"), ":", ylw(ifEmpty(dictDir, "Not Set {Optional}")))
	}
}

func ifEmpty(if_full interface{}, if_null string) string {
	if if_full == "" {
		return if_null
	}
	return fmt.Sprint(if_full)
}

func CreateCompleter() *readline.PrefixCompleter {
	return readline.NewPrefixCompleter(
		readline.PcItem("set",
			readline.PcItem("hash"),
			readline.PcItem("hashtype"),
			readline.PcItem("mode",
			readline.PcItem("auto"),
			readline.PcItem("brute"),
			readline.PcItem("dict"),),
			readline.PcItem("charset"),
			readline.PcItem("brute-range"),
			readline.PcItem("outputfile"),
		),
		readline.PcItem("load",
			readline.PcItem("hashfile"),
			readline.PcItem("dictdir"),
			readline.PcItem("rulefile"),
		),
		readline.PcItem("identify"),
		readline.PcItem("run"),
		readline.PcItem("status"),
		readline.PcItem("help"),
		readline.PcItem("exit"),
		readline.PcItem("quit"),
	)
}

var (
	reSetHash     = regexp.MustCompile(`(?i)^set\s+hash\s+(.+)$`)
	reSetHashtype = regexp.MustCompile(`(?i)^set\s+hashtype\s+(.+)$`)
	reHashid      = regexp.MustCompile(`(?i)^identify\s+(.+)$`)
	reSetMode     = regexp.MustCompile(`(?i)^set\s+mode\s+(.+)$`)
	reSetCharset  = regexp.MustCompile(`(?i)^set\s+charset\s+(.+)$`)
	reBruteRange  = regexp.MustCompile(`^(?i)set\sbrute-range\s\d+\s*-\s*\d+$`)
	reSetRules    = regexp.MustCompile(`^(?i)load\s+rulefile\s+(.+)$`)
	reUnsetRules  = regexp.MustCompile(`^(?i)drop\s+rulefile\s*$`)
	reSetDictDir  = regexp.MustCompile(`^(?i)load\s+dictdir\s+(.+)$`)
	reSetInput    = regexp.MustCompile(`^(?i)load\s+hashfile\s+(.+)$`)
	reSetOutput   = regexp.MustCompile(`^(?i)set\s+outputfile\s+(.+)$`)
)

var targetHash, hashtype, mode, hashFile, ruleFile, outputFile, dictDir, charset string
var startLen, endLen int
var usingHashFile bool

func main() {

fmt.Print("\n")
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "> ",
		HistoryFile:  "tmp/passcrax_history",
		HistoryLimit: 1000,
		AutoComplete: CreateCompleter(),
		UniqueEditLine:  true,
		HistorySearchFold: true,
	})
	if err != nil {
		fmt.Println(bred("[!] Error: "), err)
		return
	}
	defer rl.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		fmt.Println(bred("\n   Program Terminated!"))
	}()

	utils.Banner()
	Status()

	for {
	START:
		fmt.Print("\n> ")
      input, err := rl.Readline()
      
		if err != nil {
		    if err == readline.ErrInterrupt {
		        fmt.Println(bred("\n   Program Terminated!"))
		        break
		    }
		    fmt.Println(bred("\n[!] Program Terminated ('Ctrl+D pressed')"))
		    return
		}
		input = strings.TrimSpace(input)

		switch {

		case input == "exit", input == "quit":
			fmt.Println(bred("\n   Program Terminated!"))
			fmt.Println("\n")
			return

		case input == "help":
			help.Help()

		case input == "status":
			if usingHashFile {
				FileStatus()
			} else {
				Status()
			}

		case reSetHash.MatchString(input):
			usingHashFile = false
			hashFile = ""
			outputFile = ""
			targetHash = strings.TrimSpace(reSetHash.FindStringSubmatch(input)[1])
			fmt.Println(grn("\n[~] Hash set to:"), bylw(targetHash))

		case reSetHashtype.MatchString(input):
			hashtype = strings.TrimSpace(reSetHashtype.FindStringSubmatch(input)[1])
			hashtype = strings.ToLower(hashtype)
			if hashtype == analyzer.CheckValidHashType(hashtype) {
				fmt.Println(grn("\n[~] Hash Type set to:"), bylw(hashtype))
			} else {
				fmt.Println(bred("\n[!] Hashtype Value Is Invalid Or Unsupported!"))
				fmt.Println(bgrn("\n[~] Use '"), bylw("identify <hashstring>"), bgrn("' if you don't know what hashtype to use"))
				hashtype = ""
			}

		case reHashid.MatchString(input):
			if len(input) == 0 && len(targetHash) != 0 {
				input = targetHash
			}
			targetHash = strings.TrimSpace(reHashid.FindStringSubmatch(input)[1])
			if utils.IsFile(targetHash) {
				usingHashFile = true
				hashFile = targetHash

				fmt.Println(analyzer.FileAnalyze(hashFile))
				fmt.Println(bblu("\n[[[ "), bcyn("End Of File of '"), bgrn(hashFile), bcyn("'"), bblu("]]]"))

				fmt.Println(grn("\n[~] Hash File set to:"), bylw(hashFile))

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
						goto PRACK
					}
					fmt.Println(grn("\n[~] Output file set to:"), bylw(outputFile))
				}
			} else {
				usingHashFile = false
				result := analyzer.PassAnalyze(targetHash)
				if len(result) == 0 {
					targetHash = ""
				} else {
					fmt.Println(analyzer.PassAnalyze(targetHash))
					fmt.Println(grn("\n[~] Hash set to:"), bylw(targetHash))
				}
			}

		case reSetMode.MatchString(input):
			mode = strings.TrimSpace(reSetMode.FindStringSubmatch(input)[1])
			mode = strings.ToLower(mode)
			if mode == analyzer.CheckValidMode(mode) {
				fmt.Println(grn("\n[~] Mode set to:"), bylw(mode))
			} else {
				fmt.Println(bred("\n[!] Mode Value Is Invalid!"))
				fmt.Print(bgrn("\n [~] These are the list of supported modes to use:"))
				fmt.Println(bylw(analyzer.CheckValidMode(mode)))
				mode = ""
			}

		case reSetCharset.MatchString(input):
			charset = strings.TrimSpace(reSetCharset.FindStringSubmatch(input)[1])
			parsedstr := crack.ParseCharset(charset)
			fmt.Print(bgrn("\n[~] Char Count:"))
			fmt.Println(borng(len(parsedstr)))
			fmt.Print(bgrn("\n[~] Parsed Charset:"))
			fmt.Println(borng(string(parsedstr)))
			fmt.Println(grn("\n[~] Charset set to:"), bylw(charset))

		case reSetDictDir.MatchString(input):
			dictDir = strings.TrimSpace(reSetDictDir.FindStringSubmatch(input)[1])
			_, err := os.Stat(dictDir)
			if err != nil {
				fmt.Print(bred("\n[!] Error: "), dictDir, " does not exist")
				fmt.Println(borng("\n[+] Cross check if it's not a file path typographical error"))
				continue
			}
			var dictNum int
			dict_files, err := filepath.Glob(filepath.Join(dictDir, "*.txt"))
			if err != nil {
				fmt.Printf(bred("\n[!] Error Scanning Directory %s: "), dictDir, err)
				return
			}
			if len(dict_files) == 0 {
				fmt.Print(red("\n[!] Error: No Files Found In "), dictDir)
				return
			}
			for dictNum, _ = range dict_files {
			}
			dictNumber := dictNum + 1

			fmt.Print(bgrn("\n[~] Found "), borng(dictNumber), bgrn(" wordlist files from "), borng(dictDir))
			fmt.Println(grn("\n[~] Dict path set to:"), bylw(dictDir))

		case reBruteRange.MatchString(input):
			downcase := strings.ToLower(input)
			trimRange := strings.TrimPrefix(downcase, "set brute-range")
			if strings.Contains(trimRange, "-") {
				num := strings.Split(trimRange, "-")
				startNum := strings.TrimSpace(num[0])
				endNum := strings.TrimSpace(num[1])

				value, err := strconv.Atoi(startNum)
				if err != nil {
					fmt.Printf(bred("[!] Error: %v"), err)
				}
				startLen = value

				val, err := strconv.Atoi(endNum)
				if err != nil {
					fmt.Printf(bred("[!] Error: %v"), err)
				}
				endLen = val

				if startLen >= endLen {
					fmt.Print(bred("\n[!] Error: Minimum length cannot be greater than Maximum length"))
					startLen = 0
					endLen = 0
				} else {
					fmt.Println(grn("\n[~] Brute Min Length set to:"), bylw(startLen))
					fmt.Println(grn("\n[~] Brute Max Length set to:"), bylw(endLen))
				}
			} else {
				trimRange = strings.TrimSpace(trimRange)
				value, err := strconv.Atoi(trimRange)
				if err != nil {
					fmt.Printf(bred("[!] Error: %v"), err)
				}
				startLen = value
				endLen = startLen
				fmt.Println(grn("\n[~] Brute Length set to:"), bylw(startLen))
			}

		case reSetInput.MatchString(input):
			usingHashFile = true
			targetHash = ""
			hashFile = strings.TrimSpace(reSetInput.FindStringSubmatch(input)[1])
			_, err := os.Stat(hashFile)
			if err != nil {
				fmt.Print(bred("\n[!] Error: "), hashFile, " does not exist")
				fmt.Println(borng("\n[+] Cross check if it's not a file path typographical error"))
				hashFile = ""
				continue
			}
			totalLines := file.FileCount(hashFile)
			fmt.Println(grn("\n[~] Hash file set to:"), bylw(hashFile))

			fmt.Print(bgrn("\n[~] Found "), borng(totalLines), bgrn(" hashes in "), borng(hashFile))
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
				fmt.Println(grn("\n[~] Output file set to:"), bylw(outputFile))
			}

		case reSetOutput.MatchString(input):
			usingHashFile = true
			outputFile = strings.TrimSpace(reSetOutput.FindStringSubmatch(input)[1])
			fmt.Println(grn("\n[~] Output file set to:"), bylw(outputFile))

		case reSetRules.MatchString(input):
			ruleFile = strings.TrimSpace(reSetRules.FindStringSubmatch(input)[1])
			if _, err := os.Stat(ruleFile); errors.Is(err, fs.ErrNotExist) {
				fmt.Print(bred("\n[!] Error: "), ruleFile, " does not exist")
				fmt.Println(bylw("\n[+] Cross check if it's not a file path typographical error"))
				ruleFile = ""
				continue
			}
			fmt.Println(grn("\n[~] Rule file set to:"), bylw(ruleFile))

		case reUnsetRules.MatchString(input):
			if len(ruleFile) == 0 {
				fmt.Println(bred("[!] Error: No rule file loaded!"))
			} else {
				ruleFile = ""
				fmt.Println(bgrn("\n[~] Rule file dropped successfully!"))
			}

		case input == "run":
			if len(targetHash) == 0 && len(hashFile) == 0 {
				cond.HashConditions(targetHash, hashtype, mode, charset, dictDir, startLen, endLen)
			} else if len(hashFile) != 0 && len(outputFile) != 0 {
				cond.FileConditions(hashFile, hashtype, mode, charset, dictDir, startLen, endLen, outputFile)
			} else if len(ruleFile) != 0 && len(hashFile) == 0 && usingHashFile == false {
				if len(targetHash) == 0 {
					fmt.Println(bred("\n[!] Error: No hash set!"), bgrn("Use"), bylw("'set hash <hashstring>'"))
				} else if len(hashtype) == 0 {
					fmt.Println(bred("\n[!] Error: No hash type set!"), bgrn("Use"), bylw("'set hashtype <value>'"))
				} else if len(mode) == 0 {
					fmt.Println(bred("\n[!] Error: No mode set!"), bgrn("Use"), bylw("'set mode <value>'"))
				} else if mode == "dict" {
					crack.PassCrack(dictDir, targetHash, hashtype, ruleFile)
				} else if mode == "brute" && len(charset) != 0 {
					crack.BruteGen(targetHash, hashtype, charset, startLen, endLen)
					goto START
				}
			} else if len(targetHash) != 0 {
				cond.HashConditions(targetHash, hashtype, mode, charset, dictDir, startLen, endLen)
			} else {
				continue
			}

		default:
			dym := utils.DidYouMean(input)
			fmt.Print(bgrn("\nDid you mean: "), bcyn(dym), bgrn(" ?\n"))
		}
	}
}
