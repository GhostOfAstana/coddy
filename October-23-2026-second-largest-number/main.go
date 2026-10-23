func secondLargest(nums []int) (int) {
    largest := nums[0]
    second := 0
    hasSecond := false
    for i := 1; i < len(nums); i++ {
        if nums[i] > largest {
            second = largest
            largest = nums[i]
            hasSecond = true
        } else if nums[i] < largest && (!hasSecond || nums[i] > second) {
            second = nums[i]
            hasSecond = true
        }
    }
    return second
}
