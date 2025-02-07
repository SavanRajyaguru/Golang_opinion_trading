package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func EncryptAES(value, key, iv string) (string, error) {
	if value == "" {
		return "", errors.New("value to encrypt cannot be empty")
	}

	// Convert key and IV to byte slices
	keyBytes := []byte(key)
	ivBytes := []byte(iv)
	plaintext := []byte(value)

	// Create AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// Apply PKCS7 Padding
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedText := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Encrypt data in CBC mode
	ciphertext := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, ivBytes)
	mode.CryptBlocks(ciphertext, paddedText)

	// Return Base64 encoded ciphertext
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
