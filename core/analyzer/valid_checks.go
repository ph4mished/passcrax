package analyzer

import (
	"strings"
)

// Problem using "strings.Contains" accepts 'bl'
// and 'au' instead of 'blake2s' and 'auto'.
// it gets accepted because they are part of
// the spellings of such word. in the end they
// end up giving errors. there should be a way
// to accept full spellings so that no partial
// spellings are accepted.
func CheckValidHashType(hashtype string) string {
	valid_ones := []string{"md4", "md5", "sha1", "sha224", "sha256", "sha384", "sha512", "sha512_224", "sha512_256", "ripemd160", "adler32", "blake2b", "blake2s", "crc32", "crc64", "fnv1_32", "fnv1_64", "fnv1a_32", "fnv1a_64"}
	together := strings.Join(valid_ones, ", ")
	checkTrue := strings.Contains(together, hashtype)

	if checkTrue {
		return hashtype
	} else {
		return together
	}
	return hashtype
}

func CheckValidMode(mode string) string {
	var altogether string
	valid_mode := []string{"brute", "dict", "auto"}
	altogether = strings.Join(valid_mode, ", ")

	//this will be too long for the valid hastype check. theres gotta be another way. ill be back after refreshing to see this damn thing out
	if valid_mode[0] == mode || valid_mode[1] == mode || valid_mode[2] == mode {
		return mode
	} else {
		return altogether
	}
	return mode
}
