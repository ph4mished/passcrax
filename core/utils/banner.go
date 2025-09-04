package utils

import (
//	"fmt"

	"github.com/fatih/color"
)

var (
	bgrn = color.New(color.FgGreen, color.Bold)
	bred = color.New(color.FgRed, color.Bold)
	bblu = color.New(color.FgBlue, color.Bold)
	bcyn = color.New(color.FgCyan, color.Bold)
	bylw = color.New(color.FgYellow, color.Bold)
	grn  = color.New(color.FgGreen)
	red  = color.New(color.FgRed)
	blu  = color.New(color.FgBlue)
	cyn  = color.New(color.FgCyan)
	ylw  = color.New(color.FgYellow)
)

func Banner() {
	bcyn.Printf("################################################################################\n")
	//next version should be. 2.3.0
	//this tool really needs a huge code clean up. so much hanging, dead and dirty code. These damn code written by me some months ago.
	//This shows I've improved greatly enough to tell my former codes were bad
	//all redundant functions will be cleared. all things just stuck up for later fixes will be trimmed.
	bylw.Printf("                               PassCrax v2.2.2\n")
	bylw.Printf("                                by ph4mished\n")
	bylw.Printf("                             Copyright (C) 2025\n")
	bblu.Printf("             PassCrax is a simpler, and interactive hash cracker\n")
	bred.Printf("DISCLAIMER: ")
	bylw.Printf("The Developer Is Not In Any Way Responsible Or Liable For Any\n")
	bylw.Printf("                     Misuse Or Damage Caused With PassCrax\n")
	bcyn.Printf("################################################################################\n\n")
	
	grn.Printf("Type '")
	bylw.Printf("help")
	grn.Printf("' to see available commands.\n")
	
	grn.Printf("If the content of '")
	bylw.Printf("help")
	grn.Printf("' got you confused, just type '")
	bylw.Printf("run")
	grn.Printf("' to get started.\n")
}
