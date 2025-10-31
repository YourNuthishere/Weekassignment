package crack


import (
"crypto/sha1"
"encoding/hex"
)


// SHA1Of returns the hex-encoded SHA1 of the provided input string.
func SHA1Of(s string) string {
h := sha1.New()
h.Write([]byte(s))
return hex.EncodeToString(h.Sum(nil))
}


// CompareHexSHA1 compares a target hex-encoded sha1 with the sha1 of candidate.
func CompareHexSHA1(targetHex string, candidate string) bool {
return SHA1Of(candidate) == targetHex
}