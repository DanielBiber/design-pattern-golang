package main

import "fmt"

type Pizza interface {
	getName() string
}

type APizza struct {
	size string
}

func (a APizza) getName()  string{
	return "APizza"
}
type BPizza struct {
	size string
}

func (b BPizza) getName()  string{
	return "BPizza"
}
type PizzaStore struct {
	pizzaFactory PizzaFactory
}

func (p *PizzaStore) orderPizza(t int) Pizza {
	return p.pizzaFactory.cookPizza(t)
}
type PizzaFactory struct {
	
}

func (p *PizzaFactory) cookPizza(t int) Pizza {
	if t == 1{
		return APizza{"10"}
	}
	return BPizza{"5"}
}
func main() {
	ps := &PizzaStore{PizzaFactory{}}
	fmt.Println(ps.orderPizza(1).getName())
	fmt.Println(ps.orderPizza(2).getName())
}
