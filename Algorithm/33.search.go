/*
 * 因为是部分有序数组, 且进阶要求logn复杂度实现,考虑可以使用二分查找来做   
 * 时间复杂度: O(logn) 空间复杂度: O(1)
 */
func search(nums []int, target int) int {
    if len(nums) < 2 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }
    
    left, right := 0, len(nums)-1
    for left < right {
        mid := (left + right) >> 1 
        if target == nums[mid] {
            return mid
        }
        if nums[left] <= nums[mid] {
            if nums[left] <= target && target < nums[mid] {
                right = mid
            } else {
                left = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid
            }
        }
    }
    return -1
}