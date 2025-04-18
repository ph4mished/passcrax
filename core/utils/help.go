package utils

import "fmt"

const (
	ylw  = "\033[33m"
	grn  = "\033[32m"
	red  = "\033[31m"
	blu  = "\033[34m"
	rst  = "\033[0m"
	bcyn = "\033[1;36m"
)

func Help() {

	fmt.Printf("\n\t\t%sAvailable Commands:%s", bcyn, rst)
	fmt.Printf("\n%sset hash <value>%s - %sUsed to set hash (eg.,set hash 21232f297a57a5a743894a0e4a801fc3%s)", blu, rst, grn, rst)
	fmt.Printf("\n%sset hashtype <value>%s   -%sUsed to set hashtype (e.g. set hashtype  md5)%s", blu, rst, grn, rst)
	fmt.Printf("\n%srun%s             - %sExecute password cracking%s", blu, rst, grn, rst)
	fmt.Printf("\n%sstatus%s             - %sShow current settings%s", blu, rst, grn, rst)
	fmt.Printf("\n%sexit%s%s or%s%s Ctrl+c%s               - %sQuit program%s", blu, rst, grn, rst, blu, rst, grn, rst)
	fmt.Printf("\n%shashid <value>%s     - %sIdentify hash type (eg. hashid 49f68a5c8493ec2c0bf489821c21fc3b )%s", blu, rst, grn, rst)
	fmt.Printf("\n%sset mode auto%s   -%sUsed to set mode to automatic (first picks up wordlist then moves to bruteforce )%s", blu, rst, grn, rst)
	fmt.Printf("\n%shelp%s               - %sShow this help%s", blu, rst, grn, rst)
	fmt.Printf("\n%sset min <number>%s        - %sbrute force minimum length%s", blu, rst, grn, rst)
	fmt.Printf("\n%sset max <number>%s        -%sbrute force maximum length%s", blu, rst, grn, rst)
	fmt.Printf("\n%sset mode <value>     - %smode%s", blu, grn, rst)

	fmt.Printf("\n\t\t%sValues To Use:%s", bcyn, rst)
	fmt.Printf("\n%s set hashtype <value>     - %svalues = md5, sha1, sha224, sha256, sha384, sha512, sha512_224, sha512_256%s", grn, blu, rst)
	fmt.Printf("\n%sset mode <value>     - %svalues = brute, dict, auto%s", grn, blu, rst)

	fmt.Printf("\n%s*Only Brute Mode requires minimum and maximum length%s\n", ylw, rst)
}
