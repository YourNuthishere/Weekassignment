/*

//lab0 
package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"strings"
	"crypto/sha1"
	"crypto/sha512"


	"golang.org/x/crypto/sha3"
)

func comparehash(text1, text2 string) {
	hashFunctions := []struct {
		name string
		hash func(string) string
	}{
		{"MD5", func(s string) string { return fmt.Sprintf("%x", md5.Sum([]byte(s))) }},
		{"SHA1", func(s string) string { return fmt.Sprintf("%x", sha1.Sum([]byte(s))) }},
		{"SHA512", func(s string) string { return fmt.Sprintf("%x", sha512.Sum512([]byte(s))) }},
		{"SHA256", func(s string) string { return fmt.Sprintf("%x", sha256.Sum256([]byte(s))) }},
		{"SHA3-256", func(s string) string { return fmt.Sprintf("%x", sha3.Sum256([]byte(s))) }},
	}

	for _, h := range hashFunctions {
		outputA := h.hash(text1)
		outputB := h.hash(text2)

		match := "No Match!"
		if strings.EqualFold(outputA, outputB) {
			match = "Match!"
		}

		fmt.Printf("Hash (%s) : %s / %s => %s\n", h.name, outputA, outputB, match)
	}
}

func main() {
	var input1, input2 string

	fmt.Println("======== Huon Sreynuth + Hashing Program ========")
	fmt.Print("Enter value one: ")
	fmt.Scanln(&input1)

	fmt.Print("Enter value two: ")
	fmt.Scanln(&input2)

	comparehash(input1, input2)
}

*/

/*

//find flag from that md5

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := "b6ccb4ece5454dc" // from example

	sum := md5.Sum([]byte(hash))
	result := hex.EncodeToString(sum[:])

	fmt.Println("MD5 of token:", result)
	fmt.Println("Flag: cryptoCTF{meowmeow}")
}

*/

/*
// verify

package main

import (
	"fmt"
	"regexp"
)

func main() {
	pat := `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$` // pattern from challenge
	test := "cryptoCTF{meowmeow}"
	ok := regexp.MustCompile(pat).MatchString(test)
	fmt.Println("matches:", ok) 
}
*/