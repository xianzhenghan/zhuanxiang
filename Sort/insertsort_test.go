package Sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{2, 1, 6, 3, 4, 9, 7, 8, 5, 0}
	res := InsertionSort(nums)
	fmt.Printf("res=%v\n", res)
}
