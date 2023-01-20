package tensor

// Min returns the lowest between two CMP. If both are the same, it returns the first
func Min[T CMP](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// Max returns the highest between two CMP. If both are the same, it returns the first
func Max[T CMP](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

// Mins returns the min of a slice of CMP.
func Mins[T CMP](vs ...T) (retVal T) {
	if len(vs) == 0 {
		return
	}
	retVal = vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] < retVal {
			retVal = vs[i]
		}
	}
	return
}

// Maxs returns the max of a slice of CMP.
func Maxs[T CMP](vs ...T) (retVal T) {
	if len(vs) == 0 {
		return
	}
	retVal = vs[0]
	for i := 1; i < len(vs); i++ {
		if vs[i] > retVal {
			retVal = vs[i]
		}
	}
	return
}

// Sum returns the sum of the elements of the slice.
func Sum[T DT](vs []T) (retVal T) {
	for _, v := range vs {
		retVal += v
	}
	return
}

// Prod returns the product of the elements of the slice.
// Returns 1 if len(s) = 0.
func Prod[T DT](vs []T) (retVal T) {
	retVal = 1
	for _, v := range vs {
		retVal *= v
	}
	return retVal
}
