package help

import (
	"fmt"
	"github.com/fatih/color"
	"passcrax/core/analyzer"
)

var (
	bgrn  = color.New(color.FgGreen, color.Bold)
	bred  = color.New(color.FgRed, color.Bold)
	borng = color.New(color.FgHiYellow, color.Bold)
	bblu  = color.New(color.FgBlue, color.Bold)
	bcyn  = color.New(color.FgCyan, color.Bold)
	bylw  = color.New(color.FgYellow, color.Bold)
	grn   = color.New(color.FgGreen)
	red   = color.New(color.FgRed)
	orng  = color.New(color.FgHiYellow)
	blu   = color.New(color.FgBlue)
	cyn   = color.New(color.FgCyan)
	ylw   = color.New(color.FgYellow)
)


func Help() {
	bcyn.Println("\n\t\t\tAvailable Commands:")

	orng.Print("set hash <hashstring>")
	grn.Println(" - This command accepts and stores the hash input that is to be cracked (eg., set hash 21232f297a57a5a743894a0e4a801fc3)")

	orng.Print("set hashtype <value>")
	grn.Println(" - This command accepts and stores the selected hashtype for cracking (e.g. set hashtype md5)")

	orng.Print("run")
	grn.Println(" - This command begins hash cracking with the given data")

	orng.Print("status")
	grn.Println(" - This command shows the currently stored inputs")

	orng.Print("exit")
	grn.Print(" or ")
	orng.Print("Ctrl+c")
	grn.Println(" - These commands close the program")

	orng.Print("identify <hashstring>")
	grn.Print(" or ")
	orng.Print("identify <filepath>")
	grn.Println(" - This command helps analyze and identify the hashtype of a given hash or file (eg. hashid 49f68a5c8493ec2c0bf489821c21fc3b) (eg. hashid /home/my_files/hashed.txt)")

	orng.Print("set mode auto")
	grn.Println(" - This command sets the cracking mode to automatic (first picks up wordlist then moves to bruteforce)")

	orng.Print("help")
	grn.Println(" - This command serves as a guide to users")

	orng.Print("set brute-range <min-max>")
	grn.Println(" - This command accepts brute force start and end (minimum and maximum) numbers for word length (eg. set range 4-6)")

	orng.Print("set mode <value>")
	grn.Println(" - This command is used to set cracking mode to {brute (bruteforce), dict (wordlist cracking), or auto (bruteforce + wordlist cracking)}")

	orng.Print("set charset [char ranges]")
	grn.Println(" - This command accepts characters for bruteforce cracking. Characters are ranges of letters, numbers, and symbols enclosed in square brackets (eg. set charset [a-gA-R0-5&*^#])")

	orng.Print("load rulefile <filepath>")
	grn.Println(" - This command accepts and loads the filepath with word mangling rules (eg. load rulefile /home/my_files/hashed.rule)")

	orng.Print("drop rulefile")
	grn.Println(" - This command clears the rule file from the status table, allowing cracking without rules")

	orng.Print("load dictdir <dirpath>")
	grn.Println(" - This command accepts and loads the directory containing wordlist files (eg. load dictdir /home/my_files/my_dict)")

	orng.Print("load hashfile <filepath>")
	grn.Println(" - This command accepts and loads the filepath containing hashes (eg. load hashfile /home/my_files/hashed.txt)")

	orng.Print("set outputfile <file>")
	grn.Print(" - This command accepts the name of the file to save cracked and uncracked hashes")
	ylw.Println("\n   *The output file is automatically named and created after the hash file. This command is for custom naming preferences. Don't forget the filepath")

	bcyn.Println("\n\t\t\tValues To Use:")
	grn.Print("set mode <value>     - ")
	orng.Print("supported values = ")
	fmt.Print(analyzer.CheckValidMode(""))

	bylw.Println("\n*Only Brute Mode requires minimum and maximum length (range)\n\n")
}
