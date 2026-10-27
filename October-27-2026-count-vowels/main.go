func countVowels(s string) int {
    runes := []rune(s)
    count := 0
    for i := 0; i < len(runes); i++ {
        if runes[i] == 'A' || runes[i] == 'a' || runes[i] == 'E' || runes[i] == 'e' || runes[i] == 'I' || runes[i] == 'i' || runes[i] == 'O' || runes[i] == 'o' || runes[i] == 'U' || runes[i] == 'u' {
            count++
        }
    }
    return count
}
