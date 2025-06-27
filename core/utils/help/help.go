package help

import "passcrax/core/analyzer"

import (
	"fmt"
)

const (
	bcyn  = "\033[1;36m"
	bgrn  = "\033[1;32m"
	bblu  = "\033[1;34m"
	bred  = "\033[1;31m"
	borng = "\033[1;38;5;208m"
	bylw  = "\033[1;33m"
	grn   = "\033[32m"
	blu   = "\033[34m"
	ylw   = "\033[33m"
	red   = "\033[31m"
	orng  = "\033[38;5;208m"
	rst   = "\033[0m"
)

func Help() {

	fmt.Printf("\n\t\t\t%sAvailable Commands:%s", bcyn, rst)
	fmt.Printf("\n%sset hash <hashstring>%s               - %sThis command accepts and stores the hash input that is to be cracked(eg.,set hash 21232f297a57a5a743894a0e4a801fc3)%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset hashtype <value>%s           -%sThis commmand accepts and stores the selected hashtype for cracking (e.g. set hashtype  md5)%s", orng, rst, grn, rst)
	fmt.Printf("\n%srun%s                            - %sThis command begins hash cracking with the given data%s", orng, rst, grn, rst)
	fmt.Printf("\n%sstatus%s                         - %sThis command shows the currently stored inputs%s", orng, rst, grn, rst)
	fmt.Printf("\n%sexit%s%s or%s%s Ctrl+c%s                 - %sThese commands closes the program%s", orng, rst, grn, rst, orng, rst, grn, rst)
	fmt.Printf("\n%shashid <hashstring> %sor%s %shashid <filepath>%s                 - %sThis command helps analyzes and identifies the hashtype of a given hash or a given filepath (file) (eg. hashid 49f68a5c8493ec2c0bf489821c21fc3b ) (eg. hashid /home/my_files/hashed.txt)%s", orng, grn, rst, orng, rst, grn, rst)
	fmt.Printf("\n%sset mode auto%s                  -%sThis command sets the cracking mode to automatic (first picks up wordlist then moves to bruteforce )%s", orng, rst, grn, rst)
	fmt.Printf("\n%shelp%s                           - %sThis command serves as a guide to users%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset brute-range <min-max>%s      -%sThis command accepts brute force start and end (minimum and maximum) numbers. That is the length of words bruteforce begins and ends with. (eg. set range 4-6)%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset mode <value>%s               - %sThis command is used to set cracking mode to {brute (bruteforce), dict (wordlist cracking), or auto (bruteforce + wordlist cracking)}%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset charset [char ranges]%s                  -%sThis command accepts the characters for bruteforce cracking. The characters are ranges of letters (lower or upper) and numbers plus user defined symbol characters The value or character ranges should be enclosed in square brakets followed by a range of alphabets or numbers or special chars which are seperated by hyphens. \n eg.set charset [a-g A-R 0-5 &*^#]%s", orng, rst, grn, rst)
	fmt.Printf("\n%sload rulefile <filepath>%s            -%sThis command accepts and loads the filepath with the file that contains words mangling rules. (eg. load rulefile /home/my_files/hashed.rule)%s", orng, rst, grn, rst)
	fmt.Printf("\n%sdrop rulefile%s                -%sThis command clears the name of the rule file from the status table. This makes it possible to crack hashes without using rules. (eg.drop rulefile)%s\n%sNB:%sThe command for dropping rule file accepts no value (filename or filepath)%s\n", orng, rst, grn, rst, red, ylw, rst)
	fmt.Printf("\n%sload dictdir <dirpath>%s            -%sThis command accepts and loads the directory that contains wordlist files. This is to allow the use of user defined or external wordlist directory. (eg. load dictdir /home/my_files/my_dict)%s", orng, rst, grn, rst)
	fmt.Printf("\n%sload hashfile <filepath>%s            -%sThis command accepts and loads the filepath with the file that contains the hashes. (eg. load hashfile /home/my_files/hashed.txt)%s", orng, rst, grn, rst)
	fmt.Printf("\n\n%sset outputfile <file>%s               - %sThis command accepts the name of the file to which the cracked and uncracked hashes will be saved.%s \n\n %s*The output file is automatically named and created after the hash file. So this command is negligible, except for users who want to name the output file by their preferences. Do not forget the filepath%s", orng, rst, grn, rst, ylw, rst)

	fmt.Printf("\n\t\t\t%sValues To Use:%s", bcyn, rst)
	fmt.Printf("\n%s set hashtype <value>     - %ssupported values = %s%s", grn, orng, analyzer.CheckValidHashType(""), rst)
	fmt.Printf("\n%sset mode <value>     - %ssupported values = %s%s", grn, orng, analyzer.CheckValidMode(""), rst)

	fmt.Printf("\n\n%s*Only Brute Mode requires minimum and maximum length (range)%s\n", ylw, rst)
}
