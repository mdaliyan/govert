package main

import "github.com/mdaliyan/govert"
import "fmt"

func main() {

	// float64 to string
	aString := govert.ToString(3783.2882332, 3) // "3783.288"

	// bool to string
	aTrueString := govert.ToString(true)   // "true"
	aFalseString := govert.ToString(false) // "false"

	// bool to string
	aTrueInt8 := govert.ToInt8(true)   // 1
	aFalseInt8 := govert.ToInt8(false) // 0

	// string to float64
	aFloat32 := govert.ToFloat32("3783.2882332") // 3783.2883

	// string to int8
	anInt8 := govert.ToInt8("56") // 56

	fmt.Println(aString, aTrueString, aFalseString, aTrueInt8, aFalseInt8, aFloat32, anInt8)
	// 3783.288 true false 1 0 3783.2883 56
}
