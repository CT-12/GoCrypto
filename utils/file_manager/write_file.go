package file_manager

import (
	"os"
)

func WriteFile(path string, text string)(error){
	buf := []byte(text)
	err := os.WriteFile(path, buf, 0666)
	return err
}