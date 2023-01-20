package tensor

// Shape represents the dimensions of a Tensor. A (2,3) matrix has a shape of (2,3) - 2 rows, 3 columns.
// Likewise, a shape of (2,3,4) means a Tensor has 3 dimensions: 2 layers, 3 rows, 4 columns.
type Shape []uint

// TotalSize returns the number of elements expected in a Tensor of a certain shape
func (s Shape) TotalSize() uint {
	return Prod(s)
}
