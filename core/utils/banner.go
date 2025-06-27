package utils

import "fmt"

const (
	bcyn  = "\033[1;36m"
	borng = "\033[1;38;5;208m"
	bgrn  = "\033[1;32m"
	bblu  = "\033[1;34m"
	bred  = "\033[1;31m"
	bylw  = "\033[1;33m"
	grn   = "\033[32m"
	blu   = "\033[34m"
	ylw   = "\033[33m"
	red   = "\033[31m"
	orng  = "\033[38;5;208m"
	rst   = "\033[0m"
)

func Banner() {
	fmt.Printf("%s################################################################################%s\n", bcyn, rst)
	fmt.Printf("%s                               PassCrax v2.2.1%s\n", borng, rst)
	fmt.Printf("%s                                by ph4mished%s\n", borng, rst)
	fmt.Printf("%s                             Copyright (C) 2025%s\n", borng, rst)
	fmt.Printf("%s             PassCrax is a simpler, and interactive hash cracker%s\n", bblu, rst)
	fmt.Printf("%sDISCLAIMER:%s The Developer Is Not In Any Way Responsible Or Liable For Any%s\n", bred, bylw, rst)
	fmt.Printf("%s                     Misuse Or Damage Caused With PassCrax%s\n", bylw, rst)
	fmt.Printf("%s################################################################################%s\n\n", bcyn, rst)
	fmt.Printf("%sType '%shelp%s%s' to see available commands.%s\n", grn, bylw, rst, grn, rst)
	fmt.Printf("%sIf the content of '%shelp%s%s' got you confused, just type '%srun%s%s' to get started.%s\n", grn, bylw, rst, grn, bylw, rst, grn, rst)
}
