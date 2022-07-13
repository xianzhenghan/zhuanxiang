package Sort

import "fmt"

//原理：遍历数组，比较并将大的元素与下一个元素交换位置，
//在一轮的循环之后，可以让未排序i的最大元素排列到数组右侧。
//在一轮循环中，如果没有发生元素位置交换，那么说明数组已经是有序的，此时退出排序。
func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		ischang := false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				fmt.Printf("sweped arr=%v	", arr)
				ischang = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Printf("arr=%v\n", arr)
			}
		}
		if ischang == false {
			break
		}
	}
	return arr
}
