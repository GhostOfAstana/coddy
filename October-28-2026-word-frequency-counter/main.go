package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Прочитать полную строку текста (включая пробелы)
	scanner.Scan()
	text := scanner.Text()

	// Прочитать значение порога
	scanner.Scan()
	thresholdStr := scanner.Text()

	// Преобразовать threshold в целое число
	threshold, _ := strconv.Atoi(thresholdStr)

	// TODO: Напишите ваш код ниже
	// 1. Создайте функцию countWords, которая возвращает map[string]int
	//    Подсказка: используйте strings.Fields и strings.ToLower для разделения и нормализации слов
	// 2. Use the function to count word frequencies
	// 3. Отфильтруйте слова, которые появляются как минимум 'threshold' раз
	// 4. Отсортируйте отфильтрованные слова по алфавиту и отобразите их
	// 5. Вычислите и отобразите сводную статистику
    func countWords(s string) map[string]int {
        wordsCount := map[string]int{}
        words := strings.Fields(s)
        current := ""
        for i := 0; i < len(words); i++ {
            current = strings.ToLower(words[i])
        }
    }
    //func myFields(s string) []string {
    //    words := []string{}              // Слайс готовых слов
    //    current := ""                    // Текущее слово
    //    for i := 0; i < len(s); i++ {    // Перебираем байты строки
    //        if s[i] != ' ' {              // Если это не пробел
    //        current += string(s[i])  // Добавляем символ к слову
    //    } else if current != "" {    // Если пробел встретился после слова
    //        words = append(words, current) // Сохраняем слово
    //        current = ""             // Начинаем собирать новое
    //    }
    //}
    //    if current != "" {               // Если строка закончилась словом
    //        words = append(words, current) // Сохраняем последнее слово
    //    }
    //    return words                     // Возвращаем все слова
    //}
	_ = text
	_ = threshold
	fmt.Println("Word Frequency Analysis:")
}
