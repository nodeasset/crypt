package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nodeasset/crypt"
)

var mode string
var srcfile string
var destfile string
var password string

func _encrypt() {
	fmt.Println("Starting the application...")
	dat, err := ioutil.ReadFile(srcfile)
	if err != nil {
		panic(err.Error())
	}
	crypt.EncryptFile(destfile, []byte(dat), password)
	fmt.Println(string(crypt.DecryptFile(destfile, password)))

}

func _decrypt() {

	dfile := crypt.DecryptFile(srcfile, password)
	err := ioutil.WriteFile(destfile, dfile, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	mode = os.Args[1]
	srcfile = os.Args[2]
	destfile = os.Args[3]
	password = os.Args[4]

	fmt.Println("mode...", mode)
	fmt.Println("srcfile...", srcfile)
	fmt.Println("destfile...", destfile)
	fmt.Println("password...", password)

	if mode == "encrypt" {
		fmt.Println("Starting encryption...")
		_encrypt()
	}

	if mode == "decrypt" {
		fmt.Println("Starting decryption...")
		_decrypt()
	}

	fmt.Println("Starting the application...")
}
