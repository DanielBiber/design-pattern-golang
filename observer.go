/*
观察者模式定义了对象之间的一对多依赖，这样一来，当一个对象改变状态时，他的所有依赖者都会收到通知并自动更新。 
观察者模式提供了一种对象设计，让主题和观察者之间松耦合.
如下例所示
观察者ABoard，BBoard订阅了主题WeatherData。一旦WeatherData变化，观察者立刻收到信息。
*/
package main

import (
	"fmt"
	"math/rand"
)

type observable interface {
	registerObserver(o observer)
	removeObserver(i int)
	notifyObserver()
}
type observer interface {
	update(a int)
}

type WeatherData struct {
	info int
	boardList []observer
}

func (w WeatherData) init() {
	w.info = 0
	w.boardList = []observer{}
}
func (w *WeatherData) registerObserver(o observer){
	w.boardList =append(w.boardList, o)
}
func (w *WeatherData) removeObserver(i int){
	target := w.boardList[:0]
	for j := 0;j < len(w.boardList);j++{
		for index, item := range w.boardList{
			if index != i{
				target = append(target,item)
			}
		}
	}
	w.boardList = target
}
func (w WeatherData) notifyObserver() {
	a := rand.Intn(100)
	for _,item := range w.boardList{
		item.update(a)
	}
}
type ABoard struct {
	subject *WeatherData
}

func (a ABoard) init(w *WeatherData)  {
	a.subject = w
}
func (a ABoard) update(b int)  {
	fmt.Println("ABoard update ", b)
}

type BBorad struct {
	subject *WeatherData
}

func (b BBorad) init(w *WeatherData)  {
	b.subject = w
}
func (b BBorad) update(a int)  {
	fmt.Println("BBorad update ",a)
}

func main() {
	w := new(WeatherData)
	a := &ABoard{}
	a.init(w)
	w.registerObserver(a)
	b := &BBorad{}
	b.init(w)
	w.registerObserver(b)
	w.notifyObserver()
}
