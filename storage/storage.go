package storage

import (
	"errors"
	"fmt"
)

// ООПшное

type Employee struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e Employee) error
	get(id int) (Employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]Employee
}

func newMemoryStorage() *memoryStorage {
	// Это инициализатор или конструктор для memoryStorage, который возвращает указатель на структуру и инициализирует мапу при создании объекта
	// Если не инициализировать мапу, то получим ошибку при попытке записи в не инициализированную мапу
	return &memoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *memoryStorage) insert(e Employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("Employee with such id doesn't exist")
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

func (d dumbStorate) insert(e Employee) error {
	fmt.Printf("Вставка пользователя с id: %d прошла успешно\n", e.id)
	return nil
}

func (d dumbStorate) get(id int) (Employee, error) {
	e := Employee{
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

func main() {
	printType("stroka")
}
