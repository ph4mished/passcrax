package analyzer
import (
"strings"
)
//Problem using "string" accepts 'bl' and 'au' instead of 'blake2s' and 'auto'. it gets accepted because they are part of the spellings of such word. in the end they end up giving errors. there should be a way to accept full spellings so that no partial spellings are accepted.
func CheckValidHashType(hashtype string)string{
  valid_ones :=[]string{"md4", "md5", "sha1", "sha224", "sha256", "sha384", "sha512", "sha512_224", "sha512_256", "ripemd160", "adler32", "blake2b", "blake2s", "crc32", "crc64", "fnv1_32", "fnv1_64", "fnv1a_32", "fnv1a_64"}
  together := strings.Join(valid_ones, ", ")
  checkTrue := strings.Contains(together, hashtype)
  
   if checkTrue{
  return hashtype
  }else{
  return together
}
return hashtype
}

func CheckValidMode(mode string)string{
var altogether string
valid_mode :=[]string{"brute", "dict", "auto"}
  //altogether := strings.Join(valid_mode, ", ")
 // split := strings.Split(altogether, ",")
  //string.Contains isnt foolproof because words like 'br', 'dic' will be accepted although they arent valid
//  for i:= 0; i<=2; i++{
  //show = valid_mode[i]
  altogether = strings.Join(valid_mode, ", ")
 // checkTrue = strings.Compare(valid_mode[i], mode)
// }
  
 //this will be too long for the valid hastype check. theres gotta be another way. ill be back after refreshing to see this damn thing out 
  if strings.Compare(valid_mode[0], mode) == 0 || strings.Compare(valid_mode[1], mode) == 0 || strings.Compare(valid_mode[2], mode) == 0{
  return mode
  }else{ 
  return altogether
}
return mode
}
