package tensor

// Shape represents the dimensions of a Tensor. A (2,3) matrix has a shape of (2,3) - 2 rows, 3 columns.
// Likewise, a shape of (2,3,4) means a Tensor has 3 dimensions: 2 layers, 3 rows, 4 columns.
type Shape []int

// TotalSize returns the number of elements expected in a Tensor of a certain shape
func (s Shape) TotalSize() int {
	return Prod(s)
}

// CalcStrides calculates the default strides for a shape
func (s Shape) CalcStrides() []int {
	if s.IsScalar() {
		return nil
	}
	retVal := make([]int, len(s))
	stride := 1
	for i := len(s) - 1; i >= 0; i-- {
		retVal[i] = stride
		d := s[i]
		if d < 0 {
			panic("negative dimension size does not make sense")
		}
		stride *= d
	}
	return retVal
}

// IsScalar returns true if the access pattern indicates it's a scalar value
func (s Shape) IsScalar() bool {
	return len(s) == 0 || (len(s) == 1 && s[0] == 1)
}

// IsScalarEquiv returns true if the access pattern indicates it's a scalar-like value
func (s Shape) IsScalarEquiv() bool {
	if len(s) == 0 {
		return true
	}
	isEquiv := true
	for i := range s {
		if s[i] != 1 {
			return false
		}
	}
	return isEquiv
}

// IsVector returns whether the access pattern falls into one of three possible definitions of vectors:
//
//	vanilla vector (not a row or a col)
//	column vector
//	row vector
func (s Shape) IsVector() bool { return s.IsColVec() || s.IsRowVec() || (len(s) == 1 && s[0] > 1) }

// IsColVec returns true when the access pattern has the shape (x, 1)
func (s Shape) IsColVec() bool { return len(s) == 2 && (s[1] == 1 && s[0] > 1) }

// IsRowVec returns true when the access pattern has the shape (1, x)
func (s Shape) IsRowVec() bool { return len(s) == 2 && (s[0] == 1 && s[1] > 1) }

// IsMatrix returns true if it's a matrix. This is mostly a convenience method. RowVec and ColVecs are also considered matrices
func (s Shape) IsMatrix() bool { return len(s) == 2 }

// Dims returns the number of dimensions in the shape
func (s Shape) Dims() int { return len(s) }
