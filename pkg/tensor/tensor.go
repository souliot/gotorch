package tensor

type tensor[T DT] struct {
	AP
	array[T]
}

func (m tensor[T]) String() string {
	for len(m.shape) <= 1 {
		return m.array.String()
	}
	return m.array.String()
}

func (m *tensor[T]) Shape() Shape {
	return m.shape
}

func (m *tensor[T]) Strides() []int {
	return m.strides
}

func (m *tensor[T]) Dtype() Dtype {
	return m.t
}

func (m *tensor[T]) Dims() int {
	return len(m.shape)
}

func (m *tensor[T]) Size() int {
	return m.shape.TotalSize()
}

func (m *tensor[T]) DataSize() int {
	return m.array.Len()
}

func (m *tensor[T]) IsScalar() bool {
	return m.shape.IsScalar()
}

func Zero[T DT](s Shape) Tensor {
	dt := GetDtype[T]()
	size := s.TotalSize()
	return &tensor[T]{
		AP:    NewAP(s),
		array: NewArray(dt, make([]T, size)),
	}
}
