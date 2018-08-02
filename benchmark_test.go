package govert

import (
	"reflect"
	"strconv"
	"testing"
)

var (
	i  int         = 101
	ii interface{} = int(101)

	s  string      = "25.41"
	si interface{} = "25.41"

	tmp interface{}
)

func e(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
}

/* =========================================================================
 *  converting from int
 * =========================================================================
 */
func Benchmark_InterfaceInt_ToString_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = String(ii)
	}
}

func Benchmark_InterfaceInt_ToFloat64_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = Float64(ii)
	}
}

func Benchmark_Int_ToString_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = String(i)
	}
}

func Benchmark_Int_ToFloat64_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = String(i)
	}
}

func Benchmark_InterfaceInt_ToString_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		if reflect.TypeOf(ii).Kind() == reflect.Int {
			tmp = strconv.FormatInt(int64(ii.(int)), 10)
		}
	}
}

func Benchmark_InterfaceInt_ToFloat64_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		if reflect.TypeOf(ii).Kind() == reflect.Int {
			tmp = float64(ii.(int))
		}
	}
}

func Benchmark_Int_ToString_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = strconv.FormatInt(int64(i), 10)
	}
}

func Benchmark_Int_ToFloat64_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = float64(i)
	}
}

/* =========================================================================
 *  converting from string
 * =========================================================================
 */
func Benchmark_InterfaceString_ToInt_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = Int(si)
	}
}

func Benchmark_InterfaceString_ToFloat64_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = Float64(si)
	}
}

func Benchmark_String_ToInt_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = Int(s)
	}
}

func Benchmark_String_ToFloat64_govert(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp = Int(s)
	}
}

func Benchmark_InterfaceString_ToInt_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		if reflect.TypeOf(si).Kind() == reflect.String {
			t, _ := strconv.ParseInt(si.(string), 10, 64)
			tmp = int(t)
		}
	}
}

func Benchmark_InterfaceString_ToFloat64_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		if reflect.TypeOf(si).Kind() == reflect.String {
			tmp, _ = strconv.ParseFloat(si.(string), 64)
		}
	}
}

func Benchmark_String_ToInt_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		t1, _ := strconv.ParseInt(s, 10, 64)
		tmp = int(t1)
	}
}

func Benchmark_String_ToFloat64_manual(b *testing.B) {
	e(b)
	for n := 0; n < b.N; n++ {
		tmp, _ = strconv.ParseFloat(s, 64)
	}
}
