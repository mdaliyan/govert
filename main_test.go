package govert

import (
	"reflect"
	"testing"
	"encoding/json"
)

func assert(t *testing.T, converting, to string, expected, got interface{}) {
	if expected != got {
		t.Errorf("converting %s to %s, expected %v, got %v", converting, to, expected, got)
	}
}

func TestConvertBool(t *testing.T) {

	var Val bool = true
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "true", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(1), Int(Val))
	assert(t, typ, "Int8", int8(1), Int8(Val))
	assert(t, typ, "Int16", int16(1), Int16(Val))
	assert(t, typ, "Int32", int32(1), Int32(Val))
	assert(t, typ, "Int64", int64(1), Int64(Val))

	assert(t, typ, "Uint", uint(1), Uint(Val))
	assert(t, typ, "Uint8", uint8(1), Uint8(Val))
	assert(t, typ, "Uint16", uint16(1), Uint16(Val))
	assert(t, typ, "Uint32", uint32(1), Uint32(Val))
	assert(t, typ, "Uint64", uint64(1), Uint64(Val))

	assert(t, typ, "Float32", float32(1), Float32(Val))
	assert(t, typ, "Float64", float64(1), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(1), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(1), 0), Complex128(Val))

	Val = false

	assert(t, typ, "String", "false", String(Val))

	assert(t, typ, "Bool", false, Bool(Val))

	assert(t, typ, "Int", int(0), Int(Val))
	assert(t, typ, "Int8", int8(0), Int8(Val))
	assert(t, typ, "Int16", int16(0), Int16(Val))
	assert(t, typ, "Int32", int32(0), Int32(Val))
	assert(t, typ, "Int64", int64(0), Int64(Val))

	assert(t, typ, "Uint", uint(0), Uint(Val))
	assert(t, typ, "Uint8", uint8(0), Uint8(Val))
	assert(t, typ, "Uint16", uint16(0), Uint16(Val))
	assert(t, typ, "Uint32", uint32(0), Uint32(Val))
	assert(t, typ, "Uint64", uint64(0), Uint64(Val))

	assert(t, typ, "Float32", float32(0), Float32(Val))
	assert(t, typ, "Float64", float64(0), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(0), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(0), 0), Complex128(Val))

}

func TestConvertNil(t *testing.T) {

	typ := "nil"

	assert(t, typ, "String", "", String(nil))

	assert(t, typ, "Bool", false, Bool(nil))

	assert(t, typ, "Int", int(0), Int(nil))
	assert(t, typ, "Int8", int8(0), Int8(nil))
	assert(t, typ, "Int16", int16(0), Int16(nil))
	assert(t, typ, "Int32", int32(0), Int32(nil))
	assert(t, typ, "Int64", int64(0), Int64(nil))

	assert(t, typ, "Uint", uint(0), Uint(nil))
	assert(t, typ, "Uint8", uint8(0), Uint8(nil))
	assert(t, typ, "Uint16", uint16(0), Uint16(nil))
	assert(t, typ, "Uint32", uint32(0), Uint32(nil))
	assert(t, typ, "Uint64", uint64(0), Uint64(nil))

	assert(t, typ, "Float32", float32(0), Float32(nil))
	assert(t, typ, "Float64", float64(0), Float64(nil))

	assert(t, typ, "Complex64", complex64(complex(0, 0)), Complex64(nil))
	assert(t, typ, "Complex128", complex(0, 0), Complex128(nil))
}

func TestConvertNilPointer(t *testing.T) {

	var Val *int
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "", String(Val))

	assert(t, typ, "Bool", false, Bool(Val))

	assert(t, typ, "Int", int(0), Int(Val))
	assert(t, typ, "Int8", int8(0), Int8(Val))
	assert(t, typ, "Int16", int16(0), Int16(Val))
	assert(t, typ, "Int32", int32(0), Int32(Val))
	assert(t, typ, "Int64", int64(0), Int64(Val))

	assert(t, typ, "Uint", uint(0), Uint(Val))
	assert(t, typ, "Uint8", uint8(0), Uint8(Val))
	assert(t, typ, "Uint16", uint16(0), Uint16(Val))
	assert(t, typ, "Uint32", uint32(0), Uint32(Val))
	assert(t, typ, "Uint64", uint64(0), Uint64(Val))

	assert(t, typ, "Float32", float32(0), Float32(Val))
	assert(t, typ, "Float64", float64(0), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(0, 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(0, 0), Complex128(Val))
}

func TestConvertJsonNumber(t *testing.T) {

	var Val json.Number = json.Number("371")
	typ := reflect.TypeOf(Val).Kind().String()

	in, _ := Val.Int64()

	assert(t, typ, "String", "371", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(in), Int(Val))
	assert(t, typ, "Int8", int8(in), Int8(Val))
	assert(t, typ, "Int16", int16(in), Int16(Val))
	assert(t, typ, "Int32", int32(in), Int32(Val))
	assert(t, typ, "Int64", int64(in), Int64(Val))

	assert(t, typ, "Uint", uint(in), Uint(Val))
	assert(t, typ, "Uint8", uint8(in), Uint8(Val))
	assert(t, typ, "Uint16", uint16(in), Uint16(Val))
	assert(t, typ, "Uint32", uint32(in), Uint32(Val))
	assert(t, typ, "Uint64", uint64(in), Uint64(Val))

	assert(t, typ, "Float32", float32(in), Float32(Val))
	assert(t, typ, "Float64", float64(in), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(in), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(in), 0), Complex128(Val))
}

func TestConvertInt(t *testing.T) {

	var Val int = 371
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "371", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertInt8(t *testing.T) {

	var Val int8 = 32
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "32", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertInt16(t *testing.T) {

	var Val int16 = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "3211", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertInt32(t *testing.T) {

	var Val int32 = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "3211", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertInt64(t *testing.T) {

	var Val int64 = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "3211", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertUint(t *testing.T) {

	var Val uint = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "3211", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertUint8(t *testing.T) {

	var Val uint8 = 32
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "32", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertUint16(t *testing.T) {

	var Val uint16 = 321
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "321", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertUint32(t *testing.T) {

	var Val uint32 = 321
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "321", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertUint64(t *testing.T) {

	var Val uint64 = 321
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "321", String(Val))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertFloat32(t *testing.T) {

	var Val float32 = 371.528194
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "371.528", String(Val, 3))
	assert(t, typ, "String", "371.53", String(Val, 2))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}

func TestConvertFloat64(t *testing.T) {

	var Val float64 = 371.528194
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "String", "371.528", String(Val, 3))
	assert(t, typ, "String", "371.53", String(Val, 2))

	assert(t, typ, "Bool", true, Bool(Val))

	assert(t, typ, "Int", int(Val), Int(Val))
	assert(t, typ, "Int8", int8(Val), Int8(Val))
	assert(t, typ, "Int16", int16(Val), Int16(Val))
	assert(t, typ, "Int32", int32(Val), Int32(Val))
	assert(t, typ, "Int64", int64(Val), Int64(Val))

	assert(t, typ, "Uint", uint(Val), Uint(Val))
	assert(t, typ, "Uint8", uint8(Val), Uint8(Val))
	assert(t, typ, "Uint16", uint16(Val), Uint16(Val))
	assert(t, typ, "Uint32", uint32(Val), Uint32(Val))
	assert(t, typ, "Uint64", uint64(Val), Uint64(Val))

	assert(t, typ, "Float32", float32(Val), Float32(Val))
	assert(t, typ, "Float64", float64(Val), Float64(Val))

	assert(t, typ, "Complex64", complex64(complex(float64(Val), 0)), Complex64(Val))
	assert(t, typ, "Complex128", complex(float64(Val), 0), Complex128(Val))
}
