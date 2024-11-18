package cmd

import (
	"fmt"

	"github.com/CT-12/Go-Crypto/utils/crypto"
	"github.com/CT-12/Go-Crypto/utils/file_manager"
	"github.com/spf13/cobra"
)

var(
	encrypted_file_read_path string
	decrypted_file_output_path string
	decrypt_key string
)

var decryptCmd = &cobra.Command{
	Use:  "decrypt",
	Short: "Decrypt a text",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[INFO] Reading file from: ", encrypted_file_read_path)
		encrypted_text, err := file_manager.ReadFile(encrypted_file_read_path)
		if err != nil{
			panic(err)
		}

		fmt.Println("[INFO] Decrypting text...")
		result, err := crypto.DecryptAES(encrypted_text, decrypt_key)
		if err != nil{
			panic(err)
		}

		fmt.Println("[INFO] Writing decrypted text to: ", decrypted_file_output_path)
		file_manager.WriteFile(decrypted_file_output_path, result)
	},
}

func init(){
	decryptCmd.Flags().StringVarP(&encrypted_file_read_path, "read", "r", "encrypted_text.txt", "The file path to read the encrypted text")
	decryptCmd.Flags().StringVarP(&decrypted_file_output_path, "output", "o", "decrypted_text.txt", "The output path to save the decrypted text")
	decryptCmd.Flags().StringVarP(&decrypt_key, "key", "k", "", "The key to decrypt the text")
	decryptCmd.MarkFlagRequired("key")
}