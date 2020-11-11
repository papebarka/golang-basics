package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello World")
	var x int
	var y int = 70
	z := 65
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	if y > z {
		fmt.Println("Hello")
	} else if y < z {
		fmt.Println("There")
	} else {
		fmt.Println("Else")
	}

	var arrays [5]int
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println(arrays)
	fmt.Println(a)

	b := []int{5, 4, 3, 2, 1}
	b = append(b, 7)
	fmt.Println(b)

	human := make(map[string]int)
	human["age"] = 26
	human["year"] = 1994
	fmt.Println(human)
	fmt.Println(human["year"])
	delete(human, "year")
	fmt.Println(human)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	c := 0
	for c < 5 {
		fmt.Println(c)
		c++
	}

	myarr := []string{"a", "b", "c", "d"}
	for index, value := range myarr {
		fmt.Println("index:", index, "value:", value)
	}

	mymap := make(map[string]string)
	mymap["one"] = "Africa"
	mymap["two"] = "Europe"
	mymap["three"] = "Asia"

	for key, value := range mymap {
		fmt.Println("Key:", key, "Value:", value)
	}

	addition := sum(3, 7)
	fmt.Println(addition)

	squareroot, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SquareRoot:", squareroot)
	}

	persons()

	pointersExample()
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func persons() {
	p := person{name: "Isaac", age: 65}
	fmt.Println(p)
}

func pointersExample() {
	i := 10
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}
