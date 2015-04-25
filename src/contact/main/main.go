package main

import (
	"contact"
	"fmt"
)

func main() {
	msg := contact.Say()
	fmt.Println(msg)
}
