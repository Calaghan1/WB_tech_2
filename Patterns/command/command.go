// инкапсулирование запроса в виде объекта

package main

import "fmt"

// Command interface
type Command interface {
	Execute()
}

// Receiver
type Light struct {
	isOn bool
}

func (l *Light) On() {
	fmt.Println("Light is On")
	l.isOn = true
}

func (l *Light) Off() {
	fmt.Println("Light is Off")
	l.isOn = false
}

// ConcreteCommand
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(cmd Command) {
	r.command = cmd
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}
	lightOnCmd := &LightOnCommand{light: light}
	lightOffCmd := &LightOffCommand{light: light}

	remote := &RemoteControl{}

	remote.SetCommand(lightOnCmd)
	remote.PressButton()

	remote.SetCommand(lightOffCmd)
	remote.PressButton()
}
