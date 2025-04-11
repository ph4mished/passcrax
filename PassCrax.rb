#!/usr/bin/env ruby
require './help'
require './pass_analyzer'
require './wordlist_cracker'
require './brute_gen'

begin
  grn = "\e[32m"
  ylw = "\e[33m"
  blu = "\e[34m"
  red = "\e[31m"
  rst = "\e[0m"
  bcyn = "\e[1;36m"

  
  hash = nil
  hashtype = nil
  mode = nil
  start_len = nil
  end_len = nil
  

#Interactive mode
  loop do 
    #Display current status
    puts "\n#{bcyn}CURRENT SETTINGS#{rst}"
    puts "#{grn}Hash#{rst}: #{ylw}#{hash || 'Not set'}#{rst}"
    puts "#{grn}Hash Type#{rst}: #{ylw}#{hashtype || 'Not set'}#{rst}"
    puts "#{grn}Mode#{rst}: #{ylw}#{mode || 'Not set'}#{rst}"
    puts "#{grn}Brute Start Length#{rst}: #{ylw}#{start_len || 'Not set'}#{rst}"
    puts "#{grn}Brute End Length#{rst}: #{ylw}#{end_len || 'Not set'}#{rst}"
    
    print "\n> "
    input = gets.chomp.strip
    
    #Handle commands
    case input
    when "exit"
    puts "\n#{red}Terminated By User!#{rst}"
    puts"\n"
        break

    when "help"
         help()

    when /^set hash (.+)$/i
    #[1] is addded so hash id will be able to capture input
        hash = input.match(/^set hash (.+)$/i)[1]
        puts "\n#{grn}Hash set to:#{rst} #{ylw} #{hash}#{rst}"

        
    when /^set hashtype (.+)$/i
        hashtype = input.match(/^set hashtype (.+)$/i)[1]
        puts "\n#{grn}Hash Type set to:#{rst} #{ylw} #{hashtype}#{rst}"

    when /^hashid (.+)$/i
                   hash = input.match(/^hashid (.+)$/i)[1]
                  pass_analyze(hash)

     when /^set mode (.+)$/i
                  mode = input.match(/^set mode (.+)$/i)[1]
                  puts "\n#{grn}Mode set to:#{rst} #{ylw} #{mode}#{rst}"

    
    when /^start (\d+)$/i
                  start_len = input.match(/^start (\d+)$/i)[1]
                  puts "\n#{grn}Brute Start Length set to:#{rst} #{ylw} #{start_len}#{rst}"

    when /^end (\d+)$/i
                  end_len = input.match(/^end (\d+)$/i)[1]
               puts "\n#{grn}Brute End Length set to:#{rst} #{ylw} #{end_len}#{rst}"

                      
    when "run"
        if !hash
            puts "\n#{red}Error: No hash set!#{rst} #{grn}Use #{rst}#{ylw}'set hash <value>'#{rst}"
        elsif !hashtype
            puts "\n#{red}Error: No hash type set!#{rst} #{grn}Use #{rst}#{ylw}'set hashtype <value>'#{rst}"

        elsif mode == "brute"
              brute_gen(hash, hashtype, start_len, end_len)
                      
        elsif mode == "dict"
           pass_crack(hash, hashtype)
        end

   when "status"

    else
        puts "\n#{red}Unknown command: #{rst} #{grn}Type#{rst} #{ylw}'help'#{rst} #{grn}for options.#{rst}"
    end
end
rescue Interrupt
    puts "\n#{red}Terminated By User!#{rst}"
        puts "\n"
end

