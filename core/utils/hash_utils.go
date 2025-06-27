package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

func HashFormats(word, hashType string) (string, error) {

	switch hashType {
	case "md4":
		md4 := md4.New()
		md4.Write([]byte(word))
		data := md4.Sum(nil)
		return hex.EncodeToString(data[:]), nil

	case "md5":
		data := md5.Sum([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha1":
		data := sha1.Sum([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha224":
		data := sha256.Sum224([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha256":
		data := sha256.Sum256([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha384":
		data := sha512.Sum384([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha512":
		data := sha512.Sum512([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha3-224":
		data := sha3.Sum224([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha3-256":
		data := sha3.Sum256([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha3-384":
		data := sha3.Sum384([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha3-512":
		data := sha3.Sum512([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha512-224":
		data := sha512.Sum512_224([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha512-256":
		data := sha512.Sum512_256([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "adler32":
		data := adler32.Checksum([]byte(word))
		return fmt.Sprintf("%08x", data), nil

	case "crc32":
		data := crc32.ChecksumIEEE([]byte(word))
		return fmt.Sprintf("%08x", data), nil

	case "crc64":
		table := crc64.MakeTable(crc64.ECMA)
		data := crc64.Checksum([]byte(word), table)
		return fmt.Sprintf("%016x", data), nil

	case "fnv1-32":
		hasher := fnv.New32()
		hasher.Write([]byte(word))
		data := hasher.Sum32()
		return fmt.Sprintf("%08x", data), nil

	case "fnv1-64":
		hasher := fnv.New64()
		hasher.Write([]byte(word))
		data := hasher.Sum64()
		return fmt.Sprintf("%016x", data), nil

	case "fnv1a-32":
		hasher := fnv.New32a()
		hasher.Write([]byte(word))
		data := hasher.Sum32()
		return fmt.Sprintf("%08x", data), nil

	case "fnv1a-64":
		hasher := fnv.New64a()
		hasher.Write([]byte(word))
		data := hasher.Sum64()
		return fmt.Sprintf("%016x", data), nil

		/*case "blake2b-128":
			 blake2b_128 := blake2b.New128()
			    blake2b_128.Write([]byte(word))
				data := blake2b_128.Sum(nil)
		     	return hex.EncodeToString(data[:]), nil*/

	case "blake2b-256":
		data := blake2b.Sum256([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "blake2b-512":
		data := blake2b.Sum512([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "bcrypt":
		data, _ := bcrypt.GenerateFromPassword([]byte(word), bcrypt.DefaultCost)
		//ok := bcrypt.CompareHashAndPassword(hash, []byte(word)) == nil
		return string(data), nil
		//hex.EncodeToString(data[:]), nil

	default:
		fmt.Printf("\n%sError: Hash Type Is Invalid: %s%s\n", bred, hashType, rst)
		fmt.Printf("\n%sType%s %s'help'%s %sfor supported hashtypes.%s\n", bgrn, rst, bylw, rst, bgrn, rst)
		return "", nil

	}
}
