package main

import "fmt"

type HDD struct {

}


func (h *HDD)StartHdd(){
	fmt.Println("Startin Hdd")
}

type DDR struct {

}

func (d *DDR)StartDDR() {
	fmt.Println("Starting DDR")
}

type CPU struct {

}

func (c *CPU)StartCPU() {
	fmt.Println("Starting CPU")
}

type PC struct{
	h HDD
	d DDR
	c CPU
}

func (p *PC)StartPC() {
	p.h.StartHdd()
	p.d.StartDDR()
	p.c.StartCPU()
}

func main() {
	p := PC{HDD{},DDR{},CPU{}}
	p.StartPC()
}