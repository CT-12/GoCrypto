package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

/*
去除 PKCS7Padding 的填充
text 為待去除填充的數據
*/
func unPKCS7Padding(text []byte)([]byte){
	// 取出填充的數據，以此得知填充了多少數據進去
	unPadding := int(text[len(text)-1])

	return text[:len(text)-unPadding]
}

func DecryptAES(encrypted_text string, key string)(string, error){
	encrypted_bytes, err := hex.DecodeString(encrypted_text)
	if err != nil{
		return "", err
	}

	hashKey := hash([]byte(key))
	block, err := aes.NewCipher(hashKey)
	if err != nil{
		return "", err
	}

	iv := hashKey[:block.BlockSize()]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(encrypted_bytes))

	blockMode.CryptBlocks(result, encrypted_bytes)
	result = unPKCS7Padding(result)
	decrypted_text := string(result)

	return decrypted_text, nil
}