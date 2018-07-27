package govert

func ToString(in interface{}, params ...interface{}) (out string) {
	To(in, &out, params...)
	return
}

func ToBool(in interface{}) (out bool) {
	To(in, &out)
	return
}

func ToInt(in interface{}) (out int) {
	To(in, &out)
	return
}

func ToInt8(in interface{}) (out int8) {
	To(in, &out)
	return
}

func ToInt16(in interface{}) (out int16) {
	To(in, &out)
	return
}

func ToInt32(in interface{}) (out int32) {
	To(in, &out)
	return
}

func ToInt64(in interface{}) (out int64) {
	To(in, &out)
	return
}

func ToUint(in interface{}) (out uint) {
	To(in, &out)
	return
}

func ToUint8(in interface{}) (out uint8) {
	To(in, &out)
	return
}

func ToUint16(in interface{}) (out uint16) {
	To(in, &out)
	return
}

func ToUint32(in interface{}) (out uint32) {
	To(in, &out)
	return
}

func ToUint64(in interface{}) (out uint64) {
	To(in, &out)
	return
}

func ToFloat32(in interface{}) (out float32) {
	To(in, &out)
	return
}

func ToFloat64(in interface{}) (out float64) {
	To(in, &out)
	return
}

func ToComplex64(in interface{}) (out complex64) {
	To(in, &out)
	return
}

func ToComplex128(in interface{}) (out complex128) {
	To(in, &out)
	return
}
