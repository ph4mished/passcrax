#!/usr/bin/ruby
require 'digest'

puts "Enter text to hash:"
text = gets.chomp

puts "Enter text for salt:"
salt = gets.chomp
pure = Digest::SHA1.hexdigest(text)
enc_front = Digest::SHA1.hexdigest(salt+text)
enc_back = Digest::SHA1.hexdigest(text+salt)
enc_omni = Digest::SHA1.hexdigest(salt+text+salt)

puts "Result of #{text} (Pure hash): #{pure}"
puts "Result of #{text} (Salt infront) in hash: #{enc_front}"
puts "Result of #{text} (Salt behind) in hash: #{enc_back}"
puts "Result of #{text} (Salt on both) in hash: #{enc_omni}"