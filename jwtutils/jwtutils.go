package jwtutils

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"os"
	"strings"
)

type JWT struct {
	Header    string
	Payload   string
	Signature string
}

func ParseToken(token string) (JWT, error) {
	token_slice := strings.Split(token, ".")
	if len(token_slice) != 3 {
		return JWT{}, errors.New("Invalid token")
	}
	var jwt = JWT{token_slice[0], token_slice[1], token_slice[2]}
	return jwt, nil
}

func CalculateSignature(jwt JWT, key string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	token := jwt.Header + "." + jwt.Payload
	h.Write([]byte(token))
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))
	signature = strings.Replace(signature, "=", "", -1)
	return signature, nil
}
func BruteForceToken(jwtstring, wordlist string) (string, error) {
	file, err := os.Open(wordlist)
	if err != nil {
		return "", err
	}
	jwt, err := ParseToken(jwtstring)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s1, err := CalculateSignature(jwt, scanner.Text())
		if err != nil {
			return "", err
		}
		if s1 == jwt.Signature {
			return scanner.Text(), nil
		}
	}
	return "", errors.New("Could not find key in list")
}
