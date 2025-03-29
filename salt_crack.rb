#!/usr/bin/env ruby

require 'digest'

def salt_crack(hash)
   grn = "\e[32m"
   blu = "\e[34m"
   ylw = "\e[33m"
   red = "\e[31m"
   rst = "\e[0m"

    puts "\n#{blu}Enter Salt:#{rst}"
   salt = gets.chomp.strip.downcase
   puts "\n#{blu}Enter Salted Hash String:#{rst}"
    hash = gets.chomp.strip.downcase
   puts "\n#{blu}SELECT HASH FUNCTIONS\n#{rst} #{grn}[1] MD5\n [2] SHA-1\n [3] SHA-256\n [4] SHA-384\n [5] SHA-512\n#{rst}"
   input = gets.chomp.to_i

   case input
    when 1
        hash_type = Digest::MD5
    
    when 2
        hash_type = Digest::SHA1

    when 3
        hash_type = Digest::SHA256

    when 4
        hash_type = Digest::SHA384

    when 5
        hash_type = Digest::SHA512

    else 
        puts "\n#{red}Error: Invalid Input!#{rst}"
        return
    end

    wordlist_dir = "Wordlist/"
    wordlist_files = Dir.glob("#{wordlist_dir}*.txt")
    if wordlist_files.empty?
        puts "\n#{red}Error: No Files Found In #{wordlist_dir}#{rst}"
    return
   end
    wordlist_files.each do |file|
        puts "\n#{ylw}Scanning File: #{file}...#{rst}"
        
       File.foreach(file) do |word|
           word.chomp!
           encoded = hash_type.hexdigest(word + salt)
           if encoded == hash
            puts "\n#{grn}Password Found:#{rst} #{ylw} #{word} #{rst}"
            return word
           
          end   
        end
    end
    puts "\n#{red}Password Not Found!#{rst}"
end
  
#UPDATES
   #make it possible for salt cracking in the back and both sides too eg. (salt+word+salt) and (word+salt)
   
  
   