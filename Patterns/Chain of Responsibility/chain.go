// Паттерн "Цепочка вызовов" (Chain of Responsibility) - это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков. 
// Каждый обработчик решает, может ли он обработать запрос, и либо обрабатывает его, либо передает дальше по цепочке.

package main

package main

import "fmt"

// Handler interface
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request int)
}

// ConcreteHandlerA
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerA) HandleRequest(request int) {
	if request < 10 {
		fmt.Println("Request handled by ConcreteHandlerA")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// ConcreteHandlerB
type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerB) HandleRequest(request int) {
	if request >= 10 && request < 20 {
		fmt.Println("Request handled by ConcreteHandlerB")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// ConcreteHandlerC
type ConcreteHandlerC struct {
	next Handler
}

func (h *ConcreteHandlerC) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerC) HandleRequest(request int) {
	if request >= 20 {
		fmt.Println("Request handled by ConcreteHandlerC")
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
handlerC := &ConcreteHandlerC{}
handlerA.SetNext(handlerB)
handlerB.SetNext(handlerC)

requests := []int{5, 15, 25}
for _, request := range requests {
handlerA.HandleRequest(request)
}
}