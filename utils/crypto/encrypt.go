package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

func hash(key []byte)([]byte){
	hasher := sha256.New()
	// 輸入數據
	hasher.Write(key)
	// 計算 hash 
	bytes := hasher.Sum(nil)
	// 將 byte 編碼為 16 進制的字符串，他會把每個 byte 轉換成兩個16進制字符 ex. 235 = 11101011 = eb
	// hashCode := hex.EncodeToString(bytes)

	return bytes
}

/*
PKCS7Padding: 填充模式
text: 明文內容
blockSize: 分組內容
*/
func pkcs7Padding(text []byte, blockSize int)([]byte){
	// 計算待填充的長度
	padding := blockSize - len(text) % blockSize
	var paddingText []byte
	if padding == 0{ // 已對齊，填充一整塊數據，每個數據為 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	}else{ // 未對齊，填充 padding 個數據，每個數據為 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}

	return append(text, paddingText...)
}

/*
AES CBC 加密
plain_text: 待加密的明文
key: 密鑰
*/
func EncryptAES(plain_text string, key string)(string, error){
	hashKey := hash([]byte(key))
	block, err := aes.NewCipher(hashKey)
	if err != nil{
		return "", err
	}

	// 填充
	paddingText := pkcs7Padding([]byte(plain_text), block.BlockSize())

	iv := hashKey[:block.BlockSize()]
	CBC := cipher.NewCBCEncrypter(block, iv)
	
	// 加密
	result := make([]byte, len(paddingText))
	CBC.CryptBlocks(result, paddingText)

	// 轉成 16 進制字符串
	encrypted_text := hex.EncodeToString(result)

	return encrypted_text, nil
}
