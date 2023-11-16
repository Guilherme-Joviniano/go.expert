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

func (c *Client) disable() {
	c.active = false
	fmt.Printf("O client %s foi desativado \n", c.name)
}

func (c *Client) enable() {
	c.active = true
	fmt.Printf("O client %s foi ativado \n", c.name)
}

func main() {

	client := Client{
		name:   "Guilherme",
		age:    17,
		active: true,
		cars:   []Car{{name: "Toyota Hilux", price: 300000, year: 2022}},
	}

	client.disable()

	client.enable()
}
