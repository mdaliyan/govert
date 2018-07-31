package main

import to "github.com/mdaliyan/govert"
import "fmt"

func main() {

	var i interface{} = 3783.2882332

	// interface to string
	aString := to.String(i, 3) // "3783.288"

	// float64 to string
	anotherString := to.String(3783.2882332, 3) // "3783.288"

	// bool to string
	aTrueString := to.String(true)   // "true"
	aFalseString := to.String(false) // "false"

	// bool to int8
	aTrueInt8 := to.Int8(true)   // 1
	aFalseInt8 := to.Int8(false) // 0

	// string to float64
	aFloat32 := to.Float32("3783.2882332") // 3783.2883

	// string to int8
	anInt8 := to.Int8("56") // 56

	fmt.Println(aString, anotherString, aTrueString, aFalseString, aTrueInt8, aFalseInt8, aFloat32, anInt8)
	// 3783.288 3783.288 true false 1 0 3783.2883 56
}
