package main

import (
	"fmt"
	"math"
)

func main() {
	testMap()
}

func testMap() {
	ages := make(map[string]int)

	ages["Максим"] = 30
	ages["Лолита"] = 30
	ages["Виктория"] = 17

	fmt.Println("ages:", ages)
	fmt.Printf("Максиму %d лет\n", ages["Максим"])

	newAges := map[string]int{
		"Олег":   32,
		"Виктор": 22,
	}
	age, exists := newAges["Антон"]
	if !exists {
		fmt.Println("Антона нет в списке")
	} else {
		fmt.Printf("Антону %d лет", age)
	}

	ages2 := ages

	delete(ages, "Максим")

	fmt.Println("ageeees", ages)
	fmt.Println("ages2", ages2)

}

func ChangeOrder() {
	arr := []string{"a1", "b1", "c1", "d1"}

	var new_arr []string

	for i := len(arr) - 1; i >= 0; i-- {
		new_arr = append(new_arr, arr[i])
	}
}

func testCopyAndMake() {
	arr := []string{"a1", "b1", "c1", "d1"}

	new_arr := make([]string, 4)

	copy(new_arr, arr)

	fmt.Println("arr:", arr)
	fmt.Println("new_arr:", new_arr)
}

func RaisingItems(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("2 в степени %d = %.f\n", i, math.Pow(2, 3))
	}
}

func fillArray(arr []string) []string {
	return append(arr, "e1")
}
