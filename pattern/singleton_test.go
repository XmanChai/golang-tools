package pattern

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	expectedCounter := GetInstance()
	fmt.Printf("%v", expectedCounter)
	if expectedCounter == nil {
		t.Error("Wrong initial constructor of counter")
	}
	expectedValue := expectedCounter.AddOne()
	fmt.Printf("%v", expectedValue)
	if expectedValue != 1 {
		t.Error("wrong counter method result")
	}
}
