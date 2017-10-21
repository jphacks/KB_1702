package util

import (
	"encoding/hex"
	"hash/fnv"
	"time"
)

var h = fnv.New64a()

func GetHash(str string) string {
	h.Write([]byte(time.Now().Round(time.Hour).Add(-1 * time.Hour).String()))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
