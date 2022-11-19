package main

import "fmt"

type Duck struct {
	QuackAble
	FlyAble

	Name string
}

type QuackAble interface {
	Quack() string
}

type FlyAble interface {
	Fly() string
}

type Quack struct{}
type Squea struct{}
type MuteQuack struct{}

func (d *Quack) Quack() string {
	return "Quack quack"
}

func (d *Squea) Quack() string {
	return "Squea quack"
}

func (d *MuteQuack) Quack() string {
	return "MuteQuack quack"
}

type FlyWithWings struct{}
type FlyNoWay struct{}

func (d *FlyWithWings) Fly() string {
	return "FlyWithWings fly"
}

func (d *FlyNoWay) Fly() string {
	return "FlyNoWay fly"
}

func main() {
	duck := new(Duck)
	duck.Name = "木头鸭子"
	if duck.Name == "普通鸭子" {
		duck.QuackAble = new(Quack)
		duck.FlyAble = new(FlyWithWings)
	} else if duck.Name == "木头鸭子" {
		duck.QuackAble = new(Squea)
		duck.FlyAble = new(FlyNoWay)
	}

	fmt.Println(duck.Quack())
	fmt.Println(duck.Fly())
}
