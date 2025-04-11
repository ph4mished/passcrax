#!/usr/bin/env ruby
def help()
        ylw = "\e[33m"
        grn = "\e[32m"
        red = "\e[31m"
        blu = "\e[34m"
        rst = "\e[0m"
        bcyn = "\e[1;36m"
puts "\n\t\t#{bcyn}Available Commands:#{rst}"
puts "#{blu}set hash <value>#{rst} - #{grn}Used to set hash (eg.,set hash 21232f297a57a5a743894a0e4a801fc3#{rst} )"
puts "#{blu}set hashtype <value>#{rst}   -#{grn}Used to set hashtype (e.g.set hashtype  md5)#{rst}"
puts "\n#{blu}run#{rst}             - #{grn}Execute password cracking#{rst}"
puts "#{blu}status#{rst}             - #{grn}Show current settings#{rst}"
puts "#{blu}exit#{rst}#{grn} or#{rst}#{blu} Ctrl+c#{rst}               - #{grn}Quit program#{rst}"
puts "\n#{blu}hashid <value>#{rst}     - #{grn}Identify hash type (eg. hashid 49f68a5c8493ec2c0bf489821c21fc3b )#{rst}"
puts "#{blu}help#{rst}               - #{grn}Show this help#{rst}"
puts "#{blu}start <number>#{rst}        - #{grn}brute force start length#{rst}"
puts "#{blu}end <number>#{rst}        -#{grn}brute force end length#{rst}"
puts "#{blu}set mode <value>     - #{grn}mode#{rst}"

puts "\n\t\t#{bcyn} Values To Use:#{rst}"
puts "#{grn}set hashtype <value>     - #{blu}values = md5, sha1, sha256, sha384, sha512#{rst}"
puts "#{grn}set mode <value>     - #{blu}values = brute, dict#{rst}"

puts "\n#{ylw}*Only Brute Mode requires start and end length#{rst}"   
end
