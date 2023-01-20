package tensor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestZero(t *testing.T) {
	tests := []struct {
		shape Shape
		size  int
	}{
		{
			shape: Shape{3},
			size:  3,
		},
		{
			shape: Shape{3, 1},
			size:  3,
		},
		{
			shape: Shape{3, 3},
			size:  9,
		},
		{
			shape: Shape{3, 3, 3},
			size:  27,
		},
	}
	Convey("Size", t, func() {
		for _, v := range tests {
			tensor := Zero[int8](v.shape)
			size := tensor.Size()
			So(size, ShouldEqual, v.size)
			t.Log(tensor, tensor.Strides())
		}
	})
}
