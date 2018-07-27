# govert
    
[![Build Status](https://travis-ci.org/mdaliyan/govert.svg?branch=master)](https://travis-ci.org/mdaliyan/govert)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE)
    
`govert` is simple package that provides you some helpers to convert 
golang basic data types to one another easily.
 
This package uses reflection to detect data types so it may not be what 
you want if you have so many ops and you need code performance to be high. 
Otherwise it’s ok for simple use cases, build prototypes or learning purpose.
    
##install:
````bash
$ go get github.com/mdaliyan/govert
````
    
## usage:
 
first, import the package as following,then you have 2 ways to use `govert`. 
 
````go
import "github.com/mdaliyan/govert"
````
    
##### 1. Converting by passing destination pointer
    
````go
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
```` 
    
    
##### 2. Converting inline
````go
package main
    
import "github.com/mdaliyan/govert"
import "fmt"
    
func main() {
    
    // float64 to string
    aString := govert.ToString(3783.2882332, 3) // "3783.288"
    
    // bool to string
    aTrueString := govert.ToString(true)   // "true"
    aFalseString := govert.ToString(false) // "false"
    
    // bool to int8
    aTrueInt8 := govert.ToInt8(true)   // 1
    aFalseInt8 := govert.ToInt8(false) // 0
    
    // string to float64
    aFloat32 := govert.ToFloat32("3783.2882332") // 3783.2883
    
    // string to int8
    anInt8 := govert.ToInt8("56") // 56
    
    fmt.Println(aString, aTrueString, aFalseString, aTrueInt8, aFalseInt8, aFloat32, anInt8)
    // 3783.288 true false 1 0 3783.2883 56
}
```` 
    
## Todo:
    
 - Support converting from complex type to other types. (At this point I'm thinking of just removing imaginary part.)
 - Adding benchmarks
    
## Contributing
    
Feel free to send PR.
