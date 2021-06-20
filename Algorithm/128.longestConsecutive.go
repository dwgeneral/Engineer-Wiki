/*
 * 使用哈希表记录所有整数是否出现过 O(1)
 * 然后遍历数组球连续序列的长度, x-1如果计算过, 则没必要计算x;
 * O(n) O(n)
 */
 func longestConsecutive(nums []int) (res int) {
    hashmap := make(map[int]bool, len(nums))
    for i := 0; i < len(nums); i++ {
        hashmap[nums[i]] = true
    }
    for i := 0; i < len(nums); i++ {
        if !hashmap[nums[i]-1] {
            curr := nums[i]
            currSeq := 1
            for hashmap[curr+1] {
                curr++
                currSeq++
            }
            if currSeq > res {
                res = currSeq
            }
        }
    }
    return res
}