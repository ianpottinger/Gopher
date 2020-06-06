package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const purpose string = "Strong statically typed, easy concurrency friendly, garbaage collected language, with fast compilation."
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

type Element struct {
	row    int
	column int
	value  int
}

var fixedArray [5]Element

var element int = 0
var array1D [3]int = [3]int{1, 5, 9}

// Comments
type CaseType struct {
	//struct, not class for case types
	text   string
	number int
}

// Key Value Pair
type KeyValue struct {
	key   string
	value string
}

func (keyValuePair KeyValue) getKey() string {
	return keyValuePair.key
}

func (keyValuePair KeyValue) getValue() string {
	return keyValuePair.value
}

func (keyValuePair *KeyValue) setKey(key string) {
	keyValuePair.key = key
}

func (keyValuePair *KeyValue) setValue(value string) {
	keyValuePair.value = value
}

var (
	varblock string = "var block"
)

var numbers = [...]int{1, 2, 3}
var pointer = &numbers
var slice = []string{"Ready", "Steady", "Go!"}

func isOdd(numeral int) bool {
	return numeral%2 == 1
}

func isEven(numeral int) bool {
	return numeral%2 == 0
}

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

	var evens map[string]int = map[string]int{
		"zero":  0,
		"two":   2,
		"four":  4,
		"six":   6,
		"eight": 8,
	}
	fmt.Println(evens)
	fmt.Println(evens["four"])
	delete(evens, "four")
	even, present := evens["four"]
	fmt.Println(evens, even, present)
	if !present && isEven(4) {
		evens["four"] = 4
	}
	fmt.Println(evens)

	var odds map[int]string = map[int]string{
		1: "one",
		3: "three",
		5: "five",
		7: "seven",
		9: "nine",
	}
	fmt.Println(odds)
	fmt.Println(odds[5])
	delete(odds, 5)
	odd, present := odds[5]
	fmt.Println(odds, odd, present)
	if !present && isOdd(5) {
		odds[5] = "five"
	}
	fmt.Println(odds)

	fmt.Println(&pointer)

	kv := KeyValue{"The Key", "The Value"}
	fmt.Printf("\nPointer %v to %v and %v", &kv, kv.key, kv.value)
	fmt.Printf("\nPointer %t to %t and %t", &kv, kv.key, kv.value)
	kv.setKey("New Key")
	kv.setValue("New Value")
	fmt.Printf("\nPointer %v to %v and %v", &kv, kv.getKey(), kv.getValue())
	fmt.Printf("\nPointer %t to %t and %t", &kv, kv.getKey(), kv.getValue())

}
