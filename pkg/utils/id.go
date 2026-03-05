package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateID(prefix string) string {
	bytes := make([]byte, 4) 
	if _, err := rand.Read(bytes); err != nil {
		return fmt.Sprintf("%s-fallback", prefix)
	}
	return fmt.Sprintf("%s-%s", prefix, hex.EncodeToString(bytes))
}
