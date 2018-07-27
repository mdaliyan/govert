package govert

import (
	"testing"
	"reflect"
)

func assert(t *testing.T, converting, to string, expected, got interface{}) {
	if expected != got {
		t.Errorf("converting %s to %s, expected %v, got %v", converting, to, expected, got)
	}
}

func TestConvertBool(t *testing.T) {

	var Val bool = true
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "true", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(1), ToInt(Val))
	assert(t, typ, "ToInt8", int8(1), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(1), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(1), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(1), ToInt64(Val))

	assert(t, typ, "ToUint", uint(1), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(1), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(1), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(1), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(1), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(1), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(1), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(1), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(1), 0), ToComplex128(Val))

	Val = false

	assert(t, typ, "ToString", "false", ToString(Val))

	assert(t, typ, "ToBool", false, ToBool(Val))

	assert(t, typ, "ToInt", int(0), ToInt(Val))
	assert(t, typ, "ToInt8", int8(0), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(0), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(0), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(0), ToInt64(Val))

	assert(t, typ, "ToUint", uint(0), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(0), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(0), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(0), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(0), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(0), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(0), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(0), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(0), 0), ToComplex128(Val))

}

func TestConvertInt(t *testing.T) {

	var Val int = 371
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "371", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertInt8(t *testing.T) {

	var Val int8 = 32
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "32", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertInt16(t *testing.T) {

	var Val int16 = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "3211", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertInt32(t *testing.T) {

	var Val int32 = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "3211", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertInt64(t *testing.T) {

	var Val int64 = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "3211", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertUint(t *testing.T) {

	var Val uint = 3211
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "3211", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertUint8(t *testing.T) {

	var Val uint8 = 32
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "32", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertUint16(t *testing.T) {

	var Val uint16 = 321
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "321", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertUint32(t *testing.T) {

	var Val uint32 = 321
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "321", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertUint64(t *testing.T) {

	var Val uint64 = 321
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "321", ToString(Val))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertFloat32(t *testing.T) {

	var Val float32 = 371.528194
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "371.528", ToString(Val, 3))
	assert(t, typ, "ToString", "371.53", ToString(Val, 2))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}

func TestConvertFloat64(t *testing.T) {

	var Val float64 = 371.528194
	typ := reflect.TypeOf(Val).Kind().String()

	assert(t, typ, "ToString", "371.528", ToString(Val, 3))
	assert(t, typ, "ToString", "371.53", ToString(Val, 2))

	assert(t, typ, "ToBool", true, ToBool(Val))

	assert(t, typ, "ToInt", int(Val), ToInt(Val))
	assert(t, typ, "ToInt8", int8(Val), ToInt8(Val))
	assert(t, typ, "ToInt16", int16(Val), ToInt16(Val))
	assert(t, typ, "ToInt32", int32(Val), ToInt32(Val))
	assert(t, typ, "ToInt64", int64(Val), ToInt64(Val))

	assert(t, typ, "ToUint", uint(Val), ToUint(Val))
	assert(t, typ, "ToUint8", uint8(Val), ToUint8(Val))
	assert(t, typ, "ToUint16", uint16(Val), ToUint16(Val))
	assert(t, typ, "ToUint32", uint32(Val), ToUint32(Val))
	assert(t, typ, "ToUint64", uint64(Val), ToUint64(Val))

	assert(t, typ, "ToFloat32", float32(Val), ToFloat32(Val))
	assert(t, typ, "ToFloat64", float64(Val), ToFloat64(Val))

	assert(t, typ, "ToComplex64", complex64(complex(float64(Val), 0)), ToComplex64(Val))
	assert(t, typ, "ToComplex128", complex(float64(Val), 0), ToComplex128(Val))
}
