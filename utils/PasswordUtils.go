package utils

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"strings"
)

func PasswordEncrypt(username string, password string, salt string) string {
	ok, _ := Encrypt(password, 0, username, []byte(salt))
	return ok
}

func getDerivedKey(password string, salt []byte, count int) ([]byte, []byte) {
	key := md5.Sum([]byte(password + string(salt)))
	for i := 0; i < count-1; i++ {
		key = md5.Sum(key[:])
	}
	return key[:8], key[8:]
}

func Encrypt(password string, obtenationIterations int, plainText string, salt []byte) (string, error) {
	padNum := byte(8 - len(plainText)%8)
	for i := byte(0); i < padNum; i++ {
		plainText += string(padNum)
	}
	dk, iv := getDerivedKey(password, salt, obtenationIterations)
	block, err := des.NewCipher(dk)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(plainText))
	encrypter.CryptBlocks(encrypted, []byte(plainText))
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func Decrypt(password string, obtenationIterations int, cipherText string, salt []byte) (string, error) {
	msgBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	dk, iv := getDerivedKey(password, salt, obtenationIterations)
	block, err := des.NewCipher(dk)
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(msgBytes))
	decrypter.CryptBlocks(decrypted, msgBytes)
	decryptedString := strings.TrimRight(string(decrypted), "\x01\x02\x03\x04\x05\x06\x07\x08")
	return decryptedString, nil
}
