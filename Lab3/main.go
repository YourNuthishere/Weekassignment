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
	// Hardcoded target SHA512 (concatenated)
	target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	// Hardcoded path to the password file
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
	fmt.Println("Target SHA512:", target)
	fmt.Println("Beginning cracking loop...")

	for scanner.Scan() {
		attempts++
		word := scanner.Text()

		// verbose line for every attempt
		fmt.Printf("[%d] trying: %s -> sha512=%s\n", attempts, word, crack.SHA512Of(word))

		// compare computed sha512 with target
		if crack.CompareHexSHA512(target, word) {
			elapsed := time.Since(start)
			fmt.Println("=== MATCH FOUND ===")
			fmt.Println("Password:", word)
			fmt.Println("Attempts:", attempts)
			fmt.Println("Elapsed:", elapsed)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return
	}

	fmt.Println("No match found after", attempts, "attempts.")
	fmt.Println("End time:", time.Now().Format(time.RFC3339))
}
