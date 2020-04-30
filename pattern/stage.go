package pattern

import "fmt"

//MEMO 接口
type ITravel interface {
	Travel()
}

type Airplane struct {
}

type Train struct {
}

type Walk struct {
}

//MEMO 具体策略
func (a Airplane) Travel() {
	fmt.Println("airplane travel")
}

func (t Train) Travel() {
	fmt.Println("train travel")
}

func (w Walk) Travel() {
	fmt.Println("walk travel")
}

//MEMO 上下文
type person struct {
	Name    string
	ITravel //MEMO 可以直接赋值接口
}

func (p person) travel() {
	fmt.Println(p.Name)
	p.Travel()
}

func DoTravel() {
	p := person{"xmanchai", Airplane{}}
	p.travel()
}
