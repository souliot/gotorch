package tensor

import (
	"reflect"
	"unsafe"
)

// Dtype represents a data type of a Tensor. Concretely it's implemented as an embedded reflect.Type
// which allows for easy reflection operations. It also implements hm.Type, for type inference in Gorgonia
type Dtype struct {
	reflect.Type
}

// oh how nice it'd be if I could make them immutable
var (
	Int        = Dtype{reflect.TypeOf(int(1))}
	Int8       = Dtype{reflect.TypeOf(int8(1))}
	Int16      = Dtype{reflect.TypeOf(int16(1))}
	Int32      = Dtype{reflect.TypeOf(int32(1))}
	Int64      = Dtype{reflect.TypeOf(int64(1))}
	Uint       = Dtype{reflect.TypeOf(uint(1))}
	Uint8      = Dtype{reflect.TypeOf(uint8(1))}
	Uint16     = Dtype{reflect.TypeOf(uint16(1))}
	Uint32     = Dtype{reflect.TypeOf(uint32(1))}
	Uint64     = Dtype{reflect.TypeOf(uint64(1))}
	Float32    = Dtype{reflect.TypeOf(float32(1))}
	Float64    = Dtype{reflect.TypeOf(float64(1))}
	Complex64  = Dtype{reflect.TypeOf(complex64(1))}
	Complex128 = Dtype{reflect.TypeOf(complex128(1))}

	// aliases
	Byte = Uint8

	// extras
	Uintptr       = Dtype{reflect.TypeOf(uintptr(0))}
	UnsafePointer = Dtype{reflect.TypeOf(unsafe.Pointer(&Uintptr))}
)

func GetDtype[T DT]() Dtype {
	return Dtype{reflect.TypeOf(T(1))}
}
