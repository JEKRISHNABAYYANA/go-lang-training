package main

import (
	"fmt"
	"regexp"
)

type Item struct {
	ID       int
	Product  string
	Cost     float64
	Quantity int
}

type Inventory struct {
	Items []Item
}

type CartItem struct {
	Item   Item
	Amount int
}

type Cart struct {
	Items []CartItem
}

type EcommerceApp struct {
	Inventory Inventory
	Cart      Cart
}

func (app *EcommerceApp) AddToCart(item Item, amount int) {
	cartItem := CartItem{Item: item, Amount: amount}
	app.Cart.Items = append(app.Cart.Items, cartItem)
}

func (app *EcommerceApp) CalculateTotalBill() float64 {
	total := 0.0
	for _, cartItem := range app.Cart.Items {
		total += cartItem.Item.Cost * float64(cartItem.Amount)
	}
	return total
}

func (app *EcommerceApp) ListAllItems() {
	fmt.Println(" Items Available :->")
	for _, item := range app.Inventory.Items {
		fmt.Printf("ID: %d , Product: %s , Quantity: %d, Cost: %.2f\n", item.ID, item.Product, item.Quantity, item.Cost)
	}

}

func (app *EcommerceApp) SearchItems(query string) {
	fmt.Println(" Search Results :->")
	regex, err := regexp.Compile(query)
	if err != nil {
		fmt.Println("Invalid regular expression:", err)
		return
	}

	for _, item := range app.Inventory.Items {
		if regex.MatchString(item.Product) {
			fmt.Printf("ID: %d , Product: %s ,Quantity: %d, Cost: %.2f\n", item.ID, item.Product, item.Quantity, item.Cost)
		}
	}

}

func main() {
	app := EcommerceApp{
		Inventory: Inventory{
			Items: []Item{
				{ID: 1, Product: "Shirt", Quantity: 1, Cost: 66.99},
				{ID: 2, Product: "Trousers", Quantity: 1, Cost: 77.99},
				{ID: 3, Product: "Shoes", Quantity: 1, Cost: 88.99},
				{ID: 4, Product: "Hat", Quantity: 1, Cost: 19.99},
			},
		},
		Cart: Cart{},
	}

	app.ListAllItems()
	app.AddToCart(app.Inventory.Items[0], 2)
	app.AddToCart(app.Inventory.Items[2], 1)
	fmt.Println("Cart Total:", app.CalculateTotalBill())

	app.SearchItems("Ha")
}
