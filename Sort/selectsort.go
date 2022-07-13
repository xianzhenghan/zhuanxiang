package Sort

// 选择排序 N^2 时间复杂度 不稳定
//原理：遍历数组， 从中选择最小元素，将它与数组的第一个元素交换位置。
//继续从数组剩下的元素中选择出最小的元素，将它与数组的第二个元素交换位置。循环以上过程，直到将整个数组排序。
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}
