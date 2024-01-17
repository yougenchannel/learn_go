package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Animal interface {
	eat()
}

/**
1. 指针接收器和值接收器差异探究, 尽量保持一个struct的接收器统一（统一为值接收器或指针接收器）
2. 指针接收器能修改变量的值， 值接收器不能
3. 原则上值变量只能调用值接收器，但是go的语法糖允许调用指针接收器
4. 指针调用值接收器会将之争自动转换为值
*/
func main() {
	p1 := Person{age: 1, name: "yougen"}
	p1.sayHello()
	fmt.Println(p1.age) // 1
	//p1.setName("new Name")
	fmt.Println(p1.name)
	p2 := &p1
	p2.setName("p2 set name")
	printPerson(p2)
}

func (p *Person) sayHello() {
	fmt.Println("Hello")
	p.age = 2
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) String() string {
	return fmt.Sprintf("x {age:%d`	name:%s}", p.age, p.name)
}

func printPerson(v fmt.Stringer) {
	fmt.Println(v)
}
