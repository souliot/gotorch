package tensor

import (
	"fmt"
	"unsafe"
)

type array[T DT] struct {
	t Dtype // Dtype for tensor
	v []T   // RawData for tensor
}

func NewArray[T DT](t Dtype, v []T) array[T] {
	return array[T]{
		t: t,
		v: v,
	}
}

func (m array[T]) String() string {
	return fmt.Sprintf("%v\ntype:%v", m.v, m.t)
}

func (m array[T]) Len() int {
	return len(m.v)
}

func (m array[T]) Data() string {
	return fmt.Sprintf("%v", m.v)
}

func (m array[T]) Info() string {
	return fmt.Sprintf("type:%v, ptr:%v, len:%d, cap:%d", m.t, unsafe.SliceData(m.v), len(m.v), cap(m.v))
}

func (m array[T]) Split() string {
	return fmt.Sprintf("type:%v, ptr:%v, len:%d, cap:%d", m.t, unsafe.SliceData(m.v), len(m.v), cap(m.v))
}
