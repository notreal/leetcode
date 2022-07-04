//1. Two Sum
//Runtime: 34 ms, faster than 30.04% of Go online submissions for Two Sum.
//Memory Usage: 3.7 MB, less than 72.62% of Go online submissions for Two Sum.

func twoSum(nums []int, target int) []int {
    var ret1, ret2 int
    for i, n := range nums {
        for i2, n2 := range nums[i+1:] {
            if n + n2 == target {
                ret1, ret2 = i, i + i2 + 1
                goto exit
            }
        }
    }
    exit:
    return []int{ret1, ret2}
}