package matrix

func MatrixReshape(nums [][]int, r int, c int) [][]int {

	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	m2 := make([][]int, r)
	cols := len(nums[0])
	var cRow, cCol int
	for i := 0; i < r; i++ {
		row := make([]int, c)
		for j := 0; j < c; j++ {
			row[j] = nums[cRow][cCol]
			if cCol == cols {
				cCol = 0
				cRow++
			}
		}
		m2[i] = row
	}
	return m2
}
