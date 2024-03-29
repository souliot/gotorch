package tensor

type Tensor interface {
	// Info
	Shape() Shape
	Strides() []int
	Dtype() Dtype
	Dims() int
	Size() int
	DataSize() int
	// 标量
	IsScalar() bool
	IsVector() bool
	IsColVec() bool
	IsRowVec() bool
	IsMatrix() bool
	IsZero() bool
}

type CMP interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type DT interface {
	CMP | ~complex64 | ~complex128
}
