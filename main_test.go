package main

import "testing"

func TestAdapterTransform_Execute(t *testing.T) {
	a := AdapterTransform{}
	a.Execute()
	b := Adapter{}
	b.SpecExecute()
}
