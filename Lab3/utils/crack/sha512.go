package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512Of returns the hex-encoded SHA512 of the provided input string.
func SHA512Of(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// CompareHexSHA512 compares a target hex-encoded sha512 with the sha512 of candidate.
func CompareHexSHA512(targetHex string, candidate string) bool {
	return SHA512Of(candidate) == targetHex
}
