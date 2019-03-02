/*
策略模式定义了算法族，分别封装起来，让他们之间可以互相替换，此模式让算法的变化独立于使用算法的客户。
以下为策略模式：
例子:
设置了两个算法族接口（一个飞行行为，一个为叫声行为）。
父类duck类通过组合定义了该两个行为。
当有其他子类有特定的行为时，不需要修改代码，可以修改类中的行为属性。
遵从了多用组合的原则。
*/
package main
import "fmt"

//定义的行为接口，飞行行为与叫声行为
type FlyBehavior interface {
	fly()
}
type QuackBehavior interface{
	quack()
}

type Duck struct {
	flyBehavior   FlyBehavior  //通过组合得到行为
	quackBehavior QuackBehavior
	extra         string
}

func (d Duck) performQuack(){
	d.quackBehavior.quack()
}

func (d *Duck) setQuackBehavior(qb QuackBehavior){   //用指针才能真正修改
	d.quackBehavior = qb
}

func (d Duck) performFly(){
	d.flyBehavior.fly()
}

func (d *Duck) setFlyBehavior(fb FlyBehavior){  //同上理
	d.flyBehavior = fb
}
// 具体行为的实现
type Quack struct {}
type Fly struct{}
type Squeak struct{}
type FlyNoWay struct{}
func (sq Squeak) quack(){
	fmt.Println("Squeak")
}
func (f FlyNoWay) fly(){
	fmt.Println("can not fly")
}
func (d Quack) quack(){
	fmt.Println("quack")
}
func (d Fly) fly(){
	fmt.Println("fly")
}
//继承的子类。
type MallardDuck struct{
	duck Duck
}
type ModelDuck struct{
	duck Duck
}
func NewMallardDuck() *MallardDuck {
	return &MallardDuck{Duck{Fly{},Quack{},""}}
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
