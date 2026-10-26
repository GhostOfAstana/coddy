package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Прочитать ввод
	var numOpsStr string
	var operationsStr string
	fmt.Scanln(&numOpsStr)
	fmt.Scanln(&operationsStr)
	
	// Объявить maps без инициализации (они будут nil)
	var scores map[string]int
	var grades map[string]int
	
	// Разделить operations по запятой
	operations := strings.Split(operationsStr, ",")
	
	// TODO: Напишите ваш код ниже
	// Обработать каждую операцию последовательно
	// Обработать операции "check", "init" и "add"
	// Для операций add используйте strconv.Atoi() для преобразования строки value в integer
	for i := 0; i < len(operations); i++ {
        divideOps := strings.Split(operations[i], ":")
        switch divideOps[0] {
        case "check":
            if divideOps[1] == "scores" {
                if scores == nil {
                    fmt.Printf("Map scores is nil\n")
                } else {
                    fmt.Printf("Map scores is initialized\n")
                }
            } else if divideOps[1] == "grades" {
                if grades == nil {
                    fmt.Printf("Map grades is nil\n")
                } else {
                    fmt.Printf("Map grades is initialized\n")
                }
            }
        case "init":
            if divideOps[1] == "scores" {
                scores = make(map[string]int)
                fmt.Printf("Initialized map scores\n")
            } else if divideOps[1] == "grades" {
                grades = make(map[string]int)
                fmt.Printf("Initialized map grades\n")
            }
        case "add":
            if divideOps[1] == "scores" {
                if scores == nil {
                    fmt.Printf("Cannot add to nil map %s\n", divideOps[1])
                } else {
                    value, err := strconv.Atoi(divideOps[3])
                    if err != nil {
                        fmt.Printf("Invalid case %v\n", err)
                        return
                    }
                    scores[divideOps[2]] = value
                    fmt.Printf("Added %s:%d to scores\n", divideOps[2], value)
                }
            } else if divideOps[1] == "grades" {
                if grades == nil {
                    fmt.Printf("Cannot add to nil map %s\n", divideOps[1])
                } else {
                    value, err := strconv.Atoi(divideOps[3])
                    if err != nil {
                        fmt.Printf("Invalid case %v\n", err)
                        return
                    }
                    grades[divideOps[2]] = value
                    fmt.Printf("Added %s:%d to grades\n", divideOps[2], value)
                }
            }
        }   
    }
	// Вывести конечное состояние обеих maps
	// Проверить, являются ли maps nil или инициализированными, и вывести соответственно
    if scores == nil {
        fmt.Printf("Final state - scores: nil\n")
    } else {
        fmt.Printf("Final state - scores: %d entries\n", len(scores))
    }
    if grades == nil {
        fmt.Printf("Final state - grades: nil\n")
    } else {
        fmt.Printf("Final state - grades: %d entries\n", len(grades))
    }
}
