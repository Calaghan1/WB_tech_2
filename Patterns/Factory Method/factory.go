// Паттерн "Фабричный метод" (Factory Method) - это порождающий паттерн проектирования, который предоставляет интерфейс для создания объектов, 
// но позволяет подклассам выбирать класс создаваемого объекта. Фабричный метод делегирует ответственность по созданию объектов подклассам, предоставляя интерфейс 
// для создания объектов в суперклассе, но оставляя реализацию создания конкретных объектов для подклассов.




package main

import "fmt"

// Product interface
type Product interface {
	Use() string
}

// ConcreteProductA
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A"
}

// ConcreteProductB
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B"
}

// Creator interface
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreatorA
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	creatorA := &ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	fmt.Println("Product A:", productA.Use())

	creatorB := &ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	fmt.Println("Product B:", productB.Use())
}