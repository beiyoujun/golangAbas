package builder

import "fmt"

type Builder interface {
	head()
	body()
	foot()
	hand()
}

type Thin struct {
}

func (*Thin) head() {
	fmt.Println("我的头很瘦")
}

func (*Thin) body() {
	fmt.Println("我的身体很瘦")
}

func (*Thin) foot() {
	fmt.Println("我的脚很瘦")
}

func (*Thin) hand() {
	fmt.Println("我的手很瘦")
}

type Fat struct {
}

func (*Fat) head() {
	fmt.Println("我的头很胖")
}

func (*Fat) body() {
	fmt.Println("我的身体很胖")
}

func (*Fat) foot() {
	fmt.Println("我的脚很胖")
}

func (*Fat) hand() {
	fmt.Println("我的手很胖")
}

type Director struct {
	person Builder
}

func (d *Director) CreatePerson() {
	if d == nil {
		return
	}
	d.person.hand()
	d.person.body()
	d.person.foot()
	d.person.hand()
}
