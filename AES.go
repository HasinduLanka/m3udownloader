package main

import (
	"crypto/aes"
	"encoding/hex"
)

func TestEncryption() {
	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "This is a secret"

	c := EncryptAESStringToHex(([]byte(key))[:32], pt)

	// plaintext
	println(pt)

	// ciphertext
	println(c)

	// decrypt
	d := DecryptAESHex([]byte(key)[:32], c)
	println(d)
}

func EncryptAESStringToHex(key []byte, plaintext string) string {

	out := EncryptAESString(key, plaintext)
	return hex.EncodeToString(out)
}

func EncryptAESString(key []byte, plaintext string) []byte {

	out := EncryptAES(key, []byte(plaintext))
	return out
}

func EncryptAES(key []byte, data []byte) []byte {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(data))

	c.Encrypt(out, data)

	return out
}

func DecryptAESHex(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	pt := DecryptAES(key, ciphertext)
	s := string(pt[:])
	return s
	// fmt.Println("DECRYPTED:", s)
}

func DecryptAES(key []byte, ciphertext []byte) []byte {

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	return pt
}


