func leaders(nums []int) []int {
    leader := []int{}
    maxRight := nums[len(nums)-1]

    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] >= maxRight {
            leader = append(leader, nums[i])
            maxRight = nums[i]
        }
    }

    left := 0
    right := len(leader) - 1

    for left < right {
        leader[left], leader[right] = leader[right], leader[left]
        left++
        right--
    }

    return leader
}
