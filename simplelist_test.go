package main

import (
	"testing"
	// "fmt"
)

func TestEmptyList(t *testing.T) {

    list := NewList()
    _, err := list.Find("fred")
    if(err == nil){
		t.Error("New list should be empty")
    }
}

func TestAddValue(t *testing.T) {
	list := NewList()
	list.Add("One")

	expected := "One"
	item, _ := list.Find("One")

    if(expected != item.Value){
		t.Error("AddValues: Error in value")
    }
}

func TestValues(t *testing.T) {
	list := NewList()
	list.Add("One")

	expected := []string{"One"}
	actual := list.Values()

    if(len(expected) != len(actual)){
		t.Error("AddValues: Length error")
	}

	for i, v := range actual {
        //fmt.Printf("%d,%s, %s\n", i, v, expected[i])
        if(expected[i] != v) {
        	t.Error("AddValues: Error in value")
        }
    }
}

func TestAddValues(t *testing.T) {
	list := NewList()
	list.Add("One")
	list.Add("Two")
	list.Add("Three")

	expected := []string{"One","Two","Three"}
	actual := list.Values()

	if(len(expected) != len(actual)){
		t.Error("AddValues: Length error")
	}

	for i, v := range actual {
        if(expected[i] != v) {
        	t.Error("AddValues: Error in value")
        }
    }
}

func TestDeleteValues(t *testing.T) {
	list := NewList()
	list.Add("One")
	list.Add("Two")
	list.Add("Three")

 	item, _ := list.Find("Two")

	list.Delete(item)

	expected := []string{"One","Three"}
	actual := list.Values()

	if(len(expected) != len(actual)){
		t.Error("AddValues: Length error")
	}

	for i, v := range actual {
        if(expected[i] != v) {
        	t.Error("AddValues: Error in value")
        }
    }
}
