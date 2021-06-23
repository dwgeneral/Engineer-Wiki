/*
 * 使用中序遍历可以方便的将有序数组转换为二叉树
 * 考虑到高度平衡, 所以从数组中间取值为根节点
 */
 func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums) >> 1
    left := nums[:mid]
    right := nums[mid+1:]
    
    return &TreeNode{ nums[mid], sortedArrayToBST(left), sortedArrayToBST(right) }

}