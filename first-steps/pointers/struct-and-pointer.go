package main

import (
	"errors"
	"fmt"
)

type Client struct {
	name string
}

type Account struct {
	currency float64
}

func NewAccount() *Account {
	return &Account{
		currency: 0,
	}
}

func (ac *Account) Add(value float64) (float64, error) {
	if value <= float64(0) {
		return float64(0), errors.New("Valor é menor que zero")
	}
	return ac.currency + value, nil
}

func (ac Account) Simulate(value float64) float64 {
	ac.currency += value
	return ac.currency
}

func (c *Client) walked() {
	c.name = "Wesley Willians"
	fmt.Printf("O client %s andou! \n", c.name)
}

func main() {
	wesley := Client{
		name: "Wesley",
	}
	wesley.walked()
	println(wesley.name)

	conta := NewAccount()

	saldo := conta.Simulate(200)
	println(conta.currency)
	println(saldo)
}
