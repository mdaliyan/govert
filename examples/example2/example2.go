package main

import (
	"github.com/mdaliyan/govert"
	"fmt"
)

func main() {

	var fl = 113233.247308

	fmt.Println(govert.ToString(fl))
	fmt.Println(govert.ToString(fl,2))
	fmt.Println(govert.ToString(fl,3))
	fmt.Println(govert.ToInt(fl))
	fmt.Println(govert.ToInt8(fl))
	fmt.Println()

	var in = int32(24353)

	fmt.Println(govert.ToFloat64(in))
	fmt.Println(govert.ToInt(in))
	fmt.Println(govert.ToInt8(in))
	fmt.Println()

	var st string = "113233.247308"

	fmt.Println(govert.ToFloat64(st))
	fmt.Println(govert.ToInt(st))
	fmt.Println(govert.ToInt8(st))
	fmt.Println()

}
