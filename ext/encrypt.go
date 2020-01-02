package ext

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

// Md5 md5 hash
func Md5(text string) string {
	sign := md5.New()
	sign.Write([]byte(text))
	return fmt.Sprintf("%x", sign.Sum(nil))
}

// Sh256 sha256 hash
func Sh256(text string) string {
	sha := sha256.New()
	sha.Write([]byte(text))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

// Sh256WithMd5  text -> sha256 hash -> md5 hash
func Sh256WithMd5(text string) string {
	return Md5(Sh256(text))
}

//AesEncryptCFB aes cfb decrypt data
func AesEncryptCFB(origData []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}

//AesDecryptCFB aes cfb encrypt data
func AesDecryptCFB(encrypted []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
