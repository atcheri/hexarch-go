package main

import (
	"fmt"
)

func init() {
	fmt.Println("===== BEGIN init function =====")
	fmt.Println("We could do some flag parsing here for example: -db=mysql, or -port=8080 etc...")
	fmt.Println("===== END init function =====")
}

func main() {
	fmt.Println("===== BEGIN main function =====")
	fmt.Println("http has run")
	fmt.Println("===== END main function =====")
}
