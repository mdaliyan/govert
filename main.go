package govert

import (
	"reflect"
	"strconv"
	"errors"
	"strings"
)

var invalidOutputType = errors.New("output data type is not a simple type")

func getElemKind(i interface{}) reflect.Kind {
	return reflect.TypeOf(i).Elem().Kind()
}

func To(in interface{}, out interface{}, params ...interface{}) (err error) {

	// input might be a pointer, this ensures we check input type correctly
	for reflect.ValueOf(in).Kind() == reflect.Ptr {
		in = reflect.ValueOf(in).Elem().Interface()
	}

	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		err = errors.New("second parameter needs to be pointer")
		return
	}

	var outputValue interface{}

	switch reflect.TypeOf(reflect.ValueOf(in).Interface()).Kind() {

	case reflect.String:
		err, outputValue = convertStringTo(in.(string), out)

	case reflect.Bool:
		err, outputValue = convertBoolTo(in.(bool), out)

	case reflect.Int:
		err, outputValue = convertInt64To(int64(in.(int)), out)

	case reflect.Int8:

		err, outputValue = convertInt64To(int64(in.(int8)), out)

	case reflect.Int16:
		err, outputValue = convertInt64To(int64(in.(int16)), out)

	case reflect.Int32:
		err, outputValue = convertInt64To(int64(in.(int32)), out)

	case reflect.Int64:
		err, outputValue = convertInt64To(in.(int64), out)

	case reflect.Uint:
		err, outputValue = convertUint64To(uint64(in.(uint)), out)

	case reflect.Uint8:
		err, outputValue = convertUint64To(uint64(in.(uint8)), out)

	case reflect.Uint16:
		err, outputValue = convertUint64To(uint64(in.(uint16)), out)

	case reflect.Uint32:
		err, outputValue = convertUint64To(uint64(in.(uint32)), out)

	case reflect.Uint64:
		err, outputValue = convertUint64To(in.(uint64), out)

	case reflect.Float32:
		err, outputValue = convertFloat64To(float64(in.(float32)), out, params...)

	case reflect.Float64:
		err, outputValue = convertFloat64To(in.(float64), out, params...)

	case reflect.Complex64:

	case reflect.Complex128:

	default:
		err = errors.New("input data type is not a simple type")

	}

	if err == nil {
		reflect.ValueOf(out).Elem().Set(reflect.ValueOf(outputValue))
	}

	return
}

func convertStringTo(in string, out interface{}, params ...interface{}) (err error, convertedValue interface{}) {

	toFloat64 := func(input string) (o float64) {
		o, err = strconv.ParseFloat(input, 64)
		return
	}

	switch getElemKind(out) {

	case reflect.String:
		convertedValue = in

	case reflect.Bool:
		convertedValue = strings.ToLower(in) != "true" || in == "1"

	case reflect.Int:
		convertedValue = int(toFloat64(in))

	case reflect.Int8:
		convertedValue = int8(toFloat64(in))

	case reflect.Int16:
		convertedValue = int16(toFloat64(in))

	case reflect.Int32:
		convertedValue = int32(toFloat64(in))

	case reflect.Int64:
		convertedValue = int64(toFloat64(in))

	case reflect.Uint:
		convertedValue = uint(toFloat64(in))

	case reflect.Uint8:
		convertedValue = uint8(toFloat64(in))

	case reflect.Uint16:
		convertedValue = uint16(toFloat64(in))

	case reflect.Uint32:
		convertedValue = uint32(toFloat64(in))

	case reflect.Uint64:
		convertedValue = uint64(toFloat64(in))

	case reflect.Float32:
		convertedValue = float32(toFloat64(in))

	case reflect.Float64:
		convertedValue = toFloat64(in)

	case reflect.Complex64:
		convertedValue = complex64(complex(toFloat64(in), 0))

	case reflect.Complex128:
		convertedValue = complex(toFloat64(in), 0)

	default:
		err = invalidOutputType

	}

	return
}

func convertBoolTo(in bool, out interface{}, params ...interface{}) (err error, convertedValue interface{}) {

	toInt8 := func(input bool) (int8) {
		if input {
			return 1
		}
		return 0
	}

	switch getElemKind(out) {

	case reflect.String:
		convertedValue = strconv.FormatBool(in)

	case reflect.Bool:
		convertedValue = in

	case reflect.Int:
		convertedValue = int(toInt8(in))

	case reflect.Int8:
		convertedValue = toInt8(in)

	case reflect.Int16:
		convertedValue = int16(toInt8(in))

	case reflect.Int32:
		convertedValue = int32(toInt8(in))

	case reflect.Int64:
		convertedValue = int64(toInt8(in))

	case reflect.Uint:
		convertedValue = uint(toInt8(in))

	case reflect.Uint8:
		convertedValue = uint8(toInt8(in))

	case reflect.Uint16:
		convertedValue = uint16(toInt8(in))

	case reflect.Uint32:
		convertedValue = uint32(toInt8(in))

	case reflect.Uint64:
		convertedValue = uint64(toInt8(in))

	case reflect.Float32:
		convertedValue = float32(toInt8(in))

	case reflect.Float64:
		convertedValue = float64(toInt8(in))

	case reflect.Complex64:
		convertedValue = complex64(complex(float64(toInt8(in)), 0))

	case reflect.Complex128:
		convertedValue = complex(float64(toInt8(in)), 0)

	default:
		err = invalidOutputType

	}

	return
}

func convertFloat64To(in float64, out interface{}, params ...interface{}) (err error, convertedValue interface{}) {

	switch getElemKind(out) {

	case reflect.String:
		var prec int = 4
		if len(params) >= 1 {
			if tmp := params[0]; reflect.TypeOf(tmp).Kind() == reflect.Int {
				prec = tmp.(int)
			}
		}
		convertedValue = strconv.FormatFloat(in, 'f', prec, 64)

	case reflect.Bool:
		convertedValue = in != 0

	case reflect.Int:
		convertedValue = int(in)

	case reflect.Int8:
		convertedValue = int8(in)

	case reflect.Int16:
		convertedValue = int16(in)

	case reflect.Int32:
		convertedValue = int32(in)

	case reflect.Int64:
		convertedValue = int64(in)

	case reflect.Uint:
		convertedValue = uint(in)

	case reflect.Uint8:
		convertedValue = uint8(in)

	case reflect.Uint16:
		convertedValue = uint16(in)

	case reflect.Uint32:
		convertedValue = uint32(in)

	case reflect.Uint64:
		convertedValue = uint64(in)

	case reflect.Float32:
		convertedValue = float32(in)

	case reflect.Float64:
		convertedValue = float64(in)

	case reflect.Complex64:
		convertedValue = complex64(complex(in, 0))

	case reflect.Complex128:
		convertedValue = complex(in, 0)

	default:
		err = invalidOutputType

	}

	return
}

func convertInt64To(in int64, out interface{}) (err error, convertedValue interface{}) {

	switch getElemKind(out) {

	case reflect.String:
		convertedValue = strconv.FormatInt(in, 10)

	case reflect.Bool:
		convertedValue = in != 0

	case reflect.Int:
		convertedValue = int(in)

	case reflect.Int8:
		convertedValue = int8(in)

	case reflect.Int16:
		convertedValue = int16(in)

	case reflect.Int32:
		convertedValue = int32(in)

	case reflect.Int64:
		convertedValue = int64(in)

	case reflect.Uint:
		convertedValue = uint(in)

	case reflect.Uint8:
		convertedValue = uint8(in)

	case reflect.Uint16:
		convertedValue = uint16(in)

	case reflect.Uint32:
		convertedValue = uint32(in)

	case reflect.Uint64:
		convertedValue = uint64(in)

	case reflect.Float32:
		convertedValue = float32(in)

	case reflect.Float64:
		convertedValue = float64(in)

	case reflect.Complex64:
		convertedValue = complex64(complex(float64(in), 0))

	case reflect.Complex128:
		convertedValue = complex(float64(in), 0)

	default:
		err = invalidOutputType

	}

	return
}

func convertUint64To(in uint64, out interface{}) (err error, convertedValue interface{}) {

	switch getElemKind(out) {

	case reflect.String:
		convertedValue = strconv.FormatUint(in, 10)

	case reflect.Bool:
		convertedValue = in != 0

	case reflect.Int:
		convertedValue = int(in)

	case reflect.Int8:
		convertedValue = int8(in)

	case reflect.Int16:
		convertedValue = int16(in)

	case reflect.Int32:
		convertedValue = int32(in)

	case reflect.Int64:
		convertedValue = int64(in)

	case reflect.Uint:
		convertedValue = uint(in)

	case reflect.Uint8:
		convertedValue = uint8(in)

	case reflect.Uint16:
		convertedValue = uint16(in)

	case reflect.Uint32:
		convertedValue = uint32(in)

	case reflect.Uint64:
		convertedValue = uint64(in)

	case reflect.Float32:
		convertedValue = float32(in)

	case reflect.Float64:
		convertedValue = float64(in)

	case reflect.Complex64:
		convertedValue = complex64(complex(float64(in), 0))

	case reflect.Complex128:
		convertedValue = complex(float64(in), 0)

	default:
		err = invalidOutputType

	}

	return
}
