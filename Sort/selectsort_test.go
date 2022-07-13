package Sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{2, 1, 6, 3, 4, 9, 7, 8, 5, 0}
	res := SelectionSort(nums)
	fmt.Printf("res=%v\n", res)
}
