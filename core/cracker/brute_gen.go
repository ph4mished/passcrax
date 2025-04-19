package cracker

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

const (
	grn = "\033[32m"
	blu = "\033[34m"
	ylw = "\033[33m"
	red = "\033[31m"
	rst = "\033[0m"
)

var charset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.,;!?:'*#€@_&-+()/✓[]}{><∆§×÷=°^¢$¥£~|•√π`")

func BruteGen(hash string, hashtype string, start_len int, end_len int) string {
	if start_len < 1 || end_len < 1 || start_len > end_len {
		fmt.Printf("\n%sError: Invalid Length Parameters%s", red, rst)
		return ""
	}

	fmt.Printf("\n%sBrute-Forcing From Length %d To %d...%s\n", ylw, start_len, end_len, rst)

	for length := start_len; length <= end_len; length++ {
		fmt.Printf("%sTrying Length: %d%s\n", blu, length, rst)

		total := 1
		for i := 0; i < length; i++ {
			total *= len(charset)
		}

		for i := 0; i < total; i++ {
			word := make([]rune, length)
			n := i
			for j := length - 1; j >= 0; j-- {
				word[j] = charset[n%len(charset)]
				n /= len(charset)
			}

			input := string(word)
			var hash_type string

			switch hashtype {
			case "md5":
				data := md5.Sum([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			case "sha1":
				data := sha1.Sum([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			case "sha224":
				data := sha256.Sum224([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			case "sha256":
				data := sha256.Sum256([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			case "sha384":
				data := sha512.Sum384([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			case "sha512":
				data := sha512.Sum512([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			case "sha512_224":
				data := sha512.Sum512_224([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			case "sha512_256":
				data := sha512.Sum512_256([]byte(input))
				hash_type = hex.EncodeToString(data[:])
			default:
				fmt.Printf("\n%sError: Invalid Hash Type: %s%s", red, hashtype, rst)
				fmt.Printf("\n%sType%s %s'help'%s %sfor options.%s", grn, rst, ylw, rst, grn, rst)
				return ""
			}

			if hash_type == hash {
				fmt.Printf("\n%sPassword Found:%s %s%s%s\n", grn, rst, ylw, input, rst)
				return input
			}
		}
	}

	fmt.Printf("\n%sPassword Not Found!%s\n", red, rst)
	return ""
}
