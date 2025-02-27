/**
 * 装饰模式是一种结构型设计模式， 允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为。
 * https://refactoringguru.cn/design-patterns/decorator
 */
package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {

	pizza := &VeggieMania{}

	// Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	// Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
