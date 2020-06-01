package main

import (
	"fmt"
	"math"
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
var backtick rune = '`'
var decimal float32 = (float32(number / 10))
var fourtyTwo string = (strconv.Itoa(number))
var camelCase string = "camel case is used for variables"
var state bool = true
var debug bool = false

// comments
var ubyt8x1 uint8 = 255
var ubyt8x2 uint16 = 65535
var ubyt8x4 uint32 = 4294967295
var ubyt8x8 uint64 = 18446744073709551615

// comments
var sbyt8x1 int8 = -127
var sbyt8x2 int16 = -32767
var sbyt8x4 int32 = -2147483647
var sbyt8x8 int64 = -9223372036854775807

// var pointer uintptr = *ubyt84

var ieee75432 float32 = 0.0
var ieee75464 float64 = math.Pi

var dfloat32 complex64 = -32.32
var dfloat64 complex128 = 64.64

var not bool = !false
var or bool = true || false
var and bool = true && true

/*
struct element {
	key   int
	value string
}
var fixedArray [5]element
*/

var element int = 0
var array1D [3]int = [3]int{1, 5, 9}

// comments
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
	fmt.Printf("Value %v, of type %T\n", purpose, purpose)
	fmt.Println(strings.ToUpper(welcome))
	fmt.Println(fourtyTwo)
	becomesEqualTo := "Declaration becomes equal to value\t"
	fmt.Printf("%v only inside block scopes.\n", becomesEqualTo)

	pascalCase := CaseType{
		text: "Pascal case is used for structs(i.e. classes)",
	}
	acronyms := CaseType{
		text: "acronyms(i.e. A.K.A.) should always be capitalised",
	}
	fmt.Printf("%v, while %v and %v.\n", pascalCase.text, camelCase, acronyms.text)
	var concat = fmt.Sprintf("Quoted: %q, while %q and %q.\n", pascalCase.text, camelCase, acronyms.text)
	fmt.Printf(concat)

	var base2 = fmt.Sprintf("Binary %b ", 2)
	var base8 = fmt.Sprintf("Octadecimal %o", 8)
	var base10 = fmt.Sprintf("Decimal %d ", 10)
	var base16 = fmt.Sprintf("Hexadecimal %X ", 16)
	fmt.Printf("%v\t %v\t %v\t %v\t", base2, base8, base10, base16)

	fmt.Printf("\nFloat: %e, %e", ieee75432, ieee75464)
	fmt.Printf("\nSingle: %f, %f", ieee75432, ieee75464)
	fmt.Printf("\nDouble: %g, %g", ieee75432, ieee75464)
	fmt.Printf("\nLeft padded %-09.03f", dfloat32)
	fmt.Printf("\nRight padded %018.06f", dfloat64)

	fmt.Printf("\nNOT ! false: %v", not)
	fmt.Printf("\nOR || true: %v", or)
	fmt.Printf("\nAND && true: %v", and)

	if and != or {
		fmt.Println("\nand is not equal to or")
	} else {
		fmt.Println("\nand is equal to or\n")
	}

	for loop := 0; loop <= 9; loop += 2 {
		fmt.Printf("Progress %v\n", loop)
	}

	start := 0
	limit := 60
	for {
		if !(start%3 == 0 || start%5 == 0) {
			fmt.Printf("%v ", start)
		}
		if start%3 == 0 {
			fmt.Printf("Fizz")
		}
		if start%5 == 0 {
			fmt.Printf("Buzz")
		}
		if !(start >= limit) {
			fmt.Printf("\n")
			start++
			continue
		}
		if start == limit {
			fmt.Printf("\nBuzzly all Fizzed out\n")
			break
		}

	}

	answer := 5
	switch answer {
	case 1:
		fmt.Println("One")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("Big")

	}

	fmt.Println(numbers)
	fmt.Println(array1D)

}
