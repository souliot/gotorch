package tensor

import (
	"reflect"
	"unsafe"
)

type array[T DT] struct {
	hdr *reflect.SliceHeader // SliceHeader of the storage array for tensor
	t   Dtype                // Dtype for tensor
	v   []T                  // RawData for tensor
}

func NewArray[T DT](t Dtype, v []T) array[T] {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	return array[T]{
		hdr: sh,
		t:   t,
		v:   v,
	}
}
