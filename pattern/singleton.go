package pattern

type counter struct {
	count int
}

var instance *counter

func GetInstance() *counter {
	if instance == nil {
		instance = new(counter)
	}
	return instance
}

func (c counter) AddOne() int {
	c.count++
	return c.count
}
