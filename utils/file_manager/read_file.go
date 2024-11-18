package file_manager

import (
	"io"
	"os"
)

func ReadFile(filePath string)(string, error){
	file, err := os.Open(filePath)
	if err != nil{
		return "", err
	}
	defer file.Close()

	buf := make([]byte, 256)
	var result string

	for {
		len, err := file.Read(buf)
		if err == io.EOF{
			break
		}
		result += string(buf[:len])
	}

	return result, nil
}