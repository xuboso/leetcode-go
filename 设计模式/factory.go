/**
 * 工厂模式
 */
package main

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	println("Circle Draw")
}

func NewShape(shapeType string) Shape {
	if shapeType == "circle" {
		return &Circle{}
	}
	return nil
}
