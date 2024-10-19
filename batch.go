package slicendice

// Given total length of array and batch size
// returns tuple of [startIndex, endIndex]
// Handles both odd and even
// Golang truncated division https://go.dev/ref/spec#Arithmetic_operators
// total = 12 groups = 5
// totalGroup = 12/5 + (12 % 5 != 0 ? 1 : 0) = 2 + 1
// [0, 4], [5, 9], [10, 11]
//
// total = 12, groups = 3
// totalGroups = (12 / 3) + (12 % 3) = 4
// [0, 2], [3,5], [6, 8], [9, 11]
func Batch(total int64, groups int) [][]int {
	if total <= int64(groups) {
		return [][]int{{0, int(total) - 1}}
	}
	totalGroups := (total / int64(groups))
	if total%int64(groups) != 0 {
		totalGroups += 1
	}

	batchIndexs := [][]int{}
	groupie := groups - 1 // we will be returning indexes, so groups of 4, (0, 3)

	for i := 0; i < int(totalGroups); i++ {
		lft := (i * groupie) + i
		ryt := Min(lft+groupie, int(total-int64(1)))

		batchIndexs = append(batchIndexs, []int{lft, ryt})
	}
	return batchIndexs
}
