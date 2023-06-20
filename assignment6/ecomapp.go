package main

import (
	"fmt"
	"regexp"
)

type Item struct {
	ID    int
	Name  string
	Price float64
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
		total += cartItem.Item.Price * float64(cartItem.Amount)
	}
	return total
}

func (app *EcommerceApp) ListAllItems() {
	fmt.Println("----- Items Available -----")
	for _, item := range app.Inventory.Items {
		fmt.Printf("ID: %d | Name: %s | Price: %.2f\n", item.ID, item.Name, item.Price)
	}
	fmt.Println("---------------------------")
}

func (app *EcommerceApp) SearchItems(query string) {
	fmt.Println("----- Search Results -----")
	regex, err := regexp.Compile(query)
	if err != nil {
		fmt.Println("Invalid regular expression:", err)
		return
	}

	for _, item := range app.Inventory.Items {
		if regex.MatchString(item.Name) {
			fmt.Printf("ID: %d | Name: %s | Price: %.2f\n", item.ID, item.Name, item.Price)
		}
	}
	fmt.Println("--------------------------")
}

func main() {
	app := EcommerceApp{
		Inventory: Inventory{
			Items: []Item{
				{ID: 1, Name: "Shirt", Price: 25.99},
				{ID: 2, Name: "Jeans", Price: 39.99},
				{ID: 3, Name: "Shoes", Price: 59.99},
				{ID: 4, Name: "Hat", Price: 14.99},
			},
		},
		Cart: Cart{},
	}

	app.ListAllItems()
	app.AddToCart(app.Inventory.Items[0], 2)
	app.AddToCart(app.Inventory.Items[2], 1)
	fmt.Println("Cart Total:", app.CalculateTotalBill())

	app.SearchItems("Sh")
}
