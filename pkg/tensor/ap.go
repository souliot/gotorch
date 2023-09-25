package tensor

type AP struct {
	shape   Shape
	strides []int
	fin     bool
}

func NewAP(shape Shape) AP {
	return AP{
		shape:   shape,
		strides: shape.CalcStrides(),
		fin:     true,
	}
}

// Shape returns the shape of the AP
func (ap *AP) Shape() Shape { return ap.shape }

// Strides returns the strides of the AP
func (ap *AP) Strides() []int { return ap.strides }

// Dims returns the dimensions of the shape in the AP
func (ap *AP) Dims() int { return ap.shape.Dims() }

// Size returns the expected array size of the shape
func (ap *AP) Size() int { return ap.shape.TotalSize() }

// IsVector returns whether the access pattern falls into one of three possible definitions of vectors:
//
//	vanilla vector (not a row or a col)
//	column vector
//	row vector
func (ap *AP) IsVector() bool { return ap.shape.IsVector() }

// IsColVec returns true when the access pattern has the shape (x, 1)
func (ap *AP) IsColVec() bool { return ap.shape.IsColVec() }

// IsRowVec returns true when the access pattern has the shape (1, x)
func (ap *AP) IsRowVec() bool { return ap.shape.IsRowVec() }

// IsScalar returns true if the access pattern indicates it's a scalar value
func (ap *AP) IsScalar() bool { return ap.shape.IsScalar() }

// IsMatrix returns true if it's a matrix. This is mostly a convenience method. RowVec and ColVecs are also considered matrices
func (ap *AP) IsMatrix() bool { return len(ap.shape) == 2 }

// IsZero tell us if the ap has zero size
func (ap *AP) IsZero() bool {
	return len(ap.shape) == 0 && len(ap.strides) == 0 && !ap.fin
}
