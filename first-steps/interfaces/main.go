package main

import "fmt"

func main() {
	guilherme := Client{
		name:   "Guilherme",
		active: true,
	}

	Desativar(&guilherme)
}

func Desativar(person Person) {
	person.Disable()
}

type Person interface {
	Disable()
	Enable()
}

type Client struct {
	name   string
	active bool
}

func (c *Client) Disable() {
	c.active = false
	fmt.Printf("O cliente %s foi desativado")
}
func (c *Client) Enable() {
	c.active = true
	fmt.Printf("O cliente %s foi ativado")
}
