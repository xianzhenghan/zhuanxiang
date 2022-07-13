package Sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	///nums := []int{2, 1, 6, 3, 4, 9, 7, 8, 5, 0}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := BubbleSort(nums)
	fmt.Printf("res=%v\n", res)
}
