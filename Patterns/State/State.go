// Паттерн "Состояние" (State) - это поведенческий паттерн проектирования, который позволяет объекту изменять свое поведение при изменении 
// его внутреннего состояния. Паттерн "Состояние" помещает каждое состояние в отдельный класс и делает объект состояния полностью заменяемым. 
// Это позволяет объекту менять свое состояние на лету, прозрачно для клиента.

package main

import "fmt"

// State interface
type State interface {
	HandleRequest()
}

// ConcreteStateA
type LockedState struct{}

func (s *LockedState) HandleRequest() {
	fmt.Println("Door is locked.")
}

// ConcreteStateB
type UnlockedState struct{}

func (s *UnlockedState) HandleRequest() {
	fmt.Println("Door is unlocked. You can pass.")
}

// Context
type Door struct {
	state State
}

func NewDoor() *Door {
	return &Door{state: &LockedState{}}
}

func (d *Door) SetState(state State) {
	d.state = state
}

func (d *Door) Open() {
	d.state.HandleRequest()
}

func main() {
	door := NewDoor()

	door.Open() // Output: Door is locked.

	door.SetState(&UnlockedState{})
	door.Open() // Output: Door is unlocked. You can pass.
}