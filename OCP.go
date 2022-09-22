package main

import "fmt"

type A struct {
	year int
}

func (a A) Get() {
	fmt.Println("Hi Devps to the year: ", a.year)
}

type B struct {
	A
}

func (b B) Get() {
	fmt.Println("Hi Devps to the year: ", b.year)
}

func main() {
	var a A
	a.year = 2016
	var b B
	b.year = 2018
	a.Get()
	b.Get()
}
