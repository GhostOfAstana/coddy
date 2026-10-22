func findMax(nums []int) (int) {
    maxValue := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > maxValue {
            maxValue = nums[i]
        }
    }
    return maxValue
}
