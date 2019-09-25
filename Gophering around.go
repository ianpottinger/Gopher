package main

import (
	"fmt"
	"strconv"
	"strings"
)

const purpose string = "Strong statically typed, easy concurrent friendly, garbaage collected language, with fast compilation."
const welcome string = "Ready, Steady, Welcome to Pascal 2.0 Go! Stop Gophering around."
const terminal string = "Ctrl - `"
const goRun string = "go run github.com/ianpottinger/Gopher"
const goBuild string = "go build github.com/ianpottinger/Gopher"
const goInstall string = "go install github.com/ianpottinger/Gopher"

// comments
var number int = 42
var decimal float32 = (float32(number / 10))
var fourtyTwo string = (strconv.Itoa(number))
var camelCase = "camel case is used for variables"
var state bool = true
var debug bool = false

type CaseType struct {
	//struct, not class for case types
	text string
}

var (
	varblock string = "var block"
)

var numbers = [...]int{1, 2, 3}
var pointers = &numbers
var slice = []string{"Ready", "Steady", "Go!"}

func main() {
	logical := debug == state
	fmt.Printf("Is %v equal to %v? %v %T\n", state, debug, logical, logical)
	fmt.Printf("%v, %T\n", purpose, purpose)
	fmt.Println(strings.ToUpper(welcome))
	fmt.Println(fourtyTwo)
	becomesEqualTo := "Declaration becomes equal to value"
	fmt.Printf("%v only inside block scopes.\n", becomesEqualTo)

	pascalCase := CaseType{
		text: "Pascal case is used for structs(i.e. classes)",
	}
	acronyms := CaseType{
		text: "acronyms(i.e. A.K.A.) should always be capitalised",
	}
	fmt.Printf("%v, while %v and %v.", pascalCase.text, camelCase, acronyms.text)
}
