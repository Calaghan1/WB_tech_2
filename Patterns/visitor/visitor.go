// Паттерн Посетитель был впервые описан в книге "Большие шести шаблонов проектирования" (Design Patterns: Elements of Reusable Object-Oriented Software) 
// Гаммы, Хелма, Джонсона и Влиссидеса.
//  Этот паттерн используется для обработки объектов, которые имеют сложную структуру и требуют выполнения различных операций в зависимости от их типа. 
//  Он предотвращает необходимость изменения 
//  класса объекта для добавления новой функциональности.

// Основной идеей паттерна Посетитель является выделение операции по обработке объектов из классов объектов, 
// в которых эти операции выполняются. Вместо этого создаются отдельные классы-посетители, которые реализуют различные операции над объектами. 
// Классы-посетители могут использоваться для обработки объектов разных типов, не изменяя при этом сами объекты.

package main

import (
	"fmt"
)


type Animal interface {
	Accept(visitor AnimalVisitor) string
	}

	type AnimalVisitor interface {
		DogSound(d *Dog) string
		CatSound(c *Cat) string
		LionSound(l *Lion) string
		}

type Dog struct {
	Name string
}
func (d *Dog) Accept(visitor AnimalVisitor) string {
	return visitor.DogSound(d)
}
type Cat struct {
	Name string
}

func (c *Cat) Accept(visitor AnimalVisitor) string {
	return visitor.CatSound(c)
}

type Lion struct {
	Name string
}

func (l *Lion) Accept(visitor AnimalVisitor) string {
	return visitor.LionSound(l)
}

type AnimalSoundVisitor struct{}

func (v *AnimalSoundVisitor) DogSound(d *Dog) string {
	return fmt.Sprintf("%s says Woof!", d.Name)
}

func (v *AnimalSoundVisitor) CatSound(c *Cat) string {
	return fmt.Sprintf("%s says Meow!", c.Name)
}

func (v *AnimalSoundVisitor) LionSound(l *Lion) string {
	return fmt.Sprintf("%s says Roar!", l.Name)
}

func main() {
	dog := &Dog{"Rex"}
	cat := &Cat{"Lucy"}
	lion := &Lion{"Leo"}
	
	visitor := &AnimalSoundVisitor{}
	
	fmt.Println(dog.Accept(visitor))
	fmt.Println(cat.Accept(visitor))
	fmt.Println(lion.Accept(visitor))
}