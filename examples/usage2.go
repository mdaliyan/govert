package main

import "github.com/mdaliyan/govert"
import "fmt"

func main() {

	// float64 to string
	aString := govert.ToString(3783.2882332, 3) // "3783.288"

	// bool to string
	aTrueString := govert.ToFloat32(true)  // "1"
	aFalseString := govert.ToFloat32(false)  // "0"

	// string to float64
	aFloat32 := govert.ToFloat32("3783.2882332")  // "3783.288"

	// string to int8
	anInt8 := govert.ToInt8("56")  // "56"

	fmt.Println(aString, aTrueString, aFalseString, aFloat32, anInt8)
	//3783.288 1 0 3783.2883 56
}
