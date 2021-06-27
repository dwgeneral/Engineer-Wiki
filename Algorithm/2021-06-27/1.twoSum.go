package algorithm

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := m[diff]; ok {
			return []int{i, v}
		}
		m[nums[i]] = i
	}
	return []int{}
}
