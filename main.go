package main

import (
	"fmt"
	"math"
)

type absractEmployee struct { // Объявление обычно происходит вне тела функции, чтобы можно было создать экземпляр в любом месте программы
	name   string
	sex    string
	age    int
	salary int
}

func newAbsractEmployee(name, sex string, age, salary int) absractEmployee {
	// Такие функции называются конструкторами или инициализаторами. Важно начинать название с newНазваниСтруктуры
	return absractEmployee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (e absractEmployee) getInfo() string {
	// Метод структуры absractEmployee - func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
	// e - это ресивер, хорошая практика 1-2 символа в названии
	return fmt.Sprintf("Сотрудник: %s\nВозраст: %d\nЗарплата: %d\n", e.name, e.age, e.salary)
}

func setName(e *absractEmployee, name string) {
	e.name = name
}

func main() {
	testStructure()
}

func testStructure() {
	// employee := struct {  // Объявление структуры с данными
	// 	name   string
	// 	sex    string
	// 	age    int
	// 	salary int
	// }{
	// 	name:   "Вася",
	// 	sex:    "М",
	// 	age:    25,
	// 	salary: 1500,
	// }
	// fmt.Printf("%+v\n", employee)

	type employee struct { // Объявление абстрактной структуры. За пределами функции невозможно создать экземпляр
		name   string
		sex    string
		age    int
		salary int
	}

	employee1 := employee{
		name:   "Вася",
		sex:    "М",
		age:    25,
		salary: 1500,
	}
	fmt.Printf("%+v\n", employee1)

	// Создание возможно за пределами, т.к. объявили структуру за пределами тела функций
	petyaEmployee := newAbsractEmployee("Петя", "М", 30, 2000)
	fmt.Println(petyaEmployee.getInfo())
	setName(&petyaEmployee, "Oleg")
	fmt.Println(petyaEmployee.getInfo())
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

	fmt.Println("ages", ages)
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
