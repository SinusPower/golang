package main

import (
	"fmt"
	"strconv"

	hw04 "github.com/sinuspower/golang/test/hw04_lru_cache"
)

func listTest() {
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

	for j := 0; j < 5; j++ {
		b := testList.Back().Value
		testList.MoveToFront(testList.Back())
		fmt.Printf("After MoveToFront(%d)) =>\n", b)
		l = testList.Back()
		for l != nil {
			fmt.Println("item: ", l)
			fmt.Println("data: ", l.Value)
			fmt.Println("prev: ", l.Prev)
			fmt.Println("next: ", l.Next)
			fmt.Println("___")
			l = l.Next
		}
	}
	fmt.Printf("____________________________\n\n")
}

func cacheTest() {
	cache := hw04.NewCache(3)
	fmt.Println("After initialization =>")
	fmt.Printf("cache: %v\n", cache)
	fmt.Printf("____________________________\n\n")

	var f bool
	for i := 0; i < 10; i++ {
		f = cache.Set(hw04.Key("item"+strconv.Itoa(i)), i)
		fmt.Printf("After Set(%d) =>\n", i)
		fmt.Printf("existed: %v\n", f)
		fmt.Printf("cache: %v\n", cache)
		q := cache.GetQueue()
		fmt.Printf("queue: <-- ")
		for itm := q.Front(); itm != nil; itm = itm.Prev {
			fmt.Printf("%v ", itm.Value)
		}
		fmt.Printf("<--\n")
		fmt.Printf("____________________________\n\n")
	}

	var v interface{}
	for i := 0; i < 10; i++ {
		v, f = cache.Get(hw04.Key("item" + strconv.Itoa(i)))
		fmt.Printf("After Get(%d) =>\n", i)
		fmt.Printf("item: %v\n", v)
		fmt.Printf("found: %v\n", f)
		fmt.Printf("cache: %v\n", cache)
		q := cache.GetQueue()
		fmt.Printf("queue: <-- ")
		for itm := q.Front(); itm != nil; itm = itm.Prev {
			fmt.Printf("%v ", itm.Value)
		}
		fmt.Printf("<--\n")
		fmt.Printf("____________________________\n\n")
	}

	for i := 7; i < 10; i++ {
		f = cache.Set(hw04.Key("item"+strconv.Itoa(i)), i+10)
		fmt.Printf("After Set(%d) =>\n", i)
		fmt.Printf("existed: %v\n", f)
		fmt.Printf("cache: %v\n", cache)
		q := cache.GetQueue()
		fmt.Printf("queue: <-- ")
		for itm := q.Front(); itm != nil; itm = itm.Prev {
			fmt.Printf("%v ", itm.Value)
		}
		fmt.Printf("<--\n")
		fmt.Printf("____________________________\n\n")
	}
}

func main() {
	// listTest()
	cacheTest()
}
