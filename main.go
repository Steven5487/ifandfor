package main

import "fmt"

var productName string = "Galaxy S21 Ultra"
var productPrice int = 10000
var productStock int = 100

func main() {

	fmt.Printf("Product Name: %s\n", productName)
	fmt.Printf("Product Price: %d\n", productPrice)

	var sum int
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum of first 10 natural numbers: %d\n", sum)

	key := []string{"a", "b", "c"}
	for i := 0; i < len(key); i++ {
		fmt.Println(key[i])
	}

	for {
		fmt.Printf("Product Stock: %d\n", productStock)

		var quantity int
		fmt.Print("Enter quantity to buy: ")
		fmt.Scan(&quantity)

		if quantity > productStock {
			fmt.Printf("Insufficient stock. Only %d units available.\n", productStock)
		} else if quantity <= 0 {
			fmt.Println("Invalid quantity. Please enter a positive integer.")
		} else {
			productStock -= quantity
			fmt.Printf("Remaining stock: %d\n", productStock)
		}

		if productStock == 0 {
			fmt.Println("Stock is out. Thank you for shopping!")
			break
		}
	}
}
