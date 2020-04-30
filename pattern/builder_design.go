package pattern

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	SetLogo() BuildProcess
	GetVehicle() VehicleProduct
}

//MEMO 指导者
type ManufacturingDirector struct {
	builder BuildProcess
}

//MEMO 构造流程
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels().SetLogo()
}

//MEMO 选择产品类型
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

//MEMO 产品
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
	Logo      string
}

//MEMO 产品类型
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car builder"
	return c
}

func (c *CarBuilder) SetLogo() BuildProcess {
	c.v.Logo = "BMW"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
