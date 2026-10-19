func countVowels(s string) int {
    runes := []rune(s)
    vowels := 0
    for i := 0; i < len(runes); i++ {
        if runes[i] == 'a' || runes[i] == 'e' || runes[i] == 'i' || runes[i] == 'o' || runes[i] == 'u' || runes[i] == 'A' || runes[i] == 'E' || runes[i] == 'I' || runes[i] == 'O' || runes[i] == 'U'{
            vowels++
        }
    }
    return vowels
}
