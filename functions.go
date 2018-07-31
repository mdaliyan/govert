package govert

// String converts any type to string
// second parameter is used when you try to convert float to string as number
// of digits after decimal point. note that golang will round the result.
//     fmt.Println(govert.String(224.58719,2)) // 224.59
//     fmt.Println(govert.String(224.58719,3)) // 224.587
func String(in interface{}, params ...interface{}) (out string) {
	This(in, &out, params...)
	return
}

// Bool converts basic types to bool.
//     // will return false if input interface is numeric and is equal to 0.
//     fmt.Println(govert.Bool(224.58719)) // true
//     fmt.Println(govert.Bool(0)) // false
//
//     // will return true if input interface is string and is equal to "true" or "1", compared case insensitive.
//     fmt.Println(govert.Bool("anything")) // false
//     fmt.Println(govert.Bool("false")) // false
//     fmt.Println(govert.Bool("True")) // true
//     fmt.Println(govert.Bool("true")) // true
//     fmt.Println(govert.Bool("1")) // true
func Bool(in interface{}) (out bool) {
	This(in, &out)
	return
}

// Bool converts basic types to int
func Int(in interface{}) (out int) {
	This(in, &out)
	return
}

// Bool converts basic types to in8
func Int8(in interface{}) (out int8) {
	This(in, &out)
	return
}

// Bool converts basic types to int16
func Int16(in interface{}) (out int16) {
	This(in, &out)
	return
}

// Bool converts basic types to int32
func Int32(in interface{}) (out int32) {
	This(in, &out)
	return
}

// Bool converts basic types to int64
func Int64(in interface{}) (out int64) {
	This(in, &out)
	return
}

// Bool converts basic types to uint
func Uint(in interface{}) (out uint) {
	This(in, &out)
	return
}

// Bool converts basic types to uint8
func Uint8(in interface{}) (out uint8) {
	This(in, &out)
	return
}

// Bool converts basic types to uint6
func Uint16(in interface{}) (out uint16) {
	This(in, &out)
	return
}

// Bool converts basic types to uint32
func Uint32(in interface{}) (out uint32) {
	This(in, &out)
	return
}

// Bool converts basic types to uint64
func Uint64(in interface{}) (out uint64) {
	This(in, &out)
	return
}

// Bool converts basic types to float32
func Float32(in interface{}) (out float32) {
	This(in, &out)
	return
}

// Bool converts basic types to float64
func Float64(in interface{}) (out float64) {
	This(in, &out)
	return
}

// Bool converts basic types to complex64
func Complex64(in interface{}) (out complex64) {
	This(in, &out)
	return
}

// Bool converts basic types to complex128
func Complex128(in interface{}) (out complex128) {
	This(in, &out)
	return
}
