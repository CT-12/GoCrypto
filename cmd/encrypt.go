package cmd

import (
	"github.com/CT-12/Go-Crypto/utils/crypto"
	"github.com/CT-12/Go-Crypto/utils/file_manager"
	"github.com/spf13/cobra"
)

var(
	plain_file_read_path string
	encrypted_file_output_path string
	encrypt_key string

)

var encryptCmd = &cobra.Command{
	Use:  "encrypt",
	Short: "Encrypt a text",
	Run: func(cmd *cobra.Command, args []string) {
		plain_text, err := file_manager.ReadFile(plain_file_read_path)
		if err != nil{
			panic(err)
		}

		result, err := crypto.EncryptAES(plain_text, encrypt_key)
		if err != nil{
			panic(err)
		}

		file_manager.WriteFile(encrypted_file_output_path, result)
	},
}

func init(){
	encryptCmd.Flags().StringVarP(&plain_file_read_path, "read", "r", "plain_file.txt", "The file path to read the plain file")
	encryptCmd.Flags().StringVarP(&encrypted_file_output_path, "output", "o", "encrypted_text.txt", "The output path to save the encrypted text")
	encryptCmd.Flags().StringVarP(&encrypt_key, "key", "k", "", "The key to encrypt the text")
	encryptCmd.MarkFlagRequired("key")
}