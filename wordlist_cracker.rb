#!/usr/bin/env ruby

def help()
        ylw = "\e[33m"
        grn = "\e[32m"
        red = "\e[31m"
        blu = "\e[34m"
        rst = "\e[0m"
        bcyn = "\e[1;36m"
puts "\n\t\t#{bcyn}Available Commands:#{rst}"
puts "#{blu}set hash <value>#{rst} - #{grn}Set hash (eg.,21232f297a57a5a743894a0e4a801fc3#{rst} )"
puts "#{blu}set hashtype <value>#{rst}   -#{grn} Set hashtype (e.g. md5)#{rst}"
puts "#{blu}run#{rst}             - #{grn}Execute password cracking#{rst}"
puts "#{blu}status#{rst}             - #{grn}Show current settings#{rst}"
puts "#{blu}exit#{rst}#{grn} or#{rst}#{blu} Ctrl+c#{rst}               - #{grn}Quit program#{rst}"
puts "#{blu}hashid <value>#{rst}     - #{grn}Identify hash type (eg. hashid 49f68a5c8493ec2c0bf489821c21fc3b )#{rst}"
puts "#{blu}help#{rst}               - #{grn}Show this help#{rst}"

puts "\n\t\t#{bcyn}Hash Type Values To Use:#{rst}"
puts "#{grn}set hashtype <value>     - #{blu}values = md5, sha1, sha256, sha384, sha512#{rst}"
end
