package main

import (
	"fmt"
)

type Target interface {
	Execute()
}

type Adapter struct {
}

func (a *Adapter) SpecExecute() {
	fmt.Println("chargging...")
}

type AdapterTransform struct {
	*Adapter
}

func (a AdapterTransform) Execute() {
	a.SpecExecute()
}

func main() {

}
