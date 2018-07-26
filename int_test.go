package goconvert

import (
	"testing"
	"reflect"
)

func TestConvertInt(t *testing.T) {

	var Val int = 371
	typ := reflect.TypeOf(Val).Kind().String()

	var toBool bool
	if To(Val, &toBool) != nil || !toBool {
		t.Errorf("expected converting %s to bool to be true, got \"%v\"", typ, toBool)
	}

	var toString string
	if To(Val, &toString) != nil || toString != "371" {
		t.Errorf("expected converting to String to be \"3371\", got \"%s\"", toString)
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
		t.Errorf("expected converting %s to complex128 to be (371+0i), got %v", typ, toComplex128)
	}

	var toComplex64 complex64
	if To(Val, &toComplex64) != nil || toComplex64 != complex64(complex(float64(Val), 0)) {
		t.Errorf("expected converting %s to complex64 to be (371+0i), got %v", typ, toComplex64)
	}

}

//func BenchmarkAddItemView(b *testing.B) {
//	c := NewCache(Config{PageLifetime: 2, AutoExpirePage: true}) // 1 second
//	c.Page("PageID").Item("ItemID").AddViewWithPrice(1, 350)
//	wg := 	sync.WaitGroup{}
//	ch := rune(0)
//	b.ReportAllocs()
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		wg.Add(1)
//		go func(newRune rune) {
//			item := c.Page(string(newRune%20 + 65)).Item(string(newRune%32 + 65))
//			item.AddView(1)
//			ch++
//			wg.Done()
//		}(ch)
//		ch++
//	}
//	wg.Wait()
//}
