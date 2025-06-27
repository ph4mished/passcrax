package analyzer

import (
	"strings"
)

func CheckValidHashType(hashtype string) string {
	valid_ones := []string{"md4", "md5", "sha1", "sha224", "sha256", "sha384", "sha512", "sha3-224", "sha3-256", "sha3-384", "sha3-512", "sha512-224", "sha512-256", "adler32", "crc32", "crc64", "fnv1-32", "fnv1-64", "fnv1a-32", "fnv1a-64", "bcrypt"}
	together := strings.Join(valid_ones, ", ")
	for _, iterHashtype := range valid_ones {
		if iterHashtype == hashtype {
			return hashtype
		}
	}
	return together
}

func CheckValidMode(mode string) string {
	var altogether string
	valid_mode := []string{"brute", "dict", "auto"}
	altogether = strings.Join(valid_mode, ", ")
	for _, iterMode := range valid_mode {
		if iterMode == mode {
			return mode
		}
	}
	return altogether
}
