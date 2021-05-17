package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// type AESKey struct {
// 	Method string
// 	Key    []byte
// }

func TestEncryption() {
	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "this is a test phrase for aes128"

	c := EncryptAESStringToHex(([]byte(key))[:32], pt)

	// plaintext
	println(pt)

	// ciphertext
	println(c)

	// decrypt
	d := DecryptAESHex([]byte(key)[:32], c)
	println(d)
}

func TestEncryptionKey(key []byte) bool {

	// plaintext
	pt := "this is a test phrase for aes128"

	c := EncryptAESString(key, pt)

	// plaintext
	println(pt)

	// ciphertext
	println(c)

	// decrypt
	d := DecryptAESToString(key, c)
	println(d)

	return (pt == d)

}

func EncryptAESStringToHex(key []byte, plaintext string) string {

	out := EncryptAESString(key, plaintext)
	return hex.EncodeToString(out)
}

func EncryptAESString(key []byte, plaintext string) []byte {

	out := EncryptAES(key, []byte(plaintext))
	return out
}

func DecryptAESHex(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	pt := DecryptAES(key, ciphertext)
	s := string(pt[:])
	return s
	// fmt.Println("DECRYPTED:", s)
}

func DecryptAESToString(key []byte, ciphertext []byte) string {
	return string(DecryptAES(key, ciphertext))
}

func EncryptAESBlock(key []byte, data []byte) []byte {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(data))

	c.Encrypt(out, data)

	return out
}

func DecryptAESBlock(key []byte, ciphertext []byte) []byte {

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	return pt
}

func EncryptAES(key []byte, data []byte) []byte {

	block, err := aes.NewCipher(key)
	CheckError(err)

	iv := make([]byte, block.BlockSize())

	lendata := len(data)

	if lendata%block.BlockSize() != 0 {
		panic("Data is not a multiple of the block size")
	}

	out := make([]byte, lendata)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(out, data)

	// block.Encrypt(out, data)

	return out
}

func DecryptAES(key []byte, ciphertext []byte) []byte {

	block, err := aes.NewCipher(key)
	CheckError(err)

	lenciphertext := len(ciphertext)
	iv := make([]byte, block.BlockSize())

	if lenciphertext%block.BlockSize() != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	out := make([]byte, lenciphertext)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(out, ciphertext)

	// block.Decrypt(pt, ciphertext)
	return out
}
