package main

import (
	"github.com/mdaliyan/govert"
	"fmt"
)

func main() {

	var aFloat float64
	var aString string
	var anInt int64
	var anInt8 int8

	var fl = 113233.247308

	fmt.Println(govert.To(fl, &aString), aString)
	fmt.Println(govert.To(fl, &aString, 2), aString)
	fmt.Println(govert.To(fl, &aString, 3), aString)
	fmt.Println(govert.To(fl, &anInt), anInt)
	fmt.Println(govert.To(fl, &anInt8), anInt8)
	fmt.Println()

	var in = int32(24353)
	aFloat = 0
	anInt8 = 0
	anInt = 0

	fmt.Println(govert.To(in, &aFloat), aFloat)
	fmt.Println(govert.To(in, &anInt), anInt)
	fmt.Println(govert.To(in, &anInt8), anInt8)
	fmt.Println()

	var st string = "113233.247308"
	aFloat = 0
	anInt8 = 0
	anInt = 0

	fmt.Println(govert.To(st, &aFloat), aFloat)
	fmt.Println(govert.To(st, &anInt), anInt)
	fmt.Println(govert.To(st, &anInt8), anInt8)
	fmt.Println()

}
