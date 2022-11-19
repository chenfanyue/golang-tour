/*
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
*/
func findRepeatNumber(nums []int) int {
    var res int
    for i := range nums {
        for nums[i] != i {
            v := nums[i]
            if nums[v] == v {
                res = v
                goto ret // goto只能往外层跳
            }
            nums[i], nums[v] = nums[v], nums[i]
        }
    }
    ret:
    return res
}
