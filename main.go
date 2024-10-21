package main

import (
	"errors"
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

// ООПшное

type employee struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	// Это инициализатор или конструктор для memoryStorage, который возвращает указатель на структуру и инициализирует мапу при создании объекта
	// Если не инициализировать мапу, то получим ошибку при попытке записи в не инициализированную мапу
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with such id doesn't exist")
	}
	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

type dumbStorate struct{}

func newDumbStorage() *dumbStorate {
	return &dumbStorate{}
}

func (d dumbStorate) insert(e employee) error {
	fmt.Printf("Вставка пользователя с id: %d прошла успешно\n", e.id)
	return nil
}

func (d dumbStorate) get(id int) (employee, error) {
	e := employee{
		id: id,
	}
	return e, nil
}

func (d dumbStorate) delete(id int) error {
	fmt.Printf("Удаление пользователя с id: %d прошло успешно\n", id)
	return nil
}

func printType(value interface{}) {
	// любой тип в программе на Go по умолчанию удовлетворяет пустой интерфейс
	// Работая с пустыми интерфейсами, мы имеем возможность проверять интерфейс на его динамический тип
	switch value.(type) {
	case int:
		fmt.Println("тип аргумента int")
	case string:
		fmt.Println("тип аргумента string")
	default:
		fmt.Println("тип аргумента не int и не string")
	}
}

func testStorage() {
	var s storage

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = newMemoryStorage()
	// Без проблем присвоили в переменную s типа storage объект *newMemoryStorage,
	// т.к. этот тип соответствует интерфейсу, потому что обладает всеми необходимыми методами

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)
	fmt.Println("______________")

	s = newDumbStorage()
	// Мы заменили один тип на другой, который удовлетворяет требованиям интерфейса. Такое свойствой называется ВЗАИМОЗАМЕНЯЕМОСТЬ
	// Создали 2 структуры, которые реализовывают один и тот же интерфейс - это ПОЛИМОРФИЗМ

	fmt.Println("s", s)
	fmt.Printf("type of s: %T\n", s)

	s = nil

	fmt.Println("s", s)
	fmt.Printf("type of s: %T\n", s)

	printType(5)
	printType("stroka")
	printType([]string{"слайсы", "тоже"})
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

func main() {
	testModules()
}
