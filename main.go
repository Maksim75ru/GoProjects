package main

import (
	"GoProject/shape"
	"GoProject/storage"
	"fmt"
	"math"
	"time"
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

	// type employee struct { // Объявление абстрактной структуры. За пределами функции невозможно создать экземпляр
	// 	name   string
	// 	sex    string
	// 	age    int
	// 	salary int
	// }

	employee1 := storage.Employee{
		Name:   "Вася",
		Age:    25,
		Salary: 1500,
	}
	fmt.Printf("%+v\n", employee1)

	// Создание возможно за пределами, т.к. объявили структуру за пределами тела функций
	petyaEmployee := newAbsractEmployee("Петя", "М", 30, 2000)
	fmt.Println(petyaEmployee.getInfo())
	setName(&petyaEmployee, "Oleg")
	fmt.Println(petyaEmployee.getInfo())
}

func testMap() {
	ages := make(map[string]int) // make используется для инициализации пустой мапы

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

func testModules() {
	// s := scheduler.NewScheduler()

	// s.Add(
	// 	context.Background(),
	// 	func(ctx context.Context) { fmt.Printf("Текущее время: %s\n", time.Now()) },
	// 	time.Second,
	// )

	// time.Sleep(time.Minute)

}

func spawnEmployees(s storage.Storage) {
	for i := 1; i <= 10; i++ {
		s.Insert(storage.Employee{Id: i})
	}
}

func testStoragePackage() {
	ms := storage.NewMemoryStorage()
	ds := storage.NewDumbStorage()

	spawnEmployees(ms)

	fmt.Println(ms.Get(3))

	spawnEmployees(ds)
}

func testShapePackage() {
	s := shape.NewSquare(5)
	r := shape.NewRectangle(4, 5)

	fmt.Printf("Площадь квадрата: %d\n", s.GetSquare())
	fmt.Printf("Площадь прямоугольника: %d\n", r.GetSquare())

}

func testGorutines() {
	t := time.Now()

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	result1 := make(chan int)
	result2 := make(chan int)

	go calculateSomething(300, result1)
	go calculateSomething(500, result2)

	fmt.Println(<-result1)
	fmt.Println(<-result2)

	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))
}

func calculateSomething(n int, res chan int) {
	t := time.Now()

	result := 0
	for i := 0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Результат: %d; Прошло времени: %s\n", result, time.Since(t))
	res <- result
}

func testGorutines2() {
	numbers := make(chan int)

	go generateNumbers(1000, numbers)

	for {
		number, ok := <-numbers
		fmt.Println(number, ok)

		if !ok {
			break
		}
	}

}

func generateNumbers(n int, res chan int) {
	for i := 0; i <= n; i++ {
		res <- i * 2
	}

	close(res)
}

func testGorutines3() {
	number := make(chan int)

	go func() {
		fmt.Println("РАБОТАЕТ АНОНИМНАЯ ФУНКЦИЯ1")
		number <- 42
	}()

	go func() {
		fmt.Println("РАБОТАЕТ АНОНИМНАЯ ФУНКЦИЯ2")
		number <- 55
	}()

	time.Sleep(time.Millisecond * 100)

	select {
	case n := <-number:
		fmt.Println(n)
	default:
		fmt.Println("nothing, empty, completle nothing")
	}
}

func calculateFactorial(n int, ch chan int) {
	if n == 1 {
		fromChannel := <-ch
		ch <- n * fromChannel
		close(ch)

	}

	_, ok := <-ch

	if !ok {
		ch <- n
	} else {
		fromChannel := <-ch
		ch <- n * fromChannel
	}

	go calculateFactorial(n-1, ch)
}

func main() {
	n := 45
	ch := make(chan int)

	go func() {
		for _, r := range `\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go calculateFactorial(n, ch)

	result := <-ch
	fmt.Printf("Факториал числа %d равен %d\n", n, result)
}
