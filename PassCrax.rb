#!/usr/bin/env ruby
require './wordlist_cracker'
require './pass_analyzer'
require './salt_crack'
begin
   grn = "\e[32m"
   blu = "\e[34m"
   ylw = "\e[33m"
   red = "\e[31m"
   rst = "\e[0m"
   cyn = "\e[36m"
   loop do
       
   pass_crax = <<~ART
    #{cyn} ____    __                ____    ____     __          #{rst}
    #{cyn}|  /\\\\  /  \\   ___   ___  /\\___\\  |  /\\\\   /  \\  /\\    /\\ #{rst}
    #{ylw}|_/_// / /\\ \\ /  /  /  / / /      |_/_//  / /\\ \\ \\ \\  / /#{rst}
    #{ylw}| /    \\/--\\/ \\___  \\___ | \\      | \\_    \\/__\\/  \\//\\/#{rst}
    #{cyn}/|     /    \\     \\     \\ \\/____  /|  \\/\\ /    \\ /\\/  \\/\\ #{rst}
    #{cyn}\\|     \\/  \\/ /___/ /___/  \\____/ \\|   \\/ \\/  \\/ \\/    \\/#{rst}
    ART
   puts "\n\n\n#{pass_crax}"

   pass_analyze()

  
   puts "\n#{blu}SELECT ONE\n#{rst}#{grn} [1]Salted Password Cracker\n [2]Pure Password Cracker#{rst}"
   slct = gets.chomp.to_i

       case slct
       when 1
           salt_crack(hash)

       when 2
           pass_crack(hash)

       else
           print "\n#{red}Error: Invalid Input!#{rst}"
   end
   end

rescue Interrupt
    print "\n#{red}Terminated By User!#{rst}"
    print "\n"
    end
