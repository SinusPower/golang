package main

import (
	"fmt"

	hw04 "github.com/sinuspower/golang/test/hw04_lru_cache"
)

func main() {
	testList := hw04.NewList()
	fmt.Println("After initialization =>")
	fmt.Println("list: ", testList)
	fmt.Println("len: ", testList.Len())
	fmt.Println("front: ", testList.Front())
	fmt.Println("back: ", testList.Back())
	fmt.Printf("____________________________\n\n")

	i := 1
	toFront := testList.PushFront(i)
	fmt.Printf("After PushFront(%d) =>\n", i)
	fmt.Println("list: ", testList)
	fmt.Println("len: ", testList.Len())
	fmt.Println("front: ", testList.Front())
	fmt.Println("back: ", testList.Back())
	fmt.Printf("____________________________\n\n")

	i = 2
	testList.PushFront(i)
	fmt.Printf("After PushFront(%d) =>\n", i)
	fmt.Println("list: ", testList)
	fmt.Println("len: ", testList.Len())
	fmt.Println("front: ", testList.Front())
	fmt.Println("back: ", testList.Back())
	fmt.Printf("____________________________\n\n")

	i = 3
	testList.PushFront(i)
	fmt.Printf("After PushFront(%d) =>\n", i)
	fmt.Println("list: ", testList)
	fmt.Println("len: ", testList.Len())
	fmt.Println("front: ", testList.Front())
	fmt.Println("back: ", testList.Back())
	fmt.Printf("____________________________\n\n")

	i = 0
	toRemove := testList.PushBack(i)
	fmt.Printf("After PushBack(%d) =>\n", i)
	fmt.Println("list: ", testList)
	fmt.Println("len: ", testList.Len())
	fmt.Println("front: ", testList.Front())
	fmt.Println("back: ", testList.Back())
	fmt.Printf("____________________________\n\n")

	for i := 4; i < 10; i++ {
		testList.PushFront(i)
	}

	for i := -1; i > -10; i-- {
		testList.PushBack(i)
	}

	l := testList.Back()
	for l != nil {
		fmt.Println("item: ", l)
		fmt.Println("data: ", l.Value)
		fmt.Println("prev: ", l.Prev)
		fmt.Println("next: ", l.Next)
		fmt.Println("___")
		l = l.Next
	}
	fmt.Printf("____________________________\n\n")

	testList.Remove(toRemove)
	fmt.Printf("After Remove(%d) =>\n", toRemove.Value)
	l = testList.Back()
	for l != nil {
		fmt.Println("item: ", l)
		fmt.Println("data: ", l.Value)
		fmt.Println("prev: ", l.Prev)
		fmt.Println("next: ", l.Next)
		fmt.Println("___")
		l = l.Next
	}
	fmt.Printf("____________________________\n\n")

	testList.Remove(testList.Front())
	fmt.Printf("After Remove(testList.Front()) =>\n")
	l = testList.Back()
	for l != nil {
		fmt.Println("item: ", l)
		fmt.Println("data: ", l.Value)
		fmt.Println("prev: ", l.Prev)
		fmt.Println("next: ", l.Next)
		fmt.Println("___")
		l = l.Next
	}
	fmt.Printf("____________________________\n\n")

	testList.Remove(testList.Back())
	fmt.Printf("After Remove(testList.Back()) =>\n")
	l = testList.Back()
	for l != nil {
		fmt.Println("item: ", l)
		fmt.Println("data: ", l.Value)
		fmt.Println("prev: ", l.Prev)
		fmt.Println("next: ", l.Next)
		fmt.Println("___")
		l = l.Next
	}
	fmt.Printf("____________________________\n\n")

	testList.MoveToFront(toFront)
	fmt.Printf("After MoveToFront(%d)) =>\n", toFront.Value)
	l = testList.Back()
	for l != nil {
		fmt.Println("item: ", l)
		fmt.Println("data: ", l.Value)
		fmt.Println("prev: ", l.Prev)
		fmt.Println("next: ", l.Next)
		fmt.Println("___")
		l = l.Next
	}
	fmt.Printf("____________________________\n\n")
}
