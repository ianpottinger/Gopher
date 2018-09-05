package main

import (
	"fmt"
	"strings"
)

const welcome = "Welcome to Go!"
const shout = "Shout it from the roof tops"

func main() {
	fmt.Println(welcome)
	fmt.Println(strings.ToUpper(shout))

}
