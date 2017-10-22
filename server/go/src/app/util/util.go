package util

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strconv"

	"golang.org/x/crypto/scrypt"
)

// GetRandomString 16桁のランダムな文字列を返す
func getRandomString() string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, 16)
	for i := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}

// CreateTokenHash 渡された文字列とsaltからTokenハッシュを生成する
func CreateTokenHash(str string) string {
	converted, _ := scrypt.Key([]byte(str), []byte(getRandomString()), 16384, 8, 1, 10)
	return hex.EncodeToString(converted[:])
}
