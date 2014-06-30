package main

import (
	"errors"
)

type List struct {
    Head *Item
    Last *Item
}

type Item struct {
	Value string
    Next *Item
}

func NewList()(*List){
    return &List{}
}

func NewItem(item string)(*Item){
	return &Item{Value: item, Next: nil}
}

func (l *List) Find(value string) (*Item, error) {
    var found *Item
    item := l.Head
	for item != nil {
		if(item.Value == value){
			found = item
			break
		}
		item = item.Next
	}
	if(found != nil){
		return found, nil
	}
    return nil, errors.New("not found")
}

func (l *List) Add(item string) {
   new_item :=NewItem(item)
   if(l.Last == nil){
       l.Last = new_item
       l.Head = l.Last
   }else{
	  l.Last.Next = new_item //algo con punteros here
      l.Last = new_item
   }
}

func (l *List) Values() ([]string){
	result := []string{}
	item := l.Head
	for item != nil {
		result = append(result, item.Value)
		item = item.Next
	}
	return result
}

func (l *List) Delete(item *Item) {
	previous := l.previous(item)
	if(previous != nil){
		previous.Next = item.Next
	}
}

func (l *List)previous(item *Item)(*Item){
	for p:= l.Head; p != nil; p = p.Next {
		if(p.Next == item){
			return p
		}
	}
	return nil
}

func main() {

}
