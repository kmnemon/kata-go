package main

import "fmt"

type H struct{
}

func (h H) abc(){

}

func main() {
	h := H{}
	d := H.abc
	d(h)
	
	
}
