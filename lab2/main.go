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
	// Hardcoded target hash (from lab)
	target := "6a85dfd77d9cb35770c9dc6728d73d3f"

	// Hardcoded path to the password file you downloaded
	passwordFile := filepath.Join(".", "wordlist.txt")

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
	fmt.Println("Target MD5:", target)
	fmt.Println("Beginning cracking loop...")

	for scanner.Scan() {
		attempts++
		word := scanner.Text()

		// verbose line for every attempt
		fmt.Printf("[%d] trying: %s -> md5=%s\n", attempts, word, crack.MD5Of(word))

		// compare computed md5 with target
		if crack.CompareHex(target, word) {
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
