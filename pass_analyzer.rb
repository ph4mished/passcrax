#!/usr/bin/env ruby

HASH_PATTERNS = {
  "MD5"       => /^[a-f0-9]{32}$/i,
  "SHA-1"     => /^[a-f0-9]{40}$/i,
  "SHA-224"   => /^[a-f0-9]{56}$/i,
  "SHA-256"   => /^[a-f0-9]{64}$/i,
  "SHA-384"   => /^[a-f0-9]{96}$/i,
  "SHA-512"   => /^[a-f0-9]{128}$/i,
  "NTLM"      => /^[a-f0-9]{32}$/i,
  "LM Hash"   => /^[a-f0-9]{32}$/i,
  "MySQL v3+" => /^[a-f0-9]{16}$/i,
  "MySQL v5+" => /^\*[A-F0-9]{40}$/i,
  "bcrypt"    => /^\$2[ayb]\$.{56}$/i,
  "Argon2"    => /^\$argon2[a-z]+\$.+/i,
  "DES (Unix)"=> /^.{13}$/i
}

def find_hash(hash)
    grn = "\e[32m"
    blu = "\e[34m"
    ylw = "\e[33m"
    red = "\e[31m"
    rst = "\e[0m"

puts "\n#{blu}Enter Hash String For Analysis:#{rst}"
hash = gets.chomp.strip

  if hash.nil? || hash.empty?
      return "#{red}Empty Input!#{rst}"
      #FIXIT: Should restart after invalid input or empty input.
  end

  hash = hash.strip
 
  matches = HASH_PATTERNS.select { |name, regex| hash.match?(regex) }.keys
  return matches.join(" / ") unless matches.empty?
  "#{red}Unknown Hash Format!#{rst}"
end

def pass_analyze()
   grn = "\e[32m"
    blu = "\e[34m"
    ylw = "\e[33m"
    red = "\e[31m"
    rst = "\e[0m"

puts "#{grn}\nHash Type:#{rst} #{ylw}#{find_hash(hash)}#{rst}"
end
