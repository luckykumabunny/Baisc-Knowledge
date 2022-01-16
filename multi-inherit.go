package main

import "fmt"

func main() {
	var dog *Dog = new(Dog)
	dog.Speak()
	dog.Speakto()
	var pet *Pet = new(Pet)
	pet.Speak()
}

type Pet struct{
	
}

// 隐式组合嵌套
type Dog struct{
	Pet
}


func (p *Pet) Speak(){
	fmt.Print("bunny love uki")
}

// go其实不支持多其他语言那样的多态，无法加载子类数据的方法
func (p *Pet) Speakto{
	p.Speak()
}

func (d *Dog) Speak(){
	fmt.Print("bunny love kuma")
}


