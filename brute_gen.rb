#!/usr/bin/env ruby
require 'digest'  
def brute_gen(hash, hashtype, start_len, end_len)
    grn = "\e[32m"
    blu = "\e[34m"
    ylw = "\e[33m"
    red = "\e[31m"
    rst = "\e[0m"

    start_len = start_len.to_i
    end_len = end_len.to_i
         puts "\n#{ylw}Brute-Forcing...#{rst}"
         (start_len..end_len).each do |length|
            chars = "a" * length
            end_chars = "z" * length
            while chars <= end_chars

             hash_type = case hashtype
                 when "md5"
                     Digest::MD5.hexdigest(chars)
                 
                 when "sha1"
                     Digest::SHA1.hexdigest(chars)
             
                 when "sha256"
                     Digest::SHA256.hexdigest(chars)
             
                 when "sha384"
                     Digest::SHA384.hexdigest(chars)
             
                 when "sha512"
                     Digest::SHA512.hexdigest(chars)

                 else 
                    puts "\n#{red}Error: Invalid Hash Type!#{rst}"
                    puts "\n#{grn}Type#{rst} #{ylw}'help'#{rst} #{grn}for options.#{rst}"
                    return
               end

            if hash_type == hash
                puts "\n#{grn}Password Found:#{rst} #{ylw} #{chars} #{rst}"
                            return chars
            end
                   chars = chars.next
                              end
                            end
                  puts "\n#{red}Password Not Found!#{rst}"
end
