package goconvert

import (
	"testing"
	"reflect"
)

func TestConvertFloat32(t *testing.T) {

	var Val float32 = 371.528194
	typ := reflect.TypeOf(Val).Kind().String()

	var toString string
	if To(Val, &toString, 3) != nil || toString != "371.528" {
		t.Errorf("expected converting %s to string to be \"371.528\", got \"%s\"", typ, toString)
	}
	if To(Val, &toString, 2) != nil || toString != "371.53" {
		t.Errorf("expected converting %s to string to be \"371.53\", got \"%s\"", typ, toString)
	}

	var toBool bool
	if To(Val, &toBool) != nil || !toBool {
		t.Errorf("expected converting %s to bool to be true, got \"%v\"", typ, toBool)
	}

	var toInt16 int16
	if To(Val, &toInt16) != nil || toInt16 != int16(Val) {
		t.Errorf("expected converting %s to int16 to be 3371, got %d", typ, toInt16)
	}

	var toInt64 int64
	if To(Val, &toInt64) != nil || toInt64 != int64(Val) {
		t.Errorf("expected converting %s to int64 to be 3371, got %d", typ, toInt64)
	}

	var toComplex128 complex128
	if To(Val, &toComplex128) != nil || toComplex128 != complex(float64(Val), 0) {
		t.Errorf("expected converting %s to complex128 to be (371.5281982421875+0i), got %v", typ, toComplex128)
	}

	var toComplex64 complex64
	if To(Val, &toComplex64) != nil || toComplex64 != complex64(complex(float64(Val), 0)) {
		t.Errorf("expected converting %s to complex64 to be (371.5282+0i), got %v", typ, toComplex64)
	}
}

func TestConvertFloat64(t *testing.T) {

	var Val float64 = 371.528194
	typ := reflect.TypeOf(Val).Kind().String()

	var toString string
	if To(Val, &toString, 3) != nil || toString != "371.528" {
		t.Errorf("expected converting %s to String to be \"371.528\", got \"%s\"", typ, toString)
	}
	if To(Val, &toString, 2) != nil || toString != "371.53" {
		t.Errorf("expected converting %s to String to be \"371.53\", got \"%s\"", typ, toString)
	}

	var toBool bool
	if To(Val, &toBool) != nil || !toBool {
		t.Errorf("expected converting %s to String to be true, got \"%v\"", typ, toBool)
	}

	var toInt16 int16
	if To(Val, &toInt16) != nil || toInt16 != int16(Val) {
		t.Errorf("expected converting %s to int16 to be 3371, got %d", typ, toInt16)
	}

	var toInt64 int64
	if To(Val, &toInt64) != nil || toInt64 != int64(Val) {
		t.Errorf("expected converting %s to int64 to be 3371, got %d", typ, toInt64)
	}

	var toComplex128 complex128
	if To(Val, &toComplex128) != nil || toComplex128 != complex(Val, 0) {
		t.Errorf("expected converting %s to complex128 to be (371.528194+0i), got %v", typ, toComplex128)
	}

	var toComplex64 complex64
	if To(Val, &toComplex64) != nil || toComplex64 != complex64(complex(Val, 0)) {
		t.Errorf("expected converting %s to complex64 to be (371.5282+0i), got %v", typ, toComplex64)
	}

}
