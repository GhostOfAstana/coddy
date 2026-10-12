package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numProductsStr string
	var productDataStr string
	var operationsStr string
	
	fmt.Scanln(&numProductsStr)
	fmt.Scanln(&productDataStr)
	fmt.Scanln(&operationsStr)
	
	type Product struct {
        Price float64
        Quantity int
    }
	
	product := map[string]*Product{}
	
    sliceProduct := []string{}
	
    productDividedStr := strings.Split(productDataStr, ",")
    for i := 0; i < len(productDividedStr); i++ {
        productSecondDivideStr := strings.Split(productDividedStr[i], ":")
        price, err := strconv.ParseFloat(productSecondDivideStr[1], 64)
        if err != nil {
            fmt.Printf("Invalid float: %v\n", err)
            return
        } 
        quantity, err := strconv.Atoi(productSecondDivideStr[2])
        if err != nil {
            fmt.Printf("Invalid integer: %v\n", err)
            return
        }
        firstProduct := Product{price, quantity}
        product[productSecondDivideStr[0]] = &firstProduct
    }

}
