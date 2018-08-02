# govert
    
[![Build Status](https://travis-ci.org/mdaliyan/govert.svg?branch=master)](https://travis-ci.org/mdaliyan/govert)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdaliyan/govert?style=flat-square)](https://goreportcard.com/report/github.com/mdaliyan/govert)
[![GoDoc](https://godoc.org/github.com/mdaliyan/govert?status.svg)](https://godoc.org/github.com/mdaliyan/govert)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo/master/LICENSE)
    
Package govert provides you some helpers to convert golang basic data types specially _interfaces_ to any another basic type.
   
## Install:
````bash
$ go get github.com/mdaliyan/govert
````
    
## Usage:
    
import govert as following, then you have 2 ways to use it.
    
````go
import to "github.com/mdaliyan/govert"
````
    
##### 1. Converting inline
````go
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
````
    
##### 2. Converting by passing destination pointer
    
````go
package main
    
import to "github.com/mdaliyan/govert"
import "fmt"
    
func main() {
    
    var err error
    
    // string to float64
    var aFloat float64
    err = to.This("3783.2882332", &aFloat) // 3783.2882332
    
    // string to int64
    var anInt64 int64
    err = to.This("3783.2882332", &anInt64) // 3783
    
    // float64 to int32
    var anInt32 int32
    err = to.This(3783.2882332, &anInt32) // 3783
    
    // float64 to string
    var String1, String2, String3 string
    err = to.This(3783.2882332, &String1)    // "3783.2882", 4 percs by default
    err = to.This(3783.2882332, &String2, 3) // "3783.288"
    err = to.This(3783.2882332, &String3, 2) // "3783.29"
    
    fmt.Println(err, aFloat, anInt64, anInt32, String1, String2, String3)
    // <nil> 3783.2882332 3783 3783 3783.2882 3783.288 3783.29
}
```` 
    
## Todo:
- [ ] Support converting from complex type to other types. (At this point I'm thinking of just removing imaginary part.)
- [x] Adding benchmarks
 
## Benchmarks:
    
govert uses reflection to detect data types so it may not be what 
you want if you have so many ops and you need code performance to be high. 
Otherwise it’s ok for simple use cases, build prototypes or learning purpose.
   
Converting numbers like int, in16, int32, in64, float32, float64, uint, uint, uint16, uint32, 
uint64 to each other with govert is totally non-optimal. Just convert them like this.
    
````go
 var a float32 = 331.23
 var b int64 = int64(a)
````
    
If you have an interface and you have to check type before converting it, or you are
trying to convert something to string or vice versa, using govert makes sense.
    
<table>
    <thead>
      <tr>
        <th colspan="2">Converting</th>
        <th colspan="3">govert</th>
        <th colspan="3">regular way</th>
        <th>compare</th>
      </tr>
      <tr>
        <th>From</th>
        <th>To</th>
        <th>ns/op</th>
        <th>B/op</th>
        <th>allocs/op</th>
        <th>ns/op</th>
        <th>B/op</th>
        <th>allocs/op</th>
        <th>times slower</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>int (interface)</td>
        <td>string</td>
        <td>200</td>
        <td>51</td>
        <td>4</td>
        <td>63.1</td>
        <td>19</td>
        <td>2</td>
        <td>x3.2</td>
      </tr>
      <tr>
        <td>int (interface)</td>
        <td>float64</td>
        <td>119</td>
        <td>24</td>
        <td>3</td>
        <td>22.0</td>
        <td>8</td>
        <td>1</td>
        <td>x5.4</td>
      </tr>
      <tr>
        <td>int</td>
        <td>string</td>
        <td>216 </td>
        <td>64 </td>
        <td>5</td>
        <td>59.6 </td>
        <td>19</td>
        <td>2</td>
        <td>x11.6</td>
      </tr>
      <tr>
        <td>int</td>
        <td>float64</td>
        <td>219 </td>
        <td>64</td>
        <td>5</td>
        <td>18.6</td>
        <td>8</td>
        <td>1</td>
        <td>x12</td>
      </tr>
      <tr>
        <td>string (interface)</td>
        <td>int</td>
        <td>149 </td>
        <td>24</td>
        <td>3</td>
        <td>75</td>
        <td>48</td>
        <td>1</td>
        <td>x2</td>
      </tr>
      <tr>
        <td>string (interface)</td>
        <td>float64</td>
        <td>150 </td>
        <td>24</td>
        <td>3</td>
        <td>45.3 </td>
        <td>8</td>
        <td>1</td>
        <td>x3.3</td>
      </tr>
      <tr>
        <td>string</td>
        <td>int</td>
        <td>185 </td>
        <td>40</td>
        <td>4</td>
        <td>71.5 </td>
        <td>48</td>
        <td>1</td>
        <td>x2.6</td>
      </tr>
      <tr>
        <td>string</td>
        <td>float64</td>
        <td>186 </td>
        <td>40</td>
        <td>4</td>
        <td>40.6 </td>
        <td>8</td>
        <td>1</td>
        <td>x4.5</td>
      </tr>
    </tbody> 
</table>
    
## Contributing:
    
Feel free to send PR.
