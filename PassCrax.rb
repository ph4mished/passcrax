#!/usr/bin/env ruby

require 'digest'

def pass_crack(hash, hashtype)
   grn = "\e[32m"
   blu = "\e[34m"
   ylw = "\e[33m"
   red = "\e[31m"
   rst = "\e[0m"

wordlist_dir = "Wordlist/"            
wordlist_files = Dir.glob("#{wordlist_dir}*.txt")
            if wordlist_files.empty?          
            puts "\n#{red}Error: No Files Found In #{wordlist_dir}#{rst}"
           return
           end
            wordlist_files.each do |file|
               puts "\n#{blu}Scanning File: #{file}...#{rst}"
               File.foreach(file) do |word|
               word.chomp!


   hash_type = case hashtype
    when "md5"
        Digest::MD5.hexdigest(word)
    
    when "sha1"
        Digest::SHA1.hexdigest(word)

    when "sha256"
        Digest::SHA256.hexdigest(word)

    when "sha384"
        Digest::SHA384.hexdigest(word)

    when "sha512"
        Digest::SHA512.hexdigest(word)



    else 
       puts "\n#{red}Error: Invalid Hash Type!#{rst}"
       puts "\n#{grn}Type#{rst} #{ylw}'help'#{rst} #{grn}for options.#{rst}"
       return
    end

           if hash_type == hash
            puts "\n#{grn}Password Found:#{rst} #{ylw} #{word} #{rst}"
            return word
           
          end   
        end
    end
    puts "\n#{red}Password Not Found!#{rst}"
end
   
  

