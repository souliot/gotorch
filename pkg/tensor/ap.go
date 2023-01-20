package tensor

type AP struct {
	shape   Shape
	strides []uint
}

func NewAP(shape Shape) AP {
	return AP{
		shape: shape,
	}
}
