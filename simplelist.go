package main

import (
	//"errors"
)

type List struct {

}

type Item struct {

}

func NewList()(*List){
	return &List{}
}

func (l *List) Find(item string) (Item, error) {
	//return Item {}, errors.New("empty")
	return Item {} , nil
}

func (l *List) Add(item string) {

}

func (l *List) Values() ([]string){
	return nil
}

func (l *List) Delete(Item) {

}

func main() {

}