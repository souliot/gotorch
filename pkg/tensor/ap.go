package tensor

type AP struct {
	shape   Shape
	strides []int
}

func NewAP(shape Shape) AP {
	return AP{
		shape: shape,
	}
}
