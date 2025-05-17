package utils

import "fmt"

const (
    bcyn = "\033[1;36m"
	bgrn = "\033[1;32m"
	bblu = "\033[1;34m"
	bred = "\033[1;31m"
	borng = "\033[1;38;5;208m"
	bylw = "\033[1;33m"
	grn = "\033[32m"
	blu = "\033[34m"
	ylw = "\033[33m"
	red = "\033[31m"
	orng = "\033[38;5;208m"
	rst = "\033[0m"
)

func Help() {

	fmt.Printf("\n\t\t\t%sAvailable Commands:%s", bcyn, rst)
	fmt.Printf("\n%sset hash <value>%s               - %sThis command accepts and stores the hash input that is to be cracked(eg.,set hash 21232f297a57a5a743894a0e4a801fc3)%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset hashtype <value>%s           -%sThis commmand accepts and stores the selected hashtype for cracking (e.g. set hashtype  md5)%s", orng, rst, grn, rst)
	fmt.Printf("\n%srun%s                            - %sThis command begins hash cracking with the given data%s", orng, rst, grn, rst)
	fmt.Printf("\n%sstatus%s                         - %sThis command shows the currently stored inputs%s", orng, rst, grn, rst)
	fmt.Printf("\n%sexit%s%s or%s%s Ctrl+c%s                 - %sThese commands closes the program%s", orng, rst, grn, rst, orng, rst, grn, rst)
	fmt.Printf("\n%shashid <value>%s                 - %sThis command helps analyzes and identifies the hashtype of a given hash (eg. hashid 49f68a5c8493ec2c0bf489821c21fc3b )%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset mode auto%s                  -%sThis command sets the cracking mode to automatic (first picks up wordlist then moves to bruteforce )%s", orng, rst, grn, rst)
	fmt.Printf("\n%shelp%s                           - %sThis command serves as a guide to users%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset range <min-max>%s      -%sThis command accepts brute force start and end (minimum and maximum) numbers. That is the length of words bruteforce begins and ends with. (eg. set range 4-6)%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset mode <value>%s               - %sThis command is used to set cracking mode to {brute (bruteforce), dict (wordlist cracking), or auto (bruteforce + wordlist cracking)}%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset hashfile <file>%s            -%sThis command accepts the name of the file or filepath that contains the hashes. (eg. set hashfile hashed.txt) or (eg. set hashfile /home/my_files/hashed.txt%s", orng, rst, grn, rst)
	fmt.Printf("\n%sset outputfile <file>%s               - %sThis command accepts the name of the file to which the cracked and uncracked hashes will be saved.%s \n\n %s*The output file is automatically named and created after the hash file. So there's no need to use this command, except the user wants to name the output file based on their choice%s", orng, rst, grn, rst, ylw, rst)


	fmt.Printf("\n\t\t\t%sValues To Use:%s", bcyn, rst)
	fmt.Printf("\n%s set hashtype <value>     - %ssupported values = md4, md5, sha1, sha224, sha256, sha384, sha512, sha512_224, sha512_256, ripemd160, adler32, crc32, crc64, blake2b, blake2s, fnv1_32, fnv1_64, fnv1a_32, fnv1a_64%s", grn, orng, rst)
	fmt.Printf("\n%sset mode <value>     - %ssupported values = brute, dict, auto%s", grn, orng, rst)

	fmt.Printf("\n\n%s*Only Brute Mode requires minimum and maximum length (range)%s\n", ylw, rst)
}
