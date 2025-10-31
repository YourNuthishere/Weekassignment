package main


import (
"bufio"
"fmt"
"os"
"path/filepath"
"time"


"password-cracker/utils/crack"
)


func main() {
// Hardcoded target SHA1 (from lab)
target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"


// Hardcoded path to the password file you downloaded
passwordFile := filepath.Join(".", "nord_vpn.txt")


f, err := os.Open(passwordFile)
if err != nil {
fmt.Println("Error opening password file:", err)
return
}
defer f.Close()


scanner := bufio.NewScanner(f)


attempts := 0
start := time.Now()
fmt.Println("Start time:", start.Format(time.RFC3339))
fmt.Println("Reading words from:", passwordFile)
fmt.Println("Target SHA1:", target)
fmt.Println("Beginning cracking loop...")


for scanner.Scan() {
attempts++
word := scanner.Text()


// verbose line for every attempt
fmt.Printf("[%d] trying: %s -> sha1=%s\n", attempts, word, crack.SHA1Of(word))


// compare computed sha1 with target
if crack.CompareHexSHA1(target, word) {
elapsed := time.Since(start)
fmt.Println("=== MATCH FOUND ===")
fmt.Println("Password:", word)
fmt.Println("Attempts:", attempts)
fmt.Println("Elapsed:", elapsed)
return
}
}


// scanner loop ended without match
if err := scanner.Err(); err != nil {
fmt.Println("Scanner error:", err)
return
}


fmt.Println("No match found after", attempts, "attempts.")
fmt.Println("End time:", time.Now().Format(time.RFC3339))
}

