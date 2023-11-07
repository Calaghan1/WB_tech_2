// Паттерн используются для построения сложных объектов  по компонентно
package main

import "fmt"

type Computer struct {
	CPU string
	RAM int
	MB string
}

type ComputerbuilderI interface {
	CPU(val string) ComputerbuilderI
	RAM(val int) ComputerbuilderI
	MB(val string) ComputerbuilderI

	Build() Computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb string
}

func (b *computerBuilder) CPU(val string) ComputerbuilderI {
	b.cpu = val
	return b
}
func (b *computerBuilder) RAM(val int) ComputerbuilderI {
	b.ram = val
	return b
}
func (b *computerBuilder) MB(val string) ComputerbuilderI {
	b.mb = val
	return b
}

func NewcomputerBuilder() ComputerbuilderI {
	return &computerBuilder{}
}

func (b *computerBuilder) Build() Computer {
		return Computer{
			CPU: b.cpu,
			RAM: b.ram,
			MB: b.mb,
		}
}
func main() {
	compBuilder := NewcomputerBuilder()
	comp := compBuilder.CPU("Intel i3").MB("gigabyte").RAM(16).Build()
	fmt.Println(comp)

}