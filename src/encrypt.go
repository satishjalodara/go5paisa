package go5paisa

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

func pkcs7Pad(b []byte, blocksize int) []byte {
	if blocksize <= 0 {
		return nil
	}
	if b == nil || len(b) == 0 {
		return nil
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb
}

// Encrypt implements AES encryption
func encrypt(key string, plaintext string) string {
	var iv = []byte{83, 71, 26, 58, 54, 35, 22, 11,
		83, 71, 26, 58, 54, 35, 22, 11}
	derivedKey := pbkdf2.Key([]byte(key), iv, 1000, 48, sha1.New)
	newiv := derivedKey[0:16]
	newkey := derivedKey[16:48]
	plaintextBytes := pkcs7Pad([]byte(plaintext), aes.BlockSize)
	block, err := aes.NewCipher(newkey)
	if err != nil {
		panic(err)
	}
	n := aes.BlockSize - (len(plaintext) % aes.BlockSize)
	ciphertext := make([]byte, len(plaintext)+n*2)
	mode := cipher.NewCBCEncrypter(block, newiv)
	mode.CryptBlocks(ciphertext, plaintextBytes)
	cipherstring := base64.StdEncoding.EncodeToString(ciphertext[:len(ciphertext)-n])
	return cipherstring
}
