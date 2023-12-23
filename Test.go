package main

import (
	"bufio"
	"fmt"
	"os"
)

const LoginToken string = "this is login"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the LoginToken")
	input, _ := reader.ReadString('\n')
	fmt.Println("Please Enter the Token ", input)
	fmt.Printf("Your Token id is : %T", input)
}
