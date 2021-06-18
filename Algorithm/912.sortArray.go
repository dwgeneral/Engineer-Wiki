package main

/*
 * 快速排序
 * 时间复杂度: O(nlogn) n 为数组长度 空间复杂度: O(n) n 为递归调用层数
 */
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

/*
 * 堆排序
 * 堆是一个完全二叉树
 * 大顶堆用于升序排列; 小顶堆用于降序排列
 * 创建堆, 调整堆, 交换首尾节点(为了维持一个完全二叉树才要进行首尾交换)
 * O(nlogn) O(n)
 * https://www.bilibili.com/video/BV1Eb41147dK
 */
func sortArray(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	n := len(nums)
	buildHeap(nums, n)
	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] // 把根节点和最后一个节点交换
		heapify(nums, i, 0)                 // 砍断最后一个节点, 重新调整堆
	}
	return
}

func heapify(nums []int, n, i int) {
	if i >= n {
		return
	}
	c1, c2, max := 2*i+1, 2*i+2, i
	if c1 < n && nums[c1] > nums[max] {
		max = c1
	}
	if c2 < n && nums[c2] > nums[max] {
		max = c2
	}
	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heapify(nums, n, max)
	}
}

func buildHeap(nums []int, n int) {
	lastNode := n - 1
	parent := lastNode - 1>>1
	for i := parent; i >= 0; i-- {
		heapify(nums, n, i)
	}
}
