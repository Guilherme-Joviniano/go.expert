package main

import "fmt"

type Car struct {
	name  string
	price int
	year  int
}

type Client struct {
	name   string
	age    int
	active bool
	cars   []Car
}

func (p *Client) disable() {
	p.active = false
}

func main() {

	client := Client{
		name:   "Guilherme",
		age:    17,
		active: true,
		cars:   []Car{{name: "Toyota Hilux", price: 300000, year: 2022}},
	}

	fmt.Printf("nome %s, idade %d, ativo %t \n", client.name, client.age, client.active)
	for _, car := range client.cars {
		fmt.Println("Carros: ")
		fmt.Println(car.name, car.price, car.year)
	}
}
