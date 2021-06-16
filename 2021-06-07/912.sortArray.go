// 时间复杂度: O(nlogn) n 为数组长度 空间复杂度: O(n) n 为递归调用层数
func SortArray(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

func partition(arr []int, left, right int) int {
	p := arr[left]
	for left < right {
		for left < right && p <= arr[right] {
			right--
		}
		arr[left] = arr[right]
		for left < right && p >= arr[left] {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = p
	return left
}
