# 设计模式（design pattern）

## 介绍 （introduction）
> 使用模式的最好的方式是：“把模式装进脑子里，然后在你的设计和已有的应用中，寻找何处可以使用它们”。
为了更加熟悉设计模式，本项目采取了用golang语言实现《Head First 设计模式》中介绍的各个模型的例子。

### 实现的模式有：
[策略模式](https://github.com/DanielBiber/design-pattern-golang/blob/master/strategy.go)；
### 记录的设计原则有：
[link](https://github.com/DanielBiber/design-pattern-golang/blob/master/principle.md)
封装变化；多用组合，少用继承；针对接口编程，不针对实现编程；

## 策略模式
策略模式定义了算法族，分别封装起来，让他们之间可以互相替换，此模式让算法的变化独立于使用算法的客户。
