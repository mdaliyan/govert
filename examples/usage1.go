package main

import "github.com/mdaliyan/govert"
import "fmt"

func main() {

	var err error

	// string to float64
	var aFloat float64
	err = govert.To("3783.2882332", &aFloat) // 3783.2882332

	// string to int64
	var anInt64 int64
	err = govert.To("3783.2882332", &anInt64) // 3783

	// float64 to int32
	var anInt32 int32
	err = govert.To(3783.2882332, &anInt32) // 3783

	// float64 to string
	var String1, String2, String3 string
	err = govert.To(3783.2882332, &String1)    // "3783.2882", 4 percs by default
	err = govert.To(3783.2882332, &String2, 3) // "3783.288"
	err = govert.To(3783.2882332, &String3, 2) // "3783.29"

	fmt.Println(err, aFloat, anInt64, anInt32, String1, String2, String3)
	// <nil> 3783.2882332 3783 3783 3783.2882 3783.288 3783.29
}
