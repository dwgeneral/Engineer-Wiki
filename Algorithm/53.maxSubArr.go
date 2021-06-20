/*
 * 类似于贪心的思想, 只要前一个数+当前数 > 当前数, 就把它们算在一起, 
 * 这样取得的最大值, 就是某一段子数组之和, 且是最大和
 * O(n) O(1)
 */
 func maxSubArray(nums []int) int {
    n, max := len(nums), nums[0]
    if n == 1 {
        return max
    }
    for i := 1; i < n; i++ {
        if nums[i] + nums[i-1] > nums[i] {
            nums[i] += nums[i-1]
        }
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}