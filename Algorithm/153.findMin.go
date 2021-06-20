/*
 * 二分查找
 * 时间复杂度: O(logn) 空间: O(1)
 */
 func findMin(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := (left + right) >> 1
        if nums[mid] < nums[right] {
            right = mid
        } else {
           left = mid + 1 
        }
    }
    return nums[left]
}