package main

import "fmt"

func main() {
	var gopro *GoProgrammer = new(GoProgrammer)
	var javapro *JavaProgrammer = new(JavaProgrammer)
	writeFirst(gopro)
	writeFirst(javapro)
}

type Code string

// 接口实现多态,使用接口的变量必须是指针类型
type Programmer interface{
	Write() Code
}

type GoProgrammer struct{

}

func (p *GoProgrammer) Write() Code{
	return "fmt.Println()"	
}

type JavaProgrammer struct{
	
}

func (p* JavaProgrammer) Write() Code{
	return "System.out"
}

func writeFirst(p Programmer){
	fmt.Printf("%T %v\n", p, p.Write())
}

