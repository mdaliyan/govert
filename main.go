package goconvert

import (
	"reflect"
	"strconv"
	"errors"
)

func getElemKind(i interface{}) reflect.Kind {
	return reflect.TypeOf(i).Elem().Kind()
}

func To(in interface{}, out interface{}, params ...interface{}) (err error) {
	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		err = errors.New("second parameter needs to be pointer")
		return
	}
	var outputValue interface{}
	switch reflect.TypeOf(reflect.ValueOf(in).Interface()).Kind() {
	case reflect.String:
	case reflect.Bool:
	case reflect.Int:
		err, outputValue = intConverter(in.(int), out)
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Float32:
		err, outputValue = float64Converter(float64(in.(float32)), out, params...)
	case reflect.Float64:
		err, outputValue = float64Converter(in.(float64), out, params...)
	case reflect.Complex64:
	case reflect.Complex128:
	default:
		err = errors.New("input data type is not recognized")
	}
	if err == nil {
		reflect.ValueOf(out).Elem().Set(reflect.ValueOf(outputValue))
	}
	return
}

func float64Converter(in float64, out interface{}, params ...interface{}) (err error, convertedValue interface{}) {

	switch getElemKind(out) {

	case reflect.String:
		var prec int = 2
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
		tmp := complex(in, 0)
		convertedValue = complex64(tmp)
	case reflect.Complex128:
		convertedValue = complex(in, 0)

	default:
		err = errors.New("output data type is not recognized")
	}

	return
}

func intConverter(in int, out interface{}) (err error, convertedValue interface{}) {

	switch getElemKind(out) {

	case reflect.String:
		convertedValue = strconv.FormatInt(int64(in), 10)

	case reflect.Bool:
		convertedValue = in != 0

	case reflect.Int:
		convertedValue = in
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
		tmp := complex(float64(in), 0)
		convertedValue = complex64(tmp)
	case reflect.Complex128:
		convertedValue = complex(float64(in), 0)

	default:
		err = errors.New("output data type is not recognized")
	}

	return
}
