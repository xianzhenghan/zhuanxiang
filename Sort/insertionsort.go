package Sort

//原理：数组先看成两部分，排序序列和未排序序列。
//排序序列从第一个元素开始，该元素可以认为已经被排序。遍历数组，
//每次将扫描到的元素与之前的元素相比较，插入到有序序列的适当位置。

func InsertionSort(arr []int) []int {
	for curr := 1; curr < len(arr); curr++ {
		for pre := curr - 1; pre >= 0; pre-- {
			if arr[pre+1] < arr[pre] {
				arr[pre+1], arr[pre] = arr[pre], arr[pre+1]
			}
		}
	}
	return arr
}
