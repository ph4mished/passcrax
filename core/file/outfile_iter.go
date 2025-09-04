package file

import (
	"fmt"
	"math/rand"
)

func IterName() string {
	var namet string
	names := []string{"cracked", "resolved", "pwned", "digest_match", "result", "reveal", "decrypted", "valid", "unlocked", "matched", "done", "found", "success", "completed"}
	//for _, namet = range names {
	//	for j := 1; j < 100; j++ {

	namet = names[rand.Intn(len(names))]
	num := rand.Intn(100)
	return fmt.Sprintf("%s%d", namet, num)
	//}
	//}
}
