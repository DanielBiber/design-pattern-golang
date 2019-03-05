# 设计模式（design pattern）

## 介绍 （introduction）
> 使用模式的最好的方式是：“把模式装进脑子里，然后在你的设计和已有的应用中，寻找何处可以使用它们”。
为了更加熟悉设计模式，本项目采取了用golang语言实现《Head First 设计模式》中介绍的各个模型的例子。

### 实现的模式有：
[策略模式](https://github.com/DanielBiber/design-pattern-golang/blob/master/strategy.go)；
[观察者模式](https://github.com/DanielBiber/design-pattern-golang/blob/master/observer.go);
[装饰器模式]（https://github.com/DanielBiber/design-pattern-golang/blob/master/decorator.go）;
### 记录的设计原则有：
[link](https://github.com/DanielBiber/design-pattern-golang/blob/master/principle.md)
封装变化；多用组合，少用继承；针对接口编程，不针对实现编程；

## 策略模式
策略模式定义了算法族，分别封装起来，让他们之间可以互相替换，此模式让算法的变化独立于使用算法的客户。
## 观察者模式
观察者模式定义了对象之间的一对多依赖，这样一来，当一个对象改变状态时，他的所有依赖者都会收到通知并自动更新。 观察者模式提供了一种对象设计，让主题和观察者之间松耦合
##ecorator Pattern（装饰者模式）
装饰者模式动态地将责任附加到对象上。若要扩展功能，装饰者提供了比继承更有弹性的替代方案。 装饰者模式主要有Component、ConcreteComponent、Decorator和ConcreteDecorator组成：
+ 抽象组件角色(Component)：定义一个对象接口，以规范准备接受附加责任的对象，即可以给这些对象动态地添加职责。
+ 具体组件角色(ConcreteComponent) :被装饰者，定义一个将要被装饰增加功能的类。可以给这个类的对象添加一些职责。
+ 抽象装饰者(Decorator):维持一个指向构件Component对象的实例，并定义一个与抽象组件角色Component接口一致的接口。+ 
具体装饰者角色（ConcreteDecorator):向组件添加职责。
