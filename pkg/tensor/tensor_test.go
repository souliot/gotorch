package tensor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestZero(t *testing.T) {
	tests := []struct {
		shape Shape
		size  uint
	}{
		{
			shape: Shape{3},
			size:  3,
		},
		{
			shape: Shape{1, 3},
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
			tensor := Zero(v.shape, Int32)
			size := tensor.Size()
			So(size, ShouldEqual, v.size)
		}

	})
}
