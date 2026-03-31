package main

import (
	"fmt"
	"maps"
)

func callPointers() {
	performIncrementP()
	performPushReplacingP()
	performAddReplacingP()
}

func incrementP(x *int) {
	fmt.Printf("@incrementP &x=%p\n", x)
	*x += 1
}

func performIncrementP() {
	println("increment:")
	a := 0
	fmt.Printf("@performIncrementP &a=%p\n", &a)

	incrementP(&a)
	fmt.Printf("a=%d\n", a)
	println("========================================")
}

func pushReplacingP(arr *[]int, v int) {
	fmt.Printf("@pushReplacingP &arr=%p\n", arr)
	*arr = append(*arr, v)
}

func performPushReplacingP() {
	println("pushReplacingP:")
	arr := []int{1, 2, 3}
	fmt.Printf("@performPushReplacingP &arr=%p\n", &arr)

	pushReplacingP(&arr, 4)
	fmt.Printf("arr=%v\n", arr)
	println("========================================")
}

func addReplacingP(obj *map[string]string, k string, v string) {
	fmt.Printf("@addReplacing &obj=%p\n", obj)
	var obj2 map[string]string = make(map[string]string)
	maps.Copy(obj2, *obj)
	obj2[k] = v
	*obj = obj2
	fmt.Printf("@addReplacing &obj=%p; &obj2=%p\n", obj, &obj2)
}

func performAddReplacingP() {
	println("addReplacingP:")
	obj := map[string]string{"apple": " fruit", "tomato": "fruit"}
	fmt.Printf("@performAddReplacingP &obj=%p\n", &obj)

	addReplacingP(&obj, "potato", "vegetable")
	fmt.Printf("obj=%#v\n", obj)
	println("========================================")
}

// https://goplay.tools/snippet/4kDSU7tTPKy
