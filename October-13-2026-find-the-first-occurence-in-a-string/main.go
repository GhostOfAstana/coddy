package main

import "fmt"

func strStr(haystack string, needle string) (int) {
    for i := 0; i + len(needle) <= len(haystack); i++ {
        if haystack[i:i+len(needle)] == needle {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(strStr("bananarama", "ana")) // 1
    fmt.Println(strStr("banana", "xyz"))     // -1
    fmt.Println(strStr("hello", "ll"))       // 2
}
