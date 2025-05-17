package utils
import (
//	"golang.org/x/crypto/blake2b"
//	"golang.org/x/crypto/blake2s"
//	"golang.org/x/crypto/md4"
	//"golang.org/x/crypto/ripemd160"
)

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

func HashFormats(word, hashType string) (string, error) {

	switch hashType {

/*	case "md4":
		hasher := md4.New()
		hasher.Write([]byte(word))
		data := hasher.Sum(nil)
		return hex.EncodeToString(data[:]), nil
*/
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

	case "sha512_224":
		data := sha512.Sum512_224([]byte(word))
		return hex.EncodeToString(data[:]), nil

	case "sha512_256":
		data := sha512.Sum512_256([]byte(word))
		return hex.EncodeToString(data[:]), nil

	//case "ripemd160":
	//	hasher := ripemd160.New()
	//	hasher.Write([]byte(word))
	//	data := hasher.Sum(nil)
	//	return hex.EncodeToString(data[:]), nil

/*	case "blake2b":
		hasher, err := blake2b.New512(nil)
		if err != nil{
		              return "", err
		          }
		hasher.Write([]byte(word))
		data := hasher.Sum()
		return hex.EncodeToString(data), nil

	case "blake2s":
		hasher, err := blake2s.New256(nil)
		if err != nil{
			return "", err
		}
		hasher.Write([]byte(word))
		data := hasher.Sum()
		return hex.EncodeToString(data[:]), nil*/
		
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

	case "fnv1_32":
		hasher := fnv.New32()
		hasher.Write([]byte(word))
		data := hasher.Sum32()
		return fmt.Sprintf("%08x", data), nil

	case "fnv1_64":
		hasher := fnv.New64()
		hasher.Write([]byte(word))
		data := hasher.Sum64()
		return fmt.Sprintf("%016x", data), nil

	case "fnv1a_32":
		hasher := fnv.New32a()
		hasher.Write([]byte(word))
		data := hasher.Sum32()
		return fmt.Sprintf("%08x", data), nil

	case "fnv1a_64":
		hasher := fnv.New64a()
		hasher.Write([]byte(word))
		data := hasher.Sum64()
		return fmt.Sprintf("%016x", data), nil

	default:
		fmt.Printf("\n%sError: Hash Type Is Invalid: %s%s\n", bred, hashType, rst)
		fmt.Printf("\n%sType%s %s'help'%s %sfor options.%s\n", bgrn, rst, bylw, rst, bgrn, rst)
		//goto NEXT
		return "", nil
	}
	//NEXT:
	//return "", nil
}
