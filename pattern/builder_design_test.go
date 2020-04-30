package pattern

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()
	fmt.Println(car)
	if car.Wheels != 4 {
		t.Error("not real car")
	}
}
