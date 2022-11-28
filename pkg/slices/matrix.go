package slices

func Make2dMatrix[T any](xSize, ySize int) (matrix [][]T) {
	matrix = make([][]T, xSize)
	for x := range matrix {
		matrix[x] = make([]T, ySize)
	}
	return
}

func Make3dMatrix[T any](xSize, ySize, zSize int) (matrix [][][]T) {
	matrix = make([][][]T, xSize)
	for x := range matrix {
		matrix[x] = make([][]T, ySize)
		for y := range matrix[x] {
			matrix[x][y] = make([]T, zSize)
		}
	}
	return
}
