package main

import "fmt"

type Person struct {
	name   string `json:"username"`
	age    int    `json:"age"`
	active bool   `json:"active"`
}

func (p *Person) disable() {
	p.active = false
}

func main() {
	client := Person{
		name:   "Guilherme",
		age:    17,
		active: true,
	}
	client.disable()

	fmt.Printf("nome %s, idade %d, ativo %t \n", client.name, client.age, client.active)
}
