package main

import (
	"PassCrax/core/analyzer"
	"PassCrax/core/cracker"
	"PassCrax/core/utils"
)

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	grn  = "\033[32m"
	ylw  = "\033[33m"
	blu  = "\033[34m"
	red  = "\033[31m"
	rst  = "\033[0m"
	bcyn = "\033[1;36m"
)

var (
	reSetHash     = regexp.MustCompile(`(?i)^set\s+hash\s+(.+)$`)
	reSetHashtype = regexp.MustCompile(`(?i)^set\s+hashtype\s+(.+)$`)
	reHashid      = regexp.MustCompile(`(?i)^hashid\s*(.+)$`)
	reSetMode     = regexp.MustCompile(`(?i)^set\s+mode\s+(.+)$`)
	reStartLen    = regexp.MustCompile(`(?i)^set min\s+(\d+)$`)
	reEndLen      = regexp.MustCompile(`(?i)^set max\s+(\d+)$`)
)

func ifEmpty(if_full, if_empty string) string {
	if if_full == "" {
		return if_empty
	}
	return if_full
}

func main() {
	var hash, hashtype, mode string
	var start_len, end_len int
	scanner := bufio.NewScanner(os.Stdin)

	// Ctrl+C handling for user to exit
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		fmt.Printf("\n%sTerminated By User!%s\n", red, rst)
		os.Exit(0)
	}()

	for {

		fmt.Printf("\n%sCURRENT SETTINGS%s", bcyn, rst)
		fmt.Printf("\n%sHash%s: %s%s%s", grn, rst, ylw, ifEmpty(hash, "Not Set"), rst)
		fmt.Printf("\n%sHash Type%s: %s%s%s", grn, rst, ylw, ifEmpty(hashtype, "Not Set"), rst)
		fmt.Printf("\n%sMode%s: %s%s%s", grn, rst, ylw, ifEmpty(mode, "Not Set"), rst)
		fmt.Printf("\n%sBrute Min Length%s: %s%d %s", grn, rst, ylw, start_len, rst)
		fmt.Printf("\n%sBrute Max Length%s: %s%d%s\n", grn, rst, ylw, end_len, rst)

		fmt.Print("\n> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch {
		case input == "exit":
			fmt.Printf("\n%sTerminated By User!%s\n", red, rst)
			fmt.Println("\n")
			return
		case input == "help":
			utils.Help()
		case input == "status":
			continue

		case reSetHash.MatchString(input):
			hash = strings.TrimSpace(reSetHash.FindStringSubmatch(input)[1])
			fmt.Printf("\n%sHash set to:%s %s%s%s\n", grn, rst, ylw, hash, rst)

		case reSetHashtype.MatchString(input):
			hashtype = strings.TrimSpace(reSetHashtype.FindStringSubmatch(input)[1])
			fmt.Printf("\n%sHash Type set to:%s %s%s%s\n", grn, rst, ylw, hashtype, rst)

		case reHashid.MatchString(input):
			hash = strings.TrimSpace(reHashid.FindStringSubmatch(input)[1])
			analyzer.PassAnalyze(hash)

		case reSetMode.MatchString(input):
			mode = strings.TrimSpace(reSetMode.FindStringSubmatch(input)[1])
			fmt.Printf("\n%sMode set to:%s %s%s%s\n", grn, rst, ylw, mode, rst)

		case reStartLen.MatchString(input):
			if value, err := strconv.Atoi(reStartLen.FindStringSubmatch(input)[1]); err == nil {
				start_len = value
				fmt.Printf("\n%sBrute Min Length set to:%s %s%d%s\n", grn, rst, ylw, start_len, rst)
			}

		case reEndLen.MatchString(input):
			if value, err := strconv.Atoi(reEndLen.FindStringSubmatch(input)[1]); err == nil {
				end_len = value
				fmt.Printf("\n%sBrute Max Length set to:%s %s%d%s\n", grn, rst, ylw, end_len, rst)
			}

		case input == "run":
			if len(hash) == 0 {
				fmt.Printf("\n%sError: No hash set!%s %sUse %s%s'set hash <value>'%s\n", red, rst, grn, rst, ylw, rst)
			} else if len(hashtype) == 0 {
				fmt.Printf("\n%sError: No hash type set!%s %sUse %s %s'set hashtype <value>'%s\n", red, rst, grn, rst, ylw, rst)
			} else if mode == "brute" {
				if end_len == 0 && start_len == 0 {
					fmt.Printf("\n%sError: Start and End Length Cannot Be Empty In BruteForce Mode%s\n", red, rst)
					continue
				}
				cracker.BruteGen(hash, hashtype, start_len, end_len)
			} else if mode == "dict" {
				cracker.PassCrack(hash, hashtype)
			} else if mode == "auto" {
				cracked_password := cracker.PassCrack(hash, hashtype)
				if cracked_password != "" {
					continue
				}
				fmt.Printf("%sPassword Not Found In Wordlist....%s\n%sSwitching To Bruteforce (1-7 characters)...%s\n", ylw, rst, blu, rst)
				time.Sleep(2 * time.Second)
				start_len = 1
				end_len = 7
				cracker.BruteGen(hash, hashtype, start_len, end_len)
			} else {
				fmt.Printf("\n%sError: Incorrect Mode {Choose '%sbrute%s' or '%sdict%s' or '%sauto%s' %sfor mode}%s\n", red, ylw, rst, ylw, rst, ylw, rst, red, rst)
			}

		default:
			fmt.Printf("\n%sUnknown command:%s %sType%s %s'help'%s %sfor options.%s\n", red, rst, grn, rst, ylw, rst, grn, rst)
		}
	}
}
