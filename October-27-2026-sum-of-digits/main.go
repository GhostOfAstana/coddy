func sumOfDigits(n int) int {
    sum := 0
    for n > 0 {
        sum += n % 10
        n = n / 10
    }
    return sum
}
