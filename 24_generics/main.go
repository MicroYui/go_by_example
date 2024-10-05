package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("name: %s, age: %d", p.name, p.age)
}

type Result[T any] struct {
	code    int
	message string
	data    T
}

func (r Result[T]) success() Result[T] {
	r.code = 200
	r.message = "response successfully"
	return r
}

func (r Result[T]) fail() Result[T] {
	r.code = 400
	r.message = "response failed"
	return r
}

func (r Result[T]) Data(data T) Result[T] {
	r.data = data
	return r
}

func (r Result[T]) String() string {
	return fmt.Sprintf("Result={code: %d message: %s data: {%v}}", r.code, r.message, r.data)
}

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {

	fmt.Println(new(Result[person]).success().Data(person{name: "MicroYui", age: 22}))

	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SliceIndex(s, "zoo"))

	_ = SliceIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())

}
