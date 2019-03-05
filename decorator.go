/*
装饰者模式动态地将责任附加到对象上。若要扩展功能，装饰者提供了比继承更有弹性的替代方案。 装饰者模式主要有Component、ConcreteComponent、Decorator和ConcreteDecorator组成：

抽象组件角色(Component)：定义一个对象接口，以规范准备接受附加责任的对象，即可以给这些对象动态地添加职责。
具体组件角色(ConcreteComponent) :被装饰者，定义一个将要被装饰增加功能的类。可以给这个类的对象添加一些职责。
抽象装饰者(Decorator):维持一个指向构件Component对象的实例，并定义一个与抽象组件角色Component接口一致的接口。(golang 不一定需要该接口)
具体装饰者角色（ConcreteDecorator):向组件添加职责。

具体实例如下所示：
抽象组件为beverage，定义了描述和价格接口。具体组件有HouserBlend,Espresso。
抽象装饰者CondimenDecorator,增加了一个额外的方法donothing，具体装饰着角色为：Mocha,Soy
该过程为在各种饮料中添加各种配料，并且能很简便的计算出加了各种配料后的饮料价格。很容易增加饮料种类与配料种类。
*/
package main

import "fmt"

type Beverage interface {
	getDescription() string
	cost() float32
}
type CondimenDecorator interface {
	getDescription() string
	doNothing()
}
type HouseBlend struct {
	description string
}
func (h HouseBlend) getDescription() string{
	return h.description
}
func (h HouseBlend) cost() float32 {
	return 1
}
func NewHouseBlend() HouseBlend {
	return HouseBlend{"houseBlend"}
}
type Espresso struct {
	description string
}
func (e Espresso) getDescription() string{
	return e.description
}
func (e Espresso) cost() float32 {
	return 2
}
func NewEspresso() Espresso{
	return Espresso{"espresso	"}
}
type Mocha struct {
	b Beverage
}

func (m Mocha) getDescription()  string{
    return	m.b.getDescription() + ", Mocha"
}
func (m Mocha) cost() float32 {
	return m.b.cost() + 0.2
}
func (m Mocha) doNothing(){

}
func NewMocha(b Beverage) Mocha{
	return Mocha{b}
}
type Soy struct{
	b Beverage
}
func (s Soy) getDescription()  string{
	return	s.b.getDescription() + ", Soy"
}
func (s Soy) cost() float32 {
	return s.b.cost() + 0.3
}
func (s Soy) doNothing(){

}
func NewSoy(b Beverage) Soy{
	return Soy{b}
}
func main()  {
	var b Beverage = NewEspresso()
	fmt.Println(b.getDescription()," $",b.cost())
	var b2 Beverage = NewHouseBlend()
	b2 = NewMocha(b2)
	b2 = NewSoy(b2)
	fmt.Println(b2.getDescription()," $",b2.cost())
}
