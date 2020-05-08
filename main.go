package main

import (
	"fmt"

	"github.com/midoblgsm/shoppingcart/cart"
	"github.com/midoblgsm/shoppingcart/resources"
)

//test
//test2
func main() {
	cart := cart.NewCart()

	potato := resources.Item{
		ID:       "vg1",
		Name:     "Potato",
		Price:    1,
		Quantity: 2,
	}

	cart.AddItem(potato)
	fmt.Printf("Cart %#v \n", cart)
}
