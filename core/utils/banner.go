package utils
import "fmt"

func Banner(){
	fmt.Printf("%s################################################################################%s\n", bcyn, rst) 
	fmt.Printf("%s                               PassCrax v2.1.1%s\n", borng, rst) 
	fmt.Printf("%s                                by ph4mished%s\n", borng, rst)
	fmt.Printf("%s                             Copyright (C) 2025%s\n", borng, rst)
	fmt.Printf("%s             PassCrax is a simpler, and interactive hash cracker%s\n", bblu, rst)
	fmt.Printf("%sDISCLAIMER:%s The Developer Is Not In Any Way Responsible Or Liable For Any%s\n", bred, bylw, rst)
	fmt.Printf("%s                     Misuse Or Damage Caused By PassCrax%s\n", bylw, rst)
	fmt.Printf("%s################################################################################%s\n\n",bcyn, rst) 
	fmt.Printf("%sType '%shelp%s%s' to see available commands.%s\n", grn, bylw, rst, grn, rst) 
	//fmt.Println()
	
}
