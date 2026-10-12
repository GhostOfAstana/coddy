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
        sliceProduct = append(sliceProduct, productSecondDivideStr[0])
    }

    fmt.Printf("Initial Inventory:\n")
    for i := 0; i < len(sliceProduct); i++ {
        displayOrder := sliceProduct[i]
        price := product[displayOrder].Price
        quantity := product[displayOrder].Quantity
        fmt.Printf("%s: $%.2f (Stock: %d)\n", displayOrder, price, quantity)
    }

    operation := strings.Split(operationsStr, ",")
	for i := 0; i < len(operation); i++ {
		operDescribe := strings.Split(operation[i], ":")
		if operDescribe[0] == "price" {
		    opername, err := strconv.ParseFloat(operDescribe[2], 64)
                if err != nil {
			    fmt.Printf("Invalid price: %v\n", err)
			    return
		    }
            fmt.Printf("Updated %s: price changed to %.2f\n", operDescribe[1], opername)
            product[operDescribe[1]].Price = opername
        }
        if operDescribe[0] == "quantity" {
            quantity, err := strconv.Atoi(operDescribe[2])
            if err != nil {
			    fmt.Printf("Invalid quantity: %v\n", err)
			    return
		    }
            product[operDescribe[1]].Quantity = quantity
            fmt.Printf("Updated %s: quantity changed to %d\n", operDescribe[1], quantity)
        }
	}

    var total float64
    fmt.Printf("Updated Inventory:\n")
    for i := 0; i < len(sliceProduct); i++ {
        displayOrder := sliceProduct[i]
        price := product[displayOrder].Price
        quantity := product[displayOrder].Quantity
        fmt.Printf("%s: $%.2f (Stock: %d)\n", displayOrder, price, quantity)
        total += price * float64(quantity)
    }

    fmt.Printf("Total Inventory Value: $%.2f\n", total)
}
