package file

import (
	"passcrax/core/crack"
	//	"passcrax/core/utils"
	//
	// "passcrax/core/rules"
)

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func DictFile(dict_dir string, hashFile string, hashtype string, outputFile string){
	//dangerous function.. can cause OOM if hashFile is too huge
	//	filename := utils.FileLaunch(hashFile, os.O_RDWR, 0755)
	file, err := os.Open(hashFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	outFile, _ := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer outFile.Close()
	header := fmt.Sprintf("\n\n\n###---These Are Results From %s---\n### Time: %v\n\n", hashFile, time.Now())
	outFile.Write([]byte(header))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		targetHash := strings.TrimSpace(scanner.Text())
		if len(dict_dir) == 0 {
			dict_dir = ""
		}
		if targetHash != ""{
		now := crack.PassCrack(dict_dir, targetHash, hashtype, "")
		results := fmt.Sprintf("%s : %s\n", targetHash, now)
		outFile.Write([]byte(results))
		}
	}
	return
}

func BruteFile(hashFile string, hashtype string, charset string, outputFile string, min int, max int) {
	file, err := os.Open(hashFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	outFile, _ := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	header := fmt.Sprintf("\n\n\n###---These Are Results From %s---\n### Time: %v\n\n", hashFile, time.Now())
	outFile.Write([]byte(header))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		targetHash := strings.TrimSpace(scanner.Text())
if targetHash != ""{
		non := crack.BruteGen(targetHash, hashtype, charset, min, max)
		bond := fmt.Sprintf("%s : %s\n", targetHash, non)
		outFile.Write([]byte(bond))
		}
	}
}
