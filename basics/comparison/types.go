package main

import (
	"fmt"
	"maps"
)

func callBasicTypes() {
	performIncrement()
	performConcat()
	performPush()
	performPushReplacing()
	performAdd()
	performAddReplacing()
}

func increment(x int) {
	x += 1
}

func performIncrement() {
	println("increment:")
	a := 0
	increment(a)
	fmt.Printf("a=%d\n", a)
	println("========================================")
}

func concat(a, b string) {
	a += b
}

func performConcat() {
	println("concat:")
	abc := "abc"
	concat(abc, "def")
	fmt.Printf("abc=%s\n", abc)
	println("========================================")
}

func push(arr []int, pos int, v int) {
	arr[pos] = v
	fmt.Printf("@push &arr=%p\n", &arr)
}

func performPush() {
	println("push:")
	arr := []int{1, 2, 3, 0}
	fmt.Printf("@performPush &arr=%p\n", &arr)

	push(arr, 3, 4)
	fmt.Printf("arr=%v\n", arr)
	println("========================================")
}

func pushReplacing(arr []int, v int) {
	arr = append(arr, v)
	fmt.Printf("@pushReplacing &arr=%p\n", &arr)
}

func performPushReplacing() {
	println("pushReplacing:")
	arr2 := []int{1, 2, 3}
	fmt.Printf("@performPushReplacing &arr2=%p\n", &arr2)

	pushReplacing(arr2, 4)
	fmt.Printf("arr2=%v\n", arr2)
	println("========================================")
}

func add(m map[string]string, k string, v string) {
	m[k] = v
	fmt.Printf("@add &m=%p\n", &m)
}

func performAdd() {
	println("add:")
	obj := map[string]string{"apple": " fruit", "tomato": "fruit"}
	fmt.Printf("@performAdd &obj=%p\n", &obj)

	add(obj, "potato", "vegetable")
	fmt.Printf("obj=%#v\n", obj)
	println("========================================")
}

func addReplacing(obj map[string]string, k string, v string) {
	fmt.Printf("@addReplacing &obj=%p\n", &obj)
	var obj2 map[string]string = make(map[string]string)
	maps.Copy(obj2, obj)
	obj2[k] = v
	obj = obj2
	fmt.Printf("@addReplacing &obj=%p; &obj2=%p\n", &obj, &obj2)
}

func performAddReplacing() {
	println("addReplacing:")
	obj2 := map[string]string{"car": "vehicle", "bike": "vehicle"}
	fmt.Printf("@performAddReplacing &obj2=%p\n", &obj2)

	addReplacing(obj2, "bus", "vehicle")
	fmt.Printf("obj2=%#v\n", obj2)
	println("========================================")
}

// https://goplay.tools/snippet/fcXcZyqYSE8
