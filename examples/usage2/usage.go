package main

import to "github.com/mdaliyan/govert"
import "fmt"

func main() {

	// string to float64
	var aFloat float64
	err1 := to.This("3783.2882332", &aFloat) // 3783.2882332

	// string to int64
	var anInt64 int64
	err2 := to.This("3783.2882332", &anInt64) // 3783

	// float64 to int32
	var anInt32 int32
	err3 := to.This(3783.2882332, &anInt32) // 3783

	// float64 to string
	var String1, String2, String3 string
	err4 := to.This(3783.2882332, &String1)    // "3783.2882", 4 percs by default
	err5 := to.This(3783.2882332, &String2, 3) // "3783.288"
	err6 := to.This(3783.2882332, &String3, 2) // "3783.29"

	fmt.Println(err1, err2, err3, err4, err5, err6)
	fmt.Println(aFloat, anInt64, anInt32, String1, String2, String3)
	// Output:
	//  <nil> <nil> <nil> <nil> <nil> <nil>
	//  3783.2882332 3783 3783 3783.2882 3783.288 3783.29
}
