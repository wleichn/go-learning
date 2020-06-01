package StructMethodInterface

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Shape interface {
	Area() float64
}

/**
go语言中的接口是隐式的，只要传入的类型匹配接口，则可编译正确：
Rectangle有一个返回类型为float64的方法Area，所以满足接口Shape;
...
string 没有这种方法，所以不满足这个接口
*/

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
