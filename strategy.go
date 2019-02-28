/*
策略模式定义了算法族，分别封装起来，让他们之间可以互相替换，此模式让算法的变化独立于使用算法的客户。

*/
package main
import "fmt"

type FlyBehavior interface {
	fly()
}
type QuackBehavior interface{
	quack()
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
	extra         string
}

func (d Duck) performQuack(){
	d.quackBehavior.quack()
}

func (d Duck) setQuackBehavior(qb QuackBehavior){
	d.quackBehavior = qb
}

func (d Duck) performFly(){
	d.flyBehavior.fly()
}

func (d *Duck) setFlyBehavior(fb FlyBehavior){
	d.flyBehavior = fb
}

type Quack struct {}
type Fly struct{}

func (d Quack) quack(){
	fmt.Println("quack")
}
func (d Fly) fly(){
	fmt.Println("fly")
}

type Squeak struct{}
func (sq Squeak) quack(){
	fmt.Println("Squeak")
}
type FlyNoWay struct{}
func (f FlyNoWay) fly(){
	fmt.Println("can not fly")
}
type MallardDuck struct{
	duck Duck
}
func NewMallardDuck() *MallardDuck {
	return &MallardDuck{Duck{Fly{},Quack{},""}}
}
type ModelDuck struct{
	duck Duck
}

func NewModelDuck() *ModelDuck {
	return &ModelDuck{Duck{FlyNoWay{},Quack{},""}}
}

func main(){
	mallar := NewMallardDuck()
	mallar.duck.performFly()
	model := NewModelDuck()
	model.duck.performFly()
	model.duck.setFlyBehavior(Fly{})
	model.duck.performFly()
}
