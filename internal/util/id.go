package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func NewID(prefix string) string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return fmt.Sprintf("%s_fallback", prefix)
	}
	return fmt.Sprintf("%s_%s", prefix, hex.EncodeToString(b[:]))
}
