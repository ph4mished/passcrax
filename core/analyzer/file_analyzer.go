package analyzer

import (
        "bufio"
        "errors"
        "fmt"
        "io/fs"
        "os"
        "strings"

        //"github.com/fatih/color"
)

func percentage(num, total int) float64 {
        return (float64(num) / float64(total)) * 100
}


func FileAnalyze(hashFile string) string {
    hash_counts := make(map[string]int)
    total_hashes := 0
    var allResults []string

    file, err := os.Open(hashFile)
    if err != nil {
        bred.Printf("\n[!] Error: %v", err)
        return ""
    }
    defer file.Close()
    
    if _, err = os.Stat(hashFile); errors.Is(err, fs.ErrNotExist) {
        bred.Printf("\n[!] Error: %s does not exist\n", hashFile)
        return ""
    }

    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        hashLine := strings.TrimSpace(scanner.Text())
        if len(hashLine) > 0 {
            total_hashes++
            hashtype := PassAnalyze(hashLine)
            hash_counts[hashtype]++
        }
    }

   
    bgrn.Print("\n[~] Found ")
    borng.Print(total_hashes)
    bgrn.Print(" hashes in ")
    bblu.Print(hashFile)

   
    for hashtype, count := range hash_counts {
        pct := percentage(count, total_hashes)
        
        
        result := bgrn.Sprint("\n\n{") + bylw.Sprintf("%.2f%%", pct) + bgrn.Sprint("}") + bcyn.Sprintf(" %d/%d Of The Hashes Are:\n", count, total_hashes) + fmt.Sprintf(" %s", hashtype)

     //   allResults = strings.Join(result, "\n")
        allResults = append(allResults, result)
    }
    resultStr := strings.Join(allResults, "\n")
    return resultStr
}
