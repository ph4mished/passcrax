package utils

//this function isn't a util so it's not needed in this package

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"os"
)

//a hash vérifier is needed to check the inputted hash to see if it's supported. because

func HashFormats(word, hashType string) (string, bool) {

	switch hashType {
	case "md4":
		md4 := md4.New()
		md4.Write([]byte(word))
		data := md4.Sum(nil)
		return hex.EncodeToString(data[:]), true

	case "md5":
		data := md5.Sum([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha1":
		data := sha1.Sum([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha224":
		data := sha256.Sum224([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha256":
		data := sha256.Sum256([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha384":
		data := sha512.Sum384([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha512":
		data := sha512.Sum512([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha3-224":
		data := sha3.Sum224([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha3-256":
		data := sha3.Sum256([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha3-384":
		data := sha3.Sum384([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha3-512":
		data := sha3.Sum512([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha512-224":
		data := sha512.Sum512_224([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "sha512-256":
		data := sha512.Sum512_256([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "adler32":
		data := adler32.Checksum([]byte(word))
		return fmt.Sprintf("%08x", data), true

	case "crc32":
		data := crc32.ChecksumIEEE([]byte(word))
		return fmt.Sprintf("%08x", data), true

	case "crc64":
		table := crc64.MakeTable(crc64.ECMA)
		data := crc64.Checksum([]byte(word), table)
		return fmt.Sprintf("%016x", data), true

	case "fnv1-32":
		hasher := fnv.New32()
		hasher.Write([]byte(word))
		data := hasher.Sum32()
		return fmt.Sprintf("%08x", data), true

	case "fnv1-64":
		hasher := fnv.New64()
		hasher.Write([]byte(word))
		data := hasher.Sum64()
		return fmt.Sprintf("%016x", data), true

	case "fnv1a-32":
		hasher := fnv.New32a()
		hasher.Write([]byte(word))
		data := hasher.Sum32()
		return fmt.Sprintf("%08x", data), true

	case "fnv1a-64":
		hasher := fnv.New64a()
		hasher.Write([]byte(word))
		data := hasher.Sum64()
		return fmt.Sprintf("%016x", data), true

	case "blake2b-256":
		data := blake2b.Sum256([]byte(word))
		return hex.EncodeToString(data[:]), true

	case "blake2b-512":
		data := blake2b.Sum512([]byte(word))
		return hex.EncodeToString(data[:]), true

	default:
		bred.Printf("\nError: Hash Type Is Invalid: %s\n", hashType)
		bgrn.Print("\nType")
		bylw.Print(" 'identify <hashstring>'")
		bgrn.Print(" for supported hashtypes.\n")
		os.Exit(1)
		return "", false

	}
}
