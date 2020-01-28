package auth

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// GenerateAuthCode 認証コード作成のロジック
func GenerateAuthCode(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return err.Error()
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func GenerateRequestToken(tokenLength int) string {
	token := iniRandomString(tokenLength)
	return token
}

func iniRandomString(n int) string {
	strings := hex.EncodeToString(bytes(n))
	return strings
}

func bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
