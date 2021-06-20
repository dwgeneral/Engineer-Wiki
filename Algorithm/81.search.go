/*
 * 二分查找
 */
 func search(nums []int, target int) bool {
    if len(nums) < 2 {
        if nums[0] == target {
            return true
        }
        return false
    } 
    
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right - left) >> 1
        if nums[mid] == target {
            return true
        }
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
        } else if nums[left] <= nums[mid] {
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1

            }
        } else {
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return false
}