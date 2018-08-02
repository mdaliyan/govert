package main

import (
	"fmt"
	to "github.com/mdaliyan/govert"
)

func main() {

	var aFloat float64
	var aString string
	var anInt int64
	var anInt8 int8

	var fl = 113233.247308

	fmt.Println(to.This(fl, &aString), aString)
	fmt.Println(to.This(fl, &aString, 2), aString)
	fmt.Println(to.This(fl, &aString, 3), aString)
	fmt.Println(to.This(fl, &anInt), anInt)
	fmt.Println(to.This(fl, &anInt8), anInt8)
	fmt.Println()

	var in = int32(24353)
	aFloat = 0
	anInt8 = 0
	anInt = 0

	fmt.Println(to.This(in, &aFloat), aFloat)
	fmt.Println(to.This(in, &anInt), anInt)
	fmt.Println(to.This(in, &anInt8), anInt8)
	fmt.Println()

	var st string = "113233.247308"
	aFloat = 0
	anInt8 = 0
	anInt = 0

	fmt.Println(to.This(st, &aFloat), aFloat)
	fmt.Println(to.This(st, &anInt), anInt)
	fmt.Println(to.This(st, &anInt8), anInt8)
	fmt.Println()

}
