package main

import (
	"fmt"
	to "github.com/mdaliyan/govert"
)

func main() {

	var fl = 113233.247308

	fmt.Println(to.String(fl))
	fmt.Println(to.String(fl, 2))
	fmt.Println(to.String(fl, 3))
	fmt.Println(to.Int(fl))
	fmt.Println(to.Int8(fl))
	fmt.Println()

	var in = int32(24353)

	fmt.Println(to.Float64(in))
	fmt.Println(to.Int(in))
	fmt.Println(to.Int8(in))
	fmt.Println()

	var st string = "113233.247308"

	fmt.Println(to.Float64(st))
	fmt.Println(to.Int(st))
	fmt.Println(to.Int8(st))
	fmt.Println()

}
