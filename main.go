package main

import (
	"errors"
	"fmt"
)

const pi float32 = 3.1415

func main() {
	printCircleArea(0)
	printCircleArea(6)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: S=pr2")

	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}
	return float32(radius) * float32(radius) * pi, nil
}
