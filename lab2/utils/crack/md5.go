package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Of returns the hex-encoded MD5 of the provided input string.
func MD5Of(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// CompareHex compares an already-hex-encoded md5 hash with the md5 of candidate.
// Returns true when they match.
func CompareHex(targetHex string, candidate string) bool {
	return MD5Of(candidate) == targetHex
}
