# CipherPayKit

## Overview
CipherPayKit is a versatile library providing cryptographic operations and payment validation functions.

## Status
CipherPayKit aims to provide convenient payment cryptograph related functionalities. It is still under development.

## Usage
### LIB
```bash
go get github.com/veszok/cipherpaykit
```
### CLI
```bash
go install github.com/veszok/cipherpaykit
```
Use `cipherpaykit help` to get the usage details as below:
```
USAGE:
   cipherpaykit [command [subcommand]] [flags...]

COMMANDS:
   help    Shows a list of commands or help for one command
   tdes    Encrypts/Decrypts data using triple DES algorithm
   kcv     Calculates Key Check Value for the given key

SUBCOMMANDS:
   tdes:
      enc  Encrytps data using triple DES
	  dec  Decrypts data using triple DES
FLAGS:
   --key   Key data
   --mode  Block mode: ecb, cbc. Default: ecb
   --data  Data to be encrypted/decrypted
   --iv    Initialization vector (mandatory for CBC block mode)
```
Example:
```
cipherpaykit tdes dec --mode ecb --key 303132333435303132333435303132333435303132333435 --data ba1e59ff04413cd902b85219a34d2cb4 61616161616161616262626262626262
```
Please note current version only supports hexdecimal string format for input key/data. More format (Ascii, Base64, etc.) will be added in later versions.

## License
CipherPayKit is released under the Apache 2.0 license. See [LICENSE.txt](LICENSE.txt)


