package main

import "fmt"

//建造这模式

type Color string
type Sex string

type Interface interface {
	SaySex() error
	SayColor() error
}

type Cat struct {
	color string
	sex string
}

func (c *Cat)SayColor() error{
	fmt.Printf("color is %v",c.color)
	fmt.Println()
	return nil
}

func (c *Cat)SaySex() error{
	fmt.Printf("sex is %v",c.sex)
	fmt.Println()
	return nil
}


type Builder interface {
	SetColor(color string) Builder
	SetSex(sex string) Builder
	Build() Interface
}

type CatBuilder struct {
	color string
	sex string
}

func (c *CatBuilder)SetColor(color string) Builder{
	c.color = color
	return c
}
func (c *CatBuilder)SetSex(sex string) Builder{
	c.sex = sex
	return c
}
func  (c *CatBuilder)Build() Interface{
	cat := &Cat{}
	cat.color = c.color
	cat.sex = c.sex
	return cat
}

func NewBuilder() Builder{
	return &CatBuilder{}
}

func main(){
	builder := NewBuilder()
	cat := builder.SetSex("male").SetColor("red").Build()
	cat.SayColor()
	cat.SaySex()
}
