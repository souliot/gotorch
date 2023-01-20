package tensor

import (
	"fmt"
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

func (m array[T]) String() string {
	return fmt.Sprintf("%v\ntype:%v, data:%v, len:%d, cap:%d", m.v, m.t, m.hdr.Data, m.hdr.Len, m.hdr.Cap)
}

func (m array[T]) Data() string {
	return fmt.Sprintf("%v", m.v)
}

func (m array[T]) Info() string {
	return fmt.Sprintf("type:%v, data:%v, len:%d, cap:%d", m.t, m.hdr.Data, m.hdr.Len, m.hdr.Cap)
}
