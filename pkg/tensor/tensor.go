package tensor

type tensor[T DT] struct {
	AP
	array[T]
}

func (m *tensor[T]) Shape() Shape {
	return m.shape
}

func (m *tensor[T]) Strides() []uint {
	return m.strides
}

func (m *tensor[T]) Dtype() Dtype {
	return m.t
}

func (m *tensor[T]) Dims() uint {
	return uint(len(m.shape))
}

func (m *tensor[T]) Size() uint {
	return m.shape.TotalSize()
}

func Zero(s Shape, t Dtype) Tensor {
	return &tensor[int]{
		AP: NewAP(s),
	}
}
