package cracker

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PassCrack(hash string, hashtype string) string {

	var wordlist_dir = "Wordlist/"
	wordlist_files, err := filepath.Glob(filepath.Join(wordlist_dir, "*.txt"))
	if err != nil {
		fmt.Printf("\n%sError Scanning Wordlist Directory: %v %s\n", red, err, rst)
		return ""
	}
	if len(wordlist_files) == 0 {
		fmt.Printf("\n%sError: No Files Found In %s%s\n", red, wordlist_dir, rst)
		return ""
	}
	for _, filename := range wordlist_files {

		fmt.Printf("\n%sScanning File: %s...%s", blu, filename, rst)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("\n%sError: File Cannot Be Opened!%s\n", red, rst)
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			word := scanner.Text()
			word = strings.TrimSpace(word)

			var hash_type string
			switch hashtype {
			case "md5":
				data := md5.Sum([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			case "sha1":
				data := sha1.Sum([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			case "sha224":
				data := sha256.Sum224([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			case "sha256":
				data := sha256.Sum256([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			case "sha384":
				data := sha512.Sum384([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			case "sha512":
				data := sha512.Sum512([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			case "sha512_224":
				data := sha512.Sum512_224([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			case "sha512_256":
				data := sha512.Sum512_256([]byte(word))
				hash_type = hex.EncodeToString(data[:])

			default:
				fmt.Printf("\n%sError: Hash Type Is Invalid: %s%s\n", red, hashtype, rst)
				fmt.Printf("\n%sType%s %s'help'%s %sfor options.%s\n", grn, rst, ylw, rst, grn, rst)
				file.Close()
				return ""
			}

			if hash_type == hash {
				fmt.Printf("\n%sPassword Found:%s %s%s%s\n", grn, rst, ylw, word, rst)
				file.Close()
				return word
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("\n%sError Reading File: %v %s\n", red, err, rst)
		}
		file.Close()
	}
	fmt.Printf("\n%sPassword Not Found!%s\n", red, rst)
	return ""
}
