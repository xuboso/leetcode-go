/**
 * 工厂方法模式是一种创建型设计模式， 其在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型。
 * https://refactoringguru.cn/design-patterns/factory-method
 */
package main

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}

	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
