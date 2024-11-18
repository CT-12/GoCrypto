# Crypto

## Usage

```shell
Usage:
  Go-Crypto [flags]
  Go-Crypto [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     Decrypt a text
  encrypt     Encrypt a text
  help        Help about any command

Flags:
  -h, --help      help for Go-Crypto
  -v, --version   Show the version of the application

Use "Go-Crypto [command] --help" for more information about a command.
```

## Encrypt

```shell
Encrypt a text

Usage:
  Go-Crypto encrypt [flags]

Flags:
  -h, --help            help for encrypt
  -k, --key string      The key to encrypt the text
  -o, --output string   The output path to save the encrypted text (default "encrypted_text.txt")
  -r, --read string     The file path to read the plain file (default "plain_file.txt")
```

## Decrypt

```shell
Decrypt a text

Usage:
  Go-Crypto decrypt [flags]

Flags:
  -h, --help            help for decrypt
  -k, --key string      The key to decrypt the text
  -o, --output string   The output path to save the decrypted text (default "decrypted_text.txt")
  -r, --read string     The file path to read the encrypted text (default "encrypted_text.txt")
```