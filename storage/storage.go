package storage

import (
	"errors"
	"fmt"
)

// ООПшное

type Employee struct {
	Id     int
	Name   string
	Age    int
	Salary int
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type MemoryStorage struct {
	data map[int]Employee
}

func NewMemoryStorage() *MemoryStorage {
	// Это инициализатор или конструктор для memoryStorage, который возвращает указатель на структуру и инициализирует мапу при создании объекта
	// Если не инициализировать мапу, то получим ошибку при попытке записи в не инициализированную мапу
	return &MemoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *MemoryStorage) Insert(e Employee) error {
	s.data[e.Id] = e
	return nil
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("Employee with such id doesn't exist")
	}
	return e, nil
}

func (s *MemoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}

type DumbStorate struct{}

func NewDumbStorage() *DumbStorate {
	return &DumbStorate{}
}

func (d DumbStorate) Insert(e Employee) error {
	fmt.Printf("Вставка пользователя с id: %d прошла успешно\n", e.Id)
	return nil
}

func (d DumbStorate) Get(id int) (Employee, error) {
	e := Employee{
		Id: id,
	}
	return e, nil
}

func (d DumbStorate) Delete(id int) error {
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
	var s Storage

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = NewMemoryStorage()
	// Без проблем присвоили в переменную s типа storage объект *newMemoryStorage,
	// т.к. этот тип соответствует интерфейсу, потому что обладает всеми необходимыми методами

	fmt.Println("s == nil", s == nil)
	fmt.Printf("type of s: %T\n", s)
	fmt.Println("______________")

	s = NewDumbStorage()
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
